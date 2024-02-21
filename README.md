<h1 align="center">Eniac</h1>

<!-- File: README.md -->
<!-- Author: YJ -->
<!-- Email: yj1516268@outlook.com -->
<!-- Created Time: 2023-04-19 11:19:47 -->

---

<p align="center">
  <a href="https://github.com/YHYJ/eniac/actions/workflows/release.yml"><img src="https://github.com/YHYJ/eniac/actions/workflows/release.yml/badge.svg" alt="Go build and release by GoReleaser"></a>
</p>

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

* [Install](#install)
  * [一键安装](#一键安装)
* [Usage](#usage)
* [Configuration](#configuration)
* [Compile](#compile)
  * [当前平台](#当前平台)
  * [交叉编译](#交叉编译)
    * [Linux](#linux)
    * [macOS](#macos)
    * [Windows](#windows)

<!-- vim-markdown-toc -->

---

<!------------------------------->
<!--             _             -->
<!--   ___ _ __ (_) __ _  ___  -->
<!--  / _ \ '_ \| |/ _` |/ __| -->
<!-- |  __/ | | | | (_| | (__  -->
<!--  \___|_| |_|_|\__,_|\___| -->
<!------------------------------->

---

一个 Linux 系统交互工具

## Install

### 一键安装

```bash
curl -fsSL https://raw.githubusercontent.com/YHYJ/eniac/main/install.sh | sudo bash -s
```

## Usage

- `config`子命令

  该子命令用于操作配置文件，有以下参数：

  - 'create'：创建默认内容的配置文件，可以使用全局参数'--config'指定配置文件路径
  - 'force'：当指定的配置文件已存在时，使用该参数强制覆盖原文件
  - 'print'：打印配置文件内容

- `get`子命令

  该子命令用于获取系统信息，参数用于指定获取哪部分

- `version`子命令

  查看程序版本信息

- `help`子命令

  查看程序帮助信息

## Configuration

1. 使用`config`子命令生成默认配置文件（具体使用方法执行`eniac config --help`查看）
2. 参照如下说明修改配置文件：

```toml
[genealogy] # 输出的信息的中文名

  [genealogy.cpu]     # CPU 信息
    cache_unit = "KB" # CPU 缓存的单位

  [genealogy.memory]   # 内存信息
    data_unit = "GB"   # 内存信息中数据的单位
    percent_unit = "%" # 内存信息中百分比的单位

  [genealogy.update]                             # 更新信息
    record_file = "/tmp/system-checkupdates.log" # 系统更新信息记录文件

[main]            # 主程序配置
  colorful = true # 是否启用彩色输出
```

## Compile

### 当前平台

```bash
go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/eniac/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/eniac/general.BuildTime=`date +%s` -X github.com/yhyj/eniac/general.BuildBy=$USER" -o build/eniac main.go
```

### 交叉编译

使用命令`go tool dist list`查看支持的平台

#### Linux

```bash
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/eniac/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/eniac/general.BuildTime=`date +%s` -X github.com/yhyj/eniac/general.BuildBy=$USER" -o build/eniac main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64

#### macOS

```bash
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/eniac/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/eniac/general.BuildTime=`date +%s` -X github.com/yhyj/eniac/general.BuildBy=$USER" -o build/eniac main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64

#### Windows

```powershell
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -H windowsgui -X github.com/yhyj/eniac/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/eniac/general.BuildTime=`date +%s` -X github.com/yhyj/eniac/general.BuildBy=$USER" -o build/eniac.exe main.go
```

> 使用`echo %PROCESSOR_ARCHITECTURE%`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64
