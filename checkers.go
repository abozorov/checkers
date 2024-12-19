package main

import (
	"fmt"
)

type Board struct {
	fields [8][8]int
	pawn [25]Pawn
	
}

type Pawn struct {
	damka bool
	// isWhite bool
}

func initBoard(board *Board) {

	board.fields = [8][8]int{
		{1, -2, 1, -2, 1, -2, 1, -2},
		{-2, 1, -2, 1, -2, 1, -2, 1},
		{1, -2, 1, -2, 1, -2, 1, -2},
		{-1, 1, -1, 1, -1, 1, -1, 1},
		{1, -1, 1, -1, 1, -1, 1, -1},
		{2, 1, 2, 1, 2, 1, 2, 1},
		{1, 2, 1, 2, 1, 2, 1, 2},
		{2, 1, 2, 1, 2, 1, 2, 1}}

		// -1 -черная клетка
		// -2 -черная шашка
		// -3 -черная дамка
		// 1 -белая клетка
		// 2 -белая шашка
		// 3 -белая дамка
}

func showBoard(board *Board) {	
	// [id, "■"]
	// ■□●○◎◉
	// ⬛⬜⚫⚪⦾⦿🏳️🏴

	mp := map[int]string{
		1: "⬜",
		-1: "░ ",
		-2: "⚫",
		2: "⚪",
		-3: "🏴",
		3: "🏳️",
	}


	for i := -1; i <= 8; i++ {
		
		if i == -1 || i == 8 {
			for j := 'A'; j <= 'H'; j++ {
				fmt.Print(" " + string(j))
			}
			fmt.Println()
			continue
		}

		for j := -1; j <= 8; j++ {

			if j == -1 || j == 8 {
				fmt.Print(8-i)
				continue
			}

			fmt.Print(mp[board.fields[i][j]])
		}
		fmt.Println()
	}
}

func checkStep(board *Board, s1 string, s2 string, count int) bool {
	return true
}

func move(board *Board, s1 string, s2 string) {

	mp := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,

		"1": 7,
		"2": 6,
		"3": 5,
		"4": 4,
		"5": 3,
		"6": 2,
		"7": 1,
		"8": 0,
	}

	x := board.fields[ mp[ string(s2[1]) ] ][ mp[ string(s2[0]) ] ]
	board.fields[ mp[ string(s2[1]) ] ][ mp[ string(s2[0]) ] ] = board.fields[ mp[ string(s1[1]) ] ][ mp[ string(s1[0]) ] ]
	board.fields[ mp[ string(s1[1]) ] ][ mp[ string(s1[0]) ] ] = x
}

func startGame(board *Board) {
	count := 1

	for {
		if count % 2 != 0 {
			fmt.Print("Ходят БЕЛЫЕ!\n")
		} else {
			fmt.Print("Ходят НИГГЕРЫ!\n")
		}

		var s1 string
		fmt.Scan(&s1)

		if s1 == "gg" {

			if count % 2 == 0 {
				fmt.Print("Выиграли БЕЛЫЕ!\n")
			} else {
				fmt.Print("Bыиграли НИГГЕРЫ!\n")
			}
			return
		} else if s1 == "draw" {
			var s2 string

			fmt.Print("Пусть второй игрок напечатает слово \"yes\" или \"да\"\n")
			fmt.Scan(&s2)
			
			if s2 == "yes" || s2 == "да" {
				fmt.Print("ВЫИГРАЛА ДРУЖБА!\n")
				return
			} else {
				fmt.Print("Неверная команда, пожалуйста повторите ход!\n")
				continue
			}
		} else {
			var s2 string
			fmt.Scan(&s2)

			if checkStep(board, s1, s2, count) {
				move(board, s1, s2)		
				count++
				showBoard(board)
			} else {
				fmt.Print("Неверная команда, пожалуйста повторите ход!\n")
				continue
			}
		}
	}
}

func main() {
	var board Board
	initBoard(&board)
	showBoard(&board)

	startGame(&board)
}

// ■□●○◎◉
// ⬛⬜⚫⚪⦾⦿