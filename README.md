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
ID：用户ID 
Name: 姓名 
IdentificationType: 证件类型 
Identification: 证件号码
Sex: 性别 
Birthday：生日 
BankCard:银行卡号 
PhonoNumber:手机号 
Key: 秘钥
```
###资金类
```
ID: 银行卡号 
Amount: 卡上剩余金额
```
###产品类
```
ID : 产品编号 
ProductName: 产品名称 
ProductType: 产品类型 
OrganizationID:产品所属机构ID 
Portion:产品份额
```
###机构类
```
ID：机构ID 
OrganizationName:机构名称 
OrganizationType:机构类型
```
###交易内容
```
ID：交易ID 
Trans_type:交易类型 
TransStatus:交易状态 
FromType:交易发起方类型 
FromID：交易发起方ID 
ToType:交易接收方类型 
ToID:交易接收方ID 
ProductID：交易产品ID 
Account:份额 
TransDate:交易时间 
ParentOrderNo:父订单ID
```
###入链协议类
```
SID：业务系统ID 
ReceiverSID:下游系统ID 
OriginSID：来源系统ID 
RequestSerial:来源请求流水号 
NextRequestSerial:下游请求流水号 
Time:入链时间
```

##接口设计
```
CreateUser #创建用户 
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
CreateOrgainization #创建机构 
request 参数: 
args[0] :机构ID 
args1 :机构名称 
args2: 机构类型 
response 参数: 
{” ID “:”XXX”,” OrganizationName “:”xxx”,” OrganizationType”:”xxx”}
```
```
CreateProduct #创建产品 
request 参数: 
args[0] :产品ID 
args1 :产品名称 
args2: 产品类型 
args3:产品所属机构 
args4:产品份额 
response 参数: 
{” ID “:”XXX”,” ProductName “:”xxx”,” ProductType”:”xxx” ，”OrganizationID”：”xxx”，”Portion”：”xxx” }
```
```
Transaction # 交易 
request 参数 
args[0]：交易ID 
args1 :交易类型 0，表示申购 1，表示赎回 
args2:交易状态 
args3:交易发起方类型 
args4：交易发起方ID 
args5:交易接收方类型 
args6:交易接收方ID 
args[7]：交易产品ID 
args[8]:份额 
args[9]:交易时间 
response 参数： 
{ ” ID “:”XXX”, 
“Trans_type “:”XXX”,” 
TransStatus “:”XXX”, 
“FromType “:”XXX”,” 
FromID “:”XXX”, 
“ToType “:”XXX”, 
“ToID “:”XXX”, 
“ProductID “:”XXX”, 
“Account “:”XXX”, 
“TransDate“：”XXX“}
```
```
getTransaction #获取所有交易 
request 参数
response 参数： 
{ ” ID “:”XXX”, 
“Trans_type “:”XXX”,” 
TransStatus “:”XXX”, 
“FromType “:”XXX”,” 
FromID “:”XXX”, 
“ToType “:”XXX”, 
“ToID “:”XXX”, 
“ProductID “:”XXX”, 
“Account “:”XXX”, 
“TransDate“：”XXX“}
```
```
getTransactionByID #获取某笔交易 
request 参数 
args[0]：交易ID 
response 参数： 
{ ” ID “:”XXX”, 
“Trans_type “:”XXX”,” 
TransStatus “:”XXX”, 
“FromType “:”XXX”,” 
FromID “:”XXX”, 
“ToType “:”XXX”, 
“ToID “:”XXX”, 
“ProductID “:”XXX”, 
“Account “:”XXX”, 
“TransDate“：”XXX“}
```
```
getProduct #获取产品信息 
request 参数: 
args[0] :产品ID
response 参数: 
{” ID “:”XXX”,” ProductName “:”xxx”,” ProductType”:”xxx” ，”OrganizationID”：”xxx”，”Portion”：”xxx” }
```
```
getOrganization #获取机构信息 
request 参数: 
args[0] :机构ID
response 参数: 
{” ID “:”XXX”,” OrganizationName “:”xxx”,” OrganizationType”:”xxx”}
```
```
getUser #获取用户信息
request 参数: 
args[0]：用户ID
response 参数: 
{ “ID”:”XXX”, 
” Name”:”XXX”, 
“Identification_type”:”XXX”, 
“Identification”:”XXX”, 
“Sex”:”XXX”, 
“Birthday”:”XXX”, 
“BankCard”:”XXX”, 
“PhonoNumber”:”XXX” 
}
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
```
getUserAsset #查询用户资产 
request 参数 
args[0] 用户ID 
response 参数： 
{“ID”:”XXX”, 
” Name”:”XXX”, 
“Identification_type”:”XXX”, 
“Identification”:”XXX”, 
“Sex”:”XXX”, 
“Birthday”:”XXX”, 
“BankCard”:”XXX”, 
“PhonoNumber”:”XXX”, 
” ID “:”XXX”, 
“ProductName “:”xxx”, 
” ProductType”:”xxx” ， 
“OrganizationID”： 
“xxx”，”Portion”：”xxx”}
```
