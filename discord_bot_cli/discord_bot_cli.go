package discord_bot_cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Embed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       int    `json:"color"`
}

type Context struct {
	Embed Embed `json:"embed"`
}

func getColor(level string) int {
	switch level {
	case "info":
		return 0x44aaee
	case "warn":
		return 0xffcc55
	case "error":
		return 0xff3366
	}
	return 0x44aaee
}

func Run(token string, channel string, level string, title string, description string) {
	url := fmt.Sprintf("https://discordapp.com/api/channels/%s/messages", channel)

	context := Context{
		Embed: Embed{
			Title:       title,
			Description: description,
			Color:       getColor(level),
		},
	}
	contextJson, err := json.Marshal(context)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(contextJson))
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Print("[!] " + err.Error())
	} else {
		log.Print("[*] " + resp.Status)
	}

	log.Print(bytes.NewBuffer(contextJson).String())

	respText, _ := io.ReadAll(resp.Body)
	log.Print(string(respText))
}
