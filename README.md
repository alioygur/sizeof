# sizeof

calculate memory usage base on reflect

## install

```
go get github.com/3vilive/sizeof
```

## usage

**integer**:

```go
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

// output:
// size of i8: 1 bytes
// size of i16: 2 bytes
// size of i32: 4 bytes
// size of i64: 8 bytes
```

**string**:

```go
var str1 = "hello"
fmt.Printf("size of `%s`: %d bytes\n", str1, sizeof.SizeOf(str1))
fmt.Printf("16 (size of string underlying struct) + %d (size of content)\n", len(str1) * 8)

// output:
// size of `str1`: 56 bytes
// 16 (size of string underlying struct) + 40 (size of content)
```

**pointer**:

```go
var strPtr1 = &str1
fmt.Printf("size of `strPtr1`: %d bytes\n", sizeof.SizeOf(strPtr1))
fmt.Printf("8 (size of pointer) + 56 (size of string)\n")

// output:
// size of `strPtr1`: 64 bytes
// 8 (size of pointer) + 56 (size of string)
```

**array**:

```go
var arr = [3]int64{1, 2, 3}
fmt.Printf("size of `arr`: %d bytes\n", sizeof.SizeOf(arr))

// output:
// size of `arr`: 24 bytes
```

**slice**:

```go
var (
    slice1 []int64 = nil
    slice2         = []int64{}
    slice3         = []int64{1, 2, 3}
)
fmt.Printf("size of `slice1`: %d bytes\n", sizeof.SizeOf(slice1))
fmt.Printf("size of `slice2`: %d bytes\n", sizeof.SizeOf(slice2))
fmt.Printf("size of `slice3`: %d bytes\n", sizeof.SizeOf(slice3))

// output:
// size of `slice1`: 0 bytes
// size of `slice2`: 24 bytes
// size of `slice3`: 48 bytes
```

**map**:

```go
var (
    map1 map[int64]int64 = nil
    map2                 = map[int64]int64{}
    map3                 = map[int64]int64{1: 1, 2: 2}
)
fmt.Printf("size of `map1`: %d bytes\n", sizeof.SizeOf(map1))
fmt.Printf("size of `map2`: %d bytes\n", sizeof.SizeOf(map2))
fmt.Printf("size of `map3`: %d bytes\n", sizeof.SizeOf(map3))

// output:
// size of `map1`: 0 bytes
// size of `map2`: 8 bytes
// size of `map3`: 40 bytes
```

**struct**:

```go
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

// output:
// size of `Demo1`: 1 bytes
// size of `Demo2`: 16 bytes
// size of `Demo3`: 16 bytes
// size of `Demo3 with values`: 33 bytes
// size of `Demo4 with values`: 281 bytes
```