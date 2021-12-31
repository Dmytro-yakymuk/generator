// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardavatars

// Avatar entity describes the values that make up the avatar.
type Avatar struct {
	Background int    `json:"background"`
	Heads      string `json:"heads"`
	Tshirts    string `json:"tshirts"`
	Necklace   string `json:"necklace"`
	Jacket     string `json:"jacket"`
	Hair       string `json:"hair"`
	Headwear   string `json:"headwear"`
	Glasses    string `json:"glasses"`
	Earrings   string `json:"earrings"`
}

// Config defines values needed by generate avatars.
type Config struct {
	PathToAvararsComponents string `json:"pathToAvararsComponents"`
	PathToOutputAvarars     string `json:"pathToOutputAvatars"`
	PathToOutputJSON        string `json:"pathToOutputJSON"`

	BackgroundFolder string `json:"backgroundFolder"`
	HeadsFolder      string `json:"headsFolder"`
	TshirtsFolder    string `json:"tshirtsFolder"`
	NecklaceFolder   string `json:"necklaceFolder"`
	JacketFolder     string `json:"jacketFolder"`
	HairFolder       string `json:"hairFolder"`
	HeadwearFolder   string `json:"headwearFolder"`
	GlassesFolder    string `json:"glassesFolder"`
	EarringsFolder   string `json:"earringsFolder"`

	PercentageHeadwear int `json:"percentageHeadwear"`
	PercentageNecklace int `json:"percentageNecklace"`
	PercentageGlasses  int `json:"percentageGlasses"`
	PercentageEarrings int `json:"percentageEarrings"`

	Description string `json:"description"`
	ExternalURL string `json:"externalUrl"`
	Image       string `json:"image"`
	PrefixName  string `json:"prefixName"`
}

// TypeImage defines the list of possible type of avatar image.
type TypeImage string

const (
	// TypeImagePNG indicates that the type image avatar is png.
	TypeImagePNG TypeImage = "png"
)
