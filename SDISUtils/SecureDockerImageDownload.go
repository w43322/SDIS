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

func SecureDockerImageDownload(imageBCID int) {
	start := time.Now()
	IPFSUtils.InitIPFSShell()
	conInfo := &publisherInfo{
		ID:  "yujin",
		PWD: "123456",
	}

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

	marPubInfo, _ := json.Marshal(conInfo)

	_, enDh, enKey, enCVE := smartContractUtils.RequestDataFromBlockchain(imageBCID)
	// fmt.Println(bcid, enDh, enKey, enCVE)

	conPrivPEM, err := ioutil.ReadFile("./consumer_RSAPrivateKey.pem")
	checkError(err)
	consumer_RSAPrivateKey, err := CryptoUtils.ParseRsaPrivateKeyFromPemStr(string(conPrivPEM))
	checkError(err)

	restoredCVEIPFSHash, err := CryptoUtils.RSADecryption(consumer_RSAPrivateKey, enCVE)
	checkError(err)
	restoredDimageIPFSHash, err := CryptoUtils.RSADecryption(consumer_RSAPrivateKey, enDh)
	checkError(err)
	restoredAESKey, err := CryptoUtils.RSADecryption(consumer_RSAPrivateKey, enKey)
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

	enCal := gaillier.Add(pubK, homoCVElistHashValue, homoIPFSHash)
	enCal2 := gaillier.Add(pubK, homoPubInfo, homoAESKey)
	newImageBCID := gaillier.Add(pubK, enCal, enCal2)
	decryptedImageBCID, err := gaillier.Decrypt(privK, newImageBCID)

	now := start.Format("2006-01-02 15:04:05")
	if BytesToInt(decryptedImageBCID) == imageBCID {
		fmt.Println("Verication success")
		IPFSUtils.GetFileFromIPFS(string(restoredDimageIPFSHash), "RequestedDockerFile")
		IPFSUtils.GetFileFromIPFS(string(restoredCVEIPFSHash), "RequestedCVEFile")
		CryptoUtils.AesDecryptFile("RequestedDockerFile", "con_newFile"+now+".tar", restoredAESKey)
		CryptoUtils.AesDecryptFile("RequestedCVEFile", "con_cvelist"+now, restoredAESKey)
	} else {
		fmt.Println("Do not equal")
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Docker image Download time Total cost :", elapsed)

}
