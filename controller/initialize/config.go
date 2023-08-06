package initialize

import (
	"encoding/json"
	"github.com/liCells/controller/global"
	"io"
	"os"
)

func GetConfig() {
	args := os.Args
	var file *os.File
	var err error

	if len(args) == 2 {
		file, err = os.OpenFile(args[1], os.O_CREATE|os.O_RDWR, 0666)
	} else if len(args) == 1 {
		//file, err = os.OpenFile("/etc/config.json", os.O_CREATE|os.O_RDWR, 0666)
		file, err = os.OpenFile("./config.json", os.O_CREATE|os.O_RDWR, 0666)
	} else {
		panic("Usage: <config_file_path> (default: ./config.json)")
	}
	if err != nil {
		panic("Config file read error")
	}
	defer file.Close()

	configJson, _ := io.ReadAll(file)
	err = json.Unmarshal(configJson, &global.Config)
	if err != nil {
		panic("Json parse error")
	}
}
