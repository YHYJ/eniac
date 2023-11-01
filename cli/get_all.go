/*
File: get_all.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-09-28 12:44:03

Description:
*/

package cli

import (
	"os/user"

	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

var (
	pciData, _     = ghw.PCI()           // PCI信息
	blockData, _   = ghw.Block()         // 存储设备信息
	networkData, _ = ghw.Network()       // 网络设备信息
	gpuData, _     = ghw.GPU()           // 显卡信息
	loadData, _    = load.Avg()          // 系统负载信息
	memData, _     = mem.VirtualMemory() // 内存信息
	hostData, _    = host.Info()         // 主机信息
	userData, _    = user.Current()      // 用户信息
)
