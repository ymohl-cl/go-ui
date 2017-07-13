package sinfos

// This package define the parameters to the scenes and the scenes list.

const (
	/*
	** list scenes on the uint8 values. Define by current of Scenes while
	** the game is running
	 */
	SceneUser = 0
	SceneStat = 1
	SceneGame = 2

	/* view const (global) */
	ScreenWidth        = 1280
	ScreenHeight       = 800
	FontText           = "leadcoat.ttf"
	GlobalOriginX      = 0
	GlobalOriginY      = 0
	MarginTop          = 30
	MarginLeft         = 50
	MarginRight        = ScreenWidth - MarginLeft
	MarginBot          = 25
	Padding            = 15
	HeightHeaderFooter = 105
	WidthButton        = 150
	WidthLargeButton   = 300
	HeightButton       = 40
	WidthIcon          = 30
	HeightIcon         = 30
	WidthInput         = 300
	HeightInput        = 30

	DrawFilled    = true
	DrawNotFilled = false

	/* Size text */
	SizeTitle  = 78
	SizeButton = 20
	SizeNormal = 20
	SizeInfos  = 16

	/* Color White Text - Red - Green - Blue - Opacity */
	CWTR = 255
	CWTG = 255
	CWTB = 255
	CWTO = 255
	/* Color Black Text - Red - Green - Blue - Opacity */
	CBTR = 0
	CBTG = 0
	CBTB = 0
	CBTO = 20

	/* Color Grey Light - Red - Green - Blue - Opacity */
	CGLR = 42
	CGLG = 42
	CGLB = 42
	CGLO = 155

	/* Color Grey strong - Red - Green - Blue - Opacity */
	CGSR = 57
	CGSG = 57
	CGSB = 57
	CGSO = 155
	/* Color Grey strong Over - Red - Green - Blue - Opacity */
	CGSOR = 67
	CGSOG = 67
	CGSOB = 67
	CGSOO = 155

	/* Color Yellow - Red - Green - Blue - Opacity */
	CYR = 95
	CYG = 74
	CYB = 29
	CYO = 200

	/* Color Yellow Over - Red - Green - Blue - Opacity */
	CYOR = 116
	CYOG = 108
	CYOB = 25
	CYOO = 255

	/* Color Red - Red - Green - Blue - Opacity */
	CRR = 133
	CRG = 6
	CRB = 6
	CRO = 255

	/* view const (user) */
	UserWidthBloc          = 500
	UserheightBloc         = 500
	UserHeightElemList     = 30
	UserSpaceUserElemListY = 5
)

var Current uint8
