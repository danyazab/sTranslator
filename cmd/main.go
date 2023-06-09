package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"translator/app"
	"translator/configs"
)

func main() {
	//translation, err := Translate("Привіт як справи", "uk", "ru")
	//if err != nil {
	//	fmt.Println("Translation error:", err)
	//} else {
	//	fmt.Println("Translation:", translation)
	//}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
	cfg := &configs.Config{}
	err = env.Parse(cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}
	server := app.NewBot(cfg)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
