package day21_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day21"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{121868120894282, 3582317956029}, day21.Run("day21"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day21.Run("day21")
	}
}
