/*
File: get_userinfo.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 10:37:57

Description: 获取用户信息
*/

package function

import "os/user"

// UserInfoStruct 用户信息结构体
type UserInfoStruct struct {
	User        string `json:"user"`          // 当前用户
	UserName    string `json:"user_name"`     // 当前用户名称
	UserUid     string `json:"user_uid"`      // 当前用户uid
	UserGid     string `json:"user_gid"`      // 当前用户gid
	UserHomeDir string `json:"user_home_dir"` // 当前用户主目录
}

// GetUserInfo 获取用户信息
func GetUserInfo() (userInfo UserInfoStruct, err error) {
	user, _ := user.Current()
	userInfo.User = user.Name
	userInfo.UserName = user.Username
	userInfo.UserUid = user.Uid
	userInfo.UserGid = user.Gid
	userInfo.UserHomeDir = user.HomeDir

	return userInfo, nil
}
