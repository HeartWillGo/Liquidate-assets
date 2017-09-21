__
# Liquidate-assets
这是一个关于资产清算的区块链智能合约，主要做机构，用户之间的清算交易。

##主要函数

CreateUser #创建用户

CreateOrgainization #创建机构

CreateProduct #创建产品 

Transaction # 交易 

getTransaction #获取所有交易 

getTransactionByID #获取某笔交易 

getProduct #获取产品信息 

getOrganization #获取机构信息 
getUser #获取用户信息 

writeUser #修改用户信息 

writeOrgainization #修改机构信息 

writeProduct #修改产品信息 

getUserAsset #查询用户资产

##数据结构设计

###用户
```
//用户
type User struct {
	ID                 string `json:"id"`
	Name               string `json:"Name"`
	Identificationtype int    `json:"identificationtype"`
	Identification     string `json:"identification"`
	Sex                int `json:"sex"`
	Birthday           string `json:"birthday"`
	Bankcard           string `json:"bankcard"`
	Phonenumber        string `json:"phonenumber"`
	Token              string `json:"token"`
}
```
###资金类
```
ID: 银行卡号 
Amount: 卡上剩余金额
```
###产品类
```
//产品
type Product struct {
	Productid      string  `json:"productid"`
	Productname    string  `json:"productname"`
	Producttype    int     `json:"producttype"`
	Organizationid string  `json:"organizationid"`
	Amount         float64 `json:"amount"`
	Price          float64 `json:"price"`
}
```
###机构类
```
//机构
type Organization struct {
	OrganizationID   string `json:"organizationid"`   //机构id
	OrganizationName string `json:"organizationname"` //机构名称
	OrganizationType int    `json:"organizationtype"` //机构类型
}
```
###账本信息
```
// 账本数据
type Transaction struct {

	//交易头部
	SID               string `json:"SID"`
	ReceiverSID       string `json:"ReceiverSID"`
	OriginSID         string `json:"OriginSID"`
	RequestSerial     string `json:"RequestSerial"`
	NextRequestSerial string `json:"NextRequestSerial"`
	Proposaltime      int64   `json:"Proposaltime"`
	//交易ID,区块链中的索引
	Transactionid     string `json:"transactionid"`
	Transactiondate   int64    `json:"transactiondate"`
	Parentorder       string `json:"parentorder"`
	Suborder          string `json:"suborder"`
	Payid             string `json:"payid"`
	//交易双方
	Transtype         string `json:"transtype"`
	Fromtype          int    `json:"fromtype"`
	Fromid            string `json:"fromid"`
	Totype            int    `json:"totype"`
	Toid              string `json:"toid"`
	//实际内容
	Productid         string `json:"productid"`
	Productinfo       string `json:"productinfo"`
	Organizationid    string `json:"organizationid"`
	Amount            float64    `json:"amount"`
	Price             float64    `json:"price"`
}
```

##接口设计
```
CreateUser #创建用户 
request 参数: 
args[0]:
{
    "id": "userid8",
    "Name": "JBYNCsmE",
    "identificationtype": 1,
    "identification": "23342",
    "sex": 1,
    "birthday": "20340912",
    "bankcard": "123243",
    "phonenumber": "999999",
    "token": "8"
  }
response:
    nil 
```

```
getUser #查询用户
request 参数: 
args[0] : "user.ID"
response:
 {
        "id": "userid6",
        "Name": "JBYNCsmE",
        "identificationtype": 1,
        "identification": "23342",
        "sex": 1,
        "birthday": "20340912",
        "bankcard": "123243",
        "phonenumber": "999999",
        "token": "6"
      }

  
```

