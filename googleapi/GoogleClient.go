package GoogleAPI

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/jwt"
)

//ClientConfig Google Speech-to-Text config
type ClientConfig struct {
	Email      string
	PrivateKey string
	TokenURL   string `default:"https://accounts.google.com/o/oauth2/token"`
}

//GetClient Get the client configured to send authenticated requests
func GetClient(c ClientConfig) *http.Client {
	if c.TokenURL == "" {
		c.TokenURL = "https://accounts.google.com/o/oauth2/token"
	}
	conf := &jwt.Config{
		Email: c.Email,
		// The contents of your RSA private key or your PEM file
		// that contains a private key.
		// If you have a p12 file instead, you
		// can use `openssl` to export the private key into a pem file.
		//
		//    $ openssl pkcs12 -in key.p12 -passin pass:notasecret -out key.pem -nodes
		//
		// The field only supports PEM containers with no passphrase.
		// The openssl command will convert p12 keys to passphrase-less PEM containers.
		PrivateKey: []byte(c.PrivateKey),
		TokenURL:   c.TokenURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/bigquery",
			"https://www.googleapis.com/auth/blogger",
			"https://www.googleapis.com/auth/cloud-platform",
		},
		// If you would like to impersonate a user, you can
		// create a transport with a subject. The following GET
		// request will be made on the behalf of user@example.com.
		// Optional.
	}
	// Initiate an http.Client, the following GET request will be
	// authorized and authenticated on the behalf of user@example.com.
	client := conf.Client(context.Background())
	return client
}
