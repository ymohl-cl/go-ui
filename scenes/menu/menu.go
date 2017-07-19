package menu

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/ymohl-cl/game-builder/database"
	"github.com/ymohl-cl/game-builder/objects"
	"github.com/ymohl-cl/game-builder/objects/text"
)

const (
	// order layers of scene
	layerBackground = 0
	layerStructure  = 1
	layerButton     = 2
	layerNotice     = 3
	layerText       = 4
	layerVS         = 5
	layerInput      = 6
	layerPlayers    = 7
)

type Menu struct {
	layers map[uint8][]objects.Object

	input  objects.Object
	notice *text.Text
	music  objects.Object
	vs     *text.Text
	data   *database.Data

	/* sdl objects */
	renderer *sdl.Renderer
}

/*
** Functions scene specifications
 */

/*
** Interface scene functions
 */
func (M *Menu) Init(d *database.Data, r *sdl.Renderer) error {
	var err error

	if r == nil {
		return errors.New(objects.ErrorRenderer)
	}
	M.renderer = r
	M.data = d

	M.layers = make(map[uint8][]objects.Object)

	if err = M.build(); err != nil {
		return err
	}

	if err = M.check(); err != nil {
		return err
	}
	return nil
}

func (M Menu) Run() error {
	var wg sync.WaitGroup

	if ok := M.music.IsInit(); ok {
		wg.Add(1)
		go M.music.Draw(&wg, M.renderer)
		wg.Wait()
	}
	return nil
}

