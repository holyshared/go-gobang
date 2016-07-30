package gobang

type PutStoneResult int

const (
  Failed PutStoneResult = iota
  Continue
  Reached
  Filled
)
