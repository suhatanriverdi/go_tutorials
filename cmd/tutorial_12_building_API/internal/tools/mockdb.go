package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"genesis": {
		Username:  "genesis",
		AuthToken: "gen",
	},
	"melo": {
		Username:  "melo",
		AuthToken: "melo",
	},
	"test": {
		Username:  "test",
		AuthToken: "test",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"melo": {
		Coins:    100,
		Username: "melo",
	},
	"genesis": {
		Coins:    10,
		Username: "genesis",
	},
	"test": {
		Coins:    1453,
		Username: "test",
	},
}

func (m *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (m *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
