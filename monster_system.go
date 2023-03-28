package main

import (
	"github.com/norendren/go-fov/fov"
)

func TakeMonsterAction(game *Game) {
	l := game.Map.CurrentLevel()
	playerPosition := Position{}

	for _, plr := range game.World.Query(game.WorldTags["players"]) {
		pos := plr.Components[position].(*Position)
		playerPosition.X = pos.X
		playerPosition.Y = pos.Y
	}

	for _, r := range game.World.Query(game.WorldTags["monsters"]) {
		pos := r.Components[position].(*Position)
		// mon := r.Components[monster].(*Monster)
		monsterSees := fov.New()
		monsterSees.Compute(l, pos.X, pos.Y, 8)
		if monsterSees.IsVisible(playerPosition.X, playerPosition.Y) {
			if pos.GetManhattanDistance(&playerPosition) == 1 {
				AttackSystem(game, pos, &playerPosition)
				if r.Components[health].(*Health).CurrentHealth <= 0 {
					// We Ded ... stop taking up real estate
					l.Tiles[l.GetIndexFromXY(pos.X, pos.Y)].Blocked = false
				}
			}
			astar := AStar{}
			path := astar.GetPath(l, pos, &playerPosition)
			if len(path) > 1 {
				nextTile := l.Tiles[l.GetIndexFromXY(path[1].X, path[1].Y)]
				if nextTile.IsWalkable() {
					l.Tiles[l.GetIndexFromXY(pos.X, pos.Y)].Blocked = false
					pos.X = path[1].X
					pos.Y = path[1].Y
					l.Tiles[l.GetIndexFromXY(path[1].X, path[1].Y)].Blocked = true
				}
			}
		}
	}
	game.Turn = PlayerTurn
}
