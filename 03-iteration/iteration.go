package iteration

import "strings"

// Strings in Go are immutable, meaning every concatenation, such as in our Repeat function, involves copying memory to accommodate the new string. This impacts performance, particularly during heavy string concatenation.
// The standard library provides the strings.BuilderstringsBuilder type which minimizes memory copying. It implements a WriteString method which we can use to concatenate strings
const repeatCount = 5

func Repeat(character string, times int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

// func Repeat(character string, times int) string {
// 	var repeated string
// 	for i := 0; i < times; i++ {
// 		repeated += character
// 	}
// 	return repeated
// }