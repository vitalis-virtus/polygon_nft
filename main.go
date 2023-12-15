package main

import (
	"context"
	"math/big"
	"os"

	"github.com/cosmos/go-bip39"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/tyler-smith/go-bip32"

	"github.com/vitalis-virtus/polygon_nft/contract/nirgp"
	"github.com/vitalis-virtus/polygon_nft/pkg/logger"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	ctx := context.Background()

	log := logger.New()

	// initialize Polygon network client
	client, err := ethclient.Dial(os.Getenv("POLYGON_TESTNET_URL"))
	if err != nil {
		log.Fatalf("failed to connect to the Polygon node: %v\n", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(os.Getenv("NIRGP_CONTRACT_ADDRESS"))

	contractOwnerAddress := common.HexToAddress(os.Getenv("NIRGP_CONTRACT_OWNER"))

	bal, err := client.BalanceAt(ctx, contractOwnerAddress, nil)
	if err != nil {
		log.Fatalf("failed get balance of the account: %v\n", err)
	}
	log.Printf("contract owner %s MATIC balance: %v\n", contractOwnerAddress.String(), bal)

	addressTo := common.HexToAddress("0x790Ba237C185Fd0A38941E9c822fdcB5E9fB9907")

	instance, err := nirgp.NewNirgp(contractAddress, client)
	if err != nil {
		log.Fatalf("failed to create instance: %v\n", err)
	}

	// building transaction
	// get nonce of the given account in the pending state
	nonce, err := client.PendingNonceAt(ctx, contractOwnerAddress)
	if err != nil {
		log.Fatalf("failed to get BTC nonce of the given account in the pending state: %v\n", err)
	}

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatalf("failed get chainID: %v\n", err)
	}

	// create auth bind with sender private key
	seed := bip39.NewSeed(os.Getenv("MNEMONIC"), "")
	// m, _ := hd.ComputeMastersFromSeed(seed)

	pk, err := bip32.NewMasterKey(seed)

	// converting derivation path from []byte to *ecdsa.PrivateKey
	privateKeyEDCSA, err := crypto.ToECDSA(pk.Key)
	if err != nil {
		log.Fatalf("failed to create privateKeyEDCSA: %V\n", err)
	}

	address := crypto.PubkeyToAddress(privateKeyEDCSA.PublicKey)
	log.Println(address)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEDCSA, chainID)
	if err != nil {
		log.Fatalf("failed to get transact/auth opts: %v\n", err)
	}

	// set additional data in auth
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(1)

	tx, err := instance.SafeMint(auth, addressTo)
	if err != nil {
		log.Fatalf("failed to SafeMint: %v\n", err)
	}

	log.Printf("%#v\n", tx)
}
