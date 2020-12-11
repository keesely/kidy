// array_test.go kee > 2020/12/05

package utils_test

import (
	"fmt"
	"okauth/utils"
	"testing"
)

var array utils.Array

func TestPush(t *testing.T) {
	array.Push(1, 2, 3, 4)
	fmt.Println(array.Len(), array)
}

func TestShift(t *testing.T) {
	x := array.Shift()
	array.UnShift("100")
	fmt.Println(array.Len(), array, x)
}

func TestPop(t *testing.T) {
	x := array.Pop()
	array.Push("200")
	fmt.Println(array.Len(), array, x)
}

func TestConcat(t *testing.T) {
	b := utils.Array{9, 8, 7}
	c := utils.Array{6, 5, 4}
	array.Concat(b, c)
	fmt.Println(array.Len(), array, array.Index(7))

	for i, v := range array {
		fmt.Printf("%d: %v.(%T) \r\n", i, v, v)
	}
}

func TestSearch(t *testing.T) {
	x := array.Search("200")
	fmt.Println("Search(200):", x)
	array.Unset(x)
	fmt.Println("Search(200) deleted:", array.Search("200"))
	fmt.Println(array)
}
