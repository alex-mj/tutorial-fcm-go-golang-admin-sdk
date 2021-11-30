package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	messaging "firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

func main() {
	//------------------------------------------------------------------
	fmt.Println("Hello, I'm prototype FCM sending")

	// 1) download private key (.json) from https://console.firebase.google.com/
	//  ->Project Settings ->Service accounts  save to local file /cred/cred.json
	// (dont forget add cred.json to .gitignore)

	opt := option.WithCredentialsFile("./cred/cred.json")
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		// TODO: change it
		ProjectID: "YOU'R-PROJECT-NAME"}, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
	}
	//------------------------------------------------------------------
	fmt.Println("I'll try to send a message")
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	// 2) You must add the mobile app
	// https://console.firebase.google.com/
	// ->Project Settings ->General
	// and get a registration token from the mobile device
	// example:
	// registrationToken = "cqzrxJjSQD29FAOBDzyBP5:APA91bF7A7Z_f21_m5XGgwM17lcoiNEGX5sUqfpygnd2alokQ1nWdDI23xE4q1Hv0taT12jaQ00mTWzgGuBVyDO9yyvfsxkq47iOaYcQPyfjuvpw5qmDByO7lLzr471RNHFCtMS_ALmH"

	// TODO: change it
	registrationToken := "registration_token_from_you'r_device"

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		if messaging.IsUnregistered(err) {
			//UNREGISTERED	(error code HTTP = 404) Requested entity was not found.
			logger.L.Info("registration token became invalid")
		}
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
}