func (M Menu) Close() error {
	var err error

	if ok := M.music.IsInit(); ok {
		if err = M.music.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (M Menu) GetLayers() map[uint8][]objects.Object {
	return M.layers
}

/*func (M Menu) Draw() {
	var wg sync.WaitGroup

	for _, layer := range M.layers {
		for _, object := range layer {
			fmt.Println("drawing")
			go object.Draw(&wg)
		}
		wg.Wait()
	}
}*/

/*
** Private function scene
 */

func (M Menu) check() error {
	if M.layers == nil {
		return errors.New("Objects not define for menu scene")
	}
	/*	if M.input == nil {
			return errors.New("Object to input not define")
		}
		if M.notice == nil {
			return errors.New("Object to notice not define")
		}*/

	return nil
}

/*type SUser struct {
input  string
notice string

// music
// layout map[uint8][]*objects.ObjectType
// nbLayout uint8
/*	sObjs              []*objects.ObjectType
	dObjs              []*objects.ObjectType
	inputObjs          *objects.ObjectType
	playerObjs         []*objects.ObjectType
	playerSelectedObjs []*objects.ObjectType*/
//}

//s.ProcessUpload()

/*func (s *SUser) Init(dt *database.Data) error {
	var x int32
	var y int32

	// music
	OMusic := objects.NewMusic("ambiant.wav", "ambiant")
	TMusic := objects.New(objects.TypeMusic, OMusic)
	TMusic.SetStatus(objects.StatusFix)

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

	// Container Users recorded
	y = I.MarginTop + I.Padding + I.HeightHeaderFooter
	OContainerUsers := objects.NewContainer(I.CGLR, I.CGLG, I.CGLB, I.CGLO, I.MarginLeft, y, I.UserWidthBloc, I.UserheightBloc, I.DrawFilled)
	TContainerUsers := objects.New(objects.TypeContainer, OContainerUsers)
	TContainerUsers.SetStatus(objects.StatusFix)

	// Container NewUser -- bloc right to create new player
	x = I.MarginRight - I.UserWidthBloc
	OContainerNewUser := objects.NewContainer(I.CGLR, I.CGLG, I.CGLB, I.CGLO, x, y, I.UserWidthBloc, I.UserheightBloc, I.DrawFilled)
	TContainerNewUser := objects.New(objects.TypeContainer, OContainerNewUser)
	TContainerNewUser.SetStatus(objects.StatusFix)

	// Add Button create new Users
	x = I.MarginRight - (I.UserWidthBloc / 2) - (I.WidthLargeButton / 2)
	y = I.MarginTop + I.Padding + I.HeightHeaderFooter + I.UserheightBloc - (I.UserheightBloc / 2) + I.HeightButton + I.Padding
	OButtunNew := objects.NewContainer(I.CYOR, I.CYOG, I.CYOB, I.CYOO, x, y, I.WidthLargeButton, I.HeightButton, I.DrawFilled)
	OButtunNewOver := objects.NewContainer(I.CYR, I.CYG, I.CYB, 255, x, y, I.WidthLargeButton, I.HeightButton, I.DrawFilled)
	OButtunNewClicDown := objects.NewContainer(I.CYR, I.CYG, I.CYB, 255, x+1, y+1, I.WidthLargeButton-2, I.HeightButton-2, I.DrawFilled)
	TButtunNew := objects.New(objects.TypeContainer, OButtunNew)
	TButtunNew.SetObjOver(OButtunNewOver)
	TButtunNew.SetObjClicDown(OButtunNewClicDown)
	TButtunNew.Action = s.AddUser
	TButtunNew.ActionDatas = append(TButtunNew.ActionDatas, dt)
	x += I.WidthLargeButton / 2
	y += I.HeightButton / 2
	OTextButtunNew := objects.NewText(I.CWTR, I.CWTG, I.CWTB, I.CWTO, x, y, I.SizeButton, "Create new player", I.FontText)
	TTextButtunNew := objects.New(objects.TypeText, OTextButtunNew)
	TTextButtunNew.SetStatus(objects.StatusFix)
	OStyleTextButtunNew := objects.NewText(I.CBTR, I.CBTG, I.CBTB, I.CBTO, x, y+2, I.SizeButton, "Create new player", I.FontText)
	TStyleTextButtunNew := objects.New(objects.TypeText, OStyleTextButtunNew)
	TStyleTextButtunNew.SetStatus(objects.StatusFix)
	TButtunNew.SetChilds(TStyleTextButtunNew, TTextButtunNew)

	// Add Button reset name
	x = I.MarginRight - (I.UserWidthBloc / 2) - (I.WidthLargeButton / 2)
	y = I.MarginTop + I.Padding + I.HeightHeaderFooter + I.UserheightBloc - (I.UserheightBloc / 2) + I.HeightButton*2 + I.Padding*2
	OButtunClearName := objects.NewContainer(I.CYOR, I.CYOG, I.CYOB, I.CYOO, x, y, I.WidthLargeButton, I.HeightButton, I.DrawFilled)
	OButtunClearNameOver := objects.NewContainer(I.CYR, I.CYG, I.CYB, 255, x, y, I.WidthLargeButton, I.HeightButton, I.DrawFilled)
	OButtunClearClicDown := objects.NewContainer(I.CYR, I.CYG, I.CYB, 255, x+1, y+1, I.WidthLargeButton-2, I.HeightButton-2, I.DrawFilled)
	TButtunClearName := objects.New(objects.TypeContainer, OButtunClearName)
	TButtunClearName.SetObjOver(OButtunClearNameOver)
	TButtunClearName.SetObjClicDown(OButtunClearClicDown)
	TButtunClearName.Action = s.ClearName
	x += I.WidthLargeButton / 2
	y += I.HeightButton / 2
	OTextButtunClearName := objects.NewText(I.CWTR, I.CWTG, I.CWTB, I.CWTO, x, y, I.SizeButton, "clear name new player", I.FontText)
	TTextButtunClearName := objects.New(objects.TypeText, OTextButtunClearName)
	TTextButtunClearName.SetStatus(objects.StatusFix)
	OStyleTextButtunClearName := objects.NewText(I.CBTR, I.CBTG, I.CBTB, I.CBTO, x, y+2, I.SizeButton, "clear name new player", I.FontText)
	TStyleTextButtunClearName := objects.New(objects.TypeText, OStyleTextButtunClearName)
	TStyleTextButtunClearName.SetStatus(objects.StatusFix)
	TButtunClearName.SetChilds(TStyleTextButtunClearName, TTextButtunClearName)

	// Add Button play
	x = I.ScreenWidth/2 - I.WidthButton/2
	y = I.MarginTop + I.Padding + I.HeightHeaderFooter + I.UserheightBloc - (I.HeightButton + I.Padding + I.HeightButton)
	OButtunPlay := objects.NewContainer(I.CYOR, I.CYOG, I.CYOB, I.CYOO, x, y, I.WidthButton, I.HeightButton, I.DrawFilled)
	OButtunPlayOver := objects.NewContainer(I.CYR, I.CYG, I.CYB, 255, x, y, I.WidthButton, I.HeightButton, I.DrawFilled)
	OButtunPlayClicDown := objects.NewContainer(I.CYR, I.CYG, I.CYB, 255, x+1, y+1, I.WidthButton-2, I.HeightButton-2, I.DrawFilled)
	TButtunPlay := objects.New(objects.TypeContainer, OButtunPlay)
	TButtunPlay.SetObjOver(OButtunPlayOver)
	TButtunPlay.SetObjClicDown(OButtunPlayClicDown)
	TButtunPlay.Action = s.PlayUser
	x = I.ScreenWidth / 2
	y += I.HeightButton / 2
	OTextButtunPlay := objects.NewText(I.CWTR, I.CWTG, I.CWTB, I.CWTO, x, y, I.SizeButton, "Play", I.FontText)
	TTextButtunPlay := objects.New(objects.TypeText, OTextButtunPlay)
	TTextButtunPlay.SetStatus(objects.StatusFix)
	OStyleTextButtun := objects.NewText(I.CBTR, I.CBTG, I.CBTB, I.CBTO, x, y+2, I.SizeButton, "Play", I.FontText)
	TStyleTextButtun := objects.New(objects.TypeText, OStyleTextButtun)
	TStyleTextButtun.SetStatus(objects.StatusFix)
	TButtunPlay.SetChilds(TStyleTextButtun, TTextButtunPlay)

	// Add Button reset choice player
	x = I.ScreenWidth/2 - I.WidthButton/2
	y = I.MarginTop + I.Padding + I.HeightHeaderFooter + I.UserheightBloc - (I.HeightButton)
	OButtunReset := objects.NewContainer(I.CYOR, I.CYOG, I.CYOB, I.CYOO, x, y, I.WidthButton, I.HeightButton, I.DrawFilled)
	OButtunResetOver := objects.NewContainer(I.CYR, I.CYG, I.CYB, 255, x, y, I.WidthButton, I.HeightButton, I.DrawFilled)
	OButtunResetClicDown := objects.NewContainer(I.CYR, I.CYG, I.CYB, 255, x+1, y+1, I.WidthButton-2, I.HeightButton-2, I.DrawFilled)
	TButtunReset := objects.New(objects.TypeContainer, OButtunReset)
	TButtunReset.SetObjOver(OButtunResetOver)
	TButtunReset.SetObjClicDown(OButtunResetClicDown)
	TButtunReset.Action = s.ResetChoiceToPlay
	TButtunReset.ActionDatas = append(TButtunReset.ActionDatas, dt)
	x = I.ScreenWidth / 2
	y += I.HeightButton / 2
	OTextButtunReset := objects.NewText(I.CWTR, I.CWTG, I.CWTB, I.CWTO, x, y, I.SizeButton, "reset", I.FontText)
	TTextButtunReset := objects.New(objects.TypeText, OTextButtunReset)
	TTextButtunReset.SetStatus(objects.StatusFix)
	OStyleTextButtunReset := objects.NewText(I.CBTR, I.CBTG, I.CBTB, I.CBTO, x, y+2, I.SizeButton, "reset", I.FontText)
	TStyleTextButtunReset := objects.New(objects.TypeText, OStyleTextButtunReset)
	TStyleTextButtunReset.SetStatus(objects.StatusFix)
	TButtunReset.SetChilds(TStyleTextButtunReset, TTextButtunReset)

	// signature build gomoku
	y = I.ScreenHeight - I.MarginBot - I.HeightHeaderFooter
	OContainerFooter := objects.NewContainer(I.CGLR, I.CGLG, I.CGLB, I.CGLO, I.GlobalOriginX, y, I.ScreenWidth, I.HeightHeaderFooter, I.DrawFilled)
	TContainerFooter := objects.New(objects.TypeContainer, OContainerFooter)
	TContainerFooter.SetStatus(objects.StatusFix)
	x = I.ScreenWidth - I.ScreenWidth/3
	y += I.HeightHeaderFooter - I.HeightHeaderFooter/6
	textSign :=
	OSign := objects.NewText(I.CWTR, I.CWTG, I.CWTB, I.CWTO, x, y, I.SizeInfos, textSign, I.FontText)
	TSign := objects.New(objects.TypeText, OSign)
	TSign.SetStatus(objects.StatusFix)
	TContainerFooter.Childs = append(TContainerFooter.Childs, TSign)

	s.sObjs = append(s.sObjs, TMusic)
	s.sObjs = append(s.sObjs, TBackground)
	s.sObjs = append(s.sObjs, TContainerTitle)
	s.sObjs = append(s.sObjs, TContainerUsers)
	s.sObjs = append(s.sObjs, TContainerNewUser)
	s.sObjs = append(s.sObjs, TContainerFooter)
	s.sObjs = append(s.sObjs, TButtunPlay)
	s.sObjs = append(s.sObjs, TButtunReset)
	s.sObjs = append(s.sObjs, TButtunNew)
	s.sObjs = append(s.sObjs, TButtunClearName)

	return nil
}

func (s *SUser) Update(dt *database.Data) error {
	var posX int32
	var posY int32
	var portionBlock int32

	posX = 50
	posY = 150 + I.UserSpaceUserElemListY
	// define data player by player
	for _, p := range dt.Players {
		OContainerUser := objects.NewContainer(I.CGSR, I.CGSG, I.CGSB, I.CGSO, posX, posY, I.UserWidthBloc, I.UserHeightElemList, I.DrawFilled)
		OOverContainerUser := objects.NewContainer(I.CGSOR, I.CGSOG, I.CGSOB, I.CGSOO, posX, posY, I.UserWidthBloc, I.UserHeightElemList, I.DrawFilled)
		ODownContainerUser := objects.NewContainer(I.CYOR, I.CYOG, I.CYOB, I.CYOO, posX, posY, I.UserWidthBloc, I.UserHeightElemList, I.DrawFilled)
		TContainerUser := objects.New(objects.TypeContainer, OContainerUser)
		TContainerUser.SetObjOver(OOverContainerUser)
		TContainerUser.SetObjClicDown(ODownContainerUser)
		TContainerUser.Action = s.AddCurrent
		TContainerUser.ActionDatas = append(TContainerUser.ActionDatas, p)

		x := posX + I.UserWidthBloc/2
		y := posY + I.UserHeightElemList/2
		OTextUser := objects.NewText(I.CWTR, I.CWTG, I.CWTB, I.CWTO, x, y, I.SizeNormal, p.Name, I.FontText)
		TTextUser := objects.New(objects.TypeText, OTextUser)
		TTextUser.SetStatus(objects.StatusFix)
		TContainerUser.SetChilds(TTextUser)

		x = posX + I.UserWidthBloc - (I.WidthIcon*3 + I.UserSpaceUserElemListY*2)
		y = posY
		if len(p.Saves) > 0 {
			OImgSave := objects.NewImage("disk.bmp", "disk", x, y, I.WidthIcon, I.HeightIcon)
			OImgSaveOver := objects.NewImage("diskOver.bmp", "diskOver", x, y, I.WidthIcon, I.HeightIcon)
			OImgSaveDown := objects.NewImage("diskOver.bmp", "diskOver", x+1, y+1, I.WidthIcon-2, I.HeightIcon-2)
			TImgSave := objects.New(objects.TypeImage, OImgSave)
			TImgSave.Action = s.DeleteStatUser
			TImgSave.ActionDatas = append(TImgSave.ActionDatas, p)
			TImgSave.SetObjOver(OImgSaveOver)
			TImgSave.SetObjClicDown(OImgSaveDown)
			TContainerUser.SetChilds(TImgSave)
		}

		x += I.WidthIcon + I.UserSpaceUserElemListY
		OImgTrophy := objects.NewImage("trophy.bmp", "trophy", x, y, I.WidthIcon, I.HeightIcon)
		OImgTrophyOver := objects.NewImage("trophyOver.bmp", "trophyOver", x, y, I.WidthIcon, I.HeightIcon)
		OImgTrophyDown := objects.NewImage("trophyOver.bmp", "trophyOver", x+1, y+1, I.WidthIcon-2, I.HeightIcon-2)
		TImgTrophy := objects.New(objects.TypeImage, OImgTrophy)
		TImgTrophy.Action = s.LookStats
		TImgTrophy.ActionDatas = append(TImgTrophy.ActionDatas, p)
		TImgTrophy.SetObjOver(OImgTrophyOver)
		TImgTrophy.SetObjClicDown(OImgTrophyDown)
		TContainerUser.SetChilds(TImgTrophy)

		x += I.WidthIcon + I.UserSpaceUserElemListY
		OImgDelete := objects.NewImage("delete1.bmp", "delete", x, y, I.WidthIcon, I.HeightIcon)
		OImgDeleteOver := objects.NewImage("delete2.bmp", "deleteOver", x, y, I.WidthIcon, I.HeightIcon)
		OImgDeleteDown := objects.NewImage("delete2.bmp", "deleteOver", x+1, y+1, I.WidthIcon-2, I.HeightIcon-2)
		TImgDelete := objects.New(objects.TypeImage, OImgDelete)
		TImgDelete.Action = s.DeleteUser
		TImgDelete.ActionDatas = append(TImgDelete.ActionDatas, p)
		TImgDelete.SetObjOver(OImgDeleteOver)
		TImgDelete.SetObjClicDown(OImgDeleteDown)
		TContainerUser.SetChilds(TImgDelete)

		s.dObjs = append(s.dObjs, TContainerUser)
		posY += I.UserHeightElemList + I.UserSpaceUserElemListY
	}

	// draw players selected to play the game
	posX = I.ScreenWidth/2 - I.WidthButton/2
	posY = I.MarginTop + I.Padding + I.HeightHeaderFooter
	OContainerVS := objects.NewContainer(I.CGLR, I.CGLG, I.CGLB, I.CGLO, posX, posY, I.WidthButton, I.HeightHeaderFooter, I.DrawFilled)
	TContainerVS := objects.New(objects.TypeContainer, OContainerVS)
	TContainerVS.SetStatus(objects.StatusFix)
	posX = I.ScreenWidth / 2
	portionBlock = I.HeightHeaderFooter / 6
	posY += portionBlock // 1/6 of block
	OP1 := objects.NewText(I.CWTR, I.CWTG, I.CWTB, I.CWTO, posX, posY, I.SizeNormal, dt.Current.P1.Name, I.FontText)
	TP1 := objects.New(objects.TypeText, OP1)
	TP1.SetStatus(objects.StatusFix)
	OStyleP1 := objects.NewText(I.CBTR, I.CBTG, I.CBTB, I.CBTO, posX+2, posY+2, I.SizeNormal, dt.Current.P1.Name, I.FontText)
	TStyleP1 := objects.New(objects.TypeText, OStyleP1)
	TStyleP1.SetStatus(objects.StatusFix)
	posY += portionBlock * 2 // 3/6 of block
	OVS := objects.NewText(I.CWTR, I.CWTG, I.CWTB, I.CWTO, posX, posY, I.SizeNormal, "VS", I.FontText)
	TVS := objects.New(objects.TypeText, OVS)
	TVS.SetStatus(objects.StatusFix)
	OStyleVS := objects.NewText(I.CBTR, I.CBTG, I.CBTB, I.CBTO, posX+2, posY+2, I.SizeNormal, "VS", I.FontText)
	TStyleVS := objects.New(objects.TypeText, OStyleVS)
	TStyleVS.SetStatus(objects.StatusFix)
	posY += portionBlock * 2 // 5/6 of block
	OP2 := objects.NewText(I.CWTR, I.CWTG, I.CWTB, I.CWTO, posX, posY, I.SizeNormal, dt.Current.P2.Name, I.FontText)
	TP2 := objects.New(objects.TypeText, OP2)
	TP2.SetStatus(objects.StatusFix)
	OStyleP2 := objects.NewText(I.CBTR, I.CBTG, I.CBTB, I.CBTO, posX+2, posY+2, I.SizeNormal, dt.Current.P2.Name, I.FontText)
	TStyleP2 := objects.New(objects.TypeText, OStyleP2)
	TStyleP2.SetStatus(objects.StatusFix)
	TContainerVS.SetChilds(TStyleP1, TP1, TStyleVS, TVS, TStyleP2, TP2)

	// draw create new player
	posX = I.MarginRight - (I.UserWidthBloc / 2) - (I.WidthLargeButton / 2)
	posY = I.MarginTop + I.Padding + I.HeightHeaderFooter + I.UserheightBloc - (I.UserheightBloc / 2)
	OContainerInput := objects.NewContainer(I.CWTR, I.CWTG, I.CWTB, I.CWTO, posX, posY, I.WidthLargeButton, I.HeightButton, I.DrawFilled)
	TContainerInput := objects.New(objects.TypeContainer, OContainerInput)
	OInputOver := objects.NewContainer(I.CWTR, I.CWTG, I.CWTB, I.CWTO, posX, posY, I.WidthLargeButton, I.HeightButton, I.DrawFilled)
	OInputClickDown := objects.NewContainer(I.CWTR, I.CWTG, I.CWTB, I.CWTO, posX, posY, I.WidthLargeButton, I.HeightButton, I.DrawFilled)
	TContainerInput.SetObjOver(OInputOver)
	TContainerInput.SetObjClicDown(OInputClickDown)
	TContainerInput.Action = s.InitInputText
	TContainerInput.ActionDatas = append(TContainerInput.ActionDatas, objects.ParamPosition{posX, posY})
	TContainerInput.ActionDatas = append(TContainerInput.ActionDatas, objects.ParamSize{I.WidthLargeButton, I.HeightButton})
	posX += I.WidthLargeButton / 2
	posY += I.HeightButton / 2
	OInput := objects.NewText(I.CRR, I.CRG, I.CRB, I.CRO, posX, posY, I.SizeNormal, s.input, I.FontText)
	TInput := objects.New(objects.TypeText, OInput)
	OStyleInput := objects.NewText(I.CBTR, I.CBTG, I.CBTB, I.CBTO, posX+1, posY+1, I.SizeNormal, s.input, I.FontText)
	TStyleInput := objects.New(objects.TypeText, OStyleInput)
	TContainerInput.SetChilds(TStyleInput, TInput)

	// draw notice text
	posX = I.ScreenWidth / 2
	posY = I.ScreenHeight - I.MarginBot - I.HeightHeaderFooter
	posY += I.HeightHeaderFooter / 3
	ONoticeText := objects.NewText(I.CRR, I.CRG, I.CRB, I.CRO, posX, posY, I.SizeNormal, s.notice, I.FontText)
	TNoticeText := objects.New(objects.TypeText, ONoticeText)
	//ClearNotice
	s.notice = ""

	s.dObjs = append(s.dObjs, TContainerVS)
	s.dObjs = append(s.dObjs, TContainerInput)
	s.dObjs = append(s.dObjs, TNoticeText)

	return nil
}

func (S *SUser) SetNotice(str string) {
	S.notice = str
}

func (S *SUser) GetObjs() []*objects.ObjectType {
	objs := append(S.sObjs, S.dObjs...)
	return objs
}

func (S *SUser) AddLetterToInput(txt string) {
	S.input += txt
}

func (S *SUser) GetStaticObjs() []*objects.ObjectType {
	return S.sObjs
}

func (S *SUser) GetDynamicObjs() []*objects.ObjectType {
	return S.dObjs
}

// must be define on Scene package
/*func (S *Scene) ProcessUpload() {
	S.DataObjs = S.UploadUser()
	x, y, _ := sdl.GetMouseState()
	defineIsOver(S.DataObjs, int32(x), int32(y))
}*/
