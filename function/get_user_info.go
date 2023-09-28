/*
File: get_user_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 10:37:57

Description: 获取用户信息
*/

package function

// GetUserInfo 获取用户信息
func GetUserInfo() (userInfo map[string]interface{}, err error) {
	userInfo = make(map[string]interface{})
	userInfo["User"] = userData.Name           // 用户昵称
	userInfo["UserName"] = userData.Username   // 用户名
	userInfo["UserUid"] = userData.Uid         // 用户ID
	userInfo["UserGid"] = userData.Gid         // 用户组ID
	userInfo["UserHomeDir"] = userData.HomeDir // 用户主目录

	return userInfo, err
}
