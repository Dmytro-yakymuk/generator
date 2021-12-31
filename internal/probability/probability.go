package probability

var Heads = map[string]int{
	// 35 - 1000
	// 9 - 116
	"DIAMOND-HEAD":    3,
	"IRIDESCENT-HEAD": 5,
	"ORB-HEAD":        8,
	"CRACKED-HEAD":    10,
	"ZOMBIE-HEAD":     12,
	"HELLO-HEAD":      15,
	"DAMASCUS-HEAD":   18,
	"HEART-HEAD":      20,
	"WOOD-HEAD":       25,

	// 26 - 884
	"H1-HEAD":               34,
	"H2-HEAD":               34,
	"H3-HEAD":               34,
	"H4-HEAD":               34,
	"CARBON-HEAD":           34,
	"SMILEY-HEAD":           34,
	"PAINT-HEAD":            34,
	"METALROUGH-HEAD":       34,
	"GOONGANG-HEAD":         34,
	"WAVES-HEAD":            34,
	"GREYSCALE-WAVES-HEAD":  34,
	"STICKERBOMB-HEAD":      34,
	"TAG-HEAD":              34,
	"TAG-PURPLE-HEAD":       34,
	"GOONZ-PURPLE-HEAD":     34,
	"GOONZ-YELLOW-HEAD":     34,
	"TERRAZZO-PURPLE-HEAD":  34,
	"SPLATTER-BLUE-HEAD":    34,
	"CONCRETE-HEAD":         34,
	"GLITTER-BLACK-HEAD":    34,
	"GREY-HEAD":             34,
	"TERRAZZO-BLUE-HEAD":    34,
	"GLITTER-PURPLE-HEAD":   34,
	"WORNSTEEL-HEAD":        34,
	"POLISHEDCONCRETE-HEAD": 34,
	"SILVER-GLITTER-HEAD":   34,
}

var Tshirts = map[string]int{
	// 28 - 1000
	// 13 - 325
	"MRE-TSHIRT":          10,
	"TIGER-TSHIRT":        12,
	"REAPER-TSHIRT":       15,
	"GOONZTAG-TSHIRT":     18,
	"GG-TSHIRT":           20,
	"GOONGANG-TSHIRT":     22,
	"SKULLSNAKE-TSHIRT":   25,
	"GOONPREME-TSHIRT":    28,
	"TIEDYE-TSHIRT":       30,
	"TIEDYE-WAVES-TSHIRT": 32,
	"SID-TSHIRT":          35,
	"GOONSTATION-TSHIRT":  38,
	"PG-TSHIRT":           40,

	// 15 - 675
	"CAMO-GREEN-TSHIRT":     45,
	"SNOWCAMO-TSHIRT":       45,
	"DIGICAMO-TSHIRT":       45,
	"CAMO-PURPLE-TSHIRT":    45,
	"GOONZ-LOGO-TSHIRT":     45,
	"BLACK-PLAIN-TSHIRT":    45,
	"GOONZ-TSHIRT":          45,
	"POKEGOON-TSHIRT":       45,
	"PGC-TSHIRT":            45,
	"FRIENDS-TSHIRT":        45,
	"GOONZSQUAD-TSHIRT":     45,
	"STRIPE-TSHIRT":         45,
	"PGLONDON-TSHIRT":       45,
	"GOONZ-BASEBALL-TSHIRT": 45,
	"PLAIN-TSHIRT":          45,
}

var Necklaces = map[string]int{
	// 13 - 1000
	"BITCOIN-BIG-CHAIN":        15,
	"DIAMOND-HEAD-CHAIN":       25,
	"DIAMOND-PG-CHAIN":         35,
	"IWATCH-CHAIN":             45,
	"GOONZ-GOLD-CHAIN":         55,
	"WEED-CHAIN":               65,
	"PG-SMALL-SILVER-CHAIN":    75,
	"PG-SMALL-GOLD-CHAIN":      85,
	"BITCOIN-GOLD-SMALL-CHAIN": 95,
	"GOONZLOGO-GOLD-CHAIN":     105,
	"GOONZLOGO-SILVER-CHAIN":   120,
	"CUBANGOLD-CHAIN":          130,
	"GOLDLINK-CHAIN":           150,
}

