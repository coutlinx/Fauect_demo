## Faucet 项目:

### 数据库设计:
>
>>amount表:
> >>有time和amount两个字段用于储存每天调用水龙头的转账数目
>  >> 
```mysql 
create table faucet.amount
(
time   time not null,
amount int  not null
);
```
>> transfer表
>  >>有 id address amount IP time 字段  
> 用于统计调用水龙头的地址和ip和接收eth的数目

``` mysql
create table faucet.transfer
    (
        id      int auto_increment
            primary key,
        address text not null,
        amount  int  not null,
        ip      text not null,
        time    time null
    );
```

### 合约设计:
 ```solidity
 // SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract Faucet{
    address administrator;

    uint Amount_ETH;//累计提取ETH数量
    
    uint Contract_Amount;//合约ETH数量

    mapping(address => uint)public Limit;//每个账户对应提取数量

    event withdraw_ETH(address receiver,uint value,uint time);//提取ETH事件

    constructor()payable{
        administrator = msg.sender;
        Contract_Amount = msg.value; //创建合约时就给ETH
    }
    
    function Charge()public payable{
        Contract_Amount += msg.value;
    }

    //====================核心代码======================================
    //提取ETH
    function withdraw(uint value,address payable receiver)public{
        require(Limit[receiver]<=3 ether && value <= 3 ether ,"Up to 3 eth per day");
        receiver.transfer(value);
        Limit[receiver]+= value;
        Contract_Amount-=value;
        emit withdraw_ETH(receiver,value,block.timestamp);
    }
    //====================核心代码======================================
    
    //每日清空数量
    function Empty_Limit(address receiver)public{
        Limit[receiver] = 0;
    }

    //获取合约ETH数量
    function Get_Contract_Amount() public view returns(uint){
        return Contract_Amount;
    }

    //把合约钱全部转出
    function Transfer_Other(address payable NewOne)public{
        require(msg.sender == administrator,"Only administrator can call this method!!!");
            NewOne.transfer(Contract_Amount);
            Contract_Amount = 0;
    }
}
 ```
