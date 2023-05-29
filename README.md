# GitInfo
This repository is under heavy development! This is just the beginning of a simple CLI-Tool that collects data from GitHub users and repositories and prints it to the terminal.

## Future features
- [x] UserData:     show GitHub avatar in command line (can't show *colored* ascii in GitHub README)
- [ ] General:      show (max. 10) most used languages in percentage 

## Examples
### Getting user data
```
go run GitInfo.go --info user --username menaruben
```
```
CCCCCCCCCCCCCCCCfff1CfLCCCCGGGGGGGGGGGGGGGGCCGCCCCCCCGCCCCGGLCGCGGGGGG
CCGCCCCCCCCCCCCC1fLiLLtCCCCGGGGGGGGGCCCCLLLLLGCGGGGGGGGGGGGGLCGGGGGGGG
CCLCCCCCCCCCCCCL1Cft1CfLCCCCGCCCCLLLfffffftfLCCGGCGGGGGGGGGGLCGGGGGGGG
LfffCCCCCCCCCCGf1CtL;111LLfLffffffffffLLCCCCCGGGGGGGGGGGGGGGLCGGGGGGGG
fLLCCCCCCCCCCCCtffffttffffffffLLLCCCGGGGGGGGGGCCCCCCCCGCGGGGLCGGGGGGGG
ttLLLLLLLLLfffffftttfffLLLLCGGGGGGGGGGGCCCCCGGGGGGGGGGGGGGGGLCGGGGGGGG
ttttffffffffffttLtLCCCCCCCCGGGGGGGGGGGGGCCCCCCLLLLLLLLLLLLLLfLLLLLfLLC
;tL1ttffLLCCCCffCfCCCCCCCCCCGGGGGGGGGGGCLfffLLLLLLLLLLLLLLLLLLLLLf111t
1it1ttffLCCCCC1LLfCCCCCCCCCCGGGGGGCCCCGLfffffffffttfttfffffftt1LLf1ttt
fitttftfCCCCGf1LtLCCCCCCCCCGGGGGGGGGGGGLfLftfffffffLfLCCCCCCLLtfLf11tt
Li1tfLLLLLLLL1tLftfLLLffLLLLLLLLLLLLLLLtfLffLffffffffffffffffftfLf111t
L;;i;i1fffftf1tCCftfffftttttttttttttttt1fLttLffffffffffffffffftfLf111t
Li1iiiittttff11CCftftfttf1tttttt1tt1ttt1fLftLffffffffffffffffftLLf111t
Lifttf1t1t1tft1tt1ttttttftf1tttft11i11t1fLffCCfffffffffffffffftfLf111t
f1LffLtf1ti1ffffffffffttLtLfft1i:,....,,;1tfLfffffffffLGLLCLfttfLf1111
f1fftL1t1ti1fttfffftt1iii;i;:,...........,:itLLLtLLLfttfttfftLtfLf111t
f1fttL1t1ti1tftt1;:,.........,,,.. ......,,,;tfffCCGLffLfCCCLCftLLt111
f1fttL1tft;:::,,.......,.,:;1t1111;,  .....,:;tttLfLLLLLLLfffLttLLt111
f1tttf11;,............,.:1ttt1tttttt1: .....::fLLffffffffffLfffLLLt1tt
t1tttt;:,,...... .......tCCCfLCLfffffti,....,:;i1ffffffffffftttttt1iii
;iiii;:,...... .:,....::1ttttffftt1ttfCL:.......,iiiiiiiiii;i;iiii;::;
iii;;:.......,1t:....1L:1tt11111ttfLCCft11t:. ..,ittt11111111iiiiiiii;
;;ii,.., ...:t1... .1LC;fftfLCCfttt1ii1ffttLt;,..itttttttt11111111111i
i;ii.,1:.;;1LL.,;.:tt1t;tf11titfti;itLff11tLCCLt1tfLLLLLLfttt1tttfii1i
i1t1;ii:i1;tf1,i;:CLf1i,;11t;iii11tLLtt111ttffLCLfft11111i1111i1ttii;:
tCCCCCLCLLLLLf11iitfLfti;;it;11111111iii1;i11ttfLLCt;ii;;iiii1111i1i;i
ttttttffLLLLCCCCLf1t111ti;;1iiiiiii11iiii;;iiii11tftfffffLti111tttffff
CCLffLfffffttttfCCf;;i11;::::;;;;;;;;;i;:;;iiiiii1ttttttff11ffffffffff
LLfLLLfLLLLffLtit1:,;:::::::::::,,,,,,,,::,,,,::::::::::::;;;;;iiiii11
LLLLLfLLLLffttt:i::::::::::::::::::::::::::::::::::::::::,:,,,,,,,,,,,
LLLCLtLLLCttf1i:::::::::::::::::::::::::::::::::::::::::::::::::::::::
CCCCLLCCLCtLCf:;;;;;;iiiii;;;;;:::::::::::::::::::::::::;1i1::::::::::
LLLCCCLLLCffCt111tttttttttttttt11iiiiiii;::::::::::::::;LLLf::::::::::
LLLCLLLLLCffCCLftttfLLLLLLLLLLLLfttttfffti:;;::::::::::;LCLt::::::::::
LLLLLLLLLLLLLCLfftLfffLLLLLLLLLLLftt1ttttt;;;;;;;;;;;;;:i11i:;::::::::

Username:    menaruben
Name:        Rubén Mena
Bio:         19 year old nerd spending too much time on the computer :)
Location:    Switzerland, Zürich
GitHubUrl:   https://api.github.com/users/menaruben
NumRepos:    18
Followers:   4
Repos:
	- Input-Validator
	- menaruben
	- MIDI-Keyboard-Visualizer
	- TomlForge
	- RSA-Cryptosystem
	- Cookiemonster
	- LinesOfLang
	- PowerShell-Obfuscation
	- Reciprocals-of-Primes.py
	- PyliaC
	- RandomWallpapers
	- automated-Apache-Guacamole-installation
	- Fail2Ban.py
	- GitInfo
	- menaruben.github.io
	- ClipboardSniffer
	- Fail2Go
	- GoScan
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
