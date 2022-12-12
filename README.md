<!--
parent:
  order: false
-->

<div align="center">
  <h1> DA Challenger </h1>
</div>

<div align="center">
  <a href="https://github.com/mantlenetworkio/da-challenger/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/da-challenger/da-challenger.svg" />
  </a>
  <a href="https://github.com/mantlenetworkio/da-challenger/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/da-challenger/da-challenger.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/mantlenetworkio/da-challenger">
    <img alt="GoDoc" src="https://godoc.org/github.com/mantlenetworkio/da-challenger?status.svg" />
  </a>
</div>

 Da-challenger is a project about the fraud proof for EigenDA of Mantle Network

**Tips**: need [Go 1.18+](https://golang.org/dl/)

## Install

### Install dependencies
```bash
go mod tidy
```

### build
```bash
make binding
make da-challenger
```

### start 

Config .env, You can refer to .env_example, if config finished, you can exec following command

```bash
./da-challenger
```

