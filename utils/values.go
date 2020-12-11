// values.go kee > 2020/12/05

package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Values map[string]interface{}

func getArgsFirst(args ...interface{}) interface{} {
	fmt.Sprintf("%T", args)
	if len(args) > 0 {
		return args[0]
	}
	return nil
}

func (v Values) Get(key string, defaultValue ...interface{}) interface{} {
	return ValuesGet(v, key, defaultValue...)
}

func (v Values) Set(key string, value interface{}) {
	nValues, next := ValuesSet(v, key, value)
	switch nValues.(type) {
	case Values:
		nValues.(Values)[next] = value
	case map[string]interface{}:
		nValues.(map[string]interface{})[next] = value
	case []interface{}:
		{
			index, ok := strconv.Atoi(next)
			cVal := nValues.([]interface{})
			if nil != ok || len(cVal) <= index {
				panic(fmt.Sprintf("index out of range [%d] with length %d", index, len(cVal)))
				//nValues = append(cVal, value)
			} else {
				nValues.([]interface{})[index] = value
			}
		}
	case Array:
		{
			index, e := strconv.Atoi(next)
			cVal := nValues.(Array)
			if nil != e || len(cVal) <= index {
				//nValues = append(nValues.(Array), value)
				panic(fmt.Sprintf("index out of range [%d] with length %d (Array)", index, len(cVal)))
			} else {
				nValues.(Array)[index] = value
			}
		}
	default:
		{
			panic(fmt.Sprintf("Unable to set the value (typeof: %T)", nValues))
		}
	}
}

func (v Values) Add(key string, value interface{}) {
	nValues, next := ValuesSet(v, key, value)

	switch nValues.(type) {
	case Values:
		nValues.(Values)[next] = addVal(nValues.(Values)[next], value)
	case map[string]interface{}:
		nValues.(map[string]interface{})[next] = addVal(nValues.(map[string]interface{}), value)
	case []interface{}:
		{
			index, ok := strconv.Atoi(next)
			cVal := nValues.([]interface{})
			if nil != ok || len(cVal) <= index {
				panic(fmt.Sprintf("index out of range [%d] with length %d", index, len(cVal)))
				//nValues = append(cVal, value)
			} else {
				nValues.([]interface{})[index] = addVal(cVal[index], value)
			}
		}
	case Array:
		{
			index, e := strconv.Atoi(next)
			cVal := nValues.(Array)
			if nil != e || len(cVal) <= index {
				//nValues = append(nValues.(Array), value)
				panic(fmt.Sprintf("index out of range [%d] with length %d (Array)", index, len(cVal)))
			} else {
				nValues.(Array)[index] = addVal(cVal[index], value)
			}
		}
	default:
		{
			panic(fmt.Sprintf("Unable to set the value (typeof: %T)", nValues))
		}
	}
}

func (v Values) Del(key string) {
	nValues, next := ValuesSet(v, key, nil)

	switch nValues.(type) {
	case Values:
		delete(nValues.(Values), next)
	case map[string]interface{}:
		delete(nValues.(map[string]interface{}), next)
	case []interface{}:
		{
			index, ok := strconv.Atoi(next)
			cVal := nValues.([]interface{})
			if nil != ok || len(cVal) <= index {
				panic(fmt.Sprintf("index out of range [%d] with length %d", index, len(cVal)))
			} else {
				nValues = append((cVal)[:index], (cVal)[index+1:]...)
			}
		}
	case Array:
		{
			index, e := strconv.Atoi(next)
			cVal := nValues.(Array)
			if nil != e || len(cVal) <= index {
				//nValues = append(nValues.(Array), value)
				panic(fmt.Sprintf("index out of range [%d] with length %d (Array)", index, len(cVal)))
			} else {
				cVal.Unset(index)
				nValues = cVal
			}
		}
	default:
		{
			panic(fmt.Sprintf("Unable to set the value (typeof: %T)", nValues))
		}
	}
}

