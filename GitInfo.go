package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
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
		repoMap["description"] = repo.Description
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

func printData(data interface{}) (string, error) {
	output, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return "", err
	}
	return string(output), err
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

		dataOutput, err := printData(&userData)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(dataOutput)

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

		dataOutput, err := printData(&repoData)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(dataOutput)
	}
}
