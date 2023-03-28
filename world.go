package main

import (
	"github.com/bytearena/ecs"
)

var (
	armor       *ecs.Component
	health      *ecs.Component
	meleeWeapon *ecs.Component
	monster     *ecs.Component
	name        *ecs.Component
	position    *ecs.Component
	renderable  *ecs.Component
	userMessage *ecs.Component
)

func InitializeWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	armor = manager.NewComponent()
	health = manager.NewComponent()
	meleeWeapon = manager.NewComponent()
	monster = manager.NewComponent()
	movable := manager.NewComponent()
	name = manager.NewComponent()
	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	userMessage = manager.NewComponent()

	playerImg := loadImage("assets/player.png")
	skellyImage := loadImage("assets/skelly.png")
	orcImage := loadImage("assets/orc.png")

	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()
	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			Image: playerImg,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: x,
			Y: y,
		}).
		AddComponent(health, &Health{
			MaxHealth:     30,
			CurrentHealth: 30,
		}).
		AddComponent(meleeWeapon, &MeleeWeapon{
			Name:          "Battle Axe",
			MinimumDamage: 10,
			MaximumDamage: 20,
			ToHitBonus:    3,
		}).
		AddComponent(armor, &Armor{
			Name:       "Plate Armor",
			Defense:    15,
			ArmorClass: 18,
		}).
		AddComponent(name, &Name{Label: "Hiro"}).
		AddComponent(userMessage, &UserMessage{
			AttackMessage:    "",
			DeadMessage:      "",
			GameStateMessage: "",
		})

	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			mX, mY := room.Center()

			mobSpawn := GetDiceRoll(2)
			if mobSpawn == 1 {
				manager.NewEntity().
					AddComponent(monster, &Monster{}).
					AddComponent(renderable, &Renderable{
						Image: orcImage,
					}).
					AddComponent(position, &Position{
						X: mX,
						Y: mY,
					}).
					AddComponent(health, &Health{
						MaxHealth:     30,
						CurrentHealth: 30,
					}).
					AddComponent(meleeWeapon, &MeleeWeapon{
						Name:          "Machete",
						MinimumDamage: 4,
						MaximumDamage: 8,
						ToHitBonus:    1,
					}).
					AddComponent(armor, &Armor{
						Name:       "Leather",
						Defense:    5,
						ArmorClass: 6,
					}).
					AddComponent(name, &Name{Label: "Orc"}).
					AddComponent(userMessage, &UserMessage{
						AttackMessage:    "",
						DeadMessage:      "",
						GameStateMessage: "",
					})
			} else {
				manager.NewEntity().
					AddComponent(monster, &Monster{}).
					AddComponent(renderable, &Renderable{
						Image: skellyImage,
					}).
					AddComponent(position, &Position{
						X: mX,
						Y: mY,
					}).
					AddComponent(health, &Health{
						MaxHealth:     10,
						CurrentHealth: 10,
					}).
					AddComponent(meleeWeapon, &MeleeWeapon{
						Name:          "Short Sword",
						MinimumDamage: 2,
						MaximumDamage: 6,
						ToHitBonus:    0,
					}).
					AddComponent(armor, &Armor{
						Name:       "Bone",
						Defense:    3,
						ArmorClass: 4,
					}).
					AddComponent(name, &Name{Label: "Skeleton"}).
					AddComponent(userMessage, &UserMessage{
						AttackMessage:    "",
						DeadMessage:      "",
						GameStateMessage: "",
					})
			}
		}
	}

	monsters := ecs.BuildTag(
		armor,
		health,
		meleeWeapon,
		monster,
		name,
		position,
		userMessage,
	)
	players := ecs.BuildTag(
		player,
		position,
		health,
		meleeWeapon,
		armor,
		name,
		userMessage,
	)
	renderables := ecs.BuildTag(renderable, position)
	messengers := ecs.BuildTag(userMessage)

	tags["monsters"] = monsters
	tags["players"] = players
	tags["renderables"] = renderables
	tags["messengers"] = messengers

	return manager, tags
}
