package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var client *http.Client

type UserData struct {
	UserName  string                 `json:"login"`
	Name      string                 `json:"name"`
	Bio       string                 `json:"bio"`
	Location  string                 `json:"location"`
	AvatarUrl string                 `json:"avatar_url"`
	GitHubUrl string                 `json:"url"`
	NumRepos  int                    `json:"public_repos"`
	Followers int                    `json:"followers"`
	Repos     map[int]map[string]any `json:"repos"`
}

type RepoData struct {
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	RepositoryUrl   string            `json:"html_url"`
	License         map[string]string `json:"license"`
	ForksCount      int               `json:"forks_count"`
	OpenIssuesCount int               `json:"open_issues_count"`
	Languages       map[string]int    `json:"languages"`
}

func getJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func getRepos(userName string, userData *UserData) error {
	url := "https://api.github.com/users/" + userName + "/repos"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var repos []RepoData
	err = json.Unmarshal(body, &repos)
	if err != nil {
		return err
	}

	userData.Repos = make(map[int]map[string]interface{})
	for i, repo := range repos {
		repoMap := make(map[string]interface{})
		repoMap["name"] = repo.Name
		userData.Repos[i] = repoMap
	}

	return nil
}

func GetUserData(userName string) (*UserData, error) {
	var userData UserData

	url := "https://api.github.com/users/" + userName
	err := getJson(url, &userData)
	if err != nil {
		return nil, err
	}

	err = getRepos(userName, &userData)
	if err != nil {
		return nil, err
	}

	return &userData, nil
}

func getRepoLanguageData(url string, repoData *RepoData) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var languages map[string]int
	err = json.Unmarshal(body, &languages)
	if err != nil {
		return err
	}

	repoData.Languages = languages
	return nil
}

func GetRepoData(username string, repoName string) (*RepoData, error) {
	var repoData RepoData

	url := "https://api.github.com/repos/" + username + "/" + repoName
	err := getJson(url, &repoData)
	if err != nil {
		return nil, err
	}

	languagesUrl := url + "/languages"
	err = getRepoLanguageData(languagesUrl, &repoData)
	if err != nil {
		return nil, err
	}

	return &repoData, nil
}

func printUserData(userData *UserData) {
	fmt.Printf(`
Username:    %s
Name:        %s
Bio:         %s
Location:    %s
AvatarUrl:   %s
GitHubUrl:   %s
NumRepos:    %d
Followers:   %d
Repos:
`, userData.UserName, userData.Name, userData.Bio, userData.Location, userData.AvatarUrl,
		userData.GitHubUrl, userData.NumRepos, userData.Followers)

	for _, repo := range userData.Repos {
		fmt.Printf("\t- %s\n", repo["name"])
	}
}

func formatLanguages(languages map[string]int) string {
	languageStrs := make([]string, 0, len(languages))
	for language, count := range languages {
		languageStrs = append(languageStrs, fmt.Sprintf("%s: %d", language, count))
	}
	return strings.Join(languageStrs, ", ")
}

func printRepoData(repoData *RepoData) {
	fmt.Printf(`
Name:            %s
Description:     %s
RepositoryUrl:   %s
License:         %s
ForksCount:      %d
Languages:       %s
`, repoData.Name, repoData.Description, repoData.RepositoryUrl, repoData.License["name"],
		repoData.ForksCount, formatLanguages(repoData.Languages))
}

func main() {
	info := flag.String("info", "", "specifies the type of info")
	username := flag.String("username", "", "specifies github user")
	repoName := flag.String("repo", "", "specifies github repository name")
	client = &http.Client{Timeout: 10 * time.Second}

	flag.Parse()

	switch *info {
	case "user":
		if *username == "" {
			fmt.Println("(!) Please provide a username")
			return
		}

		userData, err := GetUserData(*username)
		if err != nil {
			fmt.Println(err)
			return
		}

		printUserData(userData)

	case "repo":
		if *username == "" {
			fmt.Println("(!) Please provide a username")
			return
		}

		if *repoName == "" {
			fmt.Println("(!) Please provide a repository name")
			return
		}

		repoData, err := GetRepoData(*username, *repoName)
		if err != nil {
			fmt.Println(err)
			return
		}

		printRepoData(repoData)
	}
}
