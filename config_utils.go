package commons

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "os"
)

const (
    DebugMode   = "debug"
    ReleaseMode = "release"
    TestMode    = "test"
)

type configUtils struct {
}

var ConfigUtils *configUtils

type ConfigFormat interface {
    Version() string
}

func (cu *configUtils) LoadConfigFile(confPath string, pConfig interface{}) {
    configFile, err := ioutil.ReadFile(confPath)
    if err != nil {
        log.Fatal("Load yaml config file error: ", err)
    }

    err = yaml.Unmarshal(configFile, pConfig)
    if err != nil {
        log.Fatal("Load yaml config file error: ", err)
    }
    return
}

func (cu *configUtils) SaveConfigFile(confPath string, config ConfigFormat) {
    configBytes, err := yaml.Marshal(config)
    if err != nil {
        log.Fatal("Save yaml config file error: ", err)
    }

    err = ioutil.WriteFile(confPath, configBytes, os.ModePerm)
    if err != nil {
        log.Fatal("Save yaml config file error: ", err)
    }
    return
}
