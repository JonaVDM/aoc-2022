package day03_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day03"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{8493, 2552}, day03.Run("day03"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day03.Run("day03")
	}
}
