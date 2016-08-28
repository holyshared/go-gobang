package gobang

import (
  "encoding/json"
)

func NewGameContext(rule *GameRule, playerStone, npcPlayerStone Stone) *GameContext {
  board := NewBoard(rule.BoardSize())
  player := NewGamePlayer(playerStone, board)

  ctx := &NpcAIContext {
    rule: rule,
    board: board,
    playerStone: playerStone,
    npcPlayerStone: npcPlayerStone,
  }
  ai := NewNpcAI(ctx)
  npcPlayer := NewNpcPlayer(npcPlayerStone, ai)

  return &GameContext {
    GameRule: rule,
    board: board,
    currentPlayer: player,
    player: player,
    npcPlayer: npcPlayer,
  }
}

type GameContext struct {
  *GameRule
  board *Board
  currentPlayer Player
  player *GamePlayer
  npcPlayer *NpcPlayer
}

func (g *GameContext) CurrentBoard() *Board {
  return g.board
}

func (g *GameContext) SelectCell(point *Point) (*Cell, error) {
  if !g.board.HaveCell(point) {
    return nil, NewCellNotFoundError(point)
  }
  return g.board.SelectCell(point), nil
}

func (g *GameContext) CurrentPlayer() Player {
  return g.currentPlayer
}

func (g *GameContext) GamePlayer() *GamePlayer {
  return g.player
}

func (g *GameContext) NpcPlayer() *NpcPlayer {
  return g.npcPlayer
}

func (g *GameContext) ChangeToNextPlayer() {
  var player Player

  if (g.currentPlayer == g.npcPlayer) {
    player = g.player
  } else {
    player = g.npcPlayer
  }

  g.currentPlayer = player
}

func (g *GameContext) CheckBoard() PutStoneResult {
  player := g.CurrentPlayer()

  matcher := NewCellReachedMatcher(player.Stone(), g.ReachedStoneCount())
  result := matcher.Matches(g.board)

  if result.HasResult() {
    return Reached
  }

  if g.board.IsAllFilled() {
    return Filled
  }

  return Continue
}

func (g *GameContext) MarshalJSON() ([]byte, error) {
  jsonObject := struct {
    Rule *GameRule `json:"rule"`
    Board *Board `json:"board"`
    CurrentPlayer Player `json:"currentPlayer"`
    Player *GamePlayer `json:"player"`
    NpcPlayer *NpcPlayer `json:"npcPlayer"`
  }{
    Rule: g.GameRule,
    Board: g.board,
    CurrentPlayer: g.currentPlayer,
    Player: g.player,
    NpcPlayer: g.npcPlayer,
  }
  return json.Marshal(jsonObject)
}