var Jackets = map[string]int{
	// 12 - 1000
	// 3 - 154
	"SPIKES-JACKET":   40,
	"GARDA-JACKET":    52,
	"VERCETTI-JACKET": 62,

	// 9 - 846
	"URBAN-CAMO-JACKET":      94,
	"GLITTERBALL-JACKET":     94,
	"LOUIS-JACKET":           94,
	"DIGI-CAMO-JACKET":       94,
	"TERRAZZO-PINK-JACKET":   94,
	"TERRAZZO-PURPLE-JACKET": 94,
	"SUEDE-JACKET":           94,
	"BLUEKNIT-JACKET":        94,
	"RED-SILK-JACKET":        94,
}

var Updo string = "MOHAWK-HAIR"
var Hairs = map[string]int{
	// 26 - 1000
	// 18 - 344
	"FLAME-PURPLE-HAIR": 3,
	"FLAME-HAIR":        5,
	"HEART-HAIR":        7,
	"TIGERORANGE-HAIR":  10,
	"TIGERRED-HAIR":     12,
	"LEOPARD-HAIR":      15,
	"MARY-HAIR":         17,
	"SKULL-HAIR":        20,
	"SKULL-INVERT-HAIR": 21,
	"BLACK-FADE":        22,
	"NEON-HAIR":         23,
	"ASTRO-HAIR":        24,
	"GREYSCALE-HAIR":    25,
	"FRUITY-HAIR":       26,
	"PLAYDOUGH-HAIR":    27,
	"LAVALAMP-HAIR":     28,
	"ALIEN-HAIR":        29,
	"PINKDOT-HAIR":      30,

	// 8 - 656
	"COW":            82,
	"SULLY-HAIR":     82,
	"SYRUP-HAIR":     82,
	"STARLIGHT-HAIR": 82,
	"SKY-HAIR":       82,
	"BUBBLEGUM-HAIR": 82,
	"CARPET-HAIR":    82,
	"MOHAWK-HAIR":    82,
}

var Hats = map[string]int{
	// 33 - 1000
	// 26 - 622
	"HELECOPTER-HELICOPTER":   3,
	"SQUIGGLE-HAT":            5,
	"JACKSON-HAT":             8,
	"BLACK-LOGO-HAT-REVERSE":  10,
	"BLACK-FLIP-HAT":          12,
	"BLACK-GOONZ-HAT":         15,
	"BLACK-HAT-TAG":           18,
	"BLACK-LOGO-HAT-FRONT":    20,
	"BANDANA-PAISLEY":         21,
	"GARDA-BUCKET-HAT":        22,
	"SHERIFF-HAT":             23,
	"BANDANA-GREENCAMO":       24,
	"CAMO-VISOR":              25,
	"WHITE-SWEAT":             26,
	"RED-SWEAT":               27,
	"BLACK-SWEAT":             28,
	"BLUE-SWEAT":              29,
	"LEATHER-VISOR":           30,
	"CHROME-RED-HAT-REVERSE":  31,
	"BLUE-BASEBALL-FRONT":     32,
	"CHROME-BLUE-HAT-REVERSE": 33,
	"PURPLE-BLOB-HAT":         34,
	"PINK-BLOB-HAT":           35,
	"PURPLE-BASEBALL-FRONT":   36,
	"PURPLE-STRIPE-BACKWARD":  37,
	"PURPLE-SUEDE-VISOR":      38,

	// 7 - 378
	"SUEDE-VISOR":              57,
	"BLACK-FEDORA":             57,
	"CHROME-GREEN-HAT-REVERSE": 57,
	"GOONZ-LOGO-HAT":           57,
	"MC'D-RED-VISOR":           57,
	"BLACK-SILVER-GEO-HAT":     57,
	"BANDANA-SNOW":             57,
}

var Glasses = map[string]int{
	// 14 - 1000
	"HEART-SHADES":    10,
	"3D-SHADES":       20,
	"BITCOIN-SHADES":  30,
	"ETH-SHADES":      40,
	"GOOGLY-TRANCE":   50,
	"GOOGLY-LOOK":     60,
	"GOOGLY-CROSS":    70,
	"GOOGLY-STONED":   80,
	"BLACKOUT-SHADES": 90,
	"BLACK-SHADE":     95,
	"RED-SHADE":       100,
	"MEME-SHADE":      110,
	"SMILEY-SHADES":   120,
	"FRAMES":          125,
}

