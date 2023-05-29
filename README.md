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
CCCCCCCCCCCCCCCCCCCCCCtLfiLCtLCCCCCCCGGGGGGGGGGGGGGGCCCCCGGGCCCCCCCCCCGCCCCCCCGCLCGCGCGGGGGGG
CCCCCCCCCCCCCCCCCCCCCC1LCifCffCCCCCCGGGGGGGGGGGGGGGGGGGGGCCGGGGGGGGGGGGGGCGGGGGGLCGGGGGGGGGGG
CCCCCCCCCCCCCCCCCCCCCL1CL;1CL1CCCCCCGGGGGGGGGGGGGGGGCCLLfffCCGCGGGGCGGGGGGGGGGGCLCGCGGGGGGGGG
CCGGCCCCCCCCCCCCCCCCCt1CL1iLC1fCCCCCCGGGGGGGGGGCLLLffffLfLLLGCCGGGCCGGGCGGGGGGGCLCGGGGGGGGGGG
LCCCCCCCCCCCCCCCCCCCC1tCff1tCftCCCCCGGGGGCCCLLffffLLLLLLttLLGCCGGCCGCGGGGGGGGGGCLCGGGGGGGGGGG
LCfLfLLLCCCCCCCCCCCCCifCtLt1CC1CCCCLCCLLLLfLLLLLLLLffftfttfCGCGGCGGGGGGGGGGGGGGCLCGGGGGGGGGGG
LffLfLCCCCCCCCCCCCCCL;LCtCf:i1;tffffffffLLLLLfft1ttfLLCCGGGGCCGCCGCGGGGGGGGGGGGCLCGGGGGGGGGGG
ffLffLCCCCCCCCCCCCCGfiCL1C1,itttffLLLLLffftttffLCCGGGGGGGGGCCGGGGGGGGGGGGGGGGGGCLCGCGGGGGGGGG
fLLLLCGCCCCCCCCGGCCC11fttfffLLLLLLffttttfLLCGGGGGGCCCGGGGGGGGGGGGGGCCGCCCGCGGGGCLCGGGGGGGGGGG
ffLCCCCCCCCCCCCLLfffffffLLLffffttttffCGGGGGGGGGGGGGGGGCCCCGGGGCGCCCGGCGCCGGGGGGCLCGGCGGGGGGGG
t1tLLLLCLLLLLfffffLLLfffttttttfLLCCCGGGGGGGGGGGGGGGGGGGCCCCCGGCCCCCCCCCCCGGGGGGCLCGGGGGGGCGGG
ttttttffffffLLLLLftttttitfLCCCCCCCCCCGGGGGGGGGGGGGGGCCCCCCCCGGGGGGGCGCCCCCCCCCCLfLCCCCCCCGGGG
tttfffffLLLLLfftttt1iCLtGCCCCCCCCCCCCGGGGGGGGGGGGGGGCCGGCCCCCLLLffffffffftttffffffffffftttfLC
i1LLtLfffttttfLLCCGftGLfCCCCCCCCCCCCCGGGGGGGGGGGGGGGCCLLffffffffffLLLLLLLLLLLLLLLLLLLLLt1t111
;1fCi1itfLCCCGCCCCC1fCffCCCCCCCCCCCCCGGGGGGGGGGGGGCGLffffffLLLLLCCCCCCCCCLLCLLLLLLCLLLLt1tttt
iiititffCffLCCCCCCCiLCtLCCCCCCCCCCCCCGGGGGGGGGGGGGCGLfLffLLLCLLLfffftttt11111111i1itLLLt11ttt
t111tt1ttLffLCCCCCLiLCtCCCCCCCCCCCCCCGGGGGGGGCGGCCCGLfLfft1111111t1ttttfffLLLLLfff1iLLLt11ttt
fi1tf1tt1tLCCCCCCCf1CftCCCCCCCCCCCCCCGGGGGGGGGCCCCCGLfLff1tffLLLLLLLLLfCGGGGGGGCCCC1CLLt111tt
Ci11ttfLLfLCCCCGCLiffftLCCCCCCCCCCCCGGGGGGGGGGGGGGGGLfLfftCLfffLfLffffffLLLLLLLfftL1LLLt111tt
Cift1LLLCCCCCCCCCf;fff1tLLLLLLLLLLLLCCLLLLLLLLLLLLLLffLfftLLffLLffffffffffffffffftf1LLLf111tt
Ciiif1iftfttffttttifLLC1tttttftttttttttttttttttttttt1tLfftLLffffffffffffffffffffftf1LLLt1111t
Ci:;i;;i1fffffffffifCCCttfffffffffttttttttttttttftffttLfftLLffffffffffffffffffffftf1LLLt1111t
Ci;i;:i;;tfffftfff;fCCCttftffft1tittttftftti1ti11111itLfftLLffffffffffffffffffffftftLLLt1t11t
C1i11i1iiftfffttff;fCCCttftfff1LC1tt1tt111t1tf1ttttt1tLfftLLtfffffffffffffffffffffftLLLf1111t
C1it1i11ii1i11tfff;tLLL1tftftf1fLifftttttttitf1ttttt1tLLftLCfffffffffffffffffffffff1LLLf1t11t
C11ft1ff11f1f11ffft11111tttttf1tf1ft1tttttf1tfi111ttttLLftLGCfffffffffffffffffffftftLLLt1t11t
f1tLffLLtttit1iffffffffffttttf1fC1Lftt1tffLi:,,,,,:;;1LLLtfCLffffffffffffLCfLCLfftftLLLf1t11t
fttLftff1tt1f1;ttffffffttttttt1tfifffLLLt;:...,,....,,;1fffLffffffffffffLG0LC0Gffff1LLLf11111
ft1LftfL1tfit1;fffttfffttttfff1fCiLLLti:...,,,..,,,,.,,,;iffftffffffffffffffffffftt1LLLf1t11t
f11Lf1ff11f1t1;ttftttttfffft1i;;;,:,,...,,,,.......,,,,,::1fCCCLtLLLLffttfftttfttLL1LLLf1t11t
f11Lf1fL11fit1;tftttttf1i:,..........,,,.............,,,,,;1fLLLtCGGGftttfttfLLLLCC1LLLf1t11t
f11Lf1fL1tftf1;t1ffft1;,....,,,,,,.,....,,.............,,,:;ttfftCtLCLfLLLLLCCCCLfC1LLLf11111
f11Lf1fL1tftft:ii;:,...,,,,.........,;1fft11;,...........,::iffftCCCGGGGGGGGGCfLCCCifLLft1111
f11LftfL11tft;,,....,,..........,.:ittt1i1tttti,..........,:;ttffffffttttttttt11tttifLLf11111
f1itt1tti1t1,....,............,,.;1t1t11tt1t1111,,........,::tfttttttttfftffffffffLLLLLf11111
f11fftffit;.,,................,.;ffftt1ttttttfffCL.........::tCCCCCCCCLLLLLLLLLLLfffLLfft1111
ti1ft1tfi,,::... ...............fCCCCffCCCfffffLft1: .....,,:;1itLLLLLLLLLLLLLLLftffffft1i111
111111ti,::............,....,..,tfLCCfCCCCCLfttt111t: ...,,.,..,:11i11111i111111i;1iiiiii;:;;
;;iiii;:;,.......... .;,...,..;,111tt1ffffttttt1ttfCC;.........,:i1iiiiiiiiiiiiiiiiii111;;;;;
iiiii;;1,.,.........:t:..... ii,1ttt1111tttttttfLCCCfi;:,..,....,;iiiiiiiiiiii;;;iiii;i;:;:;;
iiii;;1...........:tC;..... iL1,ttt1i11t111ttfffLCCt11;1L: .....,;tftttttttt11tt1;iiiiii111i;
iiii;;...........;tL; .... ;CCf:Lfft1t1ttttfffLLLft1tt1f1fi. ....;ttttttt1111111tiiiiiii111i;
;;;;1,...:....,,:1t;..... :fLCL,ftttLCCCCftttt1i;:;ifCt11CCf;. ..it11tttt1111ttttt111ttti1111
;;;1t:..;;...,::1tt,.,,. :1t1tt:fff1ffLCCCLt1;;:;1fCCt11tLLLCf:::i1LLLLLfffff1111iifft11;iiii
1ii1i: :1i.,iiifCG1 .1,.:tt1tt1:fft1i1;:tLLt;;;ifCCt1it1tfLLLLCttfLLfffffffff1111i1tiitfii11i
i;i1i,,ii:,1iiiLCC;.;t.:LCLtiii,;111ti::,,:iiifCCL111if1ttfLLLLCLtttLfLLfffffffffL1itffLiii;;
iifftiii;:;ii;;tt1;,i;,iCfffti:,;it1t;;fft11i1LLti1f11tittttfLLLLLLti1i;;;i;;;;iii:i111111i;:
iLCCCLLLtftttttttti:1i,1CCftffi:;:;1t;ifttffft111111i11ittttttLLLLLLt;;i;::::;;ii1t1tt1;;;:::
fCCCCGGGGGGCCCCCCCCtt111i1fCLtt1:;:;t;;iiiii11ttt1;i1ii:i1i1tttfLLLCt;i1iiii111111111111ttiii
1ttttfffLCCCCCCGCCCGCLti1ttf111t;;;;t;iiiiii;;;;ii;:ii1:;i;;:i1tfLLL1tffffftft::;i;i1ttttffLL
ttt1t1i11ttttttffLLCCCCCttii1itt;;;i1i11111111111111i::,i1iiiiiiii11tfffffLLLft1tfffffffttttt
CCCCLtfffttttt111tttfCCGf1:ii11t;;;;:;;;;;;iiiiii1111;:ii111111111ffftttttttt1itfffffffffffLL
CCGCfCGCCCCCCCLftttt1LLf1i:1i11i::,,::,,,,,,,,,,::::::::::;;;;;;;i1111ttttttf1iffffffffffffff
CCLtLLf111fLLLfLCCft1ttt;,:;,,:::,:::::::::::::::::,,:,,,,,,,,,,,,,,,,,::::::;;;;iii1111ttttf
fftfLLLLLLLLLfttifC1;1ti,::,:::::::::::::::::::::::::::::::::::::::,:,,,,,,,,::,,,,,,,,,,,,::
CCCCLLLLLLLfLfLCf1C1:ii,::::::::::::::::::::::::::::::::::::::::::::::::::::,:::::::::::::,,,
ffffLffiffLLLfiffiti;::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
LCCCCCftCCCCCLitf11i:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
LLCLLCtLCCLLCL1CCCt:;;;;::::::::::::::::::::::::::::::::::::::::::::::::::::iii;:::::::::::::
CCCCCLLCLCCLCL1CLC1:;:::::;;;iii1i;;;;;;:::::::::::::::::::::::::::::::::::i1ii1,::::::::::::
CCCCCCCCCLCLCLtLLL;;i111ttffffftttt11tt111i:::::::::::::::::::::::::::::::tttfft,::::::::::::
LLLCCCLCCLCCCLtfCLtt1tttt11111111111111111tt1i111111i;:;::::::::::::::::::fLLCft:::::::::::::
LLLLLCCLLLLLCLiLLCLLtti11tfffLLLLLLLLLLLfftttt11fffffti:;::;::::::::::::::fLLLft,::::::::::::
LLLLLLLLLLCLLLfLCLCLfLtLttLCCCCLLLLLLCCLLCLtttt11tttttt;;;;;:::;::::::::;:tCLLf1,::::::::::::
LLLLLLLLLLLLLLCCLLCLtLtLLf1fCLLLLLLLLLLLLLLLtttt1tttttt1;;;;;;;;;;;;;:;;;:;fLfti:::::::::::::
LLLLLLLLLLLLLLLLLLCftCttLCLttLLLLLLLLLLLLLLLLttt11ttttt1;;;;;;;;;;;;;;;;;;::;;;:;;:::::::::::

