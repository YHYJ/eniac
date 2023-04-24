# README

<!-- File: README.md -->
<!-- Author: YJ -->
<!-- Email: yj1516268@outlook.com -->
<!-- Created Time: 2023-04-19 11:19:47 -->

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

* [Usage](#usage)

<!-- vim-markdown-toc -->

---

<!-- Object info -->

---

一个系统交互工具

## Usage

- `config`子命令

    该子命令用于操作配置文件，有以下参数：

    - 'create'：创建默认内容的配置文件，可以使用全局参数'--config'指定配置文件路径
    - 'force'：当指定的配置文件已存在时，使用该参数强制覆盖原文件
    - 'print'：打印配置文件内容

- `get`子命令

    该子命令用于获取系统信息，参数用于指定获取哪部分

详细信息请使用'--help'查看
