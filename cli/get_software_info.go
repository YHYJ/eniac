/*
File: get_software_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-17 09:32:49

Description: 获取软件信息
*/

package cli

import (
	"fmt"

	"github.com/Jguer/go-alpm/v2"
	"github.com/yhyj/eniac/general"
)

func GetPackageInfo() (map[string]interface{}, error) {
	packageInfo := make(map[string]interface{})

	// Alpm初始化，获取句柄
	handle, err := alpm.Initialize("/", "/var/lib/pacman")
	if err != nil {
		return nil, err
	}

	// 获取指定句柄的本地数据库
	db, err := handle.LocalDB()
	if err != nil {
		return nil, err
	}

	// 获取本地数据库中的包列表
	pkgSlice := db.PkgCache().Slice()

	// 计算已安装包的总数
	totalCount := len(pkgSlice)

	// 计算已安装包的总大小
	var totalSize float64 = 0
	for _, p := range pkgSlice {
		totalSize += float64(p.ISize())
	}

	// 释放句柄
	if err := handle.Release(); err != nil {
		return nil, err
	}

	packageInfo["PackageTotalCount"] = totalCount
	packageTotalSize, packageTotalSizeUnit := general.Human(totalSize, "B")
	packageInfo["PackageTotalSize"] = fmt.Sprintf("%.2f %s", packageTotalSize, packageTotalSizeUnit)

	return packageInfo, nil
}
