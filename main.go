package main

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
)

/*
Big Unsigned Integer implementation in Go, with these operations:
1. Add 2 numbers
2. Substract 2 numbers
3. Long Multiply 2 numbers
4. Fast Multiply 2 numbers
5. Multiply 2 numbers using FFT
6. Divide 2 numbers
7. Take modulus of 2 numbers
8. Power 2 numbers

And these utility functions:
1. Read a number (Parser)
2. Print a number (ToString)
3. Random a number
4. Benchmarks

Number representation:
- A slice of int - 32-bit word, < 1'000'000'000, so that their product can fit into 1 int64 variable
*/

func main() {
	input1 := "1123456789123456789123456789"
	input2 := "123123456789123456789"
	input3 := "00000123123456789123456789"
	number1 := ParseUnsignedBigInteger(input1)
	number2 := ParseUnsignedBigInteger(input2)
	number3 := ParseUnsignedBigInteger(input3)
	string1 := UnsignedBigIntegerToString(number1)
	string2 := UnsignedBigIntegerToString(number2)
	string3 := UnsignedBigIntegerToString(number3)
	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(number3)
	fmt.Println(input1, string1, cmp.Equal(input1, string1))
	fmt.Println(input2, string2, cmp.Equal(input2, string2))
	fmt.Println(input3, string3, cmp.Equal(strings.TrimLeft(input3, "0"), string3))
}

// Read a 1 line of number from standard input, then parse it
// Nil return means there is errors or wrong formant in the input
func ParseUnsignedBigInteger(input string) []int {
	if len(input) < 1 || input[0] == '-' {
		return nil
	}
	input = strings.TrimLeft(input, "0")
	tokens := ChunksAndReverseWord(ReverseString(input), 9)
	words := []int{}
	for _, token := range tokens {
		var word int
		_, err := fmt.Sscan(token, &word)
		if err != nil {
			fmt.Println("Error at token:", token)
			return nil
		}
		words = append(words, word)
	}
	return words
}

func UnsignedBigIntegerToString(input []int) string {
	var stringBuilder strings.Builder
	for i := len(input) - 1; i >= 0; i-- {
		stringBuilder.WriteString(fmt.Sprint(input[i]))
	}
	return stringBuilder.String()
}

func ChunksAndReverseWord(input string, chunkSize int) []string {
	fmt.Println(input)
	if len(input) == 0 {
		return nil
	}
	if chunkSize >= len(input) {
		return []string{input}
	}
	chunks := []string{}
	var stringBuilder strings.Builder
	stringBuilder.Grow(chunkSize)
	left := 0
	runes := []rune{}
	for i, right := range input {
		runes = append(runes, right)
		left++
		if left == chunkSize || i == len(input)-1 {
			for i := len(runes) - 1; i >= 0; i-- {
				stringBuilder.WriteRune(runes[i])
			}
			runes = []rune{}
			chunks = append(chunks, stringBuilder.String())
			left = 0
			stringBuilder.Reset()
			stringBuilder.Grow(chunkSize)
		}
	}
	if left > 0 {
		chunks = append(chunks, stringBuilder.String())
	}
	return chunks
}

func ReverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseList[T any](input []T) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}
