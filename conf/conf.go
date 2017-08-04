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

	ColorInputRed     = 255
	ColorInputGreen   = 255
	ColorInputBlue    = 255
	ColorInputOpacity = 150

	ColorOverInputRed     = 255
	ColorOverInputGreen   = 255
	ColorOverInputBlue    = 255
	ColorOverInputOpacity = 255

	ColorClickInputRed     = 240
	ColorClickInputGreen   = 240
	ColorClickInputBlue    = 240
	ColorClickInputOpacity = 250
)

// Menu
const (
	MenuMusic          = ressources + "menuambiant.wav"
	MenuBackground     = ressources + "background.bmp"
	MenuIconDelete     = ressources + "delete1.bmp"
	MenuIconOverDelete = ressources + "delete2.bmp"
	MenuIconTrophy     = ressources + "trophy.bmp"
	MenuIconOverTrophy = ressources + "trophyOver.bmp"
	MenuIconLoad       = ressources + "disk.bmp"
	MenuIconOverLoad   = ressources + "diskOver.bmp"

	/* size elements */
	MenuHeaderHeight             = 102
	MenuFooterHeight             = 102
	MenuContentBlockWidth        = 500
	MenuContentLargeBlockHeight  = 500
	MenuContentMediumBlockHeight = (MenuContentLargeBlockHeight - PaddingBlock) / 2
	MenuElementPlayerHeight      = 30
	MenuElementPlayerWidth       = 395
	MenuElementPadding           = 5
	MenuIconWidth                = 30
)

// Stat
const ()

// Game
const (
	GameMusic = ressources + "game_ambiant.wav"
)

// Load
const (
	LoadMusic       = ressources + "ambiant.wav"
	LoadBlockWidth  = 20
	LoadBlockHeight = 20
)

/*
** define design
 */

// Buttun
const (
	ButtonHeight = 40
	ButtonWidth  = 150

	/* Color */
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
)

// Text
const (
	Font            = ressources + "leadcoat.ttf"
	TxtLittle       = 14
	TxtMedium       = 20
	TxtLarge        = 78
	TxtUnderPadding = 1

	/* Color */
	// Color text white
	ColorTxtRed     = 255
	ColorTxtGreen   = 255
	ColorTxtBlue    = 255
	ColorTxtOpacity = 255

	// Color text black
	ColorUnderTxtRed     = 0
	ColorUnderTxtGreen   = 0
	ColorUnderTxtBlue    = 0
	ColorUnderTxtOpacity = 20
)
