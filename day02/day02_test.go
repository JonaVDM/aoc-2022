package day02_test

import (
	"testing"

	"github.com/jonavdm/aoc-2022/day02"
	_ "github.com/jonavdm/aoc-2022/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, day02.Run(), [2]int{11603, 12725})
}
