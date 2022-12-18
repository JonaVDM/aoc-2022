package day17_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day17"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{3102, 0}, day17.Run("day17"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day17.Run("day17")
	}
}
