package SDISUtils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	"prototype1.0.0/DockerUtils"
	"prototype1.0.0/IPFSUtils"

	"prototype1.0.0/CryptoUtils"
	"prototype1.0.0/gaillier"
	"prototype1.0.0/smartContractUtils"
)

type publisherInfo struct {
	ID  string `json:"id"`
	PWD string `json:"pwd"`
}

func SecureDockerImageUpload(dockerImageName string) int {
	start := time.Now()

	// 초기 설정값-------------------------------------------------------------
	IPFSUtils.InitIPFSShell()

	enHomoPubKey, err := ioutil.ReadFile("./homoPublicKey")
	enHomoPrivKey, err := ioutil.ReadFile("./homoPrivKey")

	pubK := &gaillier.PubKey{}
	pubDec := gob.NewDecoder(bytes.NewReader(enHomoPubKey))
	err = pubDec.Decode(pubK)

	privK := &gaillier.PrivKey{}
	privDec := gob.NewDecoder(bytes.NewReader(enHomoPrivKey))
	err = privDec.Decode(privK)

	// privPEM, err := ioutil.ReadFile("./RSAPrivateKey.pem")
	pubPEM, err := ioutil.ReadFile("./RSAPublicKey.pem")

	// RSAPrivateKey, err := CryptoUtils.ParseRsaPrivateKeyFromPemStr(string(privPEM))
	RSAPublicKey, err := CryptoUtils.ParseRsaPublicKeyFromPemStr(string(pubPEM))

	//---------------------------------------------------------------------------

	// 1. CVE list 를 추출해낸다
	cvelist := DockerUtils.RunDockerSynkScan(dockerImageName)
	if cvelist == "Denied" {
		fmt.Println("Denied")
		return -1
	}
	// err = ioutil.WriteFile("./cvelist", []byte(cvelist), 0777)
	// checkError(err)
	pubInfo := &publisherInfo{
		ID:  "seunggin",
		PWD: "123456",
	}

	marPubInfo, _ := json.Marshal(pubInfo)
	// 2. 도커 이미지 암호화를 위한 세션 키를 위한 값 을 만들기
	metadata := CryptoUtils.Hash(cvelist)

	// 3. 도커 이미지를 파일로 변환 및 암호화 해준다.
	now := start.Format("2006-01-02 15:04:05")
	DockerUtils.DockerImageToTar(dockerImageName+now+".tar", dockerImageName)
	AESKey := CryptoUtils.CheckKey(fmt.Sprint(metadata))

	CryptoUtils.AesEncryptFile(dockerImageName+now+".tar", "enImage", AESKey)
	CryptoUtils.AesEncryptString(cvelist, dockerImageName+now+"_cvelist", AESKey)

	IPFScvelistfileName := dockerImageName + now + "_cvelist"
	// dockerFileHash := "QmUVzrxhE58gUL5gEKiyUpSx8UJNNPfZYHkhWPxZ8a7rig"
	// cvelistHash := "QmUVzrxhE58gUL5gEKiyUpSx8UJNNPfZYHkhWPxZ8a7pig"
	dockerFileHash := IPFSUtils.AddFileOnIPFS("enimage")
	cvelistHash := IPFSUtils.AddFileOnIPFS(IPFScvelistfileName)

	// 4. 각각의 값들을 RSA 암호화 해준다.

	enDiHash, err := CryptoUtils.RSAEncryption(RSAPublicKey, []byte(dockerFileHash))
	checkError(err)
	enCveHash, err := CryptoUtils.RSAEncryption(RSAPublicKey, []byte(cvelistHash))
	checkError(err)
	enAESKey, err := CryptoUtils.RSAEncryption(RSAPublicKey, AESKey)
	checkError(err)

	UCVEHashValue := CryptoUtils.Hash(cvelist)
	CVEHashValue := int(UCVEHashValue) / 1000
	UIPFSHashValue := CryptoUtils.Hash(dockerFileHash)
	IPFSHashValue := int(UIPFSHashValue) / 1000
	UAESKeyHashValue := CryptoUtils.Hash(fmt.Sprint(AESKey))
	AESKeyHashValue := int(UAESKeyHashValue) / 1000
	UpublisherInfoHashValue := CryptoUtils.Hash(string(marPubInfo))
	publisherInfoHashValue := int(UpublisherInfoHashValue) / 1000

	// fmt.Println("=================================================================")
	// fmt.Println("CVEHashValue:\t\t\t\t\t", CVEHashValue)
	// fmt.Println("IPFSHashValue:\t\t\t\t\t", IPFSHashValue)
	// fmt.Println("AESKeyHashValue:\t\t\t\t", AESKeyHashValue)
	// fmt.Println("publisherInfoHashValue:\t\t\t\t", publisherInfoHashValue)
	// fmt.Println("=================================================================")

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
	imageBCID := gaillier.Add(pubK, enCal, enCal2)

	// homoRestoredCVElist, err := gaillier.Decrypt(privK, homoCVElistHashValue)
	// homoRestoredHash, err := gaillier.Decrypt(privK, homoIPFSHash)
	// homoRestoredAESKey, err := gaillier.Decrypt(privK, homoAESKey)
	// homoRestoredPubInfo, err := gaillier.Decrypt(privK, homoPubInfo)

	// deEnCal, err := gaillier.Decrypt(privK, enCal)
	// deEnCal2, err := gaillier.Decrypt(privK, enCal2)

	// fmt.Println("homoRestoredCVElist :\t\t\t\t", BytesToInt(homoRestoredCVElist))
	// fmt.Println("homoRestoredHash : \t\t\t\t", BytesToInt(homoRestoredHash))
	// fmt.Println("homoRestoredHomoRestoredAESKey :\t\t", BytesToInt(homoRestoredAESKey))
	// fmt.Println("homoRestoredPubInfo :\t\t\t\t", BytesToInt(homoRestoredPubInfo))
	// fmt.Println("=================================================================")

	// restoredcveIPFSHash, err := CryptoUtils.RSADecryption(RSAPrivateKey, enCveHash)
	// checkError(err)

	// smartContractUtils.RegisterDataOnBlockchain(string(imageBCID), string(enDiHash), string(enCveHash), string(enAESKey))

	// restoredCVEIPFSHash, err := CryptoUtils.RSADecryption(RSAPrivateKey, enDiHash)
	// checkError(err)
	// restoredDimageIPFSHash, err := CryptoUtils.RSADecryption(RSAPrivateKey, enDiHash)
	// checkError(err)
	// restoredAESKey, err := CryptoUtils.RSADecryption(RSAPrivateKey, enAESKey)
	// checkError(err)

	decryptedImageBCID, err := gaillier.Decrypt(privK, imageBCID)

	smartContractUtils.RegisterDataOnBlockchain(BytesToInt(decryptedImageBCID), enDiHash, enCveHash, enAESKey)

	// fmt.Println("restoredAESKey :\t\t\t\t\t", string(restoredAESKey))
	// fmt.Println("restoredIPFSHash :\t\t\t\t", string(restoredDimageIPFSHash))
	// fmt.Println("restoredCVEIPFSHash :\t\t\t\t", string(restoredCVEIPFSHash))
	// fmt.Println("=================================================================")

	// addResult := (CVEHashValue + IPFSHashValue + AESKeyHashValue + publisherInfoHashValue)
	//
	// if addResult == BytesToInt(decryptedImageBCID) {
	//
	// 	fmt.Println("Add result: \t\t\t\t\t", addResult)
	// 	fmt.Println("Decrypted imageBCID :\t\t\t\t", BytesToInt(decryptedImageBCID))
	// 	fmt.Println("=================================================================")
	// 	IPFSUtils.GetFileFromIPFS(string(restoredDimageIPFSHash), "enimage")
	// 	IPFSUtils.GetFileFromIPFS(string(restoredCVEIPFSHash), "enCvelist")
	// 	CryptoUtils.AesDecryptFile("enimage", "newFile.tar", restoredAESKey)
	// 	CryptoUtils.AesDecryptFile("enCvelist", "cvelist", restoredAESKey)
	//
	// } else {
	// 	fmt.Println("The ImageBCID doesn`t match so you are not authroized")
	// }
	// DockerUtils.RemoveDockerImage("alpine")

	// list := DockerUtils.ListDockerImage()
	//
	// for i, v := range list {
	// 	fmt.Println(i+1, " ", v.RepoTags)
	// }
	// DockerUtils.DockerImageLoad("newFile.tar")
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Docker image upload time Total cost :", elapsed)
	return BytesToInt(decryptedImageBCID)
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func IntToBytes(i int) []byte {
	if i > 0 {
		return append(big.NewInt(int64(i)).Bytes(), byte(1))
	}
	return append(big.NewInt(int64(i)).Bytes(), byte(0))
}
func BytesToInt(b []byte) int {
	if b[len(b)-1] == 0 {
		return -int(big.NewInt(0).SetBytes(b[:len(b)-1]).Int64())
	}
	return int(big.NewInt(0).SetBytes(b[:len(b)-1]).Int64())
}
