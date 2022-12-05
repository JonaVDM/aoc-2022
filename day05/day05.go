package day05

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run() [2]interface{} {
	data := utils.ReadFileRaw("day05")

	tmpStorage := make(map[int][]rune)
	storageA := make(map[int][]rune)
	storageB := make(map[int][]rune)
	startAt := 0

	for i, v := range data {
		if strings.Contains(v, "1") {
			startAt = i
			break
		}

		for p, l := range v {
			if l != ' ' && l != '[' && l != ']' {
				if tmpStorage[p] == nil {
					tmpStorage[p] = make([]rune, 0)
				}

				tmpStorage[p] = append([]rune{l}, tmpStorage[p]...)
			}
		}
	}

	// why do simple maths if you can get the data expensively from the next row
	for k, v := range tmpStorage {
		nk, err := strconv.Atoi(string(data[startAt][k]))
		if err != nil {
			panic(err)
		}
		nv := make([]rune, len(v))
		copy(nv, v)
		storageA[nk] = v
		storageB[nk] = nv
	}

	for i := startAt + 2; i < len(data); i++ {
		if data[i] == "" {
			break
		}

		inst := strings.Split(data[i], " ")
		amount, err := strconv.Atoi(inst[1])
		if err != nil {
			panic(err)
		}

		src, err := strconv.Atoi(inst[3])
		if err != nil {
			panic(err)
		}

		dest, err := strconv.Atoi(inst[5])
		if err != nil {
			panic(err)
		}

		for x := 0; x < amount; x++ {
			l := len(storageA[src]) - 1
			storageA[dest] = append(storageA[dest], storageA[src][l])
			storageA[src] = storageA[src][:l]
		}

		l := len(storageB[src]) - amount
		storageB[dest] = append(storageB[dest], storageB[src][l:]...)
		storageB[src] = storageB[src][:l]
	}

	strA := ""
	strB := ""
	for i := 1; i < len(storageA)+1; i++ {
		strA += string(storageA[i][len(storageA[i])-1])
		strB += string(storageB[i][len(storageB[i])-1])
	}

	return [2]interface{}{
		strA,
		strB,
	}
}
