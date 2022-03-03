package functions

import (
	"bufio"
	"errors"
	"flood_fill/models"
	"fmt"
	"github.com/gookit/color"
	"log"
	"math/rand"
	"os"
	"time"
	"unicode"
)

func StartBoard(board *models.Board, nColor int, nLen int) error {
	if nColor < 3 || nColor > 9 {
		return errors.New("invalid number of colors")
	}
	if nLen < 3 || nLen > 20 {
		return errors.New("invalid number in Len")
	}
	rand.Seed(time.Now().UnixNano())
	board.NumLen = nLen
	board.NumColor = nColor

	for i := 0; i < nLen; i++ {
		for j := 0; j < nLen; j++ {
			board.Matrix[i][j] = models.Letters[rand.Intn(nColor)]
		}
	}
	return nil
}

func PrintBoard(board models.Board) {
	for i := 0; i < board.NumLen; i++ {
		for j := 0; j < board.NumLen; j++ {
			s := color.S256(0, uint8(board.Matrix[i][j])%21)
			s.Printf(" %c ", board.Matrix[i][j])
		}
		fmt.Println()
	}
}

func RemoveFromArray(arr *models.Vector, a *int, b *int) {
	length := len(arr.Vec)
	*a = arr.Vec[length-2]
	*b = arr.Vec[length-1]
	arr.Vec = arr.Vec[:length-2]
}

func AppendVector(arr *models.Vector, a int, b int) {
	arr.Vec = append(arr.Vec, a)
	arr.Vec = append(arr.Vec, b)
}

func CheckNeighbors(board models.Board, arr *models.Vector, a int, b int, lastColor byte) {
	if board.Matrix[a][b+1] == lastColor {
		AppendVector(arr, a, b+1)
	}
	if b-1 > -1 && board.Matrix[a][b-1] == lastColor {
		AppendVector(arr, a, b-1)
	}
	if board.Matrix[a+1][b] == lastColor {
		AppendVector(arr, a+1, b)
	}
	if a-1 > -1 && board.Matrix[a-1][b] == lastColor {
		AppendVector(arr, a-1, b)
	}
}

func CompleteFill(board models.Board) bool {
	for i := 0; i < board.NumLen; i++ {
		for j := 0; j < board.NumLen; j++ {
			if board.Matrix[i][j] != board.Matrix[0][0] {
				return false
			}
		}
	}
	return true
}

func ReadChar() byte {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Panicln("Error on reading string!")
	}
	return byte(unicode.ToUpper(rune(input[0])))
}

func FloodFill(board *models.Board, newColor byte) {
	var arr models.Vector
	a, b := 0, 0
	lastColor := board.Matrix[0][0]

	if newColor != lastColor {
		AppendVector(&arr, a, b)
		for len(arr.Vec) != 0 {
			RemoveFromArray(&arr, &a, &b)
			board.Matrix[a][b] = newColor
			CheckNeighbors(*board, &arr, a, b, lastColor)
		}
	}
	PrintBoard(*board)
}

func VerifyIfIsInRange(input byte, len int) bool {
	for _, i := range models.Letters[:len+1] {
		if i == input {
			return true
		}
	}
	return false
}
