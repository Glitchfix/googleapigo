package GoogleAPI

import (
	"net/http"

	texttospeech "google.golang.org/api/texttospeech/v1beta1"
)

//TextConfig Google Speech-to-Text config
type TextConfig struct {
	Text          string
	LanguageCode  string
	AudioEncoding string
}

//TextToSpeech get the text to speech
func TextToSpeech(client *http.Client, c TextConfig) (string, error) {
	texttospeechService, err := texttospeech.New(client)
	if c.AudioEncoding == "" {
		c.AudioEncoding = "MP3"
	}
	if c.LanguageCode == "" {
		c.LanguageCode = "en-IN"
	}
	if err != nil {
		// fmt.Println(err)
		return "", err
	}
	ttsConfig := texttospeech.AudioConfig{
		AudioEncoding: c.AudioEncoding,
	}
	inputText := texttospeech.SynthesisInput{
		Text: c.Text,
	}
	voiceParams := texttospeech.VoiceSelectionParams{
		LanguageCode: c.LanguageCode,
	}
	ttsSpeechReq := texttospeech.SynthesizeSpeechRequest{
		AudioConfig: &ttsConfig,
		Input:       &inputText,
		Voice:       &voiceParams,
	}
	ttsSpeechCall := texttospeechService.Text.Synthesize(&ttsSpeechReq)
	syncResponse, err := ttsSpeechCall.Do()
	if err != nil {
		// fmt.Println(err)
		return "", err
	}
	resp, err := syncResponse.MarshalJSON()
	if err != nil {
		// fmt.Println(err)
		return "", err
	}
	// fmt.Println(string(resp))
	return string(resp), nil
}
