package service

import (
	"fmt"

	"github.com/tmvrus/service1/api"
	"github.com/tmvrus/service1/consumer"
)

type CreateHandler struct {
	consumer.DropHandler
}

func (h CreateHandler) UserCreate(event *api.UserCreate) error {
	fmt.Printf("got CreateUser: %v", event)
	return nil
}

func CreateUserConsumer() {
	handler := CreateHandler{}

	uc, err := consumer.NewConsumer("user", "service2", []string{"host:port"}, handler)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	uc.Stop()
}
