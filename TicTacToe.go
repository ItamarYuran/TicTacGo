package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	manageGame()
}

func manageGame() {

	board, player := initiate()
	printBoard(board)

	for q := 0; q < 9; {
		clearScreen()
		printBoard(board)
		fmt.Printf("Player %v, it is your turn!\n", player)
		row, column := getPlayerCell(board)
		fmt.Printf("row %v, column %v\n", row, column)
		if player == 1 {
			board[row-1][column-1] = "x"
			q++
		} else {
			board[row-1][column-1] = "o"
			q++
		}
		printBoard(board)

		gameover := check(board)
		if gameover != "-" {
			fmt.Printf("GAME OVER\n Player %v Won!\n", player)
			break
		}
		player = player%2 + 1

	}
	fmt.Printf("GAME OVER\n It's a tie!\n")

}

func getPlayerCell(board [3][3]string) (r, c int) {
	var row, column int

	for {
		fmt.Printf("Press 1 - 3 to choose row\n")
		_, err := fmt.Scanln(&row)
		if err != nil {
			fmt.Println("Input isn't valid capara")
			continue
		}
		fmt.Printf("Press 1 - 3 to choose column\n")
		_, err = fmt.Scanln(&column)
		if err != nil {
			fmt.Println("Input isn't valid capara")
			continue
		}
		if !(row >= 1 && row <= 3) || !(column >= 1 && column <= 3) {
			fmt.Println("Invalid input. Please enter an integer between 1 and 3.")
			continue
		}
		if board[row-1][column-1] != " " {
			fmt.Println("Cell is unavailable")
			continue
		}
		return row, column
	}
}

func printBoard(mat [3][3]string) {
	fmt.Printf("%v | %v | %v\n", mat[0][0], mat[0][1], mat[0][2])
	fmt.Printf("-- --- --\n")
	fmt.Printf("%v | %v | %v\n", mat[1][0], mat[1][1], mat[1][2])
	fmt.Printf("-- --- --\n")
	fmt.Printf("%v | %v | %v\n", mat[2][0], mat[2][1], mat[2][2])
}

func initiate() (mat [3][3]string, player int) {
	board := [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
	rand.Seed(time.Now().UnixNano())
	playeri := rand.Intn(2) + 1
	return board, playeri
}

func check(mat [3][3]string) string {
	if mat[0][0] == mat[0][1] && mat[0][1] == mat[0][2] && mat[0][1] != " " {
		return mat[0][1]
	}
	if mat[1][0] == mat[1][1] && mat[1][1] == mat[1][2] && mat[1][1] != " " {
		return mat[1][1]
	}
	if mat[2][0] == mat[2][1] && mat[2][1] == mat[2][2] && mat[2][1] != " " {
		return mat[2][1]
	}
	if mat[0][0] == mat[1][0] && mat[1][0] == mat[2][0] && mat[0][0] != " " {
		return mat[0][0]
	}
	if mat[0][1] == mat[1][1] && mat[1][1] == mat[2][1] && mat[0][1] != " " {
		return mat[0][1]
	}
	if mat[0][2] == mat[1][2] && mat[1][2] == mat[2][2] && mat[0][2] != " " {
		return mat[0][2]
	}
	if mat[0][0] == mat[1][1] && mat[1][1] == mat[2][2] && mat[0][0] != " " {
		return mat[0][0]
	}
	if mat[0][2] == mat[1][1] && mat[1][1] == mat[2][0] && mat[0][2] != " " {
		return mat[0][2]
	}
	return "-"
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
