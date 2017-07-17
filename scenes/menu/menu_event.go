package menu

func (M *Menu) AddUser(values ...interface{}) string {
	var data *database.Data
	for idx, v := range values {
		switch idx {
		case 0:
			fmt.Println("AddUser case 1")
			data = v.(*database.Data)
		}
	}

	if len(data.Players) > 10 {
		return "there are aleady 10 profile saved. It's a max"
	}
	if len(S.input) == 0 {
		return "Can't save profile without name"
	}
	if len(S.input) > 10 {
		return "Limit 10 Characters to name player"
	}
	for _, p := range data.Players {
		if p.Name == S.input {
			return "This name already exist"
		}
	}
	p := database.CreatePlayer(S.input)
	data.AddPlayer(p)
	if sdl.IsTextInputActive() == true {
		fmt.Println("Close input")
		sdl.StopTextInput()
		S.input = ""
	}
	return ""
}

func (S *SUser) ClearName(values ...interface{}) string {
	S.input = ""
	return ""
}

func (S *SUser) DeleteStatUser(values ...interface{}) string {
	fmt.Println("DeleteStatUser")
	p := values[0].(*database.Player)
	str := p.DeleteSave()
	return str
}

func (S *SUser) InitInputText(values ...interface{}) string {
	fmt.Println("InitINputText")
	pos := values[0].(objects.ParamPosition)
	size := values[1].(objects.ParamSize)
	if sdl.IsTextInputActive() != true {
		rect := sdl.Rect{pos.X, pos.Y, size.W, size.H}
		sdl.SetTextInputRect(&rect)
		sdl.StartTextInput()
		fmt.Println("OK")
	}
	return ""
}

func (S *SUser) PlayUser(values ...interface{}) string {
	sinfos.Current = sinfos.SceneGame
	fmt.Println("Go Play")
	return "Need change scene"
}

func (S *SUser) DeleteUser(values ...interface{}) string {
	var data *database.Data

	for idx, value := range values {
		switch idx {
		case 1:
			data = value.(*database.Data)
		}
	}
	fmt.Println("DeleteUser")
	str := data.DeletePlayer(values[0].(*database.Player))

	return str
}

func (S *SUser) ResetChoiceToPlay(values ...interface{}) string {
	var Data *database.Data

	for idx, value := range values {
		switch idx {
		case 0:
			Data = value.(*database.Data)
		}
	}

	fmt.Println("ResetChoiceToPlay")
	str := Data.ResetCurrent()

	return str
}

func (S *SUser) AddCurrent(values ...interface{}) string {
	var data *database.Data

	for idx, value := range values {
		switch idx {
		case 0:
			data = value.(*database.Data)
		}
	}

	fmt.Println("AddCurrent")
	p := values[0].(*database.Player)
	notice := data.UpdateCurrent(p)
	return notice
}

func (S *SUser) LookStats(values ...interface{}) string {
	sinfos.Current = sinfos.SceneStat
	fmt.Println("Go Play")
	fmt.Println("Coucou c'est moi LookStats")
	return ""
}

func (S *SUser) Quit(values ...interface{}) string {
	fmt.Println("Coucou c'est moi Quit!")
	return ""
}
*/
