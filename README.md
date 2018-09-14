## eos-go startup

### 简介

[eos-go startup](https://github.com/dabdevelop/eosgo-startup) 整理了eos-go常用代码，包括创建、查询EOS账号，EOS转账，购买RAM，出售RAM，抵押资源，取消抵押资源，生成密钥对等。

### 前提

前提是安装了[eos-go](https://github.com/eoscanada/eos-go)：

在此工程目录下，安装 `eos-go`

```bash
    go get github.com/eoscanada/eos-go
```

### 使用方法

首先，在 `cli.go` 中替换私钥。然后对相应程序的参数进行修改。

go代码运行方式为：

```bash
    go run cli.go
```