package okx

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type Okx struct {
	Url          string
	WsUrlPublic  string
	WsUrlPrivate string
	ApiKey       string
	SecretKey    string
	PassPhrase   string
	isSimulate   bool
	conn         string
}

func NewOkxClient() *Okx {

	client := &Okx{
		Url:          "https://www.okx.com",
		WsUrlPrivate: "wss://ws.okx.com:8443/ws/v5/private",
		WsUrlPublic:  "wss://ws.okx.com:8443/ws/v5/public",
	}

	return client
}

func (m *Okx) SetApiKey(apiKey string) {
	m.ApiKey = apiKey
}

func (m *Okx) SetSecretKey(secretKey string) {
	m.SecretKey = secretKey
}

func (m *Okx) SetPassPhrase(passPhrase string) {
	m.PassPhrase = passPhrase
}

func (m *Okx) SetDemoMode(demoMode bool) {
	m.isSimulate = demoMode
}

func (m *Okx) GenerateHeaders(request *http.Request, timestamp string, sign string) (header string) {

	request.Header.Add(ACCEPT, APPLICATION_JSON)
	header += ACCEPT + ":" + APPLICATION_JSON + "\n"

	request.Header.Add(CONTENT_TYPE, APPLICATION_JSON_UTF8)
	header += CONTENT_TYPE + ":" + APPLICATION_JSON_UTF8 + "\n"

	request.Header.Add(COOKIE, LOCALE+ENGLISH)
	header += COOKIE + ":" + LOCALE + ENGLISH + "\n"

	request.Header.Add(OK_ACCESS_KEY, m.ApiKey)
	header += OK_ACCESS_KEY + ":" + m.ApiKey + "\n"

	request.Header.Add(OK_ACCESS_SIGN, sign)
	header += OK_ACCESS_SIGN + ":" + sign + "\n"

	request.Header.Add(OK_ACCESS_TIMESTAMP, timestamp)
	header += OK_ACCESS_TIMESTAMP + ":" + timestamp + "\n"

	request.Header.Add(OK_ACCESS_PASSPHRASE, m.PassPhrase)
	header += OK_ACCESS_PASSPHRASE + ":" + m.PassPhrase + "\n"

	if m.isSimulate {
		request.Header.Add(X_SIMULATE_TRADING, "1")
		header += X_SIMULATE_TRADING + ":1" + "\n"
	}

	return
}

func (m *Okx) HTTPRequest(httpRequestMethod string, httpRequestEndpoint string, params string) (body []byte, err error) {

	httpRequestUrl := m.Url + httpRequestEndpoint
	if len(params) > 1 {
		httpRequestUrl = httpRequestUrl + "?" + params
	}

	timestamp := IsoTime()
	preHash := PreHashString(timestamp, httpRequestMethod, httpRequestEndpoint, "")
	signature, _ := HmacSha256Base64Signer(preHash, m.SecretKey)

	req, _ := http.NewRequest(httpRequestMethod, httpRequestUrl, nil)
	m.GenerateHeaders(req, timestamp, signature)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Error("Error:", err)
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Error("Error:", resp.Status)
	}

	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorln("Error:", err)
	}

	return
}

func (m *Okx) GetStatus() {
	httpRequestUrl := "/api/v5/system/status"
	resp, getErr := http.Get(m.Url + httpRequestUrl)
	if getErr != nil {
		logrus.Fatal(getErr)
	}
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		logrus.Fatal(readErr)
	}

	logrus.Info("GetStatus")
	logrus.Info(string(body))
}

func (m *Okx) GetTikerSpot() {
	httpRequestMethod := GET
	httpRequestEndpoint := "/api/v5/market/tickers"
	params := "instType=SPOT"

	body, err := m.HTTPRequest(httpRequestMethod, httpRequestEndpoint, params)
	if err != nil {
		logrus.Error(err)
	}

	var data map[string]interface{}
	json.Unmarshal(body, &data)
	for key, value := range data {
		logrus.Infoln(key, value)
	}
}
