package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const (
	str = "nguisfgrpgignfhnfgihnrkgiugisuygisdhfdfr"
	cnt = 10000
)

func BenchmarkPlusConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ss := ""
		for i := 0; i < cnt; i++ {
			ss += str
		}
		_ = ss
	}
}

func BenchmarkSprintfConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ss := ""
		for i := 0; i < cnt; i++ {
			ss = fmt.Sprintf("%s%s", ss, str)
		}
		_ = ss
	}
}

func BenchmarkBuilderConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for i := 0; i < cnt; i++ {
			builder.WriteString(str)
		}
		_ = builder.String()
	}
}

func BenchmarkBufferConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf = new(bytes.Buffer)
		for i := 0; i < cnt; i++ {
			buf.WriteString(str)
		}
		_ = buf.String()
	}
}

func BenchmarkAppendConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 0)
		for i := 0; i < cnt; i++ {
			buf = append(buf, str...)
		}
		_ = string(buf)
	}
}
