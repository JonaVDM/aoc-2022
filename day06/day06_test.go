package day06_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day06"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{1282, 3513}, day06.Run("day06"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day06.Run("day06")
	}
}
