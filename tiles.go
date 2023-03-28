package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type TileType int

const (
	WALL TileType = iota
	FLOOR
)

type MapTile struct {
	Blocked    bool
	Image      *ebiten.Image
	IsRevealed bool
	PixelX     int
	PixelY     int
	TileType   TileType
}

var (
	wallImg  = loadImage("assets/wall.png")
	floorImg = loadImage("assets/floor.png")
)

func WallTile(x, y int) MapTile {
	gd := NewGameData()
	return MapTile{
		Blocked:    true,
		Image:      wallImg,
		IsRevealed: false,
		PixelX:     x * gd.TileWidth,
		PixelY:     y * gd.TileHeight,
		TileType:   WALL,
	}
}

func (t *MapTile) convertToFloor() {
	if t.TileType != FLOOR {
		t.TileType = FLOOR
		t.Image = floorImg
		t.Blocked = false
	}
}

func (t MapTile) IsWalkable() bool {
	return !t.Blocked
}

func (t MapTile) IsOpaque() bool {
	return t.TileType == WALL
}

func (t *MapTile) Draw(visible bool, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.PixelX), float64(t.PixelY))
	if visible {
		op.ColorM.Translate(0, 0, 0, 0.50)
		screen.DrawImage(t.Image, op)
		t.IsRevealed = true
	} else if t.IsRevealed {
		screen.DrawImage(t.Image, op)
	}
}
