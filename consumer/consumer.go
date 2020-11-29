package consumer

import (
	"github.com/golang/protobuf/proto"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"

	"github.com/tmvrus/service2/api"
)

var log = logrus.WithField("package", "consumer")

type Consumer struct {
	c                       *nsq.Consumer
	requeueUnmarshalError   bool
	suppressValidationError bool
	logUnknownEventError    bool
	offerHandler            OfferHandler
}

func NewConsumer(topic, channel string, nsqd []string, h OfferHandler) (*Consumer, error) {
	c, err := nsq.NewConsumer(topic, channel, nil)
	if err != nil {
		return nil, err
	}

	if err := c.ConnectToNSQDs(nsqd); err != nil {
		return nil, err
	}

	userConsumer := &Consumer{
		offerHandler:            h,
		suppressValidationError: true,
		logUnknownEventError:    true,
	}

	c.AddHandler(userConsumer)
	return userConsumer, nil

}

func (h Consumer) Stop() {
	h.c.Stop()
}

func (h Consumer) HandleMessage(message *nsq.Message) error {
	var offerEvent api.OfferEvent

	if err := proto.Unmarshal(message.Body, &offerEvent); err != nil {
		log.WithError(err).Error("failed to unmarshal offer event")
		if !h.requeueUnmarshalError {
			return err
		}

		return nil
	}

	if err := offerEvent.Validate(); err != nil {
		log.WithError(err).Error("failed to validate offer event")
		if !h.suppressValidationError {
			return err
		}

		return nil
	}

	switch e := offerEvent.Event.(type) {
	case *api.OfferEvent_OfferCreate:
		return h.offerHandler.OfferCreate(e.OfferCreate)
	case *api.OfferEvent_OfferDelete:
		return h.offerHandler.OfferDelete(e.OfferDelete)
	}

	l := log.Debug
	if !h.logUnknownEventError {
		l = log.Error
	}

	l("got unknown offer event: %v", offerEvent)
	return nil
}
