package day15_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day15"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{4582667, 10961118625406}, day15.Run("day15"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day15.Run("day15")
	}
}
