/*
File: read_confile.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 13:16:41

Description:a 读取配置文件
*/

package function

import "github.com/pelletier/go-toml"

func GetTomlConfig(filename string) (*toml.Tree, error) {
	tree, err := toml.LoadFile(filename)
	if err != nil {
		return nil, err
	}
	return tree, nil
}
