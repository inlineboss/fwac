package settings

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func confFromJSON(filePath *string) error {
	file, err := os.Open(*filePath)
	if err != nil {
		return errors.New("open configure file failed")
	}

	defer file.Close()

	rawData, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.New("read configure file failed")
	}

	err = json.Unmarshal(rawData, GetInstance())
	if err != nil {
		return errors.New("parse configure file failed")
	}

	return nil
}

func confFromArgs(pass *string, root *string, port *string) error {

	if *pass == "" {
		fmt.Println("WARNING: Not input password!")
	}

	if *root == "" {
		return errors.New("ERROR: root not specified")
	}

	if *port == "" {
		return errors.New("ERROR: port not specified")
	}

	GetInstance().Password = *pass
	GetInstance().Root = *root
	GetInstance().Port = *port

	return nil
}

func init() {

	// pass := flag.String("pass", "", "Password for access to direcory")
	// root := flag.String("dir", "", "Root directory")
	// port := flag.String("port", "", "Listening port")
	// conf := flag.String("conf", "", "Path to configure file")

	pass := flag.String("pass", "", "Password for access to direcory")
	root := flag.String("dir", "/Users/inlineboss/", "Root directory")
	port := flag.String("port", "8080", "Listening port")
	conf := flag.String("conf", "", "Path to configure file")

	flag.Parse()

	if *conf != "" {

		err := confFromJSON(conf)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

	} else {

		err := confFromArgs(pass, root, port)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

	}

}
