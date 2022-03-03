package main

import (
	"flood_fill/functions"
	"flood_fill/models"
	"fmt"
	"log"
)

func main() {
	var board models.Board
	points := 0
	err := functions.StartBoard(&board, 5, 4)

	if err != nil {
		log.Panicln("Error on start board")
	}
	functions.PrintBoard(board)

	for !functions.CompleteFill(board) {
		fmt.Println("Change color:")
		upperInput := functions.ReadChar()
		if functions.VerifyIfIsInRange(upperInput, board.NumLen) {
			functions.FloodFill(&board, upperInput)
		} else {
			fmt.Println("Invalid letter! Try again!")
		}
		points++
	}
	fmt.Printf("You complete the maze! Good Job!\nYour pontuaition: %d\n", points)
}
