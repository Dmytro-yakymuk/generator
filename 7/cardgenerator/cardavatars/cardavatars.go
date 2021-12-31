// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardavatars

// Avatar entity describes the values that make up the avatar.
type Avatar struct {
	Background int    `json:"background"`
	Heads      string `json:"heads"`
	Tshirts    string `json:"tshirts"`
	Headwear   string `json:"headwear"`
	Glasses    string `json:"glasses"`
}

// Config defines values needed by generate avatars.
type Config struct {
	PathToAvararsComponents string `json:"pathToAvararsComponents"`
	PathToOutputAvarars     string `json:"pathToOutputAvatars"`
	PathToOutputJSON        string `json:"pathToOutputJSON"`

	BackgroundFolder string `json:"backgroundFolder"`
	HeadsFolder      string `json:"headsFolder"`
	TshirtsFolder    string `json:"tshirtsFolder"`
	HeadwearFolder   string `json:"headwearFolder"`
	GlassesFolder    string `json:"glassesFolder"`

	PercentageGlasses int `json:"percentageGlasses"`
}

// TypeImage defines the list of possible type of avatar image.
type TypeImage string

const (
	// TypeImagePNG indicates that the type image avatar is png.
	TypeImagePNG TypeImage = "png"
)
