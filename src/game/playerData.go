package game

type CLASS int

const (
	MORI PLAYERSTATE = iota
	LUI
)

type PlayerData struct {
	Class CLASS
	level int
}

func NewPlayerData() *PlayerData {
	return &PlayerData{}
}

func (pd *PlayerData) SetPlayerClass(c CLASS) {
	pd.Class = pd.Class
}
