package tcpserv

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

var ServerSettings settings

type settings struct {
	DebugServer bool
}

func SettingsLoad() {
	ReadXMLConfig("settings.xml", &ServerSettings)
}

func ReadXMLConfig(fileName string, strConf interface{}) {

	xmlFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	fmt.Println("Load Config :", fileName)
	byteValue, _ := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, strConf)
	if err != nil {
		fmt.Println(err)
	}
}
