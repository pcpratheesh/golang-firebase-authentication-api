package firebase

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

//
// initiate connection with firebase
func InitiateClient() (*auth.Client, error) {
	if os.Getenv("FIR_CREDENTIAL_PATH") == "" {
		return nil, fmt.Errorf("please provide firebase credential file path")
	}

	opt := option.WithCredentialsFile(os.Getenv("FIR_CREDENTIAL_PATH"))

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	//initiate firebase authentication
	auth, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return auth, nil
}
