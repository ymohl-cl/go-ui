package conf

/*
** This package define the configuration to the scenes and the scenes list.
** Define everything you need.
 */

// Current is a global value to define the current scene.
var Current uint8

// Scenes list
const (
	SMenu = 0
	SStat = 1
	SGame = 2
	Sload = 3
)

// Global
const (
	Title        = "Gomoku"
	WindowWidth  = 1280
	WindowHeight = 800
	ressources   = "Ressources/"
	ProtoBufFile = ressources + "saveGomoku.patouch"
	OriginX      = 0
	OriginY      = 0

	/* margin - padding */
	MarginTop    = 31
	MarginLeft   = 50
	MarginRight  = 50
	MarginBot    = 29
	PaddingBlock = 18

	/* Color */
	ClearRGBO = 0

	ColorBlockRed     = 42
	ColorBlockGreen   = 42
	ColorBlockBlue    = 42
	ColorBlockOpacity = 155
)

// Menu
const (
	MenuMusic      = ressources + "ambiant.wav"
	MenuBackground = ressources + "background.bmp"

	/* position elements */
	MenuHeaderHeight       = 102
	MenuFooterHeight       = 102
	MenuContentBlockWidth  = 500
	MenuContentBlockHeight = 500
)

// Stat
const ()

// Game
const ()

// Load
const ()

/*
** define design
 */

// Buttun
const ()

// Text
const ()

//////////////////////////////////////////////////////////
const (
	/*
	** list scenes on the uint8 values. Define by current of Scenes while
	** the game is running
	 */
	//SMenu     = 0
	//SceneStat = 1
	//SceneGame = 2

	/* view const (global) */
	FontText      = "leadcoat.ttf"
	GlobalOriginX = 0
	GlobalOriginY = 0
	//	HeightHeaderFooter = 105
	WidthButton      = 150
	WidthLargeButton = 300
	HeightButton     = 40
	WidthIcon        = 30
	HeightIcon       = 30
	WidthInput       = 300
	HeightInput      = 30

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
	UserHeightElemList     = 30
	UserSpaceUserElemListY = 5
)
