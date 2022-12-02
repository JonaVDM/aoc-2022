package day01_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day01"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert.Equal(t, day01.Run(), [2]int{67658, 200158})
}
