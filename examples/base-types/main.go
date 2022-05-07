package main

import (
	"fmt"

	"github.com/3vilive/sizeof"
)

func main() {
	var (
		i8  int8
		i16 int16
		i32 int32
		i64 int64
	)

	fmt.Printf("size of i8: %d bytes\n", sizeof.SizeOf(i8))
	fmt.Printf("size of i16: %d bytes\n", sizeof.SizeOf(i16))
	fmt.Printf("size of i32: %d bytes\n", sizeof.SizeOf(i32))
	fmt.Printf("size of i64: %d bytes\n", sizeof.SizeOf(i64))

	var (
		str1 = "hello"
	)
	fmt.Printf("size of `str1`: %d bytes\n", sizeof.SizeOf(str1))
	fmt.Printf("16 (size of string underlying struct) + %d (size of content)\n", len(str1)*8)

	var (
		strPtr1 = &str1
	)
	fmt.Printf("size of `strPtr1`: %d bytes\n", sizeof.SizeOf(strPtr1))
	fmt.Printf("8 (size of pointer) + 56 (size of string)\n")

	var arr = [3]int64{1, 2, 3}
	fmt.Printf("size of `arr`: %d bytes\n", sizeof.SizeOf(arr))

	var (
		slice1 []int64 = nil
		slice2         = []int64{}
		slice3         = []int64{1, 2, 3}
	)
	fmt.Printf("size of `slice1`: %d bytes\n", sizeof.SizeOf(slice1))
	fmt.Printf("size of `slice2`: %d bytes\n", sizeof.SizeOf(slice2))
	fmt.Printf("size of `slice3`: %d bytes\n", sizeof.SizeOf(slice3))

	var (
		map1 map[int64]int64 = nil
		map2                 = map[int64]int64{}
		map3                 = map[int64]int64{1: 1, 2: 2}
	)
	fmt.Printf("size of `map1`: %d bytes\n", sizeof.SizeOf(map1))
	fmt.Printf("size of `map2`: %d bytes\n", sizeof.SizeOf(map2))
	fmt.Printf("size of `map3`: %d bytes\n", sizeof.SizeOf(map3))

	type Demo1 struct {
		a int8
	}

	type Demo2 struct {
		a int8
		b int64
		// align: 7
	}

	type Demo3 struct {
		a *Demo1
		b *Demo2
	}

	type Demo4 struct {
		Demo3
		username string
		posts    []string
	}

	fmt.Printf("size of `Demo1`: %d bytes\n", sizeof.SizeOf(Demo1{}))
	fmt.Printf("size of `Demo2`: %d bytes\n", sizeof.SizeOf(Demo2{}))
	fmt.Printf("size of `Demo3`: %d bytes\n", sizeof.SizeOf(Demo3{}))
	fmt.Printf("size of `Demo3 with values`: %d bytes\n", sizeof.SizeOf(Demo3{
		a: &Demo1{},
		b: &Demo2{},
	}))
	fmt.Printf("size of `Demo4 with values`: %d bytes\n", sizeof.SizeOf(Demo4{
		Demo3: Demo3{
			a: &Demo1{},
			b: &Demo2{},
		},
		username: "3vilive",
		posts:    []string{"sizeof is awesome"},
	}))
}
