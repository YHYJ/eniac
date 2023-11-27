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
  Arch = "系统架构"
  BIOSDate = "BIOS发布"
  BIOSVendor = "BIOS厂商"
  BIOSVersion = "BIOS版本"
  BoardName = "主板名称"
  BoardVendor = "主板厂商"
  BoardVersion = "主板版本"
  BootTime = "系统启动时间"
  CPUCache = "CPU缓存"
  CPUCores = "CPU核心"
  CPUModel = "CPU型号"
  CPUNumber = "CPU插槽"
  CPUThreads = "CPU线程"
  DaemonStatus = "更新状态"
  GPUAddress = "显卡地址"
  GPUDriver = "显卡驱动"
  GPUProduct = "显卡型号"
  GPUVendor = "显卡厂商"
  Hostname = "主机名称"
  Kernel = "内核版本"
  Load1 = "1分钟负载"
  Load15 = "15分钟负载"
  Load5 = "5分钟负载"
  MemoryAvail = "可用内存"
  MemoryBuffCache = "缓冲内存"
  MemoryFree = "空闲内存"
  MemoryShared = "共享内存"
  MemoryTotal = "内存大小"
  MemoryUsed = "已用内存"
  MemoryUsedPercent = "内存占用"
  NicDriver = "网卡驱动"
  NicDuplex = "工作模式"
  NicMacAddress = "MAC 地址"
  NicName = "网卡名称"
  NicPCIAddress = "PCI 地址"
  NicProduct = "网卡型号"
  NicSpeed = "网卡速率"
  NicVendor = "网卡厂商"
  OS = "操作系统"
  Platform = "系统类型"
  Process = "进程数"
  ProductName = "设备名称"
  ProductVendor = "设备厂商"
  StartTime = "系统启动用时"
  StorageDriver = "磁盘驱动"
  StorageModel = "磁盘型号"
  StorageName = "磁盘名称"
  StorageRemovable = "磁盘可移除"
  StorageSerial = "磁盘序列号"
  StorageSize = "磁盘容量"
  StorageType = "磁盘类型"
  StorageVendor = "磁盘厂商"
  SwapDisabled = "交换空间未启用"
  SwapFree = "可用交换空间"
  SwapTotal = "交换空间大小"
  TimeZone = "系统时区"
  UpdateList = "更新列表"
  Uptime = "系统运行时长"
  User = "用户名称"
  UserGid = "属组标识"
  UserHomeDir = "用户目录"
  UserName = "用户昵称"
  UserUid = "用户标识"

  [genealogy.cpu]     # CPU 信息
    cache_unit = "KB" # CPU 缓存的单位

  [genealogy.memory]   # 内存信息
    data_unit = "GB"   # 内存信息中数据的单位
    percent_unit = "%" # 内存信息中百分比的单位

  [genealogy.update]                             # 更新信息
    record_file = "/tmp/system-checkupdates.log" # 系统更新信息记录文件

[main]         # 主程序配置
  color = true # 是否启用彩色输出

  [main.parts] # 输出的各部件的显示名称
    BIOS = "BIOS"
    Board = "主板"
    CPU = "处理器"
    Disk = "磁盘"
    GPU = "显卡"
    Load = "负载"
    Memory = "内存"
    NIC = "网卡"
    OS = "系统"
    Product = "设备"
    Swap = "交换分区"
    Time = "时间"
    Update = "更新"
    User = "用户"
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
