package day20_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day20"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{1087, 13084440324666}, day20.Run("day20"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day20.Run("day20")
	}
}

