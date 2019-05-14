package log

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
)

func converLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
        return logs.LevelWarn
	case "info":
        return logs.LevelInfo
	case "trace":
        return logs.LevelTrace
	default:
		return logs.LevelDebug
	}
}

func InitLogger(conf map[string]string) (err error) {

	config := make(map[string]interface{})
	config["filename"] = conf["log_path"]
	config["level"] = converLogLevel(conf["log_level"])

	json, err := json.Marshal(config)
	if err != nil {
		err = errors.New("json err")
        return err
	}

	logs.SetLogger(logs.AdapterFile, string(json))
	return nil

	/*logs.Debug("this is a test, my name is %s", "stu01")
	logs.Trace("this is a trace, my name is %s", "stu02")
	logs.Warn("this is a warn, my name is %s", "stu03")*/

}
