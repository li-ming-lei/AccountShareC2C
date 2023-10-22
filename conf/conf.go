package conf

import (
	"fmt"
	utils "github.com/li-ming-lei/AccountShareC2C/util"
)

type Conf struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

const (
	ConfigFile = "config.json"
)

var GlobalConf = Conf{}

func init() {
	LoadConf()
}

func LoadConf() error {
	if exist, _ := utils.FileOrDirExist(ConfigFile); !exist {
		return fmt.Errorf("config file not found")
	}
	content, err := utils.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println("read config file error")
		return fmt.Errorf("config file not found " + err.Error())
	}
	er := utils.UnmarshalJson(content, &GlobalConf)
	if er != nil {
		return er
	}
	fmt.Printf("conf is %+v \n", GlobalConf)
	return nil
}
