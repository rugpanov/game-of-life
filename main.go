package main

import (
	"math/rand"
	"time"
)
import "fmt"

func main() {
	var fieldSize int
	var generationNumber int

	fmt.Print("Provide fields size: ")
	_, err := fmt.Scan(&fieldSize)
	if err != nil {
		return
	}

	fmt.Print("Provide number of generation: ")
	_, err = fmt.Scan(&generationNumber)
	if err != nil {
		return
	}

	field := generateField(fieldSize)
	for gi := 1; gi <= generationNumber; gi++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Generation #%d\n", gi)
		field = calcNextGeneration(field)
		alive := countAliveF(field)
		fmt.Printf("Alive: %d\n", alive)
		printField(field)
	}
}

func calcNextGeneration(field [][]int) [][]int {
	n := len(field)
	next := getEmptyArray(n, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			alive := countAliveN(field, i, j)
			if alive < 2 {
				next[i][j] = 0
			} else if alive == 2 {
				next[i][j] = field[i][j]
			} else if alive == 3 {
				next[i][j] = 1
			} else {
				next[i][j] = 0
			}
		}
	}

	return next
}

func countAliveF(field [][]int) int {
	n := len(field)
	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if field[i][j] == 1 {
				result++
			}
		}
	}
	return result
}

func countAliveN(field [][]int, i0 int, j0 int) int {
	counter := 0
	for i := i0 - 1; i <= i0+1; i++ {
		for j := j0 - 1; j <= j0+1; j++ {
			if i == i0 && j == j0 {
				continue
			}
			counter += checkPoint(field, i, j)
		}
	}
	return counter
}

func checkPoint(field [][]int, i0 int, j0 int) int {
	n := len(field)
	i := i0
	j := j0
	if i0 == -1 {
		i = n - 1
	}
	if i0 == n {
		i = 0
	}
	if j0 == -1 {
		j = n - 1
	}
	if j0 == n {
		j = 0
	}

	return field[i][j]
}

func generateField(n int) [][]int {
	field := getEmptyArray(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			field[i][j] = rand.Intn(2)
		}
	}
	return field
}

func getEmptyArray(rows, cols int) (array [][]int) {
	array = make([][]int, rows)
	for i := range array {
		array[i] = make([]int, cols)
	}
	return
}

func printField(field [][]int) {
	n := len(field)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if field[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("O")
			}

		}
		fmt.Println()
	}
}