```
CreateProduct #创建产品 
request 参数: 
args[0]: 
{
    "productid": "productid0",
    "productname": "zhaocaibao",
    "producttype": 1,
    "organizationid": "pingan",
    "amount": 3,
    "price": 33
  }
  
response 参数: 
    nil
```
```
getProduct #得到产品
request: 
args[0] : product.Productid
response:
    {
        "productid": "productid0",
        "productname": "zhaocaibao",
        "producttype": 1,
        "organizationid": "pingan",
        "amount": 3,
        "price": 33
      }
```
```
CreateOrgainization #创建机构 
request 参数: 
args[0]:
{
    "organizationid": "pingan0",
    "organizationname": "pingan",
    "organizationtype": 1
  }
  
response 参数: 
    nil
```

```
Transaction # 交易 
request 参数 
args[0]:
 
{
    "SID": "txiddsf",
    "ReceiverSID": "234423",
    "OriginSID": "23423",
    "RequestSerial": "234",
    "NextRequestSerial": "243243",
    "Proposaltime": 1506005289,
    "transactionid": "transactionid7",
    "transactiondate": 1506005289,
    "parentorder": "7",
    "suborder": "7",
    "payid": "7",
    "transtype": "7",
    "fromtype": 1,
    "fromid": "userid2",
    "totype": 1,
    "toid": "1234",
    "productid": "productid0",
    "productinfo": "wegoodi%3",
    "organizationid": "pingan",
    "amount": 4,
    "price": 9
  }
  response:
        nil
```

