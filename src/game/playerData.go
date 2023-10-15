package game

type CLASS int

const (
	MORI CLASS = iota
	LUI
)

type PlayerData struct {
	class CLASS
	level int
}

func NewPlayerData() *PlayerData {
	return &PlayerData{}
}

func (pd *PlayerData) SetPlayerClass(c CLASS) {
	pd.class = pd.class
}
