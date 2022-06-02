package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}

	LoadServer(file)
	LoadData(file)
}

// LoadServer 服务器参数
func LoadServer(file *ini.File) {
	section := "server"
	AppMode = file.Section(section).Key("AppMode").MustString("debug")
	HttpPort = file.Section(section).Key("HttpPort").MustString(":3000")
}

// LoadData 数据库参数
func LoadData(file *ini.File) {
	section := "database"
	Db = file.Section(section).Key("Db").MustString("mysql")
	DbHost = file.Section(section).Key("DbHost").MustString("localhost")
	DbPort = file.Section(section).Key("DbPort").MustString(":3306")
	DbUser = file.Section(section).Key("DbUser").MustString("")
	DbPassWord = file.Section(section).Key("DbPassWord").MustString("")
	DbName = file.Section(section).Key("DbName").MustString("")
}
