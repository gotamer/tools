// rand number, random strings
package tools

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// generates a random number between min and max
func IRand(min, max int) int {

	var length int

	if min < max {
		length = min + rand.Intn(max-min)
	} else {
		length = min
	}
	return rand.Intn(length)
}

// generates a random string
func SRand(min, max int, readable bool) string {

	var length int
	var char string

	if min < max {
		length = min + rand.Intn(max-min)
	} else {
		length = min
	}

	if readable == false {
		char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	} else {
		char = "ABCDEFHJLMNQRTUVWXYZabcefghijkmnopqrtuvwxyz23479"
	}

	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = char[rand.Intn(len(char)-1)]
	}
	return string(buf)
}

// Shuffles the string
func Shuffle(s string) string {
	l := len(s)
	b := []byte(s)
	for i := len(b) - 1; i != 0; i-- {
		r := rand.Intn(l)
		b[i], b[r] = b[r], b[i]
	}
	return string(b)
}
