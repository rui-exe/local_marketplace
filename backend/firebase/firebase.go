package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func InitializeApp() *firebase.App {
	opt := option.WithCredentialsFile("keys/firebase_service_account_key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println(err)
	}
	return app
}

var App *firebase.App = InitializeApp()
