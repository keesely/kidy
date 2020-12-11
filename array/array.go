// array.go kee > 2020/12/11

package array

import (
	"reflect"
)

type Array []interface{}

// Create an `Array`
func NewArray(array ...interface{}) Array {
	a := Array{}
	if len(array) > 1 {
		a.Push(array...)
	} else if len(array) == 1 {
		v := reflect.ValueOf(array[0])
		for i := 0; i < v.Len(); i++ {
			a.Push(v.Index(i).Interface())
		}
	}
	return a
}

// Push one or more elements onto the end of `Array`
func (a *Array) Push(values ...interface{}) {
	*a = append(*a, values...)
}

// Merge one or more arrays
func (a *Array) Concat(array ...interface{}) {
	for _, val := range array {
		v := reflect.ValueOf(val)
		for i := 0; i < v.Len(); i++ {
			*a = append(*a, v.Index(i).Interface())
		}
	}
}

// Returns the `Array` length
func (a Array) Length() int {
	return len(a)
}

// Append one or more elements onto the end of `Array`
func (a *Array) Append(values ...interface{}) {
	a.Push(values...)
}

// Prepend one or more elements to the beginning of an `Array`
func (a *Array) UnShift(values ...interface{}) int {
	if len(values) == 0 {
		return len(*a)
	}
	values = append(values, *a...)
	*a = values
	return len(*a)
}

// Shift an element off the beginning of `Array`
func (a *Array) Shift() interface{} {
	var x interface{}
	x, *a = (*a)[0], (*a)[1:]
	return x
}

// Pop the element off the end of `Array`
func (a *Array) Pop() interface{} {
	ep := len(*a) - 1
	var x interface{}
	x, *a = (*a)[ep], (*a)[:ep]
	return x
}

// Checks if the given index exists in the array
func (a Array) Exists(i int) bool {
	return len(a) > i && nil != a[i]
}

// Returns given the index value in the `Array`
func (a Array) Index(i int) interface{} {
	if a.Exists(i) {
		return a[i]
	}
	return nil
}

// Searches the array for a given value and returns the `index` if successful
// If the given value no exists in the array, return -1
func (a Array) Search(val interface{}) int {
	for i, v := range a {
		if val == v {
			return i
		}
	}
	return -1
}

// destroys the given index in the array
func (a *Array) Unset(index int) bool {
	if index >= 0 && index < len(*a) {
		*a = append((*a)[:index], (*a)[index+1:]...)
		return true
	}
	return false
}

// Return an array or slice with elements in reverse order
func Reverse(data interface{}) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if reflect.Ptr == t.Kind() {
		t = t.Elem()
		v = v.Elem()
	}
	n, x, y := v.Len(), 0, 0
	switch t.Kind() {
	case reflect.Slice, reflect.Array:
		{
			for x = 0; x < n/2; x++ {
				y = n - x - 1
				vx, vy := v.Index(x).Interface(), v.Index(y).Interface()
				// data[x], data[y] = data[y], data[x]
				v.Index(x).Set(reflect.ValueOf(vy))
				v.Index(y).Set(reflect.ValueOf(vx))
			}
		}
	}
}
