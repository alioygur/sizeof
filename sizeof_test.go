package sizeof

import "testing"

func TestSizeOf(t *testing.T) {
	// base types
	var (
		i8  int8
		i16 int16
		i32 int32
		i64 int64
	)

	if got := SizeOf(i8); got != 1 {
		t.Errorf("expect `%d` but got `%d`", 1, got)
	}
	if got := SizeOf(i16); got != 2 {
		t.Errorf("expect `%d` but got `%d`", 1, got)
	}
	if got := SizeOf(i32); got != 4 {
		t.Errorf("expect `%d` but got `%d`", 1, got)
	}
	if got := SizeOf(i64); got != 8 {
		t.Errorf("expect `%d` but got `%d`", 1, got)
	}

	// strings
	var (
		str1             = ""
		sizeOfStr1 int64 = 16 + 0
		str2             = "hello"
		sizeOfStr2 int64 = 16 + 5*8
		str3             = "hello,world"
		sizeOfStr3 int64 = 16 + 11*8
	)

	if got := SizeOf(str1); got != sizeOfStr1 {
		t.Errorf("expect `%d` but got `%d`", sizeOfStr1, got)
	}
	if got := SizeOf(str2); got != sizeOfStr2 {
		t.Errorf("expect `%d` but got `%d`", sizeOfStr2, got)
	}
	if got := SizeOf(str3); got != sizeOfStr3 {
		t.Errorf("expect `%d` but got `%d`", sizeOfStr3, got)
	}
}
