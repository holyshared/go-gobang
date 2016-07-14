package main

import (
  "fmt"
  "./game"
)

func main() {
  board := game.NewBoard(10, 10)
  cell := board.At(0, 0)

  blackStone := game.Black
  blackStone.PutTo(cell)

  whiteStone := game.White
  whiteStone.PutTo(cell)

  fmt.Println(board)
}
