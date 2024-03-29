/*
File: get_user_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 10:37:57

Description: 获取用户信息
*/

package cli

// GetUserInfo 获取用户信息
//
// 返回：
//   - 用户信息
func GetUserInfo() map[string]interface{} {
	userInfo := make(map[string]interface{})
	userInfo["User"] = userData.Name           // 用户名称
	userInfo["UserName"] = userData.Username   // 用户昵称
	userInfo["UserUid"] = userData.Uid         // 用户 ID
	userInfo["UserGid"] = userData.Gid         // 用户组 ID
	userInfo["UserHomeDir"] = userData.HomeDir // 用户主目录

	return userInfo
}
