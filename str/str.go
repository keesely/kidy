// str.go kee > 2020/12/10

package str

import (
	"kidy/utils"
	"math"
	"math/rand"
	"strings"
	"time"
)

// Return the remainder of a string after a given value.
func After(subject, search string) string {
	if search == "" {
		return subject
	}
	values := strings.Split(subject, search)[1:]
	if len(values) > 0 {
		return strings.Join(values, search)
	}
	return subject
}

// Get the portion of a string before a given value.
func Before(subject, search string) string {
	if search == "" {
		return subject
	}
	return strings.Split(subject, search)[0]
}

// Return the length of the given string.
func Length(value string) int {
	return len([]rune(value))
}

// Convert a value to camel case.
func Camel(value string) string {
	return Studly(value)
}

// Convert a value to studly caps case.
func Studly(value string) (cstr string) {
	value = strings.Replace(value, "-", " ", -1)
	value = strings.Replace(value, "_", " ", -1)
	value = UcWords(value)
	return strings.Replace(value, " ", "", -1)
}

// Convert a string to snake case.
func Snake(value, delimiter string) string {
	dlt := []byte(delimiter)
	bts := []byte(strings.Replace(UcWords(value), " ", "", -1))
	var ss []byte
	for i, b := range bts {
		if 65 <= b && 90 >= b {
			b += 32
			if i != 0 {
				ss = append(ss, dlt...)
			}
		}
		ss = append(ss, b)
	}
	return string(ss)
}

// Make a string's first character uppercase.
func UcFirst(value string) string {
	str := []byte(value)
	if 97 <= str[0] && 122 >= str[0] {
		str[0] -= 32
	}
	return string(str)
}

// Uppercase the first character of each word in a string
// The optional separators contains the word separator characters
func UcWords(value string, separators ...rune) string {
	if len(separators) == 0 {
		separators = []rune(" ")
	}
	values := Split(value, separators...)
	for i, str := range values {
		values[i] = UcFirst(str)
	}
	return strings.Join(values, " ")
}

// Generates cryptographically secure pseudo-random bytes
// The length of the random string that should be returned in bytes
func RandomBytes(n int) []byte {
	const letterBytes = "qwertyuiopasdfghjklzxcvbnm-QWERTYUIOPASDFGHJKLZXCVBNM_1234567890"
	bts := make([]byte, n)
	for i := range bts {
		rand.Seed(time.Now().UnixNano() + int64(n-i))
		bts[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return bts
}

// Split a string by a string
// ...separators The boundary string
func Split(str string, separators ...rune) []string {
	return strings.FieldsFunc(str, func(r rune) (flag bool) {
		for _, s := range separators {
			if s == r {
				flag = true
			}
		}
		return
	})
}

const (
	KIDY_STR_PAD_LEFT  = 0
	KIDY_STR_PAD_RIGHT = 1
	KIDY_STR_PAD_BOTH  = 2
)

// Padding a string to a certain length with another string
// The `padStr` may be truncated if the required number of padding characters can't be evenly divided by the pad_string's length.
// If the value of `length` is negative, less than, or equal to the length of the input string, no padding takes place, and string will be returned.
// Optional argument `padType` can be
// KDIY_STR_PAD_RIGHT, KDIY_STR_PAD_LEFT, or KIDY_STR_PAD_BOTH.
// If pad_type is not specified it is assumed to be STR_PAD_RIGHT
func Padding(str, padStr string, length, padType int) string {
	LEN := Length(str)
	if LEN >= length {
		return str
	}
	offset := length - LEN
	runes, rpad := []rune(str), []rune(padStr)
	var padding []rune
	for i := 0; i < offset/len(rpad); i++ {
		padding = append(padding, rpad...)
	}
	switch padType {
	case KIDY_STR_PAD_LEFT:
		runes = append(padding, runes...)
	case KIDY_STR_PAD_RIGHT:
		runes = append(runes, padding...)
	case KIDY_STR_PAD_BOTH:
		f := float64(offset)
		l, r := int(math.Floor(f/2)), int(math.Ceil(f/2))
		runes = append(runes, padding[:r]...)
		runes = append(padding[:l], runes...)
	}
	return string(runes)
}

// Return part of a string
// If `length` is given and is negative, then that many characters will be omitted from the end of string (after the start position has been calculated when a offset is negative). If offset denotes the position of this truncation or beyond, false will be returned.
// If `length` is given and is positive, the string returned will contain at most length characters beginning from offset (depending on the length of string).
func SubString(str string, offset, length int) string {
	runes := []rune(str)
	if offset >= len(runes) {
		return ""
	}
	if -1 == length {
		length = len(runes) - offset
	}
	l := offset + length
	l = utils.Ternary(l > len(runes), len(runes), l).(int)
	return string(runes[offset:l])
}
