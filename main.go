package main

import (
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Map         GameMap
	TurnCounter int
	Turn        TurnState
	World       *ecs.Manager
	WorldTags   map[string]ecs.Tag
}

func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()

	g.Turn = PlayerTurn
	g.TurnCounter = 0

	world, tags := InitializeWorld(g.Map.CurrentLevel())
	g.WorldTags = tags
	g.World = world

	return g
}

func (g *Game) Update() error {
	g.TurnCounter++
	if g.Turn == PlayerTurn && g.TurnCounter > 6 {
		TakePlayerAction(g)
	}
	if g.Turn == MonsterTurn {
		TakeMonsterAction(g)
	}
	g.Turn = PlayerTurn
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.CurrentLevel()
	level.DrawLevel(screen)
	ProcessRenderables(g, level, screen)
	ProcessUserLog(g, screen)
	ProcessHUD(g, screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(g.Layout(0, 0))
	ebiten.SetWindowTitle("Mongo")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
