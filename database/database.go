package database

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/ymohl-cl/game-builder/conf"
)

const (
	defaultPlayer1 = "Unknow1"
	defaultPlayer2 = "Unknow2"
)

func Get() (*Data, error) {
	D := new(Data)

	f, err := os.OpenFile(conf.ProtoBufFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	if err = f.Close(); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadFile(conf.ProtoBufFile)
	if err != nil {
		return nil, err
	}

	err = proto.Unmarshal(buf, D)
	if err != nil {
		return nil, err
	}

	err = D.initSave()

	return D, err
}

func (D *Data) initSave() error {
	var err error

	if len(D.Players) == 0 {
		unknow1 := CreatePlayer(defaultPlayer1)
		unknow2 := CreatePlayer(defaultPlayer2)
		D.Players = append(D.Players, unknow1)
		D.Players = append(D.Players, unknow2)
	}

	D.Current = new(Session)
	if err = D.DefaultPlayers(); err != nil {
		return errors.New("Save file is corrupted: " + err.Error())
	}
	return nil
}

func (D *Data) UpdateCurrent(p *Player) error {
	if D.Current.P1.Name == p.Name || D.Current.P2.Name == p.Name {
		return errors.New("This player is already selected")
	}

	if D.Current.P1.Name == defaultPlayer1 {
		D.Current.P1 = p
		return nil
	}

	D.Current.P2 = p
	return nil
}

func (D *Data) DefaultPlayers() error {
	for _, p := range D.Players {
		if p.Name == defaultPlayer1 {
			D.Current.P1 = p
		} else if p.Name == defaultPlayer2 {
			D.Current.P2 = p
		}
	}
	if D.Current.P1 == nil || D.Current.P2 == nil {
		return errors.New("Default players not found")
	}
	return nil
}

func (D *Data) AddPlayer(p *Player) {
	D.Players = append(D.Players, p)
	fmt.Println("il y a ", len(D.Players), "players dans la base into database")
}

func (D *Data) DeletePlayer(p *Player) (int, error) {
	if p.Name == defaultPlayer1 || p.Name == defaultPlayer2 {
		return 0, errors.New("You can't remove defaultUser Unknow 1 and 2")
	}
	for id, pt := range D.Players {
		if pt.Name == p.Name {
			D.Players = append(D.Players[:id], D.Players[id+1:]...)
			return id, D.DefaultPlayers()
		}
	}
	return 0, errors.New("Player name not found")
}
