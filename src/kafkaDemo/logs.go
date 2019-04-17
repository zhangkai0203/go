package main

import (
	"encoding/json"
	"fmt"
	"src/github.com/astaxie/beego/logs"
)

func main() {

	config := make(map[string]interface{})
	config["filename"] = "./logcollect.log"
	config["level"] = logs.LevelDebug

	json,err := json.Marshal(config)
	if err != nil {
		fmt.Println("json err",err)
		return
	}

	logs.SetLogger(logs.AdapterFile,string(json))

	logs.Debug("this is a test, my name is %s", "stu01")
	logs.Trace("this is a trace, my name is %s", "stu02")
	logs.Warn("this is a warn, my name is %s", "stu03")

}
