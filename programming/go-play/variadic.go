package main

import (
	"fmt"
	"strings"
)

type formatter func(s string) string

func format(s string, formatters ...formatter) string {
	for _, fmtr := range formatters {
		s = fmtr(s)
	}
	return s
}

func trim(s string) string {
	return strings.Trim(s, " ")
}

func last(s string) string {
	return s[strings.LastIndexByte(s, ' ')+1:]
}

// From https://blog.learngoprogramming.com/golang-variadic-funcs-how-to-patterns-369408f19085
func toFullname(names ...string) string {
	return strings.Join(names, " ")
}

func numbers(nums ...int) []int {
	nums[0] = 100
	return nums
}

func printVariadic(msgs ...string) {
	fmt.Println("printVariadic ", msgs)
}

func printSlice(msgs []string) {
	fmt.Println("printSlice ", msgs)
}

func variableTypesOfArgs(x string, y ...interface{}) {
	fmt.Println("x ", x)
	fmt.Println("y ", y)
}

func main() {
	fmt.Println()
	fmt.Println(toFullname("oi", "tudubom"))
	// The passed slice shares the same underlying array with the slice inside the func, changing its value inside the func also effects the passed slice
	namexx := []string{"hey", "helou"}
	fmt.Println(toFullname(namexx...))
	fmt.Println(numbers(1, 2))
	/*
		Their type-identities (https://golang.org/ref/spec#Type_identity) are not the same. Let’s assign them to variables:
		variadic := printVariadic
		slicey := printSlice
		variadic = slicey // error: type mismatch
	*/
	variableTypesOfArgs("a", 1, "aaa")
	/*
		letters := []string{"a", "b"}
		variableTypesOfArgs("a", letters...)

		this is not possible because letters is of type string slice, and variableTypesOfArgs function accepts a param of type empty interface
	*/
	// but this works:
	letters := []string{"a", "b"}
	var iletters = make([]interface{}, len(letters))
	for i, letter := range letters {
		iletters[i] = letter
	}
	variableTypesOfArgs("a", iletters...)
	fmt.Println(format(" alan turing ", trim, last, strings.ToUpper))
}
