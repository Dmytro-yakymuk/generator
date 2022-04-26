// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardavatars

// Avatar entity describes the values that make up the avatar.
type Avatar struct {
	Background int `json:"background"`
	Back       int `json:"back"`
	Stand      int `json:"stand"`
	Shell      int `json:"shell"`
	Eye        int `json:"eye"`
	Accessorie int `json:"accessorie"`
	Hat        int `json:"hat"`
	Glasse     int `json:"glasse"`
	Mouth      int `json:"mouth"`
}

// Metadata ...
type Metadata struct {
	Background string `json:"background"`
	Back       string `json:"back"`
	Stand      string `json:"stand"`
	Shell      string `json:"shell"`
	Eye        string `json:"eye"`
	Accessorie string `json:"accessorie"`
	Hat        string `json:"hat"`
	Glasse     string `json:"glasse"`
	Mouth      string `json:"mouth"`
}

// Config defines values needed by generate avatars.
type Config struct {
	PathToAvararsComponents string `json:"pathToAvararsComponents"`
	PathToOutputAvarars     string `json:"pathToOutputAvatars"`
	PathToOutputJSON        string `json:"pathToOutputJSON"`

	BackgroundsFolder string `json:"backgroundsFolder"`
	BacksFolder       string `json:"backsFolder"`
	StandsFolder      string `json:"standsFolder"`
	ShellsFolder      string `json:"shellsFolder"`
	EyesFolder        string `json:"eyesFolder"`
	AccessoriesFolder string `json:"accessoriesFolder"`
	HatsFolder        string `json:"hatsFolder"`
	GlassesFolder     string `json:"glassesFolder"`
	MouthPropsFolder  string `json:"mouthPropsFolder"`

	Mouths []int `json:"mouths"`

	Description string `json:"description"`
	ExternalURL string `json:"externalUrl"`
	Image       string `json:"image"`
	Name        string `json:"name"`
}

// NFT entity describes nft token format erc-721.
type NFT struct {
	Description string      `json:"description"`
	ExternalURL string      `json:"external_url"`
	Image       string      `json:"image"`
	Name        string      `json:"name"`
	Attributes  []Attribute `json:"attributes"`
}

// Attribute entity describes attributes for the item, which will show up on the OpenSea page for the item.
type Attribute struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
	MaxValue  interface{} `json:"max_value,omitempty"`
}

// TypeImage defines the list of possible type of avatar image.
type TypeImage string

const (
	// TypeImagePNG indicates that the type image avatar is png.
	TypeImagePNG TypeImage = "png"
)
