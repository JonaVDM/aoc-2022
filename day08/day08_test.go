package day08_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day08"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{1843, 180000}, day08.Run("day08"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day08.Run("day08")
	}
}
