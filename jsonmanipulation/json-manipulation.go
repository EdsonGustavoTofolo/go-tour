package jsonmanipulation

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int     `json:"n"`
	Saldo  float32 `json:"s"`
	CPF    string  `json:"-"`
}

func RunJSONManipulation() {
	conta := Conta{1, 1000.0, "12345678909"}

	res, err := json.Marshal(conta)

	if err != nil {
		panic(err)
	}

	fmt.Println("Via Marshal:", string(res))

	fmt.Print("Via Encoder: ")
	encoder := json.NewEncoder(os.Stdout)

	err = encoder.Encode(conta)

	if err != nil {
		panic(err)
	}

	var contaX Conta

	err = json.Unmarshal(res, &contaX)

	if err != nil {
		panic(err)
	}

	fmt.Println("Conta Via Unmarshal:", contaX)
}
