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

	ColorButtonRed     = 95
	ColorButtonGreen   = 74
	ColorButtonBlue    = 29
	ColorButtonOpacity = 200

	ColorOverButtonRed     = 116
	ColorOverButtonGreen   = 108
	ColorOverButtonBlue    = 25
	ColorOverButtonOpacity = 255

	ColorClickButtonRed     = 21
	ColorClickButtonGreen   = 34
	ColorClickButtonBlue    = 33
	ColorClickButtonOpacity = 255

	ColorTxtRed     = 255
	ColorTxtGreen   = 255
	ColorTxtBlue    = 255
	ColorTxtOpacity = 255

	ColorUnderTxtRed     = 0
	ColorUnderTxtGreen   = 0
	ColorUnderTxtBlue    = 0
	ColorUnderTxtOpacity = 20

	/* Text */
	Font      = ressources + "leadcoat.ttf"
	TxtLittle = 14
	TxtMedium = 20
	TxtLarge  = 78

	/* size elements */
	ButtonHeight = 40
	ButtonWidth  = 150
)

// Menu
const (
	MenuMusic      = ressources + "ambiant.wav"
	MenuBackground = ressources + "background.bmp"

	/* size elements */
	MenuHeaderHeight             = 102
	MenuFooterHeight             = 102
	MenuContentBlockWidth        = 500
	MenuContentLargeBlockHeight  = 500
	MenuContentMediumBlockHeight = (MenuContentLargeBlockHeight - PaddingBlock) / 2
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
	GlobalOriginX = 0
	GlobalOriginY = 0
	//	HeightHeaderFooter = 105
	WidthLargeButton = 300
	WidthIcon        = 30
	HeightIcon       = 30
	WidthInput       = 300
	HeightInput      = 30

	DrawFilled    = true
	DrawNotFilled = false

	/* Size text */

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

	/* Color Red - Red - Green - Blue - Opacity */
	CRR = 133
	CRG = 6
	CRB = 6
	CRO = 255

	/* view const (user) */
	UserHeightElemList     = 30
	UserSpaceUserElemListY = 5
)
