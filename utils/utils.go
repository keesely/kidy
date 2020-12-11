// utils.go kee > 2020/02/14

package utils

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"
)

/**
 * 处理三元运算
 *
 * @param {bool}  cond  输入运算条件
 * @param {interface{}} Tval 如果符合条件则返回Tval
 * @param {interface{}} Fval 如果条件不符合，则返回Fval
 *
 * @return interface{}
 * */
func Ternary(cond bool, Tval, Fval interface{}) interface{} {
	if cond {
		return Tval
	}
	return Fval
}

/**
 * 处理深度拷贝
 *
 * @param {interface{}} value 需要深拷贝的值
 *
 * @return interface{}
 * */
func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	}

	return value
}

/**
 * 获取系统环境变量
 *
 * @param {string}  key 环境变量key
 * @param {interface{}} def 如果不存在则返回默认, 缺省为nil
 *
 * @return interface{}
 * */
func GetEnv(key string, def ...interface{}) interface{} {
	var _def interface{}
	if def != nil && def[0] != nil {
		_def = def[0]
	}

	val := os.Getenv(key)
	return Ternary(val != "", val, _def)
}

/**
 * 获取数据类型名称
 *
 * @return string
 * */
func Typeof(value interface{}) string {
	return fmt.Sprintf("%T", value)
}

/**
 * 生成指定区间随机数
 * @parans int  x 起始区间
 * @params int  y 结束区间
 *
 * @return int
 * */
func Rand(x, y int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(y-x+1) + x
}

// 接口转换到结构体
func ConverStruct(inter map[string]interface{}, conver interface{}, tag string) {
	cRef := reflect.ValueOf(conver).Elem()
	for i := 0; i < cRef.NumField(); i++ {
		fieldInfo := cRef.Type().Field(i) // a reflect.StructField
		name := fieldInfo.Name
		//typeof := fieldInfo.Type
		fTag := fieldInfo.Tag // a reflect.StructTag
		tagName := fTag.Get(tag)
		if tagName != "" {
			name = tagName
		}
		if v, ok := inter[name]; ok {
			cRef.Field(i).Set(reflect.ValueOf(v))
		}
	}
	return
}

func VarDump(args ...interface{}) {
	fmt.Println("{")
	for i, val := range args {
		fmt.Println("  ["+strconv.Itoa(i)+"]", Typeof(val), " => ", val)
	}
	fmt.Println("}")
}

func DumpDone(args ...interface{}) {
	VarDump(args...)
	os.Exit(1)
}
