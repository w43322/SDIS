package smartContractUtils

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"

	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	registerdata "prototype1.0.0/Contracts" // for demo
)

func RequestDataFromBlockchain(imageBCID int) (*big.Int, []byte, []byte, []byte) {
	fmt.Println("Request data from blockchain")
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	blkPrivateKey, err := crypto.HexToECDSA("257a5ba8a4b0d7cb363e4cb8d1a77fdc5234431f82aa12fd2666ce35944d2f65")
	if err != nil {
		log.Fatal(err)
	}

	blkPublicKey := blkPrivateKey.Public()
	publicKeyECDSA, ok := blkPublicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(blkPrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	// smart contract address
	address := common.HexToAddress("0xEF6f24dCD9713da590f5C7aBF741e9195F5545BC")

	instance, err := registerdata.NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}

	var bcid *big.Int

	newCallOpts := &bind.CallOpts{
		Pending: true,
		From:    fromAddress,
	}
	num, err := instance.GetNum(newCallOpts)
	// var reqNum int
	//
	// for i, _ := range num {
	// 	bcid, err = instance.GetBCID(newCallOpts, num[i])
	// 	if bcid == big.NewInt(int64(imageBCID)) {
	// 		reqNum = i
	// 		break
	// 	}
	//
	// }
	bcid, err = instance.GetBCID(newCallOpts, num[len(num)-1])
	enDh, err := instance.GetEnHash(newCallOpts, num[len(num)-1])
	enKey, err := instance.GetEnKey(newCallOpts, num[len(num)-1])
	enCVE, err := instance.GetEnCVE(newCallOpts, num[len(num)-1])

	// bcid, err = instance.GetBCID(newCallOpts, num[reqNum])
	// enDh, err := instance.GetEnHash(newCallOpts, num[reqNum])
	// enKey, err := instance.GetEnKey(newCallOpts, num[reqNum])
	// enCVE, err := instance.GetEnCVE(newCallOpts, num[reqNum])

	// fmt.Println("BCID: ", bcid)
	// fmt.Println("enDh:", enDh)
	// fmt.Println("enKey:", enKey)
	// fmt.Println("enCVE", enCVE)
	fmt.Println("Request data from blockchain done")
	return bcid, enDh, enKey, enCVE
}

func RegisterDataOnBlockchain(imageBCID int, enDimage_IPFSHash, enCVE_IPFSHash, enAESKey []byte) {
	fmt.Println("Register data to blockchain")
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	// blkPrivateKey, err := crypto.HexToECDSA("b6f217a0caf514fc7c6e3e086e89f28b30f9a1f5daee40581c8f5f1805af3e8a")
	blkPrivateKey, err := crypto.HexToECDSA("257a5ba8a4b0d7cb363e4cb8d1a77fdc5234431f82aa12fd2666ce35944d2f65")
	if err != nil {
		log.Fatal(err)
	}

	blkPublicKey := blkPrivateKey.Public()
	publicKeyECDSA, ok := blkPublicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(blkPrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	// smart contract address
	address := common.HexToAddress("0xEF6f24dCD9713da590f5C7aBF741e9195F5545BC")

	instance, err := registerdata.NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}

	var iBCID *big.Int
	var enDiHash []byte
	var enCVEHash []byte
	var enAKey []byte

	iBCID = big.NewInt(int64(imageBCID))
	enDiHash = enDimage_IPFSHash
	enCVEHash = enCVE_IPFSHash
	enAKey = enAESKey

	tx, err := instance.SetDockerImageData(auth, iBCID, enAKey, enDiHash, enCVEHash)
	if err != nil {
		log.Println(err)
	}
	//
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
	// tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	// 이런 식으로 출력이 된다.
	fmt.Println()
	fmt.Println("Register data to blockchain done")
}