func ValuesSet(values Values, key string, value interface{}) (interface{}, string) {
	if "" == key {
		values = value.(Values)
	}

	var (
		k       string
		nValues interface{}
	)
	nValues = values
	keys := strings.Split(key, ".")
	for len(keys) > 1 {
		k, keys = keys[0], keys[1:]
		nValues = FillValues(nValues, k)
	}
	return nValues, keys[0]
}

func ValuesGet(values Values, key string, _default ...interface{}) interface{} {
	defaultValue := getArgsFirst(_default...)

	if "" == key {
		return defaultValue
	}

	var nValues interface{}
	nValues = values
	keys := strings.Split(key, ".")
	for i, segment := range keys {
		if len(keys) == i {
			return nValues
		}

		val, ok := ValueGet(nValues, segment)
		if ok {
			nValues = val
		} else {
			return defaultValue
		}
	}

	return nValues
}

func FillValues(value interface{}, k string) interface{} {
	var s interface{}
	switch value.(type) {
	case Values:
		{
			val, ok := value.(Values)[k]
			if !ok {
				val = Values{}
			}
			s = val
			value.(Values)[k] = s
			s = value.(Values)[k]
		}
	case map[string]interface{}:
		{
			val, ok := value.(map[string]interface{})[k]
			if !ok {
				val = Values{}
			}

			s = Values{}
			for k, v := range val.(map[string]interface{}) {
				s.(Values)[k] = v.(interface{})
			}
			value.(map[string]interface{})[k] = s
			s = value.(map[string]interface{})[k]
		}
	case []interface{}:
		{
			var val interface{}
			cVal := value.([]interface{})
			if i, ok := strconv.Atoi(k); ok == nil {
				if i < len(cVal) {
					val = cVal[i]
				} else {
					val = Values{}
					i = len(cVal)
					value = append(value.([]interface{}), val)
				}
				value.([]interface{})[i] = val
				s = value.([]interface{})[i]
			} else {
				value = append(value.([]interface{}), Values{})
				s = value.([]interface{})[len(cVal)]
			}
		}
	case Array:
		{
			var val interface{}
			cVal := value.(Array)
			if i, ok := strconv.Atoi(k); ok == nil {
				if i < len(cVal) {
					val = cVal[i]
				} else {
					val = Values{}
					i = len(cVal)
					value = append(value.(Array), val)
				}
				value.(Array)[i] = val
				s = value.(Array)[i]
			} else {
				value = append(value.(Array), Values{})
				s = value.(Array)[len(cVal)]
			}
		}
	}
	return s
}

func ValueGet(value interface{}, key string) (val interface{}, ok bool) {
	arrayValue := func(value []interface{}, key string) (val interface{}, ok bool) {
		i, e := strconv.Atoi(key)
		if e != nil {
			return
		}

		val = value[i]
		if nil != val {
			ok = true
		}
		return
	}

	switch value.(type) {
	case Values:
		val, ok = value.(Values)[key]
	case map[string]interface{}:
		val, ok = value.(map[string]interface{})[key]
	case []interface{}:
		val, ok = arrayValue(value.([]interface{}), key)
	case Array:
		val, ok = arrayValue(value.(Array), key)
	default:
		tof := reflect.TypeOf(value)
		if -1 < (Array{reflect.Struct}).Search(tof.Kind()) {
			vof := reflect.ValueOf(value)
			for i := 0; i < tof.NumField(); i++ {
				field := tof.Field(i)
				if key == field.Name {
					val = vof.Field(i).Interface()
					ok = true
				}
			}
		}
		//val, ok = value, true
	}
	return
}

func addVal(cVal interface{}, value interface{}) interface{} {
	switch cVal.(type) {
	case []interface{}:
		{
			cVal = append(cVal.([]interface{}), value)
			return cVal
		}
	case Array:
		{
			val := cVal.(Array)
			val.Push(value)
			return val
		}
	default:
		return Array{cVal, value}
	}
	return nil
}
