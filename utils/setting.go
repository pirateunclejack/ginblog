package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbType     string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	DbSslMode  string
	DbTimeZone string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("Config file loading error:", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("wn2Og76swi9Cw")

}

func LoadData(file *ini.File) {
	DbType = file.Section("database").Key("DbType").MustString("postgres")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("5432")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord = file.Section("database").Key("DbPassword").MustString("admin123")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
	DbSslMode = file.Section("database").Key("DbSslMode").MustString("disable")
	DbTimeZone = file.Section("database").Key("DbTimeZone").MustString("Asia/Shanghai")
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("")
	QiniuServer = file.Section("qiniu").Key("QiniuServer").MustString("")

}
