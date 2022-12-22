package day19

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

type Blueprint struct {
	OreBot           int
	ClayBot          int
	ObsidianBotOre   int
	ObsidianBotClay  int
	GeodeBotOre      int
	GeodeBotObsidian int
}

type Inventory struct {
	Ore      int
	Clay     int
	Obsidian int
	Geode    int

	OreBot      int
	ClayBot     int
	ObsidianBot int
	GeodeBot    int

	Itteration int
}

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	blueprints := parse(data)

	score := 0

	fmt.Println()
	for i, b := range blueprints {
		out := play(b, Inventory{OreBot: 1}, 0)
		fmt.Printf("out blueprint %d: %v\n", i+1, out)
		score += (i + 1) * out
	}

	return [2]interface{}{
		score,
		0,
	}
}

func parse(data []string) []Blueprint {
	blueprints := make([]Blueprint, 0)

	for _, row := range data {
		spl := strings.Split(row, " ")

		oreBot, err := strconv.Atoi(spl[6])
		if err != nil {
			panic(err)
		}

		clayBot, err := strconv.Atoi(spl[12])
		if err != nil {
			panic(err)
		}

		obsidianBotOre, err := strconv.Atoi(spl[18])
		if err != nil {
			panic(err)
		}

		obsidianBotClay, err := strconv.Atoi(spl[21])
		if err != nil {
			panic(err)
		}

		geodeBotOre, err := strconv.Atoi(spl[27])
		if err != nil {
			panic(err)
		}

		geodeBotObsidian, err := strconv.Atoi(spl[30])
		if err != nil {
			panic(err)
		}

		blueprints = append(blueprints, Blueprint{
			OreBot:           oreBot,
			ClayBot:          clayBot,
			ObsidianBotOre:   obsidianBotOre,
			ObsidianBotClay:  obsidianBotClay,
			GeodeBotOre:      geodeBotOre,
			GeodeBotObsidian: geodeBotObsidian,
		})
	}

	return blueprints
}

func hash(i Inventory) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(i)
	return b.Bytes()
}

func play(blueprint Blueprint, inventory Inventory, itteration int) int {
	queue := make([]Inventory, 0)
	queue = append(queue, Inventory{OreBot: 1})
	history := make([][]byte, 0)

	max := 0

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		// if utils.Contains[[]byte](history, hash(item)) {
		// 	continue
		// }

		history = append(history, hash(item))

		if item.Itteration >= 24 {
			continue
		}

		// start building
		affortOre := item.Ore >= blueprint.OreBot
		affortClay := item.Ore >= blueprint.ClayBot
		affortObsidian := item.Ore >= blueprint.ObsidianBotOre && item.Clay >= blueprint.ObsidianBotClay
		affortGeode := item.Ore >= blueprint.GeodeBotOre && item.Obsidian >= blueprint.GeodeBotObsidian

		item.Ore += item.OreBot
		item.Clay += item.ClayBot
		item.Obsidian += item.ObsidianBot
		item.Geode += item.GeodeBot

		if item.Geode > max {
			max = item.Geode
		}

		// Do nothing
		{
			cop := item
			cop.Itteration++
			queue = append(queue, cop)
		}

		// Can affort ore bot
		if affortOre {
			cop := item
			cop.OreBot += 1
			cop.Ore -= blueprint.OreBot
			cop.Itteration++
			queue = append(queue, cop)
		}

		// can affort clay bot
		if affortClay {
			cop := item
			cop.ClayBot += 1
			cop.Ore -= blueprint.ClayBot
			cop.Itteration++
			queue = append(queue, cop)
		}

		// can affort obsidian bot
		if affortObsidian {
			cop := item
			cop.ObsidianBot += 1
			cop.Ore -= blueprint.ObsidianBotOre
			cop.Clay -= blueprint.ObsidianBotClay
			cop.Itteration++
			queue = append(queue, cop)
		}

		// can affort geode bot
		if affortGeode {
			cop := item
			cop.GeodeBot += 1
			cop.Ore -= blueprint.GeodeBotOre
			cop.Obsidian -= blueprint.GeodeBotObsidian
			cop.Itteration++
			queue = append(queue, cop)
		}
	}

	return max

	// branches := make([]int, 0)

	// if itteration == 24 {
	// 	return inventory.Geode
	// }

	// affortOre := inventory.Ore == blueprint.OreBot
	// affortClay := inventory.Ore == blueprint.ClayBot
	// affortObsidian := inventory.Ore >= blueprint.ObsidianBotOre && inventory.Clay >= blueprint.ObsidianBotClay
	// affortGeode := inventory.Ore >= blueprint.GeodeBotOre && inventory.Obsidian >= blueprint.GeodeBotObsidian

	// inventory.Ore += inventory.OreBot
	// inventory.Clay += inventory.ClayBot
	// inventory.Obsidian += inventory.ObsidianBot
	// inventory.Geode += inventory.GeodeBot

	// // Do nothing
	// {
	// 	cop := inventory
	// 	branches = append(branches, play(blueprint, cop, itteration+1))
	// }

	// // Can affort ore bot
	// if affortOre {
	// 	cop := inventory
	// 	cop.OreBot += 1
	// 	cop.Ore -= blueprint.OreBot
	// 	branches = append(branches, play(blueprint, cop, itteration+1))
	// }

	// // can affort clay bot
	// if affortClay {
	// 	cop := inventory
	// 	cop.ClayBot += 1
	// 	cop.Ore -= blueprint.ClayBot
	// 	branches = append(branches, play(blueprint, cop, itteration+1))
	// }

	// // can affort obsidian bot
	// if affortObsidian {
	// 	cop := inventory
	// 	cop.ObsidianBot += 1
	// 	cop.Ore -= blueprint.ObsidianBotOre
	// 	cop.Clay -= blueprint.ObsidianBotClay
	// 	branches = append(branches, play(blueprint, cop, itteration+1))
	// }

	// // can affort geode bot
	// if affortGeode {
	// 	cop := inventory
	// 	cop.GeodeBot += 1
	// 	cop.Ore -= blueprint.GeodeBotOre
	// 	cop.Obsidian -= blueprint.GeodeBotObsidian
	// 	branches = append(branches, play(blueprint, cop, itteration+1))
	// }

	// return utils.Max(branches)
}
