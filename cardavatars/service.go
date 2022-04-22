// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardavatars

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"

	"github.com/zeebo/errs"

	"generator/pkg/imageprocessing"
)

// ErrCardAvatars indicated that there was an error in service.
var ErrCardAvatars = errs.Class("card with link to avatar service error")

// Service is handling cards with link to avatars related logic.
//
// architecture: Service
type Service struct {
	config Config
}

// NewService is a constructor for card with link to avatar service.
func NewService(config Config) *Service {
	return &Service{config: config}
}

// Generate generates cards with avatar link.
func (service *Service) Generate(ctx context.Context, start int, end int) error {
	var stringComponents []string
	i := start
	for i <= end {
		var (
			avatar Avatar
			layer  image.Image
			layers []image.Image
			err    error
		)

		// Backgrounds
		pathToBackground := filepath.Join(service.config.PathToAvararsComponents, service.config.BackgroundsFolder)
		count, err := imageprocessing.LayerComponentsCount(pathToBackground)
		if err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		avatar.Background = rand.Intn(count) + 1

		if layer, err = imageprocessing.CreateLayer(pathToBackground, avatar.Background); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		layers = append(layers, layer)

		stringComponent := fmt.Sprintf("%d. %d || %d || %d || %d || %d || %d || %d || %d || %d", i, avatar.Background, avatar.Back, avatar.Stand,
			avatar.Shell, avatar.Eye, avatar.Accessorie, avatar.Hat, avatar.Glasse, avatar.Mouth)
		for _, stringComp := range stringComponents {
			if stringComp == stringComponent {
				break
			}
		}

		originalAvatar := imageprocessing.Layering(layers, 0, 0)

		if err = imageprocessing.SaveImage(service.config.PathToOutputAvarars, filepath.Join(service.config.PathToOutputAvarars, strconv.Itoa(i)+"."+string(TypeImagePNG)), originalAvatar); err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		// if err = service.GenerateNFT(avatar, i); err != nil {
		// 	return err
		// }

		stringComponents = append(stringComponents, stringComponent)
		fmt.Println(i)
		i++
	}

	file, err := json.MarshalIndent(stringComponents, "", " ")
	if err != nil {
		return err
	}

	if err = os.MkdirAll(service.config.PathToOutputJSON, os.ModePerm); err != nil {
		return err
	}

	if err = ioutil.WriteFile(filepath.Join(service.config.PathToOutputJSON, "components.json"), file, 0644); err != nil {
		return err
	}

	return nil
}

// // Generate generates values for nft token.
// func (service *Service) GenerateNFT(avatar Avatar, sequenceNumber int) error {
// 	var attributes []Attribute

// 	if avatar.Background != "" {
// 		attributes = append(attributes, Attribute{TraitType: "BACKGROUND", Value: avatar.Background})
// 	}
// 	if avatar.Heads != "" {
// 		attributes = append(attributes, Attribute{TraitType: "HEADS", Value: avatar.Heads})
// 	}
// 	if avatar.Tshirts != "" {
// 		attributes = append(attributes, Attribute{TraitType: "TSHIRT", Value: avatar.Tshirts})
// 	}
// 	if avatar.Necklace != "" {
// 		attributes = append(attributes, Attribute{TraitType: "NECKLACE", Value: avatar.Necklace})
// 	}
// 	if avatar.Jacket != "" {
// 		attributes = append(attributes, Attribute{TraitType: "JACKET", Value: avatar.Jacket})
// 	}
// 	if avatar.Hair != "" {
// 		attributes = append(attributes, Attribute{TraitType: "HAIR", Value: avatar.Hair})
// 	}
// 	if avatar.Headwear != "" {
// 		attributes = append(attributes, Attribute{TraitType: "HATS", Value: avatar.Headwear})
// 	}
// 	if avatar.Glasses != "" {
// 		attributes = append(attributes, Attribute{TraitType: "GLASSES", Value: avatar.Glasses})
// 	}
// 	if avatar.Earrings != "" {
// 		attributes = append(attributes, Attribute{TraitType: "EARRING", Value: avatar.Earrings})
// 	}

// 	nft := NFT{
// 		Attributes:  attributes,
// 		Description: service.config.Description,
// 		ExternalURL: service.config.ExternalURL,
// 		Image:       fmt.Sprintf(service.config.Image, sequenceNumber),
// 		Name:        service.config.Name,
// 	}

// 	file, err := json.MarshalIndent(nft, "", " ")
// 	if err != nil {
// 		return err
// 	}

// 	if err := os.MkdirAll(service.config.PathToOutputAvarars, os.ModePerm); err != nil {
// 		return err
// 	}

// 	return ioutil.WriteFile(filepath.Join(service.config.PathToOutputAvarars, fmt.Sprintf("%d.json", sequenceNumber)), file, 0644)
// }
