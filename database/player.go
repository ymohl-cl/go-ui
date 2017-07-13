package database

import "fmt"

func (P Player) Play() {
	fmt.Println("Player play !")
}

func CreatePlayer(name string) *Player {
	p := new(Player)

	p.Name = name
	return p
}

func (P *Player) DeleteSave() string {
	if len(P.Saves) == 0 {
		return "No session saved"
	}
	P.Saves = make([]*Session, 0)
	return ""
}
