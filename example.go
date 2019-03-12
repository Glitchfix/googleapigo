package main

import (
	"fmt"

	GoogleAPI "github.com/Glitchfix/googleapigo/googleapi"
)

func main() {
	//Google API Configuration
	apiConf := GoogleAPI.ClientConfig{
		Email:      "YOUR CONFIG EMAIL",
		PrivateKey: "YOUR RSA KEY",
	}

	//OAuth configured Google client
	googleClient := GoogleAPI.GetClient(apiConf)

	//AudioConfig for Google Speech-to-Text
	audioConf := GoogleAPI.AudioConfig{
		FilePath: "brooklyn.flac",
	}

	textResponse, err := GoogleAPI.SpeechToText(googleClient, audioConf) //Returns the resulting transcript inside json
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(audioResponse)

	textConf := GoogleAPI.TextConfig{
		Text: "Hello, darkness my old friend!",
	}

	voiceResponse, err := GoogleAPI.TextToSpeech(googleClient, textConf) //Returns the resulting base64 of the audio file inside json
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(textResponse)
}