var Earrings = map[string]int{
	// 4 - 1000
	"DIAMOND-STUD": 100,
	"GOLD-BITCOIN": 200,
	"ETH-SILVER":   300,
	"ETH-BLACK":    400,
}

//
//
//
//
//
// 2 ////////////////////////////////
var Tshirts2 = map[string]int{
	// 14 - 1000
	"1":               10,
	"10":              20,
	"11":              30,
	"14":              40,
	"12":              50,
	"goonz black0026": 60,
	"13":              70,
	"5":               80,
	"2":               90,
	"6":               95,
	"7":               100,
	"3":               110,
	"8":               120,
	"9":               125,
}

//
//
//
//
//
// 3 ////////////////////////////////
var Tshirts3 = map[string]int{
	// 3 - 1000
	"weed":       200,
	"swag2":      300,
	"red-stripe": 500,
}

//
//
//
//
//
// 4 ////////////////////////////////
var Tshirts4 = map[string]int{
	// 1 - 1000
	"layer 11": 1000,
}

//
//
//
//
//
// 5 ////////////////////////////////
var Tshirts5 = map[string]int{
	// 4 - 1000
	"2_0026": 100,
	"1_0026": 200,
	"3_0026": 300,
	"4_0026": 400,
}

//
//
//
//
//
//
// 6 ////////////////////////////////
var Tshirts6 = map[string]int{
	// 2 - 1000
	"SUEDE RED":     350,
	"SUEDE BLACK 2": 650,
}

var Hairs6 = map[string]int{
	// 25 - 1000
	// 18 - 349
	"FLAME-PURPLE-HAIR": 4,
	"FLAME-HAIR":        6,
	"HEART-HAIR":        8,
	"TIGERORANGE-HAIR":  10,
	"TIGERRED-HAIR":     13,
	"LEOPARD-HAIR":      15,
	"MARY-HAIR":         18,
	"SKULL-HAIR":        20,
	"SKULL-INVERT-HAIR": 21,
	"BLACK-FADE":        22,
	"NEON-HAIR":         23,
	"ASTRO-HAIR":        24,
	"GREYSCALE-HAIR":    25,
	"FRUITY-HAIR":       26,
	"PLAYDOUGH-HAIR":    27,
	"LAVALAMP-HAIR":     28,
	"ALIEN-HAIR":        29,
	"PINKDOT-HAIR":      30,

	// 7 - 651
	"COW":            93,
	"SULLY-HAIR":     93,
	"SYRUP-HAIR":     93,
	"STARLIGHT-HAIR": 93,
	"SKY-HAIR":       93,
	"BUBBLEGUM-HAIR": 93,
	"CARPET-HAIR":    93,
}

var Hats6 = map[string]int{
	// 33 - 1000
	"1_0022":         125,
	"layer 6":        125,
	"layer 6 copy":   125,
	"layer 6 copy 2": 125,
	"layer 7":        125,
	"layer 7 copy":   125,
	"layer 7 copy 2": 125,
	"layer 7 copy 3": 125,
}

var Glasses6 = map[string]int{
	// 12 - 1000
	"layer 4 copy":       10,
	"bitcoin0014":        30,
	"eth0014":            50,
	"google--trance0014": 60,
	"google-left0014":    70,
	"googly--cross0014":  80,
	"googly--stoned0014": 90,
	"layer1 copy":        100,
	"layer 2 copy":       110,
	"layer 2":            120,
	"layer 4":            130,
	"smiley0014":         150,
}

//
//
//
//
//
//
// 7 ////////////////////////////////
var Tshirts7 = map[string]int{
	// 2 - 1000
	"1_0026": 500,
	"3_0026": 500,
}

var Hats7 = map[string]int{
	// 1 - 1000
	"LAYER 13": 1000,
}

var Glasses7 = map[string]int{
	// 8 - 1000 + 400
	"layer 4 copy":       50,
	"google--trance0014": 75,
	"google-left0014":    100,
	"googly-cross0014":   125,
	"googy--stoned0014":  150,
	"layer 1 copty":      155,
	"layer 4":            165,
	"smiley0014":         180,
}
