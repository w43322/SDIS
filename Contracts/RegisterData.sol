//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.4.19 <0.9.0;

// contract address

contract RegisterData   {

    struct dockerData {
        uint num ;
        uint bcid ;
        bytes enKey;
        bytes enCVEHash ;
        bytes enHash;
    }

    event NewDataSet( uint num , uint bcid , bytes enKey ,bytes cveHash, bytes enHash);
    dockerData[] private dockerDatas ;

    mapping (uint => address) public imageToOwner;
    mapping (address => uint[]) public addressToBcid;

    // 도커 이미지 소유권 보장 및 데이터 저장
    function setDockerImageData( uint bcid , bytes memory enKey , bytes memory enHash , bytes memory enCVE )public{
        uint num = dockerDatas.length;
        dockerDatas.push(dockerData( num ,bcid , enKey, enCVE , enHash));
        imageToOwner[num] = msg.sender;
        emit NewDataSet(num,bcid,enKey ,enCVE, enHash);
        addressToBcid[msg.sender].push(num);
        // push 가 더 이상 숫자를 돌려주지 않기 때문에 길이에 - 1
    }
    // 사용자의 주소를 통해서 가지고 있는 BCID 를 가지고 온다.
    function getNum() view public returns( uint[] memory){
        return addressToBcid[msg.sender];
    }
    // BCID 를 돌려준다.
    function getBCID(uint num) view public returns(uint ) {
        return dockerDatas[num].bcid;
    }
    // BCID 를 통해 암호화된 키를 가지고 온다.
    function getEnKey(uint num) view public returns(bytes memory) {
        return dockerDatas[num].enKey;
    }
    // BCID 를 통해 암호화된 해시값을 가지고 온다.
    function getEnHash(uint num) view public returns(bytes memory) {
        return dockerDatas[num].enHash;
    }
    function getEnCVE( uint num) view public returns(bytes memory){
        return dockerDatas[num].enCVEHash;
    }


}
