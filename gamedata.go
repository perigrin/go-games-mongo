package main

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
	UIHeight     int
}

func NewGameData() GameData {
	g := GameData{
		ScreenWidth:  80,
		ScreenHeight: 60,
		TileWidth:    12,
		TileHeight:   12,
		UIHeight:     10,
	}
	return g
}

func (gd GameData) GameSize() int {
	return (gd.ScreenHeight - gd.UIHeight) * gd.ScreenWidth
}
