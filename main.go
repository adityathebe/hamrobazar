package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/adityathebe/hamrobazar/hamrobazar"
	"gopkg.in/yaml.v3"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Usage: hamrobazar <filename>")
	}
	configFile := args[0]

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("failed to read %s: %v", configFile, err)
	}

	var configs []hamrobazar.Config
	err = yaml.Unmarshal(data, &configs)
	if err != nil {
		log.Fatalf("failed to unmarshal config %s: %v", configFile, err)
	}

	for _, conf := range configs {
		if conf.Disabled {
			continue
		}

		fmt.Println("Searching for", conf.Label)
		if err := hamrobazar.Run(conf.Filter); err != nil {
			log.Printf("error running filter %s: %v\n", conf.Label, err)
		}
	}
}
