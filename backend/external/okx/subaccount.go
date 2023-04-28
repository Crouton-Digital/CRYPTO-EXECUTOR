package okx

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func (m *Okx) GetSubAccountList() {
	httpRequestMethod := GET
	httpRequestEndpoint := "/api/v5/users/subaccount/list"
	params := map[string]interface{}{}
	//params["enable"] = "true"

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
