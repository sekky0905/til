# Goland で Go Modules に対応する

```bash
mkdir -p modules-sample/hoge
cd modules-sample
touch main.go
mkdir hoge && cd hoge && touch hoge.go cd ..
go mod init modules-sample
go: creating new go.mod: module modules-sample
```


今の状態

```
modules-sample
  ├── main.go
  ├── go.mod
  └── hoge
       └── hoge.go
```

