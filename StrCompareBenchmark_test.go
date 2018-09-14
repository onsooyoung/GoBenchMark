package main

import "testing"


const STR_NUM = "1"

const STR_ONE_STR = "A"

const STR_PROTO_STR = "PROTOCAL1"

const STR_SHA256_HASH = "8f2af31397ffcb527306ccc26b05466ffefdf0132b04cc516ad9d222117c27d5"

func benchmarkNum(TestComapare func(string, string) bool, b *testing.B) {
	for n := 0; n < b.N; n++ {
		TestComapare(STR_NUM, STR_NUM)
	}
}

func benchmarkOneStr(TestComapare func(string, string) bool, b *testing.B) {
	for n := 0; n < b.N; n++ {
		TestComapare(STR_ONE_STR, STR_ONE_STR)
	}
}

func benchmarkProtoStr(TestComapare func(string, string) bool, b *testing.B) {
	for n := 0; n < b.N; n++ {
		TestComapare(STR_PROTO_STR, STR_PROTO_STR)
	}
}

func benchmarkSha256Str(TestComapare func(string, string) bool, b *testing.B) {
	for n := 0; n < b.N; n++ {
		TestComapare(STR_SHA256_HASH, STR_SHA256_HASH)
	}
}

func BenchmarkNum1(b *testing.B) { benchmarkNum( StrCompare, b) }
func BenchmarkNum2(b *testing.B) { benchmarkNum( StrEqual, b) }
func BenchmarkNum3(b *testing.B) { benchmarkNum( StrHasPrefix, b) }
func BenchmarkNum4(b *testing.B) { benchmarkNum( StrIndex, b) }
func BenchmarkNum6(b *testing.B) { benchmarkNum( StrContains, b) }
func BenchmarkNum7(b *testing.B) { benchmarkNum( StrContains, b) }
func BenchmarkNum8(b *testing.B) { benchmarkNum( StrCount, b) }
func BenchmarkNum5(b *testing.B) { benchmarkNum( StrEqualFold, b) }

func BenchmarkOneStr1(b *testing.B) { benchmarkOneStr( StrCompare, b) }
func BenchmarkOneStr2(b *testing.B) { benchmarkOneStr( StrEqual, b) }
func BenchmarkOneStr3(b *testing.B) { benchmarkOneStr( StrHasPrefix, b) }
func BenchmarkOneStr4(b *testing.B) { benchmarkOneStr( StrIndex, b) }
func BenchmarkOneStr5(b *testing.B) { benchmarkOneStr( StrContains, b) }
func BenchmarkOneStr6(b *testing.B) { benchmarkOneStr( StrContains, b) }
func BenchmarkOneStr7(b *testing.B) { benchmarkOneStr( StrCount, b) }
func BenchmarkOneStr8(b *testing.B) { benchmarkOneStr( StrEqualFold, b) }

func BenchmarkProtoStr1(b *testing.B) { benchmarkProtoStr( StrCompare, b) }
func BenchmarkProtoStr2(b *testing.B) { benchmarkProtoStr( StrEqual, b) }
func BenchmarkProtoStr3(b *testing.B) { benchmarkProtoStr( StrHasPrefix, b) }
func BenchmarkProtoStr4(b *testing.B) { benchmarkProtoStr( StrIndex, b) }
func BenchmarkProtoStr5(b *testing.B) { benchmarkProtoStr( StrContains, b) }
func BenchmarkProtoStr6(b *testing.B) { benchmarkProtoStr( StrContains, b) }
func BenchmarkProtoStr7(b *testing.B) { benchmarkProtoStr( StrCount, b) }
func BenchmarkProtoStr8(b *testing.B) { benchmarkProtoStr( StrEqualFold, b) }

func BenchmarkSha256Str1(b *testing.B) { benchmarkSha256Str( StrCompare, b) }
func BenchmarkSha256Str2(b *testing.B) { benchmarkSha256Str( StrEqual, b) }
func BenchmarkSha256Str3(b *testing.B) { benchmarkSha256Str( StrHasPrefix, b) }
func BenchmarkSha256Str4(b *testing.B) { benchmarkSha256Str( StrIndex, b) }
func BenchmarkSha256Str5(b *testing.B) { benchmarkSha256Str( StrContains, b) }
func BenchmarkSha256Str6(b *testing.B) { benchmarkSha256Str( StrContains, b) }
func BenchmarkSha256Str7(b *testing.B) { benchmarkSha256Str( StrCount, b) }
func BenchmarkSha256Str8(b *testing.B) { benchmarkSha256Str( StrEqualFold, b) }


//go test -bench=.


