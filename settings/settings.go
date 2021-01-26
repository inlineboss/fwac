package settings

import (
	"fmt"
	"sync"
)

type settings struct {
	Password string `json:"password"`
	Root     string `json:"root"`
	Port     string `port:"port"`
}

var instance *settings
var once sync.Once

//GetInstance -
func GetInstance() *settings {
	once.Do(func() {
		instance = &settings{}
	})
	return instance
}

func (conf *settings) Print() {
	fmt.Printf(`
Password: %s
Root: %s
Port: %s
	`, conf.Password, conf.Root, conf.Port)
}
