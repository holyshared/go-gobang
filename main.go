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

  fmt.Println(cell.Have(game.White))
  fmt.Println(cell.Have(game.Black))

  fmt.Println(board)
}
