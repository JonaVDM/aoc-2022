package day02_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day02"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]int{11603, 12725}, day02.Run())
}

func TestRun2(t *testing.T) {
	assert.Equal(t, [2]int{11603, 12725}, day02.Run2())
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day02.Run()
	}
}

func BenchmarkRun2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day02.Run2()
	}
}
