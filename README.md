# README

<!-- File: README.md -->
<!-- Author: YJ -->
<!-- Email: yj1516268@outlook.com -->
<!-- Created Time: 2023-04-19 11:19:47 -->

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

* [示例配置](#示例配置)

<!-- vim-markdown-toc -->

---

<!-- Object info -->

---

一个系统交互工具

## 示例配置

```toml
[cpu]
cache_unit = "KB" # 期望输出缓存的单位，"B"、"KB"、"MB"

[memory]
data_unit = "GB"   # 期望输出数据的单位，"B"、"KB"、"MB"、"GB"、"TB"
percent_unit = "%" # 期望输出百分比的单位，"%"

[genealogy]
# BIOS
BIOSVendor = "BIOS厂商"
BIOSVersion = "BIOS版本"
BIOSDate = "BIOS发布日期"
# Board
BoardVendor = "主板厂商"
BoardName = "主板名称"
BoardVersion = "主板版本"
# CPU
CPUModel = "CPU型号"
CPUNum = "CPU数量"
CPUCores = "CPU核心数"
CPUThreads = "CPU线程数"
CPUCache = "CPU缓存"
# OS
OS = "操作系统"
Arch = "操作系统架构"
Kernel = "内核版本"
Platform = "平台"
Hostname = "主机名"
TimeZone = "时区"
# Load
Load1 = "1分钟负载"
Load5 = "5分钟负载"
Load15 = "15分钟负载"
# Memory
MemTotal = "总物理内存"
MemUsed = "已用物理内存"
MemUsedPercent = "物理内存使用率"
MemFree = "空闲物理内存"
MemShared = "共享物理内存"
MemBuffCache = "缓冲物理内存"
MemAvail = "可用物理内存"
# Swap
SwapTotal = "总交换内存"
SwapFree = "空闲交换内存"
# Process
Procs = "进程数"
# Product
ProductVendor = "产品厂商"
ProductName = "产品名称"
# Storage
StorageName = "存储设备名称"
StorageDriver = "存储设备驱动"
StorageVendor = "存储设备厂商"
StorageModel = "存储设备型号"
StorageSerial = "存储设备序列号"
StorageSize = "存储设备容量"
# Time
BootTime = "系统启动时间"
Uptime = "系统运行时间"
# User
User = "用户昵称"
UserName = "用户名"
UserUid = "用户ID"
UserGid = "用户组ID"
UserHomeDir = "用户主目录"
```
