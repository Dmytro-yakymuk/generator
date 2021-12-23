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
func (service *Service) Generate(ctx context.Context, cardsTotal int) error {
	var stringComponents []string
	i := 0
	for i < cardsTotal {
		var (
			avatar Avatar
			layer  image.Image
			layers []image.Image
			count  int
			err    error
		)

		// Background
		pathToBackground := filepath.Join(service.config.PathToAvararsComponents, service.config.BackgroundFolder)
		if count, err = imageprocessing.LayerComponentsCount(pathToBackground); err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		if avatar.Background, err = rand.RandomInRange(count); err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		if layer, err = imageprocessing.CreateLayer(pathToBackground, avatar.Background); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		layers = append(layers, layer)

		// Heads
		pathToHeads := filepath.Join(service.config.PathToAvararsComponents, service.config.HeadsFolder)
		if count, err = imageprocessing.LayerComponentsCount(pathToHeads); err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		if avatar.Heads, err = rand.RandomInRange(count); err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		if layer, err = imageprocessing.CreateLayer(pathToHeads, avatar.Heads); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		layers = append(layers, layer)

		// Tshirts
		pathToTshirts := filepath.Join(service.config.PathToAvararsComponents, service.config.TshirtsFolder)
		if count, err = imageprocessing.LayerComponentsCount(pathToTshirts); err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		if avatar.Tshirts, err = rand.RandomInRange(count); err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		if layer, err = imageprocessing.CreateLayer(pathToTshirts, avatar.Tshirts); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		layers = append(layers, layer)

		// Headwear
		if rand.IsIncludeRange(service.config.PercentageHeadwear) {
			pathToHeadwear := filepath.Join(service.config.PathToAvararsComponents, service.config.HeadwearFolder)
			if count, err = imageprocessing.LayerComponentsCount(pathToHeadwear); err != nil {
				return ErrCardAvatars.Wrap(err)
			}

			if avatar.Headwear, err = rand.RandomInRange(count); err != nil {
				return ErrCardAvatars.Wrap(err)
			}

			if layer, err = imageprocessing.CreateLayer(pathToHeadwear, avatar.Headwear); err != nil {
				return ErrCardAvatars.Wrap(err)
			}
			layers = append(layers, layer)
		}

		// Glasses
		if rand.IsIncludeRange(service.config.PercentageGlasses) {
			pathToGlasses := filepath.Join(service.config.PathToAvararsComponents, service.config.GlassesFolder)
			if count, err = imageprocessing.LayerComponentsCount(pathToGlasses); err != nil {
				return ErrCardAvatars.Wrap(err)
			}

			if avatar.Glasses, err = rand.RandomInRange(count); err != nil {
				return ErrCardAvatars.Wrap(err)
			}

			if layer, err = imageprocessing.CreateLayer(pathToGlasses, avatar.Glasses); err != nil {
				return ErrCardAvatars.Wrap(err)
			}
			layers = append(layers, layer)
		}

		stringComponent := fmt.Sprintf("%d_%d_%d_%d_%d", avatar.Background, avatar.Heads, avatar.Tshirts, avatar.Headwear, avatar.Glasses)
		for _, stringComp := range stringComponents {
			if stringComp == stringComponent {
				break
			}
		}

		originalAvatar := imageprocessing.Layering(layers, 0, 0)

		if err = imageprocessing.SaveImage(service.config.PathToOutputAvarars, filepath.Join(service.config.PathToOutputAvarars, strconv.Itoa(i+1)+"."+string(TypeImagePNG)), originalAvatar); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		stringComponents = append(stringComponents, stringComponent)
		i++
		fmt.Println(i)
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
