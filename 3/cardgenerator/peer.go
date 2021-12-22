// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardgenerator

import (
	"context"

	"boonji/3/cardgenerator/cardavatars"
	"boonji/internal/logger"
)

// Config is the global configuration for cardgenerator.
type Config struct {
	CardAvatars struct {
		cardavatars.Config
	} `json:"cardAvatars"`
}

// Peer is the representation of a cardgenerator.
type Peer struct {
	Config Config
	Log    logger.Logger

	cardsTotal int

	// exposes avatar cards related logic.
	CardAvatars struct {
		Service *cardavatars.Service
	}
}

// New is a constructor for cardgenerator.Peer.
func New(logger logger.Logger, config Config, cardsTotal int) (peer *Peer, err error) {
	peer = &Peer{
		Log:        logger,
		Config:     config,
		cardsTotal: cardsTotal,
	}

	{ // avatar cards setup
		peer.CardAvatars.Service = cardavatars.NewService(config.CardAvatars.Config)
	}

	return peer, nil
}

// Generate initiates generation of avatar cards.
func (peer *Peer) Generate(ctx context.Context) error {
	err := peer.CardAvatars.Service.Generate(ctx, peer.cardsTotal)
	return err
}
