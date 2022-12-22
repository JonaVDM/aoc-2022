package day19_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day19"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{0, 0}, day19.Run("day19"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day19.Run("day19")
	}
}
