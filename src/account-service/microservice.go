package main;

import (
		"net/http"
		JSON "encoding/json"
		"io/ioutil"
		"bytes"
		"errors"
		"strings"
)

type Microservice struct {

}

var microservice Microservice;
var transactionBaseUrl string = "http://localhost:8091/api/transactions";

func (mic *Microservice) CreateTransaction(payload Dto.TransPayload) error {
	client := &http.Client{}

	// marshal User to json  //
	json, err := JSON.Marshal(payload);
	if err != nil {
		return errors.New(err.Error());
	}

	// set the HTTP method, url, and request body
	var transactionServiceUrl = strings.Join([]string{transactionBaseUrl, "/createTransaction"}, "");
    req, err1 := http.NewRequest(http.MethodPost, transactionServiceUrl, bytes.NewBuffer(json))
    if err1 != nil {
		return errors.New(err1.Error());
    }
	
    // set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json");
    resp, err2 := client.Do(req);
    if err2 != nil {
		return errors.New(err2.Error());
    }

	defer resp.Body.Close();
	if resp.StatusCode == 200 {
		return nil;
	}
    
	return errors.New(resp.Status + " (Transaction Api)");
}