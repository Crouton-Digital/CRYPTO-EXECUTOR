package cli_menu

import (
	"crypto-executor/external/evm/ethereum"
	"crypto-executor/external/okx"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
)

type Menu struct {
	Promt     string
	CursorPos int
	MenuItems []*MenuItems
}

type MenuItems struct {
	Text string
	ID   string
}

func (m *Menu) NewMenu(promt string) *Menu {
	return &Menu{
		Promt:     promt,
		MenuItems: make([]*MenuItems, 0),
	}
}

func (m *Menu) Start() error {
L1:
	// the answers will be written to this struct
	answers := struct {
		Command string `survey:"command"` // or you can tag fields to match a specific name
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if answers.Command == "Exit" {
		os.Exit(1)
	}

	if answers.Command == "Generate new wallet" {
		ethereum.EthGenerateWalletWithMnemonic()
		goto L1
	}

	if answers.Command == "okx" {
		okxClient := okx.NewOkxClient()
		okxClient.SetApiKey("3f58ce61-da83-443d-9390-ede0a8b02663")
		okxClient.SetSecretKey("E2FA6C24FEB2144038FE19ED54E1ADCD")
		okxClient.SetPassPhrase("9XmNhsUvkzZ8nNH&")

		okxClient.GetBalance()

		goto L1
	}

	return err
}

func (m *Menu) AddItem(option string, ID string) *Menu {
	menuItem := &MenuItems{
		Text: option,
		ID:   ID,
	}
	m.MenuItems = append(m.MenuItems, menuItem)

	return m
}
