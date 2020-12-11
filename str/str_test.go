// str_test.go kee > 2020/12/10

package str

import (
	"fmt"
	"kidy/test"
	"testing"
)

func TestStr(t *testing.T) {
	test.Equal(t, After("hello world", "wo"), "rld")
	test.UnEqual(t, After("hello world", "wc"), "rld")

	test.Equal(t, Before("hello world", "wo"), "hello ")
	test.UnEqual(t, Before("hello world", "wc"), "hello ")

	test.Equal(t, UcFirst("hello world"), "Hello world")
	test.Equal(t, UcWords("hello world"), "Hello World")
	test.UnEqual(t, UcWords("hello world", []rune("\r")...), "Hello World")
	test.Equal(t, UcWords("Hi ki.dy:\n  Hello world", []rune(" \r\n.:")...), "Hi Ki Dy Hello World")

	test.Equal(t, Camel("uc_first"), "UcFirst")
	test.Equal(t, Studly("str-test"), "StrTest")
	test.Equal(t, Snake("Hello World", "_"), "hello_world")
	test.Equal(t, Snake("UcFirst", "_"), "uc_first")
	test.Equal(t, Snake("I Live In very old town", "-"), "i-live-in-very-old-town")

	r1, r2 := string(RandomBytes(10)), string(RandomBytes(10))
	test.UnEqual(t, r1, r2)
	fmt.Printf("random_str(1): %v \n", r1)
	fmt.Printf("random_str(2): %v \n", r2)

	lpad, rpad, bpad := Padding("11", "0", 5, KIDY_STR_PAD_LEFT),
		Padding("11", "0", 5, KIDY_STR_PAD_RIGHT),
		Padding("11", "0", 5, KIDY_STR_PAD_BOTH)

	test.Equal(t, lpad, "00011")
	test.Equal(t, Length(lpad), 5)
	test.Equal(t, rpad, "11000")
	test.Equal(t, Length(rpad), 5)
	test.Equal(t, bpad, "01100")
	test.Equal(t, Length(bpad), 5)

	substr := SubString("helloworld", 5, -1)
	test.Equal(t, substr, "world")
	test.Equal(t, Length(substr), 5)

	test.Equal(t, Length("汉字长度"), 4)
	test.Equal(t, Length("汉字长度 string length"), 18)
	test.Equal(t, Length("This is a String"), 16)
}
