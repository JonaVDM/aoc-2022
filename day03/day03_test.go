package day03_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day03"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]int{8493, 2552}, day03.Run())
}
