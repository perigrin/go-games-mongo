package main

import (
	"fmt"

	"github.com/bytearena/ecs"
)

func AttackSystem(g *Game, attackerPosition *Position, defenderPosition *Position) {
	var attacker *ecs.QueryResult = nil
	var defender *ecs.QueryResult = nil

	for _, playerCombatant := range g.World.Query(g.WorldTags["players"]) {
		pos := playerCombatant.Components[position].(*Position)
		if pos.IsEqual(attackerPosition) {
			// Player is the Attacker
			attacker = playerCombatant
		} else if pos.IsEqual(defenderPosition) {
			// Player is on the defense
			defender = playerCombatant
		}
	}

	for _, cbt := range g.World.Query(g.WorldTags["monsters"]) {
		pos := cbt.Components[position].(*Position)
		if pos.IsEqual(attackerPosition) {
			attacker = cbt
		} else if pos.IsEqual(defenderPosition) {
			defender = cbt
		}
	}

	if attacker == nil || defender == nil {
		return
	}

	defenderArmor := defender.Components[armor].(*Armor)
	defenderHealth := defender.Components[health].(*Health)
	defenderName := defender.Components[name].(*Name).Label
	attackerWeapon := attacker.Components[meleeWeapon].(*MeleeWeapon)
	attackerName := attacker.Components[name].(*Name).Label
	defenderMessage := defender.Components[userMessage].(*UserMessage)
	attackerMessage := attacker.Components[userMessage].(*UserMessage)

	// attacker is dead, the fight is out of them
	if attacker.Components[health].(*Health).CurrentHealth <= 0 {
		return
	}
	toHitRoll := GetDiceRoll(10)
	if toHitRoll+attackerWeapon.ToHitBonus > defenderArmor.ArmorClass {
		// THWAP!
		damageRoll := GetRandomBetween(attackerWeapon.MinimumDamage, attackerWeapon.MaximumDamage)
		damageDone := damageRoll - defenderArmor.Defense

		if damageDone < 0 {
			damageDone = 0
		}

		defenderHealth.CurrentHealth -= damageDone
		attackerMessage.AttackMessage = fmt.Sprintf("%s swings %s at %s and hits for %d health.\n", attackerName, attackerWeapon.Name, defenderName, damageDone)
		if defenderHealth.CurrentHealth <= 0 {
			defenderMessage.DeadMessage = fmt.Sprintf("%s has died!\n", defenderName)
			if defenderName == "Player" {
				defenderMessage.GameStateMessage = fmt.Sprintf("Game Over!\n")
				g.Turn = GameOver
			}
			// g.World.DisposeEntity(defender.Entity)
		}
	} else {
		// WOOSH
		defenderMessage.AttackMessage = fmt.Sprintf("%s swings %s at %s and misses.\n", attackerName, attackerWeapon.Name, defenderName)
	}
}
