package commons

import (
    "encoding/json"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "os"
)

type AppMode string

const (
    DebugMode   AppMode = "debug"
    TestMode    AppMode = "test"
    ReleaseMode AppMode = "release"
)

type ConfigFormat interface {
    AppMode() string
}

func LoadYamlConfigFile(_confPath string, _pConfig interface{}) {
    configFile, err := ioutil.ReadFile(_confPath)
    if err != nil {
        log.Fatal("Load yaml file read error: ", err)
    }

    err = yaml.Unmarshal(configFile, _pConfig)
    if err != nil {
        log.Fatal("Load yaml config file error: ", err)
    }
    return
}

func SaveYamlConfigFile(_confPath string, _config ConfigFormat) {
    configBytes, err := yaml.Marshal(_config)
    if err != nil {
        log.Fatal("Save yaml config file error: ", err)
    }

    err = ioutil.WriteFile(_confPath, configBytes, os.ModePerm)
    if err != nil {
        log.Fatal("Save yaml config file error: ", err)
    }
    return
}

func LoadJsonConfigFile(_confPath string, _pConfig interface{}) {
    configFile, err := ioutil.ReadFile(_confPath)
    if err != nil {
        log.Fatal("Load json file read error: ", err)
    }

    err = json.Unmarshal(configFile, _pConfig)
    if err != nil {
        log.Fatal("Load json config file error: ", err)
    }
    return
}

func SaveJsonConfigFile(_confPath string, _config ConfigFormat) {
    configBytes, err := json.Marshal(_config)
    if err != nil {
        log.Fatal("Save json config file error: ", err)
    }

    err = ioutil.WriteFile(_confPath, configBytes, os.ModePerm)
    if err != nil {
        log.Fatal("Save json config file error: ", err)
    }
    return
}
