// config.go kee > 2020/02/14

package config

import (
	"encoding/json"
	"fmt"
	"kidy/utils"
	"strings"
	"sync"
)

type Config struct {
	data map[string]interface{}
	sync.RWMutex
}

func NewConfig(data map[string]interface{}) *Config {
	return &Config{data: data}
}

func (c *Config) GetAll() map[string]interface{} {
	return c.data
}

func (c *Config) Get(key string, defValue ...interface{}) interface{} {
	if 0 == len(key) {
		return nil
	}
	var value interface{}
	if nil == defValue || nil == defValue[0] {
		value = nil
	} else {
		value = defValue[0]
	}
	if val := c.getData(key); val != nil {
		return val
	}
	return value
}

func (c *Config) Set(key string, value interface{}) error {
	if 0 == len(key) {
		return fmt.Errorf("key is empty")
	}
	c.Lock()
	defer c.Unlock()

	switch value.(type) {
	case *Config:
		value = value.(*Config).GetAll()
	}

	var data = c.data
	keys := strings.Split(key, ".")
	for i, k := range keys {
		if v, ok := data[k]; ok {
			switch v.(type) {
			case map[string]interface{}:
				{
					data[k] = v
					if i == len(keys)-1 {
						v = value
						data[k] = v
					}
					data = v.(map[string]interface{})
				}
			default:
				{
					v = value
					data[k] = v
				}
			}
		} else {
			data[k] = make(map[string]interface{})
			if i == len(keys)-1 {
				data[k] = value
			} else {
				vv := make(map[string]interface{})
				data[k] = vv
				data = data[k].(map[string]interface{})
			}
		}
	}
	return nil
}

func (c *Config) getData(key string) interface{} {
	if 0 == len(key) {
		return c.data
	}

	c.RLock()
	defer c.RUnlock()

	data := utils.DeepCopy(c.data).(map[string]interface{})
	//var data = c.data
	keys := strings.Split(key, ".")

	for i, k := range keys {
		if v, ok := data[k]; ok {
			switch v.(type) {
			case map[string]interface{}:
				{
					data = v.(map[string]interface{})
					if i == len(keys)-1 {
						return NewConfig(data)
					}
				}
			default:
				{
					return v
				}
			}
		}
	}
	return nil
}

func (c *Config) ToJSON() []byte {
	jStr, _ := json.Marshal(c.data)
	return jStr
}
