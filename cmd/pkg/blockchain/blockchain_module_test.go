package blockchain

import (
	"blockchainEtherium/api"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strconv"
	"testing"
	"time"
)

func TestModuleSendMoneyArrayContract(t *testing.T) {

	moneyCreate := big.NewInt(10000)

	_, client, _, auth, instance := createSimulateContract()

	var walletsSenders []string
	var walletsRecipients []string
	var moneys []*big.Int

	for i := 0; i < 10000; i++ {
		walletsSenders = append(walletsSenders, strconv.Itoa(i))
		walletsRecipients = append(walletsRecipients, strconv.Itoa(i+1))
		moneys = append(moneys, big.NewInt(10000))
	}

	for i := 0; i < 10000; i++ {
		instance.SetWallet(auth, strconv.Itoa(i), moneyCreate)
		client.Commit()
	}

	start := time.Now().UnixMilli()
	_, err := instance.SendMoneyTenThousandTransactions(auth, walletsSenders, walletsRecipients, big.NewInt(250))
	client.Commit()
	if err != nil {
		log.Fatal(err)
	}
	end := time.Now().UnixMilli()

	//start := time.Now().UnixMilli()
	//_, err := instance.SendMoneyOneThousandTransactions(auth, walletsSenders[0:1000], walletsRecipients[0:1000], big.NewInt(250))
	//client.Commit()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//end := time.Now().UnixMilli()

	//start := time.Now().UnixMilli()
	//_, err := instance.SendMoneyOneHundredTransactions(auth, walletsSenders[0:100], walletsRecipients[0:100], big.NewInt(250))
	//client.Commit()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//end := time.Now().UnixMilli()

	//start := time.Now().UnixMilli()
	//_, err := instance.SendMoneyTenTransactions(auth, walletsSenders[0:10], walletsRecipients[0:10], big.NewInt(250))
	//client.Commit()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//end := time.Now().UnixMilli()

	//start := time.Now().UnixMilli()
	//for i := 0; i < 10000; i++ {
	//	instance.SendMoney(auth, strconv.Itoa(i), strconv.Itoa(i+1), big.NewInt(250))
	//	client.Commit()
	//}

	for i := 0; i < 10; i++ {
		fmt.Println(instance.GetWallet(nil, strconv.Itoa(i)))
	}
	fmt.Println("time send: ", end-start)

}

func createSimulateContract() (common.Address, *backends.SimulatedBackend, *types.Transaction, *bind.TransactOpts, *api.Api) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}

	blockGasLimit := uint64(471238888)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)

	address, tx, _, err := api.DeployApi(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	instance, err := api.NewApi(address, client)

	return address, client, tx, auth, instance
}
