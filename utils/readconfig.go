package utils

import (
	//"../utils"
	//"../beego/config"
	"github.com/astaxie/beego/config"
	"fmt"
)

var Config config.Configer

func Config_init(confpath string) error{
	Conf, err := config.NewConfig("ini", confpath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Config = Conf
	return nil
}
