package client

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func DoRequest(client *http.Client, url string) ([]byte, error)  {
	resp, err := client.Get(url)
	if err != nil{
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response status code " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil 
}