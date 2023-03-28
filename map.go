package main

type GameMap struct {
	Dungeons       []Dungeon
	CurrentDungeon Dungeon
}

func NewGameMap() GameMap {
	d := NewDungeon()
	dungeons := make([]Dungeon, 0)
	dungeons = append(dungeons, d)
	gm := GameMap{Dungeons: dungeons, CurrentDungeon: d}
	return gm
}

func (gm *GameMap) CurrentLevel() Level {
	return gm.CurrentDungeon.CurrentLevel
}
