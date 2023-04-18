package okx

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

//https://www.okx.com/
//wss://ws.okx.com:8443/ws/v5/public
//wss://ws.okx.com:8443/ws/v5/private
//API key: 3f58ce61-da83-443d-9390-ede0a8b02663
//Secret key: E2FA6C24FEB2144038FE19ED54E1ADCD
//Pass: 9XmNhsUvkzZ8nNH&

type Okx struct {
	Url          string
	WsUrlPublic  string
	WsUrlPrivate string
	ApiKey       string
	SecretKey    string
	PassPhrase   string
	conn         string
}

// Helper function to compute HMAC-SHA256 signature
func hmacSha256(message []byte, secret []byte) []byte {
	mac := hmac.New(sha256.New, secret)
	mac.Write(message)
	return mac.Sum(nil)
}

func NewOkxClient() *Okx {

	client := &Okx{
		Url:          "https://www.okx.com",
		WsUrlPrivate: "wss://ws.okx.com:8443/ws/v5/private",
		WsUrlPublic:  "wss://ws.okx.com:8443/ws/v5/public",
	}

	return client
}

func (m Okx) SetApiKey(apiKey string) {
	m.ApiKey = apiKey
}

func (m Okx) SetSecretKey(secretKey string) {
	m.SecretKey = secretKey
}

func (m Okx) SetPassPhrase(secretKey string) {
	m.SecretKey = secretKey
}

func (m Okx) GetBalance() {

	httpRequestMethod := "GET"
	//httpRequestUrl := "/api/v5/account/balance"

	httpRequestUrl := "/api/spot/v3/instruments"
	params := "instType=SPOT"

	//params := ""
	// Generate the request signature
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))
	message := timestamp + httpRequestMethod + m.Url + httpRequestUrl + params
	h := sha256.New()
	h.Write([]byte(m.SecretKey))
	secret_bytes := h.Sum(nil)
	h.Reset()
	h.Write([]byte(message))
	message_bytes := h.Sum(nil)
	signature := hex.EncodeToString(hmacSha256(message_bytes, secret_bytes))

	// Create the HTTP request
	req, getErr := http.NewRequest(httpRequestMethod, m.Url+httpRequestUrl+params, nil)
	if getErr != nil {
		logrus.Fatal(getErr)
	}

	req.Header.Set("OK-ACCESS-KEY", m.ApiKey)
	req.Header.Set("OK-ACCESS-SIGN", signature)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", m.PassPhrase)
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request and get the response
	client := &http.Client{}
	resp, clientErr := client.Do(req)
	if clientErr != nil {
		logrus.Error(clientErr)
	}

	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		logrus.Error(readErr)
	}

	logrus.Info(bytes.NewReader(body))

	// Parse the response
	var data []map[string]interface{}
	json.NewDecoder(bytes.NewReader(body)).Decode(&data)
	for _, item := range data {
		logrus.Println(item["instrument_id"], item["last"])
	}
	//fmt.Println("GetBalance")
	//fmt.Println(string(body))
}

func (m Okx) GetStatus() {
	httpRequestUrl := "/api/v5/system/status"
	resp, getErr := http.Get(m.Url + httpRequestUrl)
	if getErr != nil {
		logrus.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		logrus.Fatal(readErr)
	}

	fmt.Println("GetStatus")
	fmt.Println(string(body))
}

func (m Okx) GetTikerSpot() {
	httpRequestUrl := "/api/v5/market/tickers?instType=SPOT"
	resp, getErr := http.Get(m.Url + httpRequestUrl)
	if getErr != nil {
		logrus.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		logrus.Fatal(readErr)
	}

	fmt.Println("GetStatus")
	fmt.Println(string(body))
}
