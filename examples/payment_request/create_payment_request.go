package main

import (
	"context"
	"log"

	"github.com/tonicpow/go-handcash-connect"
)

func main() {

	// Create a new client (Beta ENV)
	client := handcash.NewClient(nil, nil, handcash.EnvironmentProduction)

	// Get the payment information (given AuthToken and TxID)
	payment, err := client.PaymentRequest(
		context.Background(),
		&handcash.PaymentRequestV2{
			Product: map[string]string{
				"name":        "Winchester Model 70",
				"description": "Push feed 1988",
				"imageUrl":    "https://p1.gunbroker.com/pics/952554000/952554213/pix170505224.jpg",
			},
			Receivers: []*handcash.PaymentV2{
				{
					Amount:       0.01,
					CurrencyCode: handcash.CurrencyUSD,
					Destination:  "dylan-murray",
				},
			},
			RequestedUserData: []string{"paymail"},
			Notifications: handcash.Notifications{
				Webhook: handcash.Webhook{
					CustomParameters: map[string]string{
						"year": "1988",
					},
					WebhookUrl: "https://pay.bitnetes.com/webhooks/handcash",
				},
				Email: "dylan@britevue.com",
			},
			ExpirationType: "never",
			RedirectUrl:    "https://gunbroker.com",
		},
	)
	if err != nil {
		log.Fatalln("error: ", err)
	}
	log.Println("payment: ", payment)
}