Username:    menaruben
Name:        Rubén Mena
Bio:         19 year old nerd spending too much time on the computer :)
Location:    Switzerland, Zürich
GitHubUrl:   https://api.github.com/users/menaruben
NumRepos:    18
Followers:   4
Repos:
        - PyliaC
        - Reciprocals-of-Primes.py
        - Cookiemonster
        - GitInfo
        - PowerShell-Obfuscation
        - RSA-Cryptosystem
        - automated-Apache-Guacamole-installation
        - menaruben.github.io
        - RandomWallpapers
        - Input-Validator
        - MIDI-Keyboard-Visualizer
        - ClipboardSniffer
        - Fail2Go
        - GoScan
        - TomlForge
        - Fail2Ban.py
        - LinesOfLang
        - menaruben
```

### Getting user data and adjust ascii avatar size
The default for the ```avatarRatio``` is 0.25. If you want to resize the avatar you can use the adjust its ratio with a float64 value.
```
go run .\GitInfo.go --info user --username menaruben --avatarRatio 0.125
```
```
CCCCCCCCCCLftLfCCCGGGGGGGGGGCCCCCCCCCCGCCGGGGG
CGCCCCCCCCfftffCCCGGGGGCCCLLfCGGGGGGGGGCCGGGGG
CLLCCCCCCCtfftffCCCCCLLffffffCGCGGGGGGGCCGGGGG
LfLCCCCCCCtffi1tLfffffffLLCCCGGGGGGGGGGCCGGGGG
fLCCCCCCLLfffffffffLCCGGGGGGGGGGGGGGGGGCCGGGGG
ttfLLLLfffftffLLCCGGGGGGGGGGGGGCCCCCCCCLCCCCGG
1ftffffLLtffCCCCCCGGGGGGGGCLLLLLLLLLLLLLLLfttf
it1tLLCCGfLLCCCCCCGGGGGGGCfffLLLfffffffftLL11t
t1ttfLCCCtfLCCCCCCGGGGGGGCLfttftfffLLLLLftLt1t
f1fLLCCCftffLLLLCCCCCCCCCLfffLfLfffLLLLLffLt1t
f;11tfffttCftfftttttttttttfffLfffffffffftfLt11
fii;1ftfttCftfttttttt1t1t1tfffffffffffffffLt1t
f1tt111tt1fttftttt1tftt1ttfLfCffffffffffffLt1t
tffftt11fftttftfffff1;,,,:itfLffffffCLCftfLt11
tffftt11ftffft11i1i:.......,ifLffLftftffftLt1t
tttftt11ftt;:,.....,,......,,iLfLCLffLCLCtLt11
ttfftf1::,.......:i111;. ...,:ttfLLLLLffftLt11
ttft1i,.........ift1ttft; ...:tLfffffffLfLLf1t
11t1:,.........,LCfCCffft;...,:iffffttft1tt1ii
iii;,.....:,..,;1t1ttt1tLCi,....iiiiiii;;ii;;;
ii;,....,t;..:Litt1t1tfLLt1ti. .ittt1111iiii1i
;i;.:..:1i...1L1ftLCLti;iff1LL;,iffftt111t11i1
ii::;:itC,;:fti;t1i;11;tLt11fLCftfLffftt11ttii
1Lft1t1tt;;iCf1:;1i111tf1111ttfCLti;;;i1111i;:
fLLLCCCCCLf1tft1;iiiii11i;i;ii1tLf1t1t1i111t11
ffttttttfLCCti1t;;;iiiiiii;;iiii1tfffLttffffff
CLLLfLLffttf;;;;::,,:,::::::::::;;;iiii111tttf
LfLffLfffti;:::::::::::::::::::,,,,,,,:,,,,:::
LLLtLLftti::::::::::::::::::::::::::::::::::::
CCLLCCLfC;:;;;;;;;;::::::::::::::::::;i;::::::
LLCCLCLfL111ttttt1111i;ii;;:::::::::ifLi::::::
LLCLLCLfCLtttLLLLLLLLfttfft;:;::::::iCCi::::::
LLLLLLLLCLffLfLLLLLLLLt11tt1;;;;;;;;;11;::::::

Username:    menaruben
Name:        Rubén Mena
Bio:         19 year old nerd spending too much time on the computer :)
Location:    Switzerland, Zürich
GitHubUrl:   https://api.github.com/users/menaruben
NumRepos:    18
Followers:   4
Repos:
        - Cookiemonster
        - Fail2Ban.py
        - Fail2Go
        - GitInfo
        - MIDI-Keyboard-Visualizer
        - RandomWallpapers
        - Reciprocals-of-Primes.py
        - automated-Apache-Guacamole-installation
        - ClipboardSniffer
        - GoScan
        - Input-Validator
        - LinesOfLang
        - menaruben
        - TomlForge
        - menaruben.github.io
        - PowerShell-Obfuscation
        - PyliaC
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
