package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ZipCode struct {
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

func main() {
	http.HandleFunc("/", SearchZipCodeHandler)
	http.ListenAndServe(":8080", nil)
}

func SearchZipCodeHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	zipCodeParam := request.URL.Query().Get("cep")
	if zipCodeParam == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	zipCode, error := searchZipCode(zipCodeParam)
	if error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(zipCode)
}

func searchZipCode(zipCode string) (*ZipCode, error) {
	response, error := http.Get("http://viacep.com.br/ws/" + zipCode + "/json/")
	if error != nil {
		return nil, error
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}

	var data ZipCode
	if error = json.Unmarshal(body, &data); error != nil {
		return nil, error
	}

	return &data, nil
}
