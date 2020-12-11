// values_test.go kee > 2020/12/05

package utils_test

import (
	"encoding/json"
	"fmt"
	"okauth/utils"
	"testing"
)

func TestValueSet(t *testing.T) {
	values := utils.Values{}

	json.Unmarshal([]byte(`{
		"a" : 34325,
		"b" : "hello world",
		"c" : [
			{ "s" : 1 },
			{ "b" : 33 }
		],
		"d" : [0,1,2,3,4,5,6,7,8,9],
		"mc" : {
			"dc" : "va2"
		}
	}`), &values)

	values.Set("project.name", "test Values")
	values.Set("project.env", "developer")
	values.Set("watch.deep.values", utils.Array{1, 2, 3, 4})
	values.Set("mc.dc", "v3")
	values.Set("watch.deep.values.1", 1000)
	values.Add("c", utils.Values{"c": "c2c3"})

	fmt.Println(values)
	fmt.Println("--------------------------------------------------")

	fmt.Println("project.name:", values.Get("project.name"))
	fmt.Println("project.1env:", values.Get("project.1env", 100))
	fmt.Println("mc.dc:", values.Get("mc.dc"))
	fmt.Println("c:", values.Get("c"))
	fmt.Println("c.1.b:", values.Get("c.1.b"))
	fmt.Println("c.2.c:", values.Get("c.2.c"))
	fmt.Println("watch.deep.values:", values.Get("watch.deep.values"))
	fmt.Println("c.b:", values.Get("c.b", utils.Array{1, 3, 5, 7, 9}))

	fmt.Println("--------------------------------------------------")
}
