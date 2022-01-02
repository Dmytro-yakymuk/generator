// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardavatars

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/zeebo/errs"

	"boonji/internal/probability"
	"boonji/pkg/imageprocessing"
	"boonji/pkg/rand"
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

		// Background
		pathToBackground := filepath.Join(service.config.PathToAvararsComponents, service.config.BackgroundFolder)
		avatar.Background = rand.SearchValueByPercent(probability.Background)

		if layer, err = imageprocessing.CreateLayerByFileName(pathToBackground, avatar.Background); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		layers = append(layers, layer)

		// Heads
		pathToHeads := filepath.Join(service.config.PathToAvararsComponents, service.config.HeadsFolder)
		avatar.Heads = rand.SearchValueByPercent(probability.Heads)

		if layer, err = imageprocessing.CreateLayerByFileName(pathToHeads, avatar.Heads); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		layers = append(layers, layer)

		// Tshirts
		pathToTshirts := filepath.Join(service.config.PathToAvararsComponents, service.config.TshirtsFolder)
		avatar.Tshirts = rand.SearchValueByPercent(probability.Tshirts7)

		if layer, err = imageprocessing.CreateLayerByFileName(pathToTshirts, avatar.Tshirts); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		layers = append(layers, layer)

		// Headwear
		pathToHeadwear := filepath.Join(service.config.PathToAvararsComponents, service.config.HeadwearFolder)
		avatar.Headwear = rand.SearchValueByPercent(probability.Hats7)

		if layer, err = imageprocessing.CreateLayerByFileName(pathToHeadwear, avatar.Headwear); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		layers = append(layers, layer)

		// Glasses
		if rand.IsIncludeRange(service.config.PercentageGlasses) {
			pathToGlasses := filepath.Join(service.config.PathToAvararsComponents, service.config.GlassesFolder)
			avatar.Glasses = rand.SearchValueByPercent(probability.Glasses7)

			if layer, err = imageprocessing.CreateLayerByFileName(pathToGlasses, avatar.Glasses); err != nil {
				return ErrCardAvatars.Wrap(err)
			}
			layers = append(layers, layer)
		}

		stringComponent := fmt.Sprintf("%d. %v || %v || %v || %v || %v", i, avatar.Background, avatar.Heads, avatar.Tshirts, avatar.Headwear, avatar.Glasses)
		for _, stringComp := range stringComponents {
			if stringComp == stringComponent {
				break
			}
		}

		originalAvatar := imageprocessing.Layering(layers, 0, 0)

		if err = imageprocessing.SaveImage(service.config.PathToOutputAvarars, filepath.Join(service.config.PathToOutputAvarars, strconv.Itoa(i)+"."+string(TypeImagePNG)), originalAvatar); err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		if err = service.GenerateNFT(avatar, i); err != nil {
			return err
		}

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

// Generate generates values for nft token.
func (service *Service) GenerateNFT(avatar Avatar, sequenceNumber int) error {
	var attributes []Attribute

	if avatar.Background != "" {
		attributes = append(attributes, Attribute{TraitType: "BACKGROUND", Value: avatar.Background})
	}
	if avatar.Heads != "" {
		attributes = append(attributes, Attribute{TraitType: "HEADS", Value: avatar.Heads})
	}
	if avatar.Tshirts != "" {
		attributes = append(attributes, Attribute{TraitType: "JACKET", Value: avatar.Tshirts})
	}
	if avatar.Headwear != "" {
		attributes = append(attributes, Attribute{TraitType: "HATS", Value: avatar.Headwear})
	}
	if avatar.Glasses != "" {
		attributes = append(attributes, Attribute{TraitType: "GLASSES", Value: avatar.Glasses})
	}

	nft := NFT{
		Attributes:  attributes,
		Description: service.config.Description,
		ExternalURL: service.config.ExternalURL,
		Image:       fmt.Sprintf(service.config.Image, sequenceNumber),
		Name:        service.config.Name,
	}

	file, err := json.MarshalIndent(nft, "", " ")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(service.config.PathToOutputAvarars, os.ModePerm); err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(service.config.PathToOutputAvarars, fmt.Sprintf("%d.json", sequenceNumber)), file, 0644)
}
