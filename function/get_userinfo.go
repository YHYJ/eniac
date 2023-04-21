/*
File: get_userinfo.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 10:37:57

Description: 获取用户信息
*/

package function

import "os/user"

// GetUserInfo 获取用户信息
func GetUserInfo() (userInfo map[string]interface{}, err error) {
	info, _ := user.Current()
	userInfo = make(map[string]interface{})
	userInfo["User"] = info.Name           // 用户昵称
	userInfo["UserName"] = info.Username   // 用户名
	userInfo["UserUid"] = info.Uid         // 用户ID
	userInfo["UserGid"] = info.Gid         // 用户组ID
	userInfo["UserHomeDir"] = info.HomeDir // 用户主目录

	return userInfo, err
}
