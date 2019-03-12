# googleapigo
Easy to implement minimalistic golang package for Google APIs

### Example
```golang
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
	fmt.Println(textResponse)

	textConf := GoogleAPI.TextConfig{
		Text: "Hello, darkness my old friend!",
	}

	voiceResponse, err := GoogleAPI.TextToSpeech(googleClient, textConf) //Returns the resulting base64 of the audio file inside json
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(voiceResponse)
}
```

#### Currently supported cloud APIs
1. Speech-to-Text
2. Text-to-Speech

Please feel free to contribute or as issues.
In case you need any help please feel free to contact me.

#### Contact
[GitHub](https://github.com/Glitchfix)
[LinkedIn](https://in.linkedin.com/in/shivanjan)
Email: schakravorty846@gmail.com
Phone: +919658965891, +917978129659
