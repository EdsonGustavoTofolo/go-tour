package exercicies

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

func RunBuscaCep() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json", cep))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisicao: %v\n", err)
			continue
		}

		res, err := io.ReadAll(req.Body)

		req.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta %v\n", err)
			continue
		}

		var data ViaCEP

		err = json.Unmarshal(res, &data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta %v\n", err)
			continue
		}

		fmt.Println(data)

		file, err := os.OpenFile("cidade.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo txt de cidades %v\n", err)
			continue
		}

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s\n", data.Cep, data.Localidade, data.Uf))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo txt de cidades %v\n", err)
			continue
		}

		file.Close()
	}

}
