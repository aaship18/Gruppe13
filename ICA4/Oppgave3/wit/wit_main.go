package main

import (
	"encoding/json"
	"fmt"

	speech "github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA4/Oppgave3/wit/go-speak"
)

type SpeechToText struct {
	Text     string `json:"_text"`
	Outcomes []struct {
		Confidence interface{} `json:"confidence"`
		Intent     string      `json:"intent"`
		Text       string      `json:"_text"`
		Entities   struct {
		} `json:"entities"`
	} `json:"outcomes"`
	WARNING string `json:"WARNING"`
	MsgID   string `json:"msg_id"`
}

func main() {

	speech.SetWitKey("NQKGDPHCW465PKZIAWQ3D4RM5C5KCGYE") //Wit API Key MUST be set before calling any other Wit.AI functions
	a := []byte(speech.SendWitVoice("./audio.wav"))
	//fmt.Println(a)

	var text SpeechToText
	err := json.Unmarshal(a, &text)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", text.Text)
}	

	
	
