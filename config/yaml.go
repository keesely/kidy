// yaml.go kee > 2020/02/14

package config

import (
	yaml "gopkg.in/yaml.v2"
	"kidy/utils"
	"regexp"
	"strings"
	"sync"
)

type Yaml struct {
	data map[string]interface{}
	sync.RWMutex
}

func NewYaml(yData []byte) *Config {
	data := make(map[string]interface{})
	e := yaml.Unmarshal(yData, data)
	if e != nil {
		panic("Unmarshal yaml: " + e.Error())
	}
	data = FormatValueMaps(data)
	return NewConfig(data)
}

func FormatValueMaps(m map[string]interface{}) map[string]interface{} {
	// 获取keys
	for k, v := range m {
		switch value := v.(type) {
		case string:
			m[k] = ExpandValueEnv(value)
		case map[string]interface{}:
			m[k] = FormatValueMaps(value)
		case map[interface{}]interface{}:
			_value := make(map[string]interface{})
			for _k, _v := range m[k].(map[interface{}]interface{}) {
				_value[_k.(string)] = _v
			}
			m[k] = _value
			m[k] = FormatValueMaps(_value)
		case map[string]string:
			for k2, v2 := range value {
				value[k2] = ExpandValueEnv(v2)
			}
			m[k] = value
		}
	}
	return m
}

// Convert `$(ENV)` || `$(ENV||defaultValue)` || `$(ENV||)`
// Return the env value || if env is nil return defaultValue || env is nil return ""
func ExpandValueEnv(value string) string {
	rVal := value

	dVal := ""
	regx := regexp.MustCompile(`(?U)\$\{.+\}`)

	if x := regx.FindAllString(rVal, -1); len(x) > 0 {
		for _, v := range x {
			vL := len(v)
			if vL < 3 {
				continue
			}
			if key := v[2 : vL-1]; len(key) > 0 {
				dValal := ""
				dValalIndex := strings.Index(v, "||")
				if dValalIndex > 0 {
					key = v[2:dValalIndex]
					dValal = v[dValalIndex+2 : vL-1]
				}

				eVal := utils.GetEnv(key, dValal).(string)
				rVal = strings.Replace(rVal, v, eVal, -1)
			}
		}
	}

	if rVal == "" {
		rVal = dVal
	}
	return rVal
}

func (c *Config) ToYAML() []byte {
	yStr, _ := yaml.Marshal(c.data)
	return yStr
}