```
getTransactionByID #根据交易ID获取数据
request
args[0]: "transactionid"

response 参数： 
  {
    "transactionid": "0",
    "transactiondate": 1505887743,
    "parentorder": "0",
    "suborder": "0",
    "payid": "0",
    "transtype": "0",
    "fromtype": 1,
    "fromid": "1",
    "totype": 1,
    "toid": "VjIwPrHi",
    "productid": "0",
    "productinfo": "wegoodi%3",
    "organizationid": "pingan",
    "amount": 4,
    "price": 9
  }

  
```
```
getTransactionByUserID #根据UserID获取某个用户下的所有交易
request
args[0]："userid" 
response ： 
[
  {
    "transactionid": "0",
    "transactiondate": 1505887743,
    "parentorder": "0",
    "suborder": "0",
    "payid": "0",
    "transtype": "0",
    "fromtype": 1,
    "fromid": "1",
    "totype": 1,
    "toid": "VjIwPrHi",
    "productid": "0",
    "productinfo": "wegoodi%3",
    "organizationid": "pingan",
    "amount": 4,
    "price": 9
  },
  {
    "transactionid": "1",
    "transactiondate": 1505887743,
    "parentorder": "1",
    "suborder": "1",
    "payid": "1",
    "transtype": "1",
    "fromtype": 1,
    "fromid": "1",
    "totype": 1,
    "toid": "VjIwPrHi",
    "productid": "1",
    "productinfo": "wegoodi%3",
    "organizationid": "pingan",
    "amount": 4,
    "price": 9
  },...
 ]


```
```
getUserAsset #获取某一用户的资产详情
request 参数： len(args) =1 
args[0]: "userid"
response:
{
    "statistic_date": "1505896172",
    "trading_entity_id": "1",
    "transaction_num": 3,
    "asset_type": "",
    "asset_info": "",
    "trade_start_time": 1505887743,
    "trade_end_time": 1505887743,
    "asset_balance": 108.12,
    "asset_income": 108.12,
    "asset_outcome": 0,
    "organization_Map": {
      "pingan": {
        "id": "pingan",
        "type": 0,
        "transactionnum": 3,
        "tradestarttime": 1505887743,
        "tradeendtime": 1505887743,
        "balance": 0,
        "outcome": 108.12,
        "income": 0,
        "productmap": {
          "0": {
            "id": "0",
            "tradestarttime": 1505887743,
            "tradeendtime": 1505887743,
            "transactionum": 1,
            "balance": 0,
            "outcome": 36,
            "income": 0
          },
          "1": {
            "id": "1",
            "tradestarttime": 1505887743,
            "tradeendtime": 1505887743,
            "transactionum": 1,
            "balance": 0,
            "outcome": 36,
            "income": 0
          },
          "2": {
            "id": "2",
            "tradestarttime": 1505887743,
            "tradeendtime": 1505887743,
            "transactionum": 1,
            "balance": 0,
            "outcome": 36.12,
            "income": 0
          }
          ...
        }
      }
    }
  }

```
```
getUserAllProduct #得到该用户购买的所有产品
request
args[0]:"userid"
response:
   [
     {
       "pingan": {
         "productid": "productid0",
         "productname": "zhaocaibao",
         "producttype": 1,
         "organizationid": "pingan",
         "amount": 3,
         "price": 33
       }
     {
       "productid": "productid1",
       "productname": "zhaocaibao",
       "producttype": 1,
       "organizationid": "pingan",
       "amount": 3,
       "price": 33
     },
     {
       "productid": "productid2",
       "productname": "zhaocaibao",
       "producttype": 1,
       "organizationid": "pingan",
       "amount": 3,
       "price": 33
     }
     }
   ]

```
```
getUserOrgProductid #获取某个机构下产品的所有购买情况
request
args[0]: "organizationid"
args[1]:"userid"

response
  [
   {
     "productid": "productid0",
     "productname": "zhaocaibao",
     "producttype": 1,
     "organizationid": "pingan",
     "amount": 3,
     "price": 33
   }
   {
     "productid": "productid1",
     "productname": "zhaocaibao",
     "producttype": 1,
     "organizationid": "pingan",
     "amount": 3,
     "price": 33
   }
   {
     "productid": "productid2",
     "productname": "zhaocaibao",
     "producttype": 1,
     "organizationid": "pingan",
     "amount": 3,
     "price": 33
   }
 ]
```
```
getUserFromOrganizationAsset #获取用户在某个机构下资产详情
 
request 
args[0]: "organizationid"
args[1]: "userid"

response 参数: 
{
  "id": "pingan",
  "statistic_date": "",
  "type": 0,
  "transactionnum": 3,
  "tradestarttime": 1506005289,
  "tradeendtime": 1506005289,
  "balance": 0,
  "outcome": 108,
  "income": 0,
  "productmap": {
    "productid0": {
      "id": "productid0",
      "statistic_date": "",
      "tradestarttime": 1506005289,
      "tradeendtime": 1506005289,
      "transactionum": 1,
      "balance": 0,
      "outcome": 36,
      "income": 0,
      "asset": null
    },
   ...
  },
  "asset": null
}

```
```
getProductTransactionByProductID #根据产品获取该产品的所有账本条目
request 参数: 
args[0]："productid"
response




```
```
writeUser #修改用户信息 
request 参数: 
args[0]：用户ID 
args1: 姓名 
args2_type: 证件类型 
args3: 证件号码 
args4: 性别 
args5：生日 
args6:银行卡号 
args[7]:手机号 
args[8]: 秘钥 
response 参数: 
{ “ID”:”XXX”, 
” Name”:”XXX”, 
“Identification_type”:”XXX”, 
“Identification”:”XXX”, 
“Sex”:”XXX”, 
“Birthday”:”XXX”, 
“BankCard”:”XXX”, 
“PhonoNumber”:”XXX”, 
“Key”:”XXX”}
```
```
writeOrgainization #修改机构信息 
request 参数: 
args[0] :机构ID 
args1 :机构名称 
args2: 机构类型 
response 参数: 
{” ID “:”XXX”,” OrganizationName “:”xxx”,” OrganizationType”:”xxx”}
```
```
writeProduct #修改产品信息 
request 参数: 
args[0] :产品ID 
args1 :产品名称 
args2: 产品类型 
args3:产品所属机构 
args4:产品份额 
response 参数: 
{” ID “:”XXX”,” ProductName “:”xxx”,” ProductType”:”xxx” ，”OrganizationID”：”xxx”，”Portion”：”xxx” }
```
