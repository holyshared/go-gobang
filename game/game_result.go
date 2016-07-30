package game

type GameResult int

const (
  Win GameResult = iota + 1
  Lose
  Draw
)
