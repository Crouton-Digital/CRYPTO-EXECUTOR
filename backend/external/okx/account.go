package okx

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func (m *Okx) GetAccountBalance() {
	httpRequestMethod := GET
	httpRequestEndpoint := "/api/v5/account/balance"
	params := ""

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

func (m *Okx) GetAccountPosition() {
	httpRequestMethod := GET
	httpRequestEndpoint := "/api/v5/account/positions"
	params := ""

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

func (m *Okx) GetAccountPositionHistory() {
	httpRequestMethod := GET
	httpRequestEndpoint := "/api/v5/account/positions-history"
	params := ""

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

func (m *Okx) GetAccountPositionRisk() {
	httpRequestMethod := GET
	httpRequestEndpoint := "/api/v5/account/account-position-risk"
	params := ""

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

func (m *Okx) GetAccountBills() {
	httpRequestMethod := GET
	httpRequestEndpoint := "/api/v5/account/bills"
	params := ""

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

func (m *Okx) GetAccountArchive() {
	httpRequestMethod := GET
	httpRequestEndpoint := "/api/v5/account/bills-archive"
	params := ""

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

func (m *Okx) GetAccountConfig() {
	httpRequestMethod := GET
	httpRequestEndpoint := "/api/v5/account/config"
	params := ""

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

//func (m *Okx) SetAccountPositionMode() {
//	httpRequestMethod := POST
//	httpRequestEndpoint := "/api/v5/account/set-position-mode"
//	params := "{\"posMode\":\"long_short_mode\"}"
//
//	body, err := m.HTTPRequest(httpRequestMethod, httpRequestEndpoint, params)
//	if err != nil {
//		logrus.Error(err)
//	}
//
//	var data map[string]interface{}
//	json.Unmarshal(body, &data)
//	for key, value := range data {
//		logrus.Infoln(key, value)
//	}
//}
