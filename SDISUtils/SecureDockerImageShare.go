package SDISUtils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"prototype1.0.0/IPFSUtils"

	"prototype1.0.0/CryptoUtils"
	"prototype1.0.0/gaillier"
	"prototype1.0.0/smartContractUtils"
)

func SecureDockerImageShare(imageBCID int) []byte {

	conInfo := &publisherInfo{
		ID:  "yujin",
		PWD: "123456",
	}
	marPubInfo, _ := json.Marshal(conInfo)

	start := time.Now()
	_, enDh, enKey, enCVE := smartContractUtils.RequestDataFromBlockchain(imageBCID)
	// fmt.Println(bcid, enDh, enKey, enCVE)

	// 초기 설정값-------------------------------------------------------------
	IPFSUtils.InitIPFSShell()

	enHomoPubKey, err := ioutil.ReadFile("./homoPublicKey")
	checkError(err)
	enHomoPrivKey, err := ioutil.ReadFile("./homoPrivKey")
	checkError(err)

	pubK := &gaillier.PubKey{}
	pubDec := gob.NewDecoder(bytes.NewReader(enHomoPubKey))
	err = pubDec.Decode(pubK)

	privK := &gaillier.PrivKey{}
	privDec := gob.NewDecoder(bytes.NewReader(enHomoPrivKey))
	err = privDec.Decode(privK)

	privPEM, err := ioutil.ReadFile("./RSAPrivateKey.pem")
	conPubPEM, err := ioutil.ReadFile("./consumer_RSAPublicKey.pem")

	RSAPrivateKey, err := CryptoUtils.ParseRsaPrivateKeyFromPemStr(string(privPEM))
	consumer_RSAPublicKey, err := CryptoUtils.ParseRsaPublicKeyFromPemStr(string(conPubPEM))

	//---------------------------------------------------------------------------
	restoredCVEIPFSHash, err := CryptoUtils.RSADecryption(RSAPrivateKey, enCVE)
	checkError(err)
	restoredDimageIPFSHash, err := CryptoUtils.RSADecryption(RSAPrivateKey, enDh)
	checkError(err)
	restoredAESKey, err := CryptoUtils.RSADecryption(RSAPrivateKey, enKey)
	checkError(err)

	UCVEHashValue := CryptoUtils.Hash(string(restoredCVEIPFSHash))
	CVEHashValue := int(UCVEHashValue) / 1000
	UIPFSHashValue := CryptoUtils.Hash(string(restoredDimageIPFSHash))
	IPFSHashValue := int(UIPFSHashValue) / 1000
	UAESKeyHashValue := CryptoUtils.Hash(fmt.Sprint(restoredAESKey))
	AESKeyHashValue := int(UAESKeyHashValue) / 1000
	UpublisherInfoHashValue := CryptoUtils.Hash(string(marPubInfo))
	publisherInfoHashValue := int(UpublisherInfoHashValue) / 1000

	homoCVElistHashValue, err := gaillier.Encrypt(pubK, IntToBytes(CVEHashValue))
	// homoCVElistHashValue, err := gaillier.Encrypt(pubK, CryptoUtils.I32tob(int(CVEHashValue/10)))
	checkError(err)
	homoIPFSHash, err := gaillier.Encrypt(pubK, IntToBytes(IPFSHashValue))
	checkError(err)
	homoAESKey, err := gaillier.Encrypt(pubK, IntToBytes(AESKeyHashValue))
	checkError(err)
	homoPubInfo, err := gaillier.Encrypt(pubK, IntToBytes(publisherInfoHashValue))
	checkError(err)

	enDiHash, err := CryptoUtils.RSAEncryption(consumer_RSAPublicKey, restoredDimageIPFSHash)
	checkError(err)
	enCveHash, err := CryptoUtils.RSAEncryption(consumer_RSAPublicKey, restoredCVEIPFSHash)
	checkError(err)
	enAESKey, err := CryptoUtils.RSAEncryption(consumer_RSAPublicKey, restoredAESKey)
	checkError(err)

	enCal := gaillier.Add(pubK, homoCVElistHashValue, homoIPFSHash)
	enCal2 := gaillier.Add(pubK, homoPubInfo, homoAESKey)
	newImageBCID := gaillier.Add(pubK, enCal, enCal2)

	decryptedImageBCID, err := gaillier.Decrypt(privK, newImageBCID)

	smartContractUtils.RegisterDataOnBlockchain(BytesToInt(decryptedImageBCID), enDiHash, enCveHash, enAESKey)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Docker image share time Total cost :", elapsed)
	return decryptedImageBCID
}
