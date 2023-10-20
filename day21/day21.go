package day21

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

type Solver struct {
	Data         map[string]int
	Instructions map[string]string
}

func (s *Solver) Solve(instruction string) int {
	if num, ok := s.Data[instruction]; ok {
		return num
	}

	parts := strings.Split(s.Instructions[instruction], " ")
	var num1, num2 int

	if num, ok := s.Data[parts[0]]; ok {
		num1 = num
	} else {
		num1 = s.Solve(parts[0])
	}

	if num, ok := s.Data[parts[2]]; ok {
		num2 = num
	} else {
		num2 = s.Solve(parts[2])
	}

	switch parts[1] {
	case "+":
		return num1 + num2

	case "-":
		return num1 - num2

	case "/":
		return num1 / num2

	case "*":
		return num1 * num2

	}

	return -418
}

func (s *Solver) FindHuman(instruction string) bool {
	if instruction == "humn" {
		return true
	}

	if _, ok := s.Data[instruction]; ok {
		return false
	}

	parts := strings.Split(s.Instructions[instruction], " ")
	return s.FindHuman(parts[0]) || s.FindHuman(parts[2])
}

func (s *Solver) SolveReverse(value int, instruction string) int {
	if instruction == "humn" {
		return value
	}

	var unkown int
	var hum string

	parts := strings.Split(s.Instructions[instruction], " ")

	if s.FindHuman(parts[0]) {
		answer := s.Solve(parts[2])
		hum = parts[0]

		// A is unkown
		switch parts[1] {
		case "+":
			unkown = value - answer

		case "-":
			unkown = answer + value

		case "/":
			unkown = answer * value

		case "*":
			unkown = value / answer
		}
	} else {
		answer := s.Solve(parts[0])
		hum = parts[2]

		// B is unkown
		switch parts[1] {
		case "+":
			unkown = value - answer

		case "-":
			unkown = answer - value

		case "/":
			unkown = answer / value

		case "*":
			unkown = value / answer
		}
	}

	return s.SolveReverse(unkown, hum)
}

func (s *Solver) SolveHuman() int {
	parts := strings.Split(s.Instructions["root"], " ")
	var answer int
	var hum string

	if s.FindHuman(parts[0]) {
		answer = s.Solve(parts[2])
		hum = parts[0]
	} else {
		answer = s.Solve(parts[0])
		hum = parts[2]
	}

	return s.SolveReverse(answer, hum)
}

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)
	solver := parse(data)

	return [2]interface{}{
		solver.Solve("root"),
		solver.SolveHuman(),
	}
}

func parse(data []string) Solver {
	solver := Solver{
		Data:         make(map[string]int),
		Instructions: make(map[string]string),
	}

	for _, row := range data {
		parts := strings.Split(row, ": ")
		if strings.ContainsAny(parts[1], "+-/*") {
			solver.Instructions[parts[0]] = parts[1]
		} else {
			num, _ := strconv.Atoi(parts[1])
			solver.Data[parts[0]] = num
		}
	}

	return solver
}
