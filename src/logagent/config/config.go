package config

import (
	"errors"
	"github.com/astaxie/beego/config"
)

var Conf map[string]string

func LoadConf() (conf map[string]string, err error) {

	filename := "./conf/app.conf"
	cf, err := config.NewConfig("ini", filename)
	if err != nil {
		err = errors.New("NewConfig error")
		return conf,err
	}

	Conf = make(map[string]string)


	Conf["log_level"] = cf.String("logs::log_level")
	if len(Conf["log_level"]) == 0 {
		Conf["log_level"] = "debug"
	}

	Conf["log_path"] = cf.String("logs::log_path")
	if len(Conf["log_path"]) == 0 {
		Conf["log_path"] = "./logs"
	}


	Conf["collect_path"] = cf.String("collect::collect_path")
	if len(Conf["collect_path"]) == 0 {
		err = errors.New("collect::collect_path error")
		return Conf,err
	}

	Conf["topic"] = cf.String("collect::topic")
	if len(Conf["topic"]) == 0 {
		err = errors.New("topic error")
		return Conf,err
	}
	return Conf,nil
}
