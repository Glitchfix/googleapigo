package GoogleAPI

import (
	"encoding/base64"
	"net/http"

	speech "google.golang.org/api/speech/v1beta1"

	"io/ioutil"
)

//AudioConfig Google Speech-to-Text config
type AudioConfig struct {
	FilePath        string
	LanguageCode    string
	AudioEncoding   string
	AudioSampleRate int64
}

//SpeechToText get the speech to text
func SpeechToText(client *http.Client, c AudioConfig) (string, error) {
	if c.AudioEncoding == "" {
		c.AudioEncoding = "FLAC"
	}
	if c.AudioSampleRate == 0 {
		c.AudioSampleRate = 16000
	}
	if c.LanguageCode == "" {
		c.LanguageCode = "en-US"
	}
	speechService, err := speech.New(client)
	if err != nil {
		// fmt.Println(err)
		return "", err
	}
	fileDir := c.FilePath

	audioData, err := ioutil.ReadFile(fileDir)
	if err != nil {
		// fmt.Println(err)
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(audioData)
	speechRecConfig := speech.RecognitionConfig{
		SampleRate:   c.AudioSampleRate,
		Encoding:     c.AudioEncoding,
		LanguageCode: c.LanguageCode,
	}
	audio := speech.RecognitionAudio{
		Content: encoded,
	}
	// fmt.Println(encoded)
	// fmt.Println(speechRecConfig)
	speechRequest := speech.SyncRecognizeRequest{
		Audio:  &audio,
		Config: &speechRecConfig,
	}
	syncRecCall := speechService.Speech.Syncrecognize(&speechRequest)
	syncRecResponse, err := syncRecCall.Do()
	if err != nil {
		// fmt.Println(err)
		return "", err
	}
	resp, err := syncRecResponse.MarshalJSON()
	if err != nil {
		// fmt.Println(err)
		return "", err
	}
	return string(resp), nil
}
