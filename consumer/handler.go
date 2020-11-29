package consumer

import "github.com/tmvrus/service2/api"

var _ OfferHandler = DropHandler{}

type OfferHandler interface {
	OfferCreate(event *api.OfferCreate) error
	OfferDelete(event *api.OfferDelete) error
}

type DropHandler struct{}

func (h DropHandler) OfferCreate(event *api.OfferCreate) error {
	log.Debugf("got DropHandler OfferCreate call with %v", event)
	return nil
}

func (h DropHandler) OfferDelete(event *api.OfferDelete) error {
	log.Debugf("got DropHandler OfferDelete call with %v", event)
	return nil
}
