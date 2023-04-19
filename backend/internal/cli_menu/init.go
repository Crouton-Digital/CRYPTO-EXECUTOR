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
		okxClient.SetApiKey("85a19e6e-fd39-4280-b783-448c19f54ebf")
		okxClient.SetSecretKey("83FFBCECDB67AD177FE8D7E890621734")
		okxClient.SetPassPhrase("n543F8355918qw4c&")
		okxClient.SetDemoMode(false)

		okxClient.GetAccountConfig()
		//okxClient.GetTikerSpot()

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
