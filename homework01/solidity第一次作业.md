solidity第一次作业



```solidity

pragma solidity ^0.6.0;

contract UserManager {
    address payable public owner;
    
    mapping(uint8 => string) accounts;    //id -> passwd
    mapping(uint8 => address) ips;       //IP address
    
    event Login(uint8 id, uint time);
    event Register(uint8 id ,string passwd);
    event SetPassword(string passwd,uint time);
    
    constructor () public {
        owner = msg.sender;
    }

    function login(uint8 id, string memory passwd) public returns (bool) {
        require(ips[id] == msg.sender);
        if (keccak256(bytes(accounts[id])) == keccak256(bytes(passwd))) {
            emit Login(id, now);
            return true; 
        }
        return false;
    }
    
    function getIP(uint8 id) public view returns (address) {
        require(ips[id] != address(0));
        return ips[id];
    }
    
    function register(uint8 id,string memory passwd) public returns (bool) {
        if(keccak256(bytes(accounts[id]))!=''){
            accounts[id]=passwd;
            emit  Register(id,passwd);
        return true;
        }else{
            return false;
        }
    }
    
    function setPassword(uint8 id,string memory Usedpasswd,string memory passwd) public returns (bool) {
         if (keccak256(bytes(accounts[id])) == keccak256(bytes(Usedpasswd))) {
            accounts[id]=passwd;
            emit SetPassword(passwd,now);
            return true;
        }else{
            return false;
        }
    }
  }

```





