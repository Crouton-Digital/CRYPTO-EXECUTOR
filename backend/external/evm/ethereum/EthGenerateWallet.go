package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/sirupsen/logrus"
	"github.com/tyler-smith/go-bip39"
	"log"
)

func EthGenerateWalletWithMnemonic() {

	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		logrus.Fatal("NewEntropy: ", err)
	}
	logrus.Info("Entropy: ", entropy)

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		logrus.Fatal("NewMnemonic: ", err)
	}
	logrus.Info("Mnemonic: ", mnemonic)

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Private key in hex: %s\n", privateKey)

	publicKey, _ := wallet.PublicKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Public key in hex: %s\n", publicKey)

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559

}

//entropy, err := bip39.NewEntropy(256)
//if err != nil {
//	logrus.Fatal("NewEntropy: ", err)
//}
//logrus.Info("Entropy: ", entropy)
//
//mnemonic, err := bip39.NewMnemonic(entropy)
//if err != nil {
//	logrus.Fatal("NewMnemonic: ", err)
//}
//logrus.Info("Mnemonic: ", mnemonic)
//
//seed := bip39.NewSeed(mnemonic, "")
//logrus.Info("Seed: ", seed)
//
//masterKey, _ := bip32.NewMasterKey(seed)
//publicKey := masterKey.PublicKey()
//
//// Display mnemonic and keys
//fmt.Println("Mnemonic: ", mnemonic)
//fmt.Println("Master private key: ", masterKey)
//fmt.Println("Master public key: ", publicKey)
//
//// m/44'
//key, err := masterKey.NewChildKey(2147483648 + 44)
//if err != nil {
//	log.Fatal(err)
//}
//
//decoded := base58.Decode(key.B58Serialize())
//privateKey := decoded[46:78]
//fmt.Println(hexutil.Encode(privateKey)) // 0x801f14cc6b5f2b0785916685c838c8e64f7f4529a9ca7507c90e5f9078cefc07
//
//// Hex private key to ECDSA private key
//privateKeyECDSA, err := crypto.ToECDSA(privateKey)
//if err != nil {
//	log.Fatal(err)
//}
//
//// ECDSA private key to hex private key
//privateKey = crypto.FromECDSA(privateKeyECDSA)
//fmt.Println(hexutil.Encode(privateKey))
//
//
//0x2A6B94F231421AdBDC5D055E98d6b8003c11d966
//

//publicKeyBytes := crypto.FromECDSAPub(masterKey)

//logrus.Infoln("PRIVATE KEY:", hexutil.Encode()[2:])

//privateKey, err := crypto.ToECDSA(entropy)
//if err != nil {
//	logrus.Fatal("ToECDSA: ", err)
//}
//logrus.Info("Private key: ", privateKey)
//
//privateKeyBytes := crypto.FromECDSA(privateKey)
//logrus.Infoln("PRIVATE KEY:", hexutil.Encode(privateKeyBytes)[2:])
//
//publicKey := privateKey.Public()
//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//if !ok {
//	logrus.Fatal("error casting public key to ECDSA")
//}
//
//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
//logrus.Println("PUBLIC KEY:", hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05
//
//address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
//logrus.Println("WALLET ADDR:", address)

//address := hexutil.Encode(crypto.FromECDSA(privateKey))
//address2 := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
//
//logrus.Info("Wallet address: ", address)
//logrus.Info("Wallet address: ", address2)

//}

func EthGenerateWallet() {

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		logrus.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	logrus.Infoln("PRIVATE KEY:", hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logrus.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	logrus.Println("PUBLIC KEY:", hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	logrus.Println("WALLET ADDR:", address) // 0x96216849c49358B10257cb55b28eA603c874b05E

}
