package main

import (
	"fmt"
	"testing"
)

const MessageRequireDebug = "DEBUG must be set to true"
const MessageForbidDebug = "DEBUG must be set to false"

func BenchmarkSetOnceInt_Set(b *testing.B) {
	if DEBUG {
		panic(MessageForbidDebug)
	}

	for i := 0; i < b.N; i++ {
		foo := NewSetOnceIntBare()
		foo.Set(i)
	}
}

func BenchmarkPrimitiveInt_Set(b *testing.B) {
	if DEBUG {
		panic(MessageForbidDebug)
	}

	for i := 0; i < b.N; i++ {
		var foo int
		foo = i
		_ = foo
	}
}

func ExampleNewSetOnceInt() {
	if !DEBUG {
		panic(MessageRequireDebug)
	}

	i := NewSetOnceInt(42)
	fmt.Println(i.Get())
	// Output:42
}

func ExampleNewSetOnceIntBare() {
	if !DEBUG {
		panic(MessageRequireDebug)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()

	i := NewSetOnceIntBare()
	fmt.Println(i.Get())
	// Output: panic: Value has not been set
}

func ExampleSetOnceInt_Set() {
	if !DEBUG {
		panic(MessageRequireDebug)
	}

	i := NewSetOnceIntBare()
	i.Set(43)
	fmt.Println(i.Get())
	// Output: 43
}

func ExampleSetOnceInt_Set_setTwice() {
	if !DEBUG {
		panic(MessageRequireDebug)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()

	i := NewSetOnceInt(42)
	i.Set(43)
	// Output: panic: Value can only be set once
}
