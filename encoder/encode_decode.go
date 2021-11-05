package encoder

import (
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length   = 63
)

func Encode(number int64) string {
	encodedString := ""
	for number > 0 {
		encodedString = string(alphabet[number%length]) + encodedString
		number = number / length
	}
	return encodedString
}

func Decode(encoded string) int64 {
	decodedNumber := 0
	l := len([]rune(encoded))
	for i := 0; i < l; i++ {
		decodedNumber = decodedNumber*length + strings.Index(alphabet, string(encoded[i]))
	}
	return int64(decodedNumber)
}
