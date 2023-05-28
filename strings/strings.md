package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains
	fmt.Println(strings.Contains("hello world", "world")) // true
	fmt.Println(strings.Contains("hello world", "foo"))   // false

	// HasPrefix
	fmt.Println(strings.HasPrefix("hello world", "hello")) // true
	fmt.Println(strings.HasPrefix("hello world", "foo"))   // false

	// HasSuffix
	fmt.Println(strings.HasSuffix("hello world", "world")) // true
	fmt.Println(strings.HasSuffix("hello world", "foo"))   // false

	// Index
	fmt.Println(strings.Index("hello world", "o")) // 4
	fmt.Println(strings.Index("hello world", "z")) // -1

	// LastIndex
	fmt.Println(strings.LastIndex("hello world", "o")) // 7
	fmt.Println(strings.LastIndex("hello world", "z")) // -1

	// Replace
	fmt.Println(strings.Replace("hello world", "o", "a", -1)) // hella warld
	fmt.Println(strings.Replace("hello world", "o", "a", 1))  // hella world

	// Split
	fmt.Println(strings.Split("hello,world", ",")) // [hello world]

	// Join
	fmt.Println(strings.Join([]string{"hello", "world"}, " ")) // hello world

	// ToLower
	fmt.Println(strings.ToLower("Hello World")) // hello world

	// ToUpper
	fmt.Println(strings.ToUpper("Hello World")) // HELLO WORLD

	// TrimSpace
	fmt.Println(strings.TrimSpace("   hello world   ")) // hello world

	// Trim
	fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡")) // Hello, Gophers
}
