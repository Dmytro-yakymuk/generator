// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardavatars

// Avatar entity describes the values that make up the avatar.
type Avatar struct {
	Background int `json:"background"`
	Heads      int `json:"heads"`
	Tshirts    int `json:"tshirts"`
	Hair       int `json:"hair"`
	Headwear   int `json:"headwear"`
	Glasses    int `json:"glasses"`
	Earrings   int `json:"earrings"`
}

// Config defines values needed by generate avatars.
type Config struct {
	PathToAvararsComponents string `json:"pathToAvararsComponents"`
	PathToOutputAvarars     string `json:"pathToOutputAvatars"`
	PathToOutputJSON        string `json:"pathToOutputJSON"`

	BackgroundFolder string `json:"backgroundFolder"`
	HeadsFolder      string `json:"headsFolder"`
	TshirtsFolder    string `json:"tshirtsFolder"`
	HairFolder       string `json:"hairFolder"`
	HeadwearFolder   string `json:"headwearFolder"`
	GlassesFolder    string `json:"glassesFolder"`
	EarringsFolder   string `json:"earringsFolder"`

	PercentageHeadwear int `json:"percentageHeadwear"`
	WithoutHair        int `json:"withoutHair"`
	PercentageGlasses  int `json:"percentageGlasses"`
	PercentageEarrings int `json:"percentageEarrings"`
}

// TypeImage defines the list of possible type of avatar image.
type TypeImage string

const (
	// TypeImagePNG indicates that the type image avatar is png.
	TypeImagePNG TypeImage = "png"
)
