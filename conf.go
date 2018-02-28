package main

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

//data structure to database configuration 
type conf struct {
	Host string 
	User string
	Port string
	Name string
    Password string
    Cachetime string
    Serveport string
}


func (config *conf) getConf() *conf {

    yamlFile, err := ioutil.ReadFile("conf.yml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, config)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return config
}


func (config *conf) getUrlConnection(conector string) string {
	if conector == "postgres"{
		return "host=" + config.Host + " port=" + config.Port + " user=" + config.User + " password=" + config.Password + " dbname=" + config.Name
	}
	return ""
}

