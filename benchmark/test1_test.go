package benchmark

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkReturnStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := returnStruct()
		assert.Equal(b, "a", x.a)
		b.Log("")
	}
}

func BenchmarkReturnStructPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := returnStructPointer()
		assert.Equal(b, "a", x.a)
		b.Log("")
	}
}
