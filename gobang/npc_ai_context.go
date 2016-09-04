package gobang

type NpcAIContext struct {
	rule           *GameRule
	board          *Board
	playerStone    Stone
	npcPlayerStone Stone
}

func (c *NpcAIContext) Board() *Board {
	return c.board
}

func (c *NpcAIContext) ReachedStoneCount() int {
	return c.rule.ReachedStoneCount()
}

func (c *NpcAIContext) PlayerStone() Stone {
	return c.playerStone
}

func (c *NpcAIContext) NpcPlayerStone() Stone {
	return c.npcPlayerStone
}
