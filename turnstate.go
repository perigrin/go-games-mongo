package main

type TurnState int

const (
	BeforePlayerAction = iota
	PlayerTurn
	MonsterTurn
	GameOver
)

func GetNextState(state TurnState) TurnState {
	switch state {
	case BeforePlayerAction:
		return PlayerTurn
	case PlayerTurn:
		return MonsterTurn
	case MonsterTurn:
		return BeforePlayerAction
	case GameOver:
		return GameOver // Game Over Man, GAME OVER
	default:
		return PlayerTurn
	}
}
