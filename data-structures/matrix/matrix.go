package stack

import (
	"fmt"
)

func spiralTraverse(m [][]string) {
	rowSize := len(m[0])
	columnSize := len(m)
	visited := 0
	dir := 0
	row := 0
	column := 0
	for visited < rowSize * columnSize {
		fmt.Printf("%s\n", m[row][column])
		visited++
		switch dir%4 {
		case 0: //Top path
			if column < rowSize - row - 1 {
				column++ //Go right
			} else {
				row++ //Turn down
				dir++
			}
		case 1: //Right path
			if row < column {
				row++ //Go down
			} else {
				column-- //Turn left
				dir++
			}
		case 2: //Bottom path
			if column > rowSize - row - 1 {
				column-- //Go left
			} else {
				row-- //Turn up
				dir++
			}
		case 3: //Left path
			if row > column + 1 {
				row-- //Go up
			} else {
				column++ //Turn right
				dir++
			}
		}
	}
}