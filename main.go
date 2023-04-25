package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/adityathebe/hamrobazar/hamrobazar"
	"gopkg.in/yaml.v3"
)

func main() {
	var tgToken, tgChatID string
	flag.StringVar(&tgToken, "token", "", "Telegram bot token")
	flag.StringVar(&tgChatID, "chatid", "", "Telegram chat ID")
	flag.Parse()

	if tgToken == "" {
		tgToken = os.Getenv("HAMROBAZAR_TG_TOKEN")
	}
	if tgChatID == "" {
		tgChatID = os.Getenv("HAMROBAZAR_TG_CHAT_ID")
	}

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Usage: hamrobazar <filename>")
	}
	configFile := args[0]

	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("failed to read %s: %v", configFile, err)
	}

	var configs []hamrobazar.Config
	err = yaml.Unmarshal(data, &configs)
	if err != nil {
		log.Fatalf("failed to unmarshal config %s: %v", configFile, err)
	}

	hm := hamrobazar.NewHamrobazar(tgToken, tgChatID)
	for _, conf := range configs {
		if conf.Disabled {
			continue
		}

		fmt.Println("Searching for", conf.Label)
		if err := hm.Run(conf.Filter); err != nil {
			log.Printf("error running filter %s: %v\n", conf.Label, err)
		}
	}
}
