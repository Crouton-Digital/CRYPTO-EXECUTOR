//package main
//
//import (
//	"crypto/hmac"
//	"crypto/sha256"
//	"encoding/base64"
//	"encoding/hex"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"strings"
//	"time"
//)
//
//const (
//	OKEX_BASE_URL = "https://www.okex.com"
//)
//
//func main() {
//	// Set your API key and secret key here
//	api_key := "3f58ce61-da83-443d-9390-ede0a8b02663"
//	secret_key := "E2FA6C24FEB2144038FE19ED54E1ADCD"
//	passphrase := "9XmNhsUvkzZ8nNH&"
//
//	// Set the endpoint and request parameters here
//	endpoint := "/api/v5/account/balance"
//	params := ""
//
//	// Generate the request signature
//	//timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))
//
//	// Generate the request signature
//	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
//	message := ""
//	signature, err := signMessage(timestamp, "GET", "/api/v5/account/balance", "", secret_key)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	fmt.Println(timestamp)
//	fmt.Println(message)
//	fmt.Println(signature)
//	fmt.Println(api_key)
//	fmt.Println(passphrase)
//
//	// Create the HTTP request
//	url := OKEX_BASE_URL + endpoint + "?" + params
//	req, _ := http.NewRequest("GET", url, nil)
//	req.Header.Set("OK-ACCESS-KEY", api_key)
//	req.Header.Set("OK-ACCESS-SIGN", signature)
//	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
//	req.Header.Set("OK-ACCESS-PASSPHRASE", passphrase)
//	req.Header.Set("Content-Type", "application/json")
//
//	// Send the HTTP request and get the response
//	client := &http.Client{
//		Timeout: 10 * time.Second,
//	}
//	resp, _ := client.Do(req)
//	defer resp.Body.Close()
//	body, _ := ioutil.ReadAll(resp.Body)
//
//	if resp.StatusCode != http.StatusOK {
//		fmt.Println("Error:", resp.Status)
//		return
//	}
//	fmt.Println(body)
//
//	// Parse the response
//	var data []map[string]interface{}
//	json.Unmarshal(body, &data)
//	for _, item := range data {
//		fmt.Println(item["instrument_id"], item["last"])
//	}
//}
//
////// Helper function to compute HMAC-SHA256 signature
////func ComputeHmac256(message string, secret string) (string, error) {
////	key, err := hex.DecodeString(secret)
////	if err != nil {
////		return "", err
////	}
////	h := hmac.New(sha256.New, key)
////	h.Write([]byte(message))
////	return strings.ToUpper(hex.EncodeToString(h.Sum(nil))), nil
////}
//
//func signMessage(timestamp string, method string, requestPath string, requestBody string, secretKey string) (string, error) {
//	prehashString := timestamp + method + requestPath + requestBody
//	secretKeyBytes, err := base64.StdEncoding.DecodeString(secretKey)
//	if err != nil {
//		return "", err
//	}
//	mac := hmac.New(sha256.New, secretKeyBytes)
//	mac.Write([]byte(prehashString))
//	signature := mac.Sum(nil)
//	return base64.StdEncoding.EncodeToString(signature), nil
//}
