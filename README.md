# GitInfo
This repository is under heavy development! This is just the beginning of a simple CLI-Tool that collects data from GitHub users and repositories and prints it to the terminal.

## Examples
### Getting user data
```
go run GitInfo.go --info user --username menaruben
```
The output is just a huge json as of right now. In the future the output will be displayed in a beautiful table.
```
{
        "login": "menaruben",
        "name": "Rubén Mena",
        "bio": "19 year old nerd spending too much time on the computer :)",
        "location": "Switzerland, Zürich",
        "avatar_url": "https://avatars.githubusercontent.com/u/74914224?v=4",
        "url": "https://api.github.com/users/menaruben",
        "public_repos": 17,
        "followers": 4,
        "repos": {
                "0": {
                        "description": "This repository contains an automated installation of apache guacamole 1.5.0 (on Centos)",
                        "name": "automated-Apache-Guacamole-installation"
                },
                "1": {
                        "description": "This repository stores a clipboard sniffer and API written in Go. Whenever something is copied to the clipboard it will be sent to the API. ",
                        "name": "ClipboardSniffer"
                },
                "10": {
                        "description": "This was a school project I coded. The code should light up the corresponding LED with a color based on note velocity and octave. ",
                        "name": "MIDI-Keyboard-Visualizer"
                },
                "11": {
                        "description": "This repository features some examples on how you can obfuscate PowerShell code. Note: for educational purposes only",
                        "name": "PowerShell-Obfuscation"
                },
                "12": {
                        "description": "PyliaC - an interface that opens the possibility to run your julia commands inside your python code. ",
                        "name": "PyliaC"
                },
                "13": {
                        "description": "",
                        "name": "RandomWallpapers"
                },
                "14": {
                        "description": "This small project was inspired by the mathematician William Shanks and his big table of reciprocals of primes and their \"loop length\". I saw this in a video of Numberphile and was intintrigued enough to try and rebuild the script used in the video without looking at his code of course :). ",
                        "name": "Reciprocals-of-Primes.py"
                },
                "15": {
                        "description": "RSA (Rivest–Shamir–Adleman) is a public-key cryptosystem that is widely used for secure data transmission. It is also one of the oldest. The acronym \"RSA\" comes from the surnames of Ron Rivest, Adi Shamir and Leonard Adleman, who publicly described the algorithm in 1977. ",
                        "name": "RSA-Cryptosystem"
                },
                "16": {
                        "description": "An (unfinished) automation tool that uses toml and PowerShell. ",
                        "name": "TomlForge"
                },
                "2": {
                        "description": "",
                        "name": "Cookiemonster"
                },
                "3": {
                        "description": "This project is about my own implementation of a Fail2Ban (SSH) service for Windows 10+. ",
                        "name": "Fail2Ban.py"
                },
                "4": {
                        "description": "This repository features my implementation of Fail2Ban (SSH) for Windows 10+ in Go called Fail2Go! :) It isn't finished yet but will be soon. For more information check my Fail2Ban.py repository since the functionality is going to be almost the same :)!",
                        "name": "Fail2Go"
                },
                "5": {
                        "description": "GoScan is a blanzingly fast network/port scanner written in Go. ",
                        "name": "GoScan"
                },
                "6": {
                        "description": "Never use regex ever again inside of C#. You can now use the input validator to check if given inputs are in the correct format! ",                        "name": "Input-Validator"
                },
                "7": {
                        "description": "LinesOfLang - also called LoL - is a cli tool that tells you the how much lines of code of (a) language(s) you have written inside a specified path (including subfolders). ",
                        "name": "LinesOfLang"
                },
                "8": {
                        "description": "",
                        "name": "menaruben"
                },
                "9": {
                        "description": "",
                        "name": "menaruben.github.io"
                }
        }
}
```
### Getting repository data
```
go run GitInfo.go --info repo --username menaruben --repo GoScan
```
The output is just a json as of right now. In the future the output will be displayed in a beautiful table.
```
{
        "name": "GoScan",
        "description": "GoScan is a blanzingly fast network/port scanner written in Go. ",
        "html_url": "https://github.com/menaruben/GoScan",
        "license": {
                "key": "mit",
                "name": "MIT License",
                "node_id": "MDc6TGljZW5zZTEz",
                "spdx_id": "MIT",
                "url": "https://api.github.com/licenses/mit"
        },
        "forks_count": 0,
        "open_issues_count": 0,
        "languages": {
                "Go": 13191
        }
}
```
