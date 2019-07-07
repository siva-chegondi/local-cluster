package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type vaultStore struct {
	endpoint string
	token    string
	prefix   string
}

func prepareRequest(method, url string, data []byte) *http.Request {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err == nil {
		return req
	}
	panic(err)
}

// ReadData method to read data from store
func (store *vaultStore) readData(key string) {
	readURL := store.endpoint + "/v1/" + store.prefix + "/data"
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req := prepareRequest("GET", readURL, nil)
	if res, err := client.Do(req); err == nil {
		// write to file "/docker-node/join-token"
		data, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		// write data to path for joining to docker swarm
		f, _ := os.Open("/docker-node/join-token")
		if _, err = f.Write(data); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func main() {
	// connect to Vault
	store := &vaultStore{
		endpoint: os.Getenv("VAULT_ENDPOINT"),
		token:    os.Getenv("VAULT_TOKEN"),
		prefix:   os.Getenv("KV_STORE_PREFIX"),
	}
	store.readData("cluster-join-token")
}
