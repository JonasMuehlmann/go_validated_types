package main

import (
	"fmt"
	"testing"
)

const MessageRequireDebug = "DEBUG must be set to true"
const MessageForbidDebug = "DEBUG must be set to false"

func BenchmarkSetOnce_Set(b *testing.B) {
	if DEBUG {
		panic(MessageForbidDebug)
	}

	for i := 0; i < b.N; i++ {
		foo := NewSetOnceBare[int]()
		foo.Set(i)
	}
}

func BenchmarkPrimitiveNumber_Set(b *testing.B) {
	if DEBUG {
		panic(MessageForbidDebug)
	}

	for i := 0; i < b.N; i++ {
		var foo int
		foo = i
		_ = foo
	}
}

func ExampleNewSetOnce() {
	if !DEBUG {
		panic(MessageRequireDebug)
	}

	i := NewSetOnce(42)
	fmt.Println(i.Get())
	// Output:42
}

func ExampleNewSetOnceBare() {
	if !DEBUG {
		panic(MessageRequireDebug)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()

	i := NewSetOnceBare[int]()
	fmt.Println(i.Get())
	// Output: panic: Value has not been set
}

func ExampleSetOnce_Set() {
	if !DEBUG {
		panic(MessageRequireDebug)
	}

	i := NewSetOnceBare[int]()
	i.Set(43)
	fmt.Println(i.Get())
	// Output: 43
}

func ExampleSetOnce_Set_setTwice() {
	if !DEBUG {
		panic(MessageRequireDebug)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()

	i := NewSetOnce(42)
	i.Set(43)
	// Output: panic: Value can only be set once
}
