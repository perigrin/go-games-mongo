package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/norendren/go-fov/fov"
)

var levelHeight int

type Level struct {
	Tiles         []MapTile
	Rooms         []Rect
	PlayerVisible *fov.View
}

func NewLevel() Level {
	l := Level{}
	rooms := make([]Rect, 0)
	l.Rooms = rooms
	l.GenerateLevelTiles()
	l.PlayerVisible = fov.New()
	return l
}

func (level *Level) GetIndexFromXY(x, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

func (level Level) InBounds(x, y int) bool {
	idx := level.GetIndexFromXY(x, y)
	if idx < len(level.Tiles) {
		return true
	}
	return false
}

func (level Level) IsOpaque(x, y int) bool {
	return level.Tiles[level.GetIndexFromXY(x, y)].IsOpaque()
}

func (level *Level) createRoom(room Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].convertToFloor()
		}
	}
}

func (level *Level) createHorizontalTunnel(x1 int, x2 int, y int) {
	// gd := NewGameData()
	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		index := level.GetIndexFromXY(x, y)
		if index >= 0 && index < len(level.Tiles) {
			level.Tiles[index].convertToFloor()
		}
	}
}

func (level *Level) createVerticalTunnel(y1 int, y2 int, x int) {
	// gd := NewGameData()
	for y := min(y1, y2); y < max(y1, y2)+1; y++ {
		index := level.GetIndexFromXY(x, y)
		if index > 0 && index < len(level.Tiles) {
			level.Tiles[index].convertToFloor()
		}
	}
}

func (level *Level) createTiles() []MapTile {
	gd := NewGameData()
	tiles := make([]MapTile, levelHeight*gd.ScreenWidth)

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < levelHeight; y++ {
			tiles[level.GetIndexFromXY(x, y)] = WallTile(x, y)
		}
	}

	return tiles
}

func (level *Level) GenerateLevelTiles() {
	MIN_SIZE := 5
	MAX_SIZE := 10
	MAX_ROOMS := 30

	gd := NewGameData()
	levelHeight = gd.ScreenHeight - gd.UIHeight

	tiles := level.createTiles()
	level.Tiles = tiles

	for idx := 0; idx < MAX_ROOMS; idx++ {
		w := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		h := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		x := GetDiceRoll(gd.ScreenWidth - w - 1)
		y := GetDiceRoll(levelHeight - h - 1)

		new_room := NewRect(x, y, w, h)
		okToAdd := true
		for _, r := range level.Rooms {
			if new_room.Intersect(r) {
				okToAdd = false
				break
			}
		}
		if okToAdd {
			level.createRoom(new_room)
			if len(level.Rooms) != 0 {
				last_room := level.Rooms[len(level.Rooms)-1]
				level.connectRooms(new_room, last_room)
			}
			level.Rooms = append(level.Rooms, new_room)
		}
	}
}

func (level *Level) connectRooms(r1, r2 Rect) {
	r1X, r1Y := r1.Center()
	r2X, r2Y := r2.Center()
	coin := GetDiceRoll(2)
	if coin == 2 {
		level.createHorizontalTunnel(r1X, r2X, r1Y)
		level.createVerticalTunnel(r1Y, r2Y, r2X)
	} else {
		level.createHorizontalTunnel(r1X, r2X, r2Y)
		level.createVerticalTunnel(r1Y, r2Y, r1X)
	}
}

func (level *Level) DrawLevel(screen *ebiten.Image) {
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < levelHeight; y++ {
			index := level.GetIndexFromXY(x, y)
			isVisible := level.PlayerVisible.IsVisible(x, y)
			// MapTile.Draw mutates so we need the pointer
			level.Tiles[index].Draw(isVisible, screen)
		}
	}
}
