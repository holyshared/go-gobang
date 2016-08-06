package gobang

type GameProgressResult int

const (
  Win GameProgressResult = iota + 1
  Lose
  Draw
  Next
  Retry
)
