package handlers

import (
	"github.com/colly-scraper/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	baseURL        string = "http://localhost:8080"
	getEndpoint    string = "/coffees/"
	postEndpoint   string = getEndpoint
	deleteEndpoint string = getEndpoint
)

func GetAll() []models.CoffeeResponse {
	url := baseURL + getEndpoint
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var resp []models.CoffeeResponse
	err = json.Unmarshal(responseData, &resp)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func Add(coffee models.Coffee) {
	url := baseURL + postEndpoint
	postBody, _ := json.Marshal(coffee)
	responseBody := bytes.NewBuffer(postBody)
	response, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		log.Fatalf("An error occured %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)
}

func Delete(coffee models.CoffeeResponse) {
	url := fmt.Sprintf("%s%s%s/deletepermamently", baseURL, deleteEndpoint, strconv.Itoa(coffee.ID))
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return
	}
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
