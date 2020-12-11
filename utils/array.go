// array.go kee > 2020/12/05

package utils

type Array []interface{}

func NewArray(array ...interface{}) Array {
	a := Array{}
	a.Concat(array)
	return a
}

func (a Array) Len() int {
	return len(a)
}

func (a *Array) Push(values ...interface{}) {
	*a = append(*a, values...)
}

func (a *Array) Append(values ...interface{}) {
	a.Push(values...)
}

func (a *Array) UnShift(values ...interface{}) int {
	if len(values) == 0 {
		return len(*a)
	}
	values = append(values, *a...)
	*a = values
	return len(*a)
}

func (a *Array) Shift() interface{} {
	var x interface{}
	x, *a = (*a)[0], (*a)[1:]
	return x
}

func (a *Array) Pop() interface{} {
	ep := len(*a) - 1
	var x interface{}
	x, *a = (*a)[ep], (*a)[:ep]
	return x
}

func (a *Array) Concat(arrays ...[]interface{}) {
	for _, v := range arrays {
		*a = append(*a, v...)
	}
}

func (a Array) Exists(i int) bool {
	return len(a) > i && nil != a[i]
}

func (a Array) Index(i int) interface{} {
	if a.Exists(i) {
		return a[i]
	}
	return nil
}

func (a Array) Search(val interface{}) int {
	for i, v := range a {
		if val == v {
			return i
		}
	}
	return -1
}

func (a *Array) Unset(index int) bool {
	if index >= 0 && index < len(*a) {
		*a = append((*a)[:index], (*a)[index+1:]...)
		return true
	}
	return false
}
