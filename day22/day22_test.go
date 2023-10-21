package day22_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day22"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{58248, 0}, day22.Run("day22"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day22.Run("day22")
	}
}
