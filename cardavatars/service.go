// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardavatars

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	randMath "math/rand"
	"os"
	"path/filepath"
	"strconv"

	"github.com/zeebo/errs"

	"generator/internal/probability"
	"generator/pkg/imageprocessing"
	"generator/pkg/rand"
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
	countComponents := map[string]map[string]int{
		"backgrounds": {},
		"backs":       {},
		"stands":      {},
		"shells":      {},
		"eyes":        {},
		"accessories": {},
		"hats":        {},
		"glasses":     {},
		"mouths":      {},
	}

	var tattos []int
	if end-start >= 50 {
		for i := 0; i < 50; i++ {
			var isRepit bool
			tat := randMath.Intn(end-start) + 1

			for _, tatto := range tattos {
				if tat == tatto {
					isRepit = true
				}
			}

			if !isRepit {
				tattos = append(tattos, tat)
			}
		}
	}

	i := start
	for i <= end {
		var (
			avatar Avatar
			// metadata Metadata
			layer  image.Image
			layers []image.Image
			err    error
		)

		// Backgrounds
		pathToBackground := filepath.Join(service.config.PathToAvararsComponents, service.config.BackgroundsFolder)
		avatar.Background = rand.SearchValueByPercent(probability.Background)

		if layer, err = imageprocessing.CreateLayerByFileName(pathToBackground, avatar.Background); err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		// resultBackgrounds := strings.Split(avatar.Background, "_")
		// metadata.Background = resultBackgrounds[0]
		countComponents["backgrounds"][avatar.Background]++
		layers = append(layers, layer)

		// Backs
		if rand.IsIncludeRange(probability.DropoutPercentage) {
			pathToBack := filepath.Join(service.config.PathToAvararsComponents, service.config.BacksFolder)
			avatar.Back = rand.SearchValueByPercent(probability.Back)

			if layer, err = imageprocessing.CreateLayerByFileName(pathToBack, avatar.Back); err != nil {
				return ErrCardAvatars.Wrap(err)
			}

			countComponents["backs"][avatar.Back]++
			layers = append(layers, layer)
		}

		// Stands
		pathToStand := filepath.Join(service.config.PathToAvararsComponents, service.config.StandsFolder)
		avatar.Stand = rand.SearchValueByPercent(probability.Stand)

		if layer, err = imageprocessing.CreateLayerByFileName(pathToStand, avatar.Stand); err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		countComponents["stands"][avatar.Stand]++
		layers = append(layers, layer)

		// Shells
		pathToShell := filepath.Join(service.config.PathToAvararsComponents, service.config.ShellsFolder)
		shellNumber := rand.SearchValueByPercent(probability.ShellNumbers)
		shellType := rand.SearchValueByPercent(probability.ShellTypes)
		avatar.Shell = shellNumber + shellType + probability.ShellNames[shellNumber]

		if layer, err = imageprocessing.CreateLayerByFileName(pathToShell, avatar.Shell); err != nil {
			return ErrCardAvatars.Wrap(err)
		}

		countComponents["shells"][avatar.Shell]++
		layers = append(layers, layer)

		// Eyes
		var eyeNumber string
		if shellType != "smirk" && shellType != "smoke" {
			pathToEye := filepath.Join(service.config.PathToAvararsComponents, service.config.EyesFolder)
			eyeNumber := rand.SearchValueByPercent(probability.EyeNumbers)
			avatar.Eye = eyeNumber + shellType + probability.EyeNames[eyeNumber]

			if layer, err = imageprocessing.CreateLayerByFileName(pathToEye, avatar.Eye); err != nil {
				return ErrCardAvatars.Wrap(err)
			}

			countComponents["eyes"][avatar.Eye]++
			layers = append(layers, layer)
		}

		// Accessories
		if rand.IsIncludeRange(probability.DropoutPercentage) {
			var isTotto bool
			for _, totto := range tattos {
				if totto+start == i {
					isTotto = true
				}
			}

			pathToAccessorie := filepath.Join(service.config.PathToAvararsComponents, service.config.AccessoriesFolder)
			if isTotto {
				avatar.Accessorie = "08_"
			} else {
				avatar.Accessorie = rand.SearchValueByPercent(probability.Accessories)
			}

			if layer, err = imageprocessing.CreateLayerByFileName(pathToAccessorie, avatar.Accessorie); err != nil {
				return ErrCardAvatars.Wrap(err)
			}

			countComponents["accessories"][avatar.Accessorie]++
			layers = append(layers, layer)
		}

		// Hats
		if rand.IsIncludeRange(probability.DropoutPercentage) {
			pathToHat := filepath.Join(service.config.PathToAvararsComponents, service.config.HatsFolder)
			avatar.Hat = rand.SearchValueByPercent(probability.Hats)

			if layer, err = imageprocessing.CreateLayerByFileName(pathToHat, avatar.Hat); err != nil {
				return ErrCardAvatars.Wrap(err)
			}

			countComponents["hats"][avatar.Hat]++
			layers = append(layers, layer)
		}

		// Glasses
		if rand.IsIncludeRange(probability.DropoutPercentage) {
			if eyeNumber != "07_" {
				pathToGlasse := filepath.Join(service.config.PathToAvararsComponents, service.config.GlassesFolder)
				avatar.Glasse = rand.SearchValueByPercent(probability.Glasses)

				if layer, err = imageprocessing.CreateLayerByFileName(pathToGlasse, avatar.Glasse); err != nil {
					return ErrCardAvatars.Wrap(err)
				}

				countComponents["glasses"][avatar.Glasse]++
				layers = append(layers, layer)
			}
		}

		// Mouths
		pathToMouth := filepath.Join(service.config.PathToAvararsComponents, service.config.MouthPropsFolder)

		if shellType == "angry" || shellType == "cheeky" || shellType == "laugh" || shellType == "smirk" || shellType == "smoke" || shellType == "wavy" {
			avatar.Mouth = rand.SearchValueByPercent(probability.Mouth)
			if layer, err = imageprocessing.CreateLayerByFileName(pathToMouth, avatar.Mouth); err != nil {
				return ErrCardAvatars.Wrap(err)
			}

			countComponents["mouths"][avatar.Mouth]++
			layers = append(layers, layer)
		}

		stringComponent := fmt.Sprintf("%s || %s || %s || %s || %s || %s || %s || %s || %s", avatar.Background, avatar.Back, avatar.Stand,
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

		// shellResult := strings.Split(shell, ".")
		// shellInt, err := strconv.Atoi(shellResult[0])
		// if err != nil {
		// 	return ErrCardAvatars.Wrap(err)
		// }

		// eyeResult := strings.Split(eye, ".")
		// eyeInt, err := strconv.Atoi(eyeResult[0])
		// if err != nil {
		// 	return ErrCardAvatars.Wrap(err)
		// }
		// avatar.Shell = shellInt
		// avatar.Eye = eyeInt
		if err = service.GenerateNFT(Metadata(avatar), i); err != nil {
			return err
		}

		fmt.Println(i)
		i++
	}

	if err := os.MkdirAll(service.config.PathToOutputJSON, os.ModePerm); err != nil {
		return err
	}

	file, err := json.MarshalIndent(stringComponents, "", " ")
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(filepath.Join(service.config.PathToOutputJSON, "components.json"), file, 0644); err != nil {
		return err
	}

	file2, err := json.MarshalIndent(countComponents, "", " ")
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(filepath.Join(service.config.PathToOutputJSON, "countComponents.json"), file2, 0644); err != nil {
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
