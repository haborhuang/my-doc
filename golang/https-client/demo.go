package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func main() {
  // CA证书加载
	caData, err := ioutil.ReadFile("ca.crt")
	if nil != err {
		panic(errors.Wrap(err, "Read CA cert error"))
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(caData) {
		panic(errors.New("Append CA cert error"))
	}

	// Client证书和秘钥加载
	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		panic(errors.Wrap(err, "Load client cert and key error"))
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	tlsConfig.BuildNameToCertificate()

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	resp, err := client.Get(
		"https://xxxxxx")
	if nil != err {
		panic(errors.Wrap(err, "HTTP GET error"))
	}
	defer resp.Body.Close()

	log.Println("status code:", resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response body:", string(body))
}
