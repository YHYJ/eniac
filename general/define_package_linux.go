//go:build linux

/*
File: define_package_linux.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-07-19 15:25:25

Description: 已安装包信息
*/

package general

import "github.com/Jguer/go-alpm/v2"

type PackageData struct {
	PackageTotalCount int
	AsDependencyCount int
	AsExplicitCount   int
	PackageTotalSize  float64
	PackageTotalUnit  string
}

// GetInstalledPackageData 获取已安装包的数据
//
// 返回：
//   - 最新内核版本号
func GetInstalledPackageData() (PackageData, error) {
	id, _ := GetSystemID()

	var (
		packageData PackageData
		err         error
	)
	switch id {
	case "arch":
		packageData, err = getInstalledPackageDataForArch()
	case "debian":
		packageData, err = getInstalledPackageDataForDebian()
	default:
		packageData, err = getInstalledPackageDataForUnknown()
	}

	return packageData, err
}

// getInstalledPackageDataForArch 获取已安装包的数据，Arch 系专用
//
// 返回：
//   - 最新内核版本号
func getInstalledPackageDataForArch() (PackageData, error) {
	var packageData PackageData

	// Alpm初始化，获取句柄
	handle, err := alpm.Initialize("/", "/var/lib/pacman")
	if err != nil {
		return packageData, err
	}

	// 获取指定句柄的本地数据库
	db, err := handle.LocalDB()
	if err != nil {
		return packageData, err
	}

	// 获取本地数据库中的包列表
	pkgSlice := db.PkgCache().Slice()

	// 计算已安装包的总数
	totalCount := len(pkgSlice)

	// 计算已安装包的总大小
	var (
		totalSize          float64
		asExplicitQuantity int
		asDepsQuantity     int
	)
	for _, pkg := range pkgSlice {
		totalSize += float64(pkg.ISize())
		if pkg.Reason().String() == "Explicitly installed" {
			asExplicitQuantity++
		} else if pkg.Reason().String() == "Installed as a dependency of another package" {
			asDepsQuantity++
		}
	}

	// 释放句柄
	if err := handle.Release(); err != nil {
		return packageData, err
	}

	packageTotalSize, packageTotalUnit := Human(totalSize, "B")

	packageData.PackageTotalCount = totalCount
	packageData.AsDependencyCount = asDepsQuantity
	packageData.AsExplicitCount = asExplicitQuantity
	packageData.PackageTotalSize = packageTotalSize
	packageData.PackageTotalUnit = packageTotalUnit

	return packageData, nil
}

// getInstalledPackageDataForDebian 获取已安装包的数据，Debian 系专用
//
// 返回：
//   - 最新内核版本号
func getInstalledPackageDataForDebian() (PackageData, error) {
	var packageData PackageData
	// TODO: 待实现 <07-06-24, YJ> //
	return packageData, nil
}

// getInstalledPackageDataForUnknown 获取已安装包的数据，未支持系统专用
//
// 返回：
//   - 最新内核版本号
func getInstalledPackageDataForUnknown() (PackageData, error) {
	var packageData PackageData
	return packageData, nil
}
