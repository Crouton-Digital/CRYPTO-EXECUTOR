package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	OKEX_BASE_URL = "https://www.okx.com"
)

//apikey = "85a19e6e-fd39-4280-b783-448c19f54ebf"
//secretkey = "83FFBCECDB67AD177FE8D7E890621734"
//IP = ""
//API name = "DEV-API"
//Permissions = "Read/Withdraw/Trade"

func main() {
	// Set your API key and secret key here

	// Set the endpoint and request parameters here
	endpoint := "/api/v5/account/balance"
	//params := ""
	method := "GET"
	url := OKEX_BASE_URL + endpoint

	// Generate the request signature
	//timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	timestamp := IsoTime()
	//prehash := timestamp + "GET" + endpoint + params
	preHash := PreHashString(timestamp, method, endpoint, "")

	signature, _ := HmacSha256Base64Signer(preHash, secret_key)
	//signature, _ := ComputeHmac256(preHash, secret_key)

	fmt.Println(api_key)
	fmt.Println(secret_key)
	fmt.Println(passphrase)
	fmt.Println(endpoint)
	fmt.Println(timestamp)
	fmt.Println(preHash)
	fmt.Println(signature)

	// Create the HTTP request
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Set("OK-ACCESS-KEY", api_key)
	req.Header.Set("OK-ACCESS-SIGN", signature)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", passphrase)
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request and get the response
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.Status)
		return
	}

	// Parse the response
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	for key, value := range data {
		fmt.Println(key, value)
	}
}

// Helper function to compute HMAC-SHA256 signature
func ComputeHmac256(message string, secret string) (string, error) {
	key, err := hex.DecodeString(secret)
	if err != nil {
		return "", err
	}
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	hash := h.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(hash)
	return signature, nil
}

func HmacSha256Base64Signer(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

func IsoTime() string {
	utcTime := time.Now().UTC()
	iso := utcTime.String()
	isoBytes := []byte(iso)
	iso = string(isoBytes[:10]) + "T" + string(isoBytes[11:23]) + "Z"
	return iso
}

func PreHashString(timestamp string, method string, requestPath string, body string) string {
	return timestamp + strings.ToUpper(method) + requestPath + body
}
