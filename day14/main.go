package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
	"strconv"
	"strings"
)

const (
	mapWidth, mapHeight = 101, 103
)

type Robot struct {
	position aoc.Point
	velocity aoc.Point
}

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")
	robots := parseRobots(input)

	// Part 1
	for range 100 {
		moveRobots(robots)
	}
	q0, q1, q2, q3 := calculateQuads(robots)
	fmt.Println(q0 * q1 * q2 * q3)

	// Part 2 -> Tedious process of glazing at the terminal and waiting for the christmas tree to appear (7572 urghh...)
	robots = parseRobots(input)
	for i := 1; i < 10000; i++ {
		for k, robot := range robots {
			robots[k] = robot.move()
		}
		printRobots(robots)
		fmt.Println(i)
		fmt.Println()
	}
}

func calculateQuads(robots []Robot) (int, int, int, int) {
	q0, q1, q2, q3 := 0, 0, 0, 0
	for _, robot := range robots {
		borderX, borderY := mapWidth/2, mapHeight/2
		if robot.position.X < borderX && robot.position.Y < borderY {
			q0++
		} else if robot.position.X < borderX && robot.position.Y > borderY {
			q1++
		} else if robot.position.X > borderX && robot.position.Y < borderY {
			q2++
		} else if robot.position.X > borderX && robot.position.Y > borderY {
			q3++
		}
	}
	return q0, q1, q2, q3
}

func moveRobots(robots []Robot) {
	for k, robot := range robots {
		robots[k] = robot.move()
	}
}

func parseRobots(input []string) []Robot {
	var robots []Robot
	for _, line := range input {
		components := strings.Split(line, " ")

		ps := strings.Split(strings.Split(components[0], "=")[1], ",")
		px, _ := strconv.Atoi(ps[0])
		py, _ := strconv.Atoi(ps[1])

		vs := strings.Split(strings.Split(components[1], "=")[1], ",")
		vx, _ := strconv.Atoi(vs[0])
		vy, _ := strconv.Atoi(vs[1])

		robots = append(robots, Robot{position: aoc.Point{X: px, Y: py}, velocity: aoc.Point{X: vx, Y: vy}})
	}
	return robots
}

func printRobots(robots []Robot) {
	robotMap := make(map[aoc.Point]int)
	for _, robot := range robots {
		robotMap[robot.position] += 1
	}

	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			robotCount := robotMap[aoc.Point{X: x, Y: y}]
			if robotCount == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(robotCount)
			}
		}
		fmt.Println()
	}
}

func (r Robot) move() Robot {
	r.position = r.position.Add(r.velocity)

	if r.position.X < 0 {
		r.position.X += mapWidth
	}

	if r.position.Y < 0 {
		r.position.Y += mapHeight
	}

	if r.position.X >= mapWidth {
		r.position.X = r.position.X - mapWidth
	}

	if r.position.Y >= mapHeight {
		r.position.Y = r.position.Y - mapHeight
	}

	return r
}
