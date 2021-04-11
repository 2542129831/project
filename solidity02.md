```solidity
pragma solidity ^0.6.0;
pragma experimental ABIEncoderV2;

contract PersonDID {
    
    modifier Onlyadmin(){                                   //给予管理员权限
           require(Admin[msg.sender] == true);
        _;
    }
    
    struct Person {
        address ip;
        uint8 id;
        uint8 age;
        string name;
        string info;
    }
    
    event AddPerson(uint8 id,uint8 age,string name,string info,uint timestamp);
    
    Person admin;
    Person[] persons;
    mapping(address => bool) public Admin;                    //ip与管理员权限的映射
    mapping(address => Person) public PersonInfo;             //ip与成员信息的映射
    mapping(address => bool) public isPersonExsist;           //ip与成员是否存在的映射
    mapping(address => mapping(address => bool)) Jurisdiction;//自身ip与他人ip的映射与授权
    uint8 SeeNumber;                                          //授权的查看次数
        
    constructor (address ip,uint8 id, uint8 age, string memory name) public {
        admin = Person(ip,id,age,name,"");
        Person memory p = Person(ip,id,age,name,"");
        persons.push(p);
        PersonInfo[msg.sender] = p;
        isPersonExsist[msg.sender] = true;
        Admin[msg.sender] = true;
    }
    
    function getNumberOfPersons() view public returns (uint256) {
        return persons.length;
    }
    
    function addPerson(address ip,uint8 id, uint8 age, string memory name,string memory info) public Onlyadmin{    //只允许管理员添加成员信息,返回是否操作成功
        require(!((id == 0) || (age == 0)), "persons info can not be empty!!");
        require(!isPersonExsist[ip], "person can not exsist !!");
        Person memory p = Person(ip,id, age, name,info);
        persons.push(p);
        PersonInfo[ip] = Person(ip,id, age, name,info);
        isPersonExsist[ip] = true;
        emit AddPerson(id,age,name,info,now);
    }
    
    function setPersonAgeMem(address ip, uint8 age) public {
        Person storage p = PersonInfo[ip];
        p.age = age;
    }
    
    function getPersonAge(address ip) public view returns (uint8) {
        return PersonInfo[ip].age;
    }
    
    function Empower(address SeePersonip,uint8 SeeNumber)public returns (bool){
    //授权ip与被授权ip和授权查看的次数
        Jurisdiction[msg.sender][SeePersonip]=true;                                            //给与查看权限
        SeeNumber=SeeNumber;                                                                  //给予查看次数
        return true;                                                                          //返回是否操作成功
    }

    function getPersonInfo(address SeeIp)  public returns (string memory){          //被查看人ip与你的ip
        require(Jurisdiction[SeeIp][msg.sender]);                                    //判断你是否有权限查看该成员信息
        if(SeeNumber>0){                                                            //判断你是否还有查看次数
            SeeNumber--;                                                            //减少一次查看次数
            return PersonInfo[SeeIp].info;                                          //返回查看成员的经历
        }
    }
}
```

