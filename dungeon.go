package main

type Dungeon struct {
	Name         string
	Levels       []Level
	CurrentLevel Level
}

func NewDungeon() Dungeon {
	l := NewLevel()
	levels := make([]Level, 0)
	levels = append(levels, l)
	d := Dungeon{Name: "default", Levels: levels, CurrentLevel: l}
	return d
}
