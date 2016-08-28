package main

import (
	"github.com/DanielDanteDosSantosViana/nuveo/robot"
	"log"
	"os"
)

func main() {

	if len(os.Args) > 2 || len(os.Args) <= 1 {
		log.Fatal("É necessário apenas um argumento(URL de destino), Conforme o comando :'go nuveo [url de destino]'")
	}
	url := os.Args[1]
	rb := robot.New()
	data, err := rb.GetData(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range data {
		log.Printf("Nome: %s", value.Nome)
		log.Printf("Email: %s", value.Email)
		log.Printf("Idade: %s", value.Idade)
		log.Printf("Sexo: %s", value.Sexo)
		log.Printf("Outros campos: %s", value.Outros)

	}

}
