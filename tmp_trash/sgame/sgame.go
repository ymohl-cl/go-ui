package sgame

/*type SGame struct {
	input  string
	notice string

	// sObjs defines static object to this scene.
	sObjs []*objects.ObjectType
	// dObjs define dynamical object to this scene.
	dObjs []*objects.ObjectType
}

func (s *SGame) Init(data *database.Data) error {
	var x int32
	var y int32

	// background
	OBackground := objects.NewImage("background.bmp", "background", I.GlobalOriginX, I.GlobalOriginY, I.ScreenWidth, I.ScreenHeight)
	TBackground := objects.New(objects.TypeImage, OBackground)
	TBackground.SetStatus(objects.StatusFix)

	// header
	OContainerTitle := objects.NewContainer(I.CGLR, I.CGLG, I.CGLB, I.CGLO, I.GlobalOriginX, I.MarginTop, I.ScreenWidth, I.HeightHeaderFooter, I.DrawFilled)
	TContainerTitle := objects.New(objects.TypeContainer, OContainerTitle)
	TContainerTitle.SetStatus(objects.StatusFix)
	x = I.ScreenWidth / 2
	y = I.MarginTop + I.HeightHeaderFooter/2
	OTitle := objects.NewText(I.CWTR, I.CWTG, I.CWTB, I.CWTO, x, y, I.SizeTitle, "Gomoku game", I.FontText)
	TTitle := objects.New(objects.TypeText, OTitle)
	TTitle.SetStatus(objects.StatusFix)
	OStyleTitle := objects.NewText(I.CBTR, I.CBTG, I.CBTB, I.CBTO, x+2, y+2, I.SizeTitle, "Gomoku game", I.FontText)
	TStyleTitle := objects.New(objects.TypeText, OStyleTitle)
	TStyleTitle.SetStatus(objects.StatusFix)
	TContainerTitle.SetChilds(TStyleTitle, TTitle)

	s.sObjs = append(s.sObjs, TBackground)
	s.sObjs = append(s.sObjs, TContainerTitle)
	return nil
}

func (s *SGame) Update(data *database.Data) error {
	return nil
}

func (s *SGame) GetObjs() []*objects.ObjectType {
	return s.sObjs
}

func (s *SGame) AddLetterToInput(str string) {
	return
}

func (s *SGame) GetStaticObjs() []*objects.ObjectType {
	return s.sObjs
}

func (s *SGame) GetDynamicObjs() []*objects.ObjectType {
	return nil
}

func (s *SGame) SetNotice(str string) {
	return
}
*/
