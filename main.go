package main

import (
	"github.com/Manjit2003/bot-framework/service"
	"github.com/Manjit2003/bot-framework/types"
)

func main() {

	s := service.NewWhatsAppService(&types.ServiceConfig{
		Host:       service.DefaultHost,
		Token:      "7b9d6f0ac1fc7016fc99725865d1a85ab6c2b2b0",
		WaBaNumber: "919011331959",
		UserId:     "tclwa",
		Password:   "Tclwa@123",
	})

	res, err := s.SendContactMessage(&types.ContactMessageParams{
		TextMessageParams: &types.TextMessageParams{
			To:   "918788889688",
			Text: "Hello",
		},
		ContactData: &types.ContactData{
			ContactName:  "Manjeet",
			ContactPhone: "918788889688",
		},
	})

	if err != nil {
		panic(err)
	}

	println(string(res))

}
