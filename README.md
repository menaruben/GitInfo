# GitInfo
This repository is under heavy development! This is just the beginning of a simple CLI-Tool that collects data from GitHub users and repositories and prints it to the terminal.

## Examples
### Getting user data
```
go run GitInfo.go --info user --username menaruben
```
```
Username:    menaruben
Name:        Rubén Mena
Bio:         19 year old nerd spending too much time on the computer :)
Location:    Switzerland, Zürich
AvatarUrl:   https://avatars.githubusercontent.com/u/74914224?v=4
GitHubUrl:   https://api.github.com/users/menaruben
NumRepos:    18
Followers:   4
Repos:
        - Fail2Ban.py
        - GitInfo
        - GoScan
        - menaruben
        - RandomWallpapers
        - Reciprocals-of-Primes.py
        - automated-Apache-Guacamole-installation
        - LinesOfLang
        - menaruben.github.io
        - MIDI-Keyboard-Visualizer
        - PyliaC
        - TomlForge
        - ClipboardSniffer
        - Cookiemonster
        - Fail2Go
        - Input-Validator
        - PowerShell-Obfuscation
        - RSA-Cryptosystem
```
### Getting repository data
```
go run GitInfo.go --info repo --username menaruben --repo RSA-Cryptosystem
```
```
Name:            RSA-Cryptosystem
Description:     RSA (Rivest–Shamir–Adleman) is a public-key cryptosystem that is widely used for secure data transmission.
RepositoryUrl:   https://github.com/menaruben/RSA-Cryptosystem
License:
ForksCount:      0
Languages:       Julia: 1616, C: 4829, Rust: 3601, PowerShell: 2563, Go: 2247, Python: 1941
```
