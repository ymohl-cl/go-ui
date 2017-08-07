package scripter

/*
** private functions to the scripter
 */

func (s *Script) buildNewScene(index uint8) error {
	var err error

	if err = s.list[index].Build(); err != nil {
		return err
	}
	if err = s.list[index].Init(); err != nil {
		return err
	}

	return nil
}
