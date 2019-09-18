
`mkdir kls && cd kls`

`go get github.com/spf13/cobra/cobra`

configure your `~/.cobra.yaml` like below for example

```
author: Tummala Dhanvi <dhanvicse at gmail.com>
license: MIT
```

Cobra Init
`cobra init --pkg-name github.com/k8sBLR/24August2019/assignments/dhanvi/kls`

Add commands 

```
cobra add pods
cobra add deployment
```

Add go modules & running
```
go mod init github.com/k8sBLR/24August2019/assignments/dhanvi/kls
go install 
go build
kls -h
```

Installing the go client if the current one isn't working for you

`go get k8s.io/client-go@kubernetes-1.15.3`


Reference:
* https://github.com/surajssd/lspods
* https://github.com/surajssd/cobrademo
* https://docs.google.com/presentation/d/1NgAxNrUxcOnodm9VA4plidrKradQM_kAbrTrF844gec/
