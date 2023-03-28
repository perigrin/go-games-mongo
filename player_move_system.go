package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func TakePlayerAction(g *Game) {
	players := g.WorldTags["players"]
	x := 0
	y := 0
	turnTaken := false

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		turnTaken = true
	}

	level := g.Map.CurrentLevel()

	for _, result := range g.World.Query(players) {
		pos := result.Components[position].(*Position)
		index := level.GetIndexFromXY(pos.X+x, pos.Y+y)
		next := level.Tiles[index]

		if next.IsWalkable() {
			level.Tiles[level.GetIndexFromXY(pos.X, pos.Y)].Blocked = false
			pos.X += x
			pos.Y += y
			level.PlayerVisible.Compute(level, pos.X, pos.Y, 8)
			level.Tiles[index].Blocked = true
		} else if x != 0 || y != 0 {
			// We bumped into something
			if next.TileType != WALL {
				// And it's not a wall ...
				monsterPosition := Position{X: pos.X + x, Y: pos.Y + y}
				// ATTACK!
				AttackSystem(g, pos, &monsterPosition)
			}
		}

		if x != 0 || y != 0 || turnTaken {
			g.Turn = GetNextState(g.Turn)
			g.TurnCounter = 0
		}
	}
}
