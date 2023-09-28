/*
File: get_all.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-09-28 12:44:03

Description:
*/

package function

import (
	"os/user"

	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

var pciData, _ = ghw.PCI()           // PCI信息
var blockData, _ = ghw.Block()       // 存储设备信息
var gpuData, _ = ghw.GPU()           // 显卡信息
var loadData, _ = load.Avg()         // 系统负载信息
var memData, _ = mem.VirtualMemory() // 内存信息
var hostData, _ = host.Info()        // 主机信息
var userData, _ = user.Current()     // 用户信息
