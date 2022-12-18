package day18_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day18"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{4244, 0}, day18.Run("day18"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day18.Run("day18")
	}
}
