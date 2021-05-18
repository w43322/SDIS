//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.4.19 <0.9.0;

// contract address 

contract RegisterData   {
    
    struct dockerData {
        uint bcid ; 
        string enKey;
        string enHash;
    }
    
    event NewDataSet( uint bcid , string enKey , string enHash);
    dockerData[] private dockerDatas;

    mapping (uint => address) public imageToOwner;
    mapping (address => uint[]) public addressToBcid;

    // 도커 이미지 소유권 보장 및 데이터 저장
    function setDockerImageData( uint  bcid , string memory enKey , string memory enHash )public{   
        // uint bcid = dockerDatas.length;
        dockerDatas.push(dockerData(bcid , enKey, enHash));
        imageToOwner[bcid] = msg.sender;
        emit NewDataSet(bcid ,enKey , enHash);
        addressToBcid[msg.sender].push(bcid);
        // push 가 더 이상 숫자를 돌려주지 않기 때문에 길이에 - 1 
    }
    // 사용자의 주소를 통해서 가지고 있는 BCID 를 가지고 온다. 
    function getBCID() view public returns(uint[] memory) { 
        return addressToBcid[msg.sender];
    }
    // BCID 를 통해 암호화된 키를 가지고 온다. 
    function getEnKey(uint bcid) view public returns(string memory) { 
        return dockerDatas[bcid].enKey;
    }
    // BCID 를 통해 암호화된 해시값을 가지고 온다. 
    function getEnHash(uint bcid) view public returns(string memory) { 
        return dockerDatas[bcid].enHash;
    }   
    

}