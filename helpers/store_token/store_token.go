package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type vaultStore struct {
	endpoint string
	token    string
	prefix   string
}

type reqData struct {
	Data map[string]string
}

func prepareRequest(method, url string, data []byte) *http.Request {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err == nil {
		return req
	}
	panic(err)
}

func (store *vaultStore) storeKey(key string, value []byte) {
	url := "/v1/" + store.prefix + "/data"

	// prepare request payload
	data := make(map[string]string)
	data[key] = string(value)
	reqD := &reqData{
		Data: data,
	}
	byteData, _ := json.Marshal(reqD)

	// POST request for storing
	req := prepareRequest("POST", url, byteData)
	client := &http.Client{}
	if res, err := client.Do(req); err == nil {
		fmt.Print(res)
	} else {
		panic(err)
	}
}

func main() {
	token := make([]byte, 100)

	// read token file
	if file, err := os.Open("/docker-master/join-token.log"); err != nil {
		panic(err)
	} else {
		if _, e := file.Read(token); e == nil {
			// store key
			store := &vaultStore{
				endpoint: os.Getenv("VAULT_ENDPOINT"),
				token:    os.Getenv("VAULT_TOKEN"),
				prefix:   os.Getenv("KV_STORE_PREFIX"),
			}
			store.storeKey("cluster-join-token", token)
		} else {
			fmt.Println(e)
		}
	}
}
