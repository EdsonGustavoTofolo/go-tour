package http

import (
	"encoding/json"
	"fmt"
	"go-tour/exercicies"
	"io"
	"net/http"
)

func RunHttpServer() {
	http.HandleFunc("/", BuscaCepHandler)

	http.ListenAndServe(":8080", nil)
}

func RunHttpServerMux() {
	mux := http.NewServeMux()

	mux.Handle("/blog", Blog{"My BLog"})

	mux.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":8080", mux)
}

type Blog struct {
	title string
}

func (b Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello Mux"))
}

func RunHttpRequests() {
	req, err := http.Get("https://google.com")

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	println(string(res))
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, error := buscaCep(cepParam)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cep)
}

func buscaCep(cep string) (*exercicies.ViaCEP, error) {
	resp, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json", cep))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var c exercicies.ViaCEP

	err = json.Unmarshal(body, &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
