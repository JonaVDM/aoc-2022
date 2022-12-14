package day07_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day07"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{1350966, 6296435}, day07.Run("day07"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day07.Run("day07")
	}
}
