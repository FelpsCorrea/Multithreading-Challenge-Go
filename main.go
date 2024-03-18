package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BrasilAPICEP struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func GetBrasilCepData(cep string, ch chan<- BrasilAPICEP) {

	// Faz a requisição com o cep passado como argumento
	req, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)

	if err != nil {
		return
	}

	// Para fechar o Body assim que executar todo o código
	defer req.Body.Close()

	// Para ler o Body da requisição
	res, err := io.ReadAll(req.Body)

	if err != nil {
		return
	}

	// Convertendo o body recebido para Struct
	var data BrasilAPICEP

	err = json.Unmarshal(res, &data)

	if err != nil {
		return
	}

	// Escreve no canal
	ch <- data
}

func GetViaCepData(cep string, ch chan<- ViaCEP) {

	// Faz a requisição com o cep passado como argumento
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")

	if err != nil {
		return
	}

	// Para fechar o Body assim que executar todo o código
	defer req.Body.Close()

	// Para ler o Body da requisição
	res, err := io.ReadAll(req.Body)

	if err != nil {
		return
	}

	// Convertendo o body recebido para Struct
	var data ViaCEP

	err = json.Unmarshal(res, &data)

	if err != nil {
		return
	}

	// Escreve no canal
	ch <- data
}

func main() {

	for _, cep := range os.Args[1:] {

		// Cria os canais
		c1 := make(chan ViaCEP)
		c2 := make(chan BrasilAPICEP)

		// Instancia as duas goroutines
		go GetViaCepData(cep, c1)
		go GetBrasilCepData(cep, c2)

		select {
		case msg := <-c1:
			// Se a primeira goroutine retornar, imprime o resultado
			fmt.Println("Dado recebido de ViaCEP: " + fmt.Sprintf("%v", msg))

		case msg := <-c2:
			// Se a primeira goroutine retornar, imprime o resultado
			fmt.Println("Dado recebido de BrasilCEP: " + fmt.Sprintf("%v", msg))

		case <-time.After(1 * time.Second):
			fmt.Println("Timeout")
		}

	}
}
