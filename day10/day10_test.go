package day10_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day10"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	output := `
####..##..###...##....##.####...##.####.
...#.#..#.#..#.#..#....#.#.......#....#.
..#..#....###..#..#....#.###.....#...#..
.#...#....#..#.####....#.#.......#..#...
#....#..#.#..#.#..#.#..#.#....#..#.#....
####..##..###..#..#..##..#.....##..####.`

	assert.Equal(t, [2]interface{}{17940, output}, day10.Run("day10"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day10.Run("day10")
	}
}
