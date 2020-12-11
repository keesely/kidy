// test.go kee > 2020/12/11

package array

import (
	"fmt"
	"kidy/test"
	"testing"
)

func TestReverse(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Reverse(a)
	test.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, a)

	b := [5]int{9, 8, 7, 6, 5}
	Reverse(&b)
	test.Equal(t, [5]int{5, 6, 7, 8, 9}, b)

	c := []string{"1", "2", "3", "4", "5"}
	Reverse(c)
	test.Equal(t, []string{"5", "4", "3", "2", "1"}, c)

	d := []byte("01234567")
	Reverse(d)
	test.Equal(t, []byte("76543210"), d)
	fmt.Println("Testing Reverse done.")
}

func TestArray(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []string{"0"}
	A := NewArray(b, a)
	test.Equal(t, A, Array{[]string{"0"}, []int{1, 2, 3, 4, 5}})

	A = NewArray(a)
	s := []string{"6", "7", "8"}
	A.Concat(s)
	test.Equal(t, A, Array{1, 2, 3, 4, 5, "6", "7", "8"})
	test.UnEqual(t, A, Array{1, 2, 3, 4, 5, 6, 7, 8})

	A.Unset(5)
	test.Equal(t, A, Array{1, 2, 3, 4, 5, "7", "8"})
	test.Equal(t, A.Index(5), "7")
	test.Equal(t, A.Length(), 7)
	A.Shift()
	test.Equal(t, A, Array{2, 3, 4, 5, "7", "8"})
	A.Pop()
	A.Pop()
	test.Equal(t, A, Array{2, 3, 4, 5})
	A.UnShift(0, 1)
	test.Equal(t, A, Array{0, 1, 2, 3, 4, 5})
	test.Equal(t, A.Search(0), 0)

	Reverse(A)
	test.UnEqual(t, A, Array{0, 1, 2, 3, 4, 5})
	fmt.Println(A)

	fmt.Println("Testing Array done.")
}
