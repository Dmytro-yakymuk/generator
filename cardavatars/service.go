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
	"strings"

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
			avatar   Avatar
			metadata Metadata
			layer    image.Image
			layers   []image.Image
			name     string
			err      error
		)

		// Backgrounds
		pathToBackground := filepath.Join(service.config.PathToAvararsComponents, service.config.BackgroundsFolder)
		count, err := imageprocessing.LayerComponentsCount(pathToBackground)
		if err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		avatar.Background = rand.Intn(count) + 1

		if layer, name, err = imageprocessing.CreateLayer(pathToBackground, avatar.Background); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		resultBackgrounds := strings.Split(name, ".")
		metadata.Background = resultBackgrounds[0]
		layers = append(layers, layer)

		// Backs
		pathToBack := filepath.Join(service.config.PathToAvararsComponents, service.config.BacksFolder)
		count, err = imageprocessing.LayerComponentsCount(pathToBack)
		if err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		avatar.Back = rand.Intn(count) + 1

		if layer, name, err = imageprocessing.CreateLayer(pathToBack, avatar.Back); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		resultBacks := strings.Split(name, ".")
		metadata.Back = resultBacks[0]
		layers = append(layers, layer)

		// Stands
		pathToStand := filepath.Join(service.config.PathToAvararsComponents, service.config.StandsFolder)
		count, err = imageprocessing.LayerComponentsCount(pathToStand)
		if err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		avatar.Stand = rand.Intn(count) + 1

		if layer, name, err = imageprocessing.CreateLayer(pathToStand, avatar.Stand); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		resultStands := strings.Split(name, ".")
		metadata.Stand = resultStands[0]
		layers = append(layers, layer)

		// Shells
		pathToShell := filepath.Join(service.config.PathToAvararsComponents, service.config.ShellsFolder)
		count, err = imageprocessing.LayerComponentsCount(pathToShell)
		if err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		avatar.Shell = rand.Intn(count) + 1

		if layer, name, err = imageprocessing.CreateLayer(pathToShell, avatar.Shell); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		resultShells := strings.Split(name, "_")
		metadata.Shell = resultShells[0]
		shell := resultShells[1]
		layers = append(layers, layer)

		// Eyes
		pathToEye := filepath.Join(service.config.PathToAvararsComponents, service.config.EyesFolder)
		count, err = imageprocessing.LayerComponentsCount(pathToEye)
		if err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		var isEye bool = true
		var eye string
		for isEye {
			avatar.Eye = rand.Intn(count) + 1
			if layer, name, err = imageprocessing.CreateLayer(pathToEye, avatar.Eye); err != nil {
				return ErrCardAvatars.Wrap(err)
			}
			resultEyes := strings.Split(name, "_")
			metadata.Eye = resultEyes[0]

			eye = resultEyes[1]
			if shell != eye {
				continue
			}
			isEye = false
		}
		layers = append(layers, layer)

		// Accessories
		pathToAccessorie := filepath.Join(service.config.PathToAvararsComponents, service.config.AccessoriesFolder)
		count, err = imageprocessing.LayerComponentsCount(pathToAccessorie)
		if err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		avatar.Accessorie = rand.Intn(count) + 1

		if layer, name, err = imageprocessing.CreateLayer(pathToAccessorie, avatar.Accessorie); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		resultAccessories := strings.Split(name, ".")
		metadata.Accessorie = resultAccessories[0]
		layers = append(layers, layer)

		// Hats
		pathToHat := filepath.Join(service.config.PathToAvararsComponents, service.config.HatsFolder)
		count, err = imageprocessing.LayerComponentsCount(pathToHat)
		if err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		avatar.Hat = rand.Intn(count) + 1

		if layer, name, err = imageprocessing.CreateLayer(pathToHat, avatar.Hat); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		resultHats := strings.Split(name, ".")
		metadata.Hat = resultHats[0]
		layers = append(layers, layer)

		// Glasses
		pathToGlasse := filepath.Join(service.config.PathToAvararsComponents, service.config.GlassesFolder)
		count, err = imageprocessing.LayerComponentsCount(pathToGlasse)
		if err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		avatar.Glasse = rand.Intn(count) + 1

		if layer, name, err = imageprocessing.CreateLayer(pathToGlasse, avatar.Glasse); err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		resultGlasses := strings.Split(name, ".")
		metadata.Glasse = resultGlasses[0]
		layers = append(layers, layer)

		// Mouths
		pathToMouth := filepath.Join(service.config.PathToAvararsComponents, service.config.MouthPropsFolder)
		count, err = imageprocessing.LayerComponentsCount(pathToMouth)
		if err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		var isMouth bool = true
		for _, v := range service.config.Mouths {
			shellResult := strings.Split(shell, ".")
			shellInt, err := strconv.Atoi(shellResult[0])
			if err != nil {
				return ErrCardAvatars.Wrap(err)
			}
			if shellInt == v {
				isMouth = false
			}
		}

		if !isMouth {
			avatar.Mouth = rand.Intn(count) + 1
			if layer, name, err = imageprocessing.CreateLayer(pathToMouth, avatar.Mouth); err != nil {
				return ErrCardAvatars.Wrap(err)
			}
			resultMouths := strings.Split(name, ".")
			metadata.Mouth = resultMouths[0]

			layers = append(layers, layer)
		}

		stringComponent := fmt.Sprintf("%d || %d || %d || %d || %d || %d || %d || %d || %d", avatar.Background, avatar.Back, avatar.Stand,
			avatar.Shell, avatar.Eye, avatar.Accessorie, avatar.Hat, avatar.Glasse, avatar.Mouth)
		for _, stringComp := range stringComponents {
			if stringComp == stringComponent {
				break
			}
		}
		stringComponents = append(stringComponents, stringComponent)

		originalAvatar := imageprocessing.Layering(layers, 0, 0)
		if err = imageprocessing.SaveImage(service.config.PathToOutputAvarars, filepath.Join(service.config.PathToOutputAvarars, strconv.Itoa(i)+"."+string(TypeImagePNG)), originalAvatar); err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		shellResult := strings.Split(shell, ".")
		shellInt, err := strconv.Atoi(shellResult[0])
		if err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		eyeResult := strings.Split(eye, ".")
		eyeInt, err := strconv.Atoi(eyeResult[0])
		if err != nil {
			return ErrCardAvatars.Wrap(err)
		}
		avatar.Shell = shellInt
		avatar.Eye = eyeInt
		if err = service.GenerateNFT(metadata, i); err != nil {
			return err
		}

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
func (service *Service) GenerateNFT(metadata Metadata, sequenceNumber int) error {
	var attributes []Attribute

	if metadata.Background != "" {
		attributes = append(attributes, Attribute{TraitType: "backgrounds", Value: metadata.Background})
	}
	if metadata.Back != "" {
		attributes = append(attributes, Attribute{TraitType: "back", Value: metadata.Back})
	}
	if metadata.Stand != "" {
		attributes = append(attributes, Attribute{TraitType: "stands", Value: metadata.Stand})
	}
	if metadata.Shell != "" {
		attributes = append(attributes, Attribute{TraitType: "shells", Value: metadata.Shell})
	}
	if metadata.Eye != "" {
		attributes = append(attributes, Attribute{TraitType: "eyes", Value: metadata.Eye})
	}
	if metadata.Accessorie != "" {
		attributes = append(attributes, Attribute{TraitType: "acc", Value: metadata.Accessorie})
	}
	if metadata.Hat != "" {
		attributes = append(attributes, Attribute{TraitType: "hats", Value: metadata.Hat})
	}
	if metadata.Glasse != "" {
		attributes = append(attributes, Attribute{TraitType: "glasses", Value: metadata.Glasse})
	}
	if metadata.Mouth != "" {
		attributes = append(attributes, Attribute{TraitType: "mouth", Value: metadata.Mouth})
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
