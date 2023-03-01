/**
* @description :
* @author : Jarick
* @Date : 2022-08-26
* @Url : http://CloudWebOps
 */

package model

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type DefaultConfig struct {
	Main       string     `json:"main" ,yaml:"main"`
	Version    string     `json:"version" ,yaml:"version"`
	Error      error      `json:"error" ,yaml:"error"`
	Kubedeploy Kubedeploy `json:"kubedeploy" ,yaml:"kubedeploy"`
	Kuberes    Kuberes    `json:"kuberes" ,yaml:"kuberes"`
	Kubens     Kubens     `json:"kubens" ,yaml:"kubens"`
	Kubetrace  Kubetrace  `json:"kubetrace" ,yaml:"kubetrace"`
	Kubediags  Kubediags  `json:"kubediags" ,yaml:"kubediags"`
}

type Kubedeploy struct {
	Title string `json:"title" ,yaml:"title"`
}
type Kubens struct {
	Title string `json:"title" ,yaml:"title"`
}
type Kuberes struct {
	Title string `json:"title" ,yaml:"title"`
}
type Kubetrace struct {
	Title   string `json:"title" ,yaml:"title"`
	Timeout int    `json:"timeout" ,yaml:"timeout"`
}
type HTTPServerDebug struct {
	Title string `json:"title" ,yaml:"title"`
}
type Kubediags struct {
	PODUnAvailableWarning  float64  `json:"PODUnAvailableWarning" ,yaml:"PODUnAvailableWarning"`
	PODUnAvailableError    float64  `json:"PODUnAvailableError" ,yaml:"PODUnAvailableError"`
	CPULimitInfo           float64  `json:"CPULimitInfo" ,yaml:"CPULimitInfo"`
	MEMLimitInfo           float64  `json:"MEMLimitInfo" ,yaml:"MEMLimitInfo"`
	CPULimitWarning        float64  `json:"CPULimitWarning" ,yaml:"CPULimitWarning"`
	MEMLimitWarning        float64  `json:"MEMLimitWarning" ,yaml:"MEMLimitWarning"`
	MEMOversoldNodeInfo    float64  `json:"MEMOversoldNodeInfo" ,yaml:"MEMOversoldNodeInfo"`
	MEMOversoldNodeWarning float64  `json:"MEMOversoldNodeWarning" ,yaml:"MEMOversoldNodeWarning"`
	RequestCPUInfo         float64  `json:"RequestCPUInfo" ,yaml:"RequestCPUInfo"`
	RequestCPUWarning      float64  `json:"RequestCPUWarning" ,yaml:"RequestCPUWarning"`
	RequestMemInfo         float64  `json:"RequestMemInfo" ,yaml:"RequestMemInfo"`
	RequestMemWarning      float64  `json:"RequestMemWarning" ,yaml:"RequestMemWarning"`
	PODCapacityNodeWarning float64  `json:"PODCapacityNodeWarning" ,yaml:"PODCapacityNodeWarning"`
	PODErrorNodeWarning    float64  `json:"PODErrorNodeWarning" ,yaml:"PODErrorNodeWarning"`
	SystemNamespace        []string `json:"SystemNamespace" ,yaml:"SystemNamespace"`
}

// 初始化，配置默认参数
func InitDefaultConfig() *DefaultConfig {
	return &DefaultConfig{
		Main:      "CloudOpsCMD",
		Version:   "0.2",
		Kubetrace: Kubetrace{Title: "Kubetrace", Timeout: 2},
		Kubediags: Kubediags{
			// deployment
			PODUnAvailableWarning: 1,
			PODUnAvailableError:   5,

			CPULimitInfo:    3,
			CPULimitWarning: 5,
			MEMLimitInfo:    3,
			MEMLimitWarning: 5,

			// node
			MEMOversoldNodeInfo:    2,
			MEMOversoldNodeWarning: 4,
			RequestCPUInfo:         0.75,
			RequestCPUWarning:      0.85,
			RequestMemInfo:         0.75,
			RequestMemWarning:      0.85,

			PODCapacityNodeWarning: 80,
			PODErrorNodeWarning:    100,
			SystemNamespace:        []string{"cpaas-system", "kube-system", "tcnp", "tke", "cert-manager"},
		},
	}
}

// 获取配置文件参数
func (c *DefaultConfig) GetConf(configDir string) *DefaultConfig {

	// 判断路径
	var configPath string
	if configDir == "" {
		configPath = "./config.yaml"
	} else {
		configPath = configDir + "/config.yaml"
	}

	// yamlFile, err := ioutil.ReadFile(configPath)
	yamlFile, err := ioutil.ReadFile(filepath.Clean("./" + configPath))
	if err != nil {
		c.Error = err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		c.Error = err
	}
	return c
}

// y := model.InitDefaultConfig().GetConf(configDir).Main
// fmt.Println("yyyyyyyyyyyyyyyyy:", y)

// var c model.DefaultConfig
// x := c.GetConf(configDir).Main
// fmt.Println("xxxxxxxxxxxxxxxx:", x)

// 生成默认配置文件
func Genconfig(configDir string) (err error) {
	// Marshal yaml
	data := InitDefaultConfig()
	confData, err := yaml.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	// write the config File
	err = ioutil.WriteFile(configDir+"/cloudopscmd.yaml", confData, 0600)
	return
}
