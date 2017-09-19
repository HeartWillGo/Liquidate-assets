package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"bytes"
	"google.golang.org/genproto/googleapis/bigtable/admin/table/v1"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

//用户
type User struct {
	ID                 string `json:"id"`                 //用户ID
	Name               string `json:"name"`               //用户名字
	IdentificationType int    `json:"identificationType"` // 证件类型
	Identification     string `json:"identification"`     //证件号码
	Sex                int    `json:"sex"`                //性别
	Birthday           string `json:"birthday"`           //生日
	BankCard           string `json:"bankcard"`           //银行卡号
	PhoneNumber        string `json:"phonoumber"`         //手机号
	Token              string `json:"token"`              //密钥

	ProductMap     map[string]Product     `json:"productmap"`     //产品
	TransactionMap map[string]Transaction `json:"transactionmap"` //交易
}
type ProductProcess struct {
	ProcessType   int `json:"processtype"`   //操作类型
	ProcessAmount int `json:"processamount"` //操作份额
	ProcessPrice  int `json:"ProcessPrice"`  //价格
}

//资金
type Fund struct {
	CardID string `json:"cardid"` //银行卡id
	Amount int    `json:"amount"` //卡上剩余金额

}

//产品
type Product struct {
	ProductID      string `json:"productid"`      //产品id
	ProductName    int    `json:"productname"`    //产品名称
	ProductType    int    `json:"producttype"`    //产品类型
	OrganizationID string `json:"organizationid"` //产品所属机构id
	Portion        int    `json:"portion"`        //产品份额
	Price          int    `json:"price"`          //单价

}

//机构
type Organizaton struct {
	OrganizationID   string `json:"organizationid"`   //机构id
	OrganizationName string `json:"organizationname"` //机构名称
	OrganizationType int    `json:"organizationtype"` //机构类型
	ProductMap 	  map[string](map[string]Transaction) `json:"productmap"`

}

//交易内容
type Transaction struct {
	TransID          string `json:"id"`         //交易id
	TransDate          string `json:"transdate"`       //交易时间
	TransType     int    `json:"transtype"`  //交易类型 0 表示申购，1，表示赎回， 2，表示入金
	FromType      int    `json:"fromtype"`   //发送方角色
	FromID        string `json:"fromid"`     //发送方 ID
	ToType        int    `json:"totype"`     //接收方角色
	ToID          string `json:"toid"`       //接收方 ID


	//
	ProductID     string `json:"productid"`  //交易产品ID
	OrganizatonID string `json:"organizationid"` //机构ID
	Account       int    `json:"account"`    //交易 份额
	Price         int    `json:"price"`      //交易 价格
	//订单号
	ParentOrderNo string `json:parentorder"` //父订 单号
	SubOrderNo  string `json:"suborder"`
	PayId 		string `json:"payid"`
}

var err error

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}
	return shim.Success(nil)
}
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function != "invoke" {
		return shim.Error("Unknown function call")
	}
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting at least 2")
	}
	switch {

	case args[0] == "CreateUser":
		return t.CreateUser(stub, args)
	case args[0] == "createOrganization":
		return t.createOrganization(stub, args)
	case args[0] == "CreateProduct":
		return t.CreateProduct(stub, args)
	case args[0] == "getTransactionByID":
		return t.getTransactionByID(stub, args)
	case args[0] == "getProduct":
		return t.getProduct(stub, args)
	case args[0] == "getOrganization":
		return t.getOrganization(stub, args)
	case args[0] == "getUser":
		return t.getUser(stub, args)
	case args[0] == "WriteUser":
		return t.WriteUser(stub, args)
	case args[0] == "WriteOrganization":
		return t.WriteOrganization(stub, args)
	case args[0] == "WriteProduct":
		return t.WriteProduct(stub, args)

	case args[0] == "transation":
		return t.transation(stub, args)
	case args[0] == "getUserAsset":
		return t.getUserAsset(stub, args)
	case args[0] == "query":
		return t.query(stub, args)
	default:
		fmt.Printf("function is not exist\n")
	}

	return shim.Error("Unknown action,")
}

//用户登录验证

func (t *SimpleChaincode) UserLogin(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 UserLogin")

	var userid string   // 用户ID
	var username string //用户名称
	var token string    //用户密钥
	var user User

	if len(args) != 4 {
		return shim.Error("UserLogin :Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode

	userid = args[1]
	username = args[2]
	token = args[3]

	userinfo, err := stub.GetState(userid)
	if err != nil {
		return shim.Error(err.Error())
	}
	if userinfo == nil {
		return shim.Success(nil)
	} else {
		err = json.Unmarshal(userinfo, &user)
		if err != nil {
			return shim.Error(err.Error())
		} else if (username == user.Name) && (token == user.Token) {
			return shim.Success(nil)

		}

	}
	return shim.Success(nil)
}

//用户查询当下的所有产品
func (t *SimpleChaincode) getUserProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 WriteUser")

	var userid string //用户ID
	var user User

	var buffer bytes.Buffer

	if len(args) != 2 {
		return shim.Error("getUserProduct：Incorrect number of arguments. Expecting 2")
	}

	userid = args[1]
	UserInfo, err := stub.GetState(userid)
	if err != nil {
		return shim.Error(err.Error())
	}
	if UserInfo != nil {
		//将byte的结果转换成struct
		err = json.Unmarshal(UserInfo, &user)
		if err != nil {
			return shim.Error(err.Error())
		}
		buffer.WriteString("{")
		bArrayMemberAlreadyWritten := false

		for key, product_value := range user.ProductMap {

			if bArrayMemberAlreadyWritten == true {
				buffer.WriteString(",")
			}
			productbytes, err := json.Marshal(product_value)
			if err != nil {
				return shim.Error("wrong value")
			}
			buffer.WriteString(string(productbytes))


			fmt.Printf("产品：", key, "产品内容：", productbytes)
			bArrayMemberAlreadyWritten = true

		}
		buffer.WriteString("}")

		return shim.Success(buffer.Bytes())

	}
	return shim.Success(nil)

}

//createUser
func (t *SimpleChaincode) CreateUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 CreateUser")
	//
	var ID string              //用户ID
	var Name string            //用户名字
	var IdentificationType int // 证件类型
	var Identification string  //证件号码
	var Sex int                //性别
	var Birthday string        //生日
	var BankCard string        //银行卡号
	var PhoneNumber string     //手机号
	var token string           //密钥

	var user User

	if len(args) != 10 {
		return shim.Error("CreateUser：Incorrect number of arguments. Expecting 10")
	}

	ID = args[1]
	Name = args[2]
	IdentificationType, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：TotalNumber ")
	}
	Identification = args[4]
	Sex, err = strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：TotalNumber ")
	}
	Birthday = args[6]
	BankCard = args[7]
	PhoneNumber = args[8]
	token = args[9]

	user.ID = ID
	user.BankCard = BankCard
	user.Birthday = Birthday
	user.Identification = Identification
	user.IdentificationType = IdentificationType
	user.Name = Name
	user.PhoneNumber = PhoneNumber
	user.Sex = Sex
	user.Token = token
	user.ProductMap = make(map[string]Product)
	user.TransactionMap = make(map[string]Transaction)


	jsons_users, err := json.Marshal(user) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}
	userbytes, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("failed to return userbytes")
	} else  if userbytes != nil {
		return shim.Error(ID + "already have")
	}
	err = stub.PutState(ID, jsons_users)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//用户查询某机构的产品
func (t *SimpleChaincode) getUserProductogOrg(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 WriteUser")

	var userid string //用户ID
	var org_id string //用户ID
	var user User
	var buffer bytes.Buffer

	if len(args) != 3 {
		return shim.Error("getUserProductogOrg number of arguments. Expecting 3")
	}

	userid = args[1]
	org_id = args[2]
	UserInfo, err := stub.GetState(userid)
	if err != nil {
		return shim.Error(err.Error())
	} else if UserInfo == nil {
		return shim.Error(string(UserInfo) + "is not exists")
	}
	err = json.Unmarshal(UserInfo, &user)
	if err != nil {
		return shim.Error("unmarshal user not successful")
	}


	buffer.WriteString("{")
	bArrayMemberAlreadyWritten := false
	
	for key, product_value := range user.ProductMap {



		if product_value.OrganizationID == org_id {

			productbytes, err := json.Marshal(product_value)
			if err != nil {
				return shim.Error("productbytes marshal error")
			}

			if bArrayMemberAlreadyWritten == true {
				buffer.WriteString(",")
			}
			buffer.WriteString(string(productbytes))
			bArrayMemberAlreadyWritten = true
		}

		fmt.Printf("产品：", key, "产品内容：", product_value)

	}
	
	return shim.Success(buffer.Bytes())

}



//createOrganization
func (t *SimpleChaincode) createOrganization(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 createOrganization")

	var OrganizationID string    //机构id
	var OrganizationName string  //机构名称
	var OrganizationType int     //机构类型
	var organization Organizaton //机构

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// Initialize the chaincode
	OrganizationID = args[1]
	OrganizationName = args[2]

	OrganizationType, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	organization.OrganizationID = OrganizationID
	organization.OrganizationName = OrganizationName
	organization.OrganizationType = OrganizationType

	jsons_organization, err := json.Marshal(organization) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}

	orginfo, err := stub.GetState(OrganizationID)
	if err != nil {
		return shim.Error("failed to get orginfo")
	} else if orginfo != nil {
		return shim.Error(string(orginfo) + "\t is already exists")
	}
	// Write the state to the ledger
	err = stub.PutState(args[1], jsons_organization)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//createProduct 创建产品
func (t *SimpleChaincode) CreateProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 CreateProduct")

	var ProductID string      //产品id
	var ProductName int       //产品名称
	var ProductType int       //产品类型
	var OrganizationID string //产品所属机构id
	var Portion int           //产品份额
	var Price int             //价格
	var product Product

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	ProductID = args[1]
	ProductName, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	ProductType, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	OrganizationID = args[4]
	Portion, err = strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	Price, err = strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	product.ProductID = ProductID
	product.ProductName = ProductName
	product.ProductType = ProductType
	product.OrganizationID = OrganizationID
	product.Portion = Portion
	product.Price = Price

	jsons_product, err := json.Marshal(product) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}
	
	productinfo, err := stub.GetState(ProductID)
	if err != nil {
		return shim.Error("failed to get productID")
	} else if productinfo != nil {
		return shim.Error(string(productinfo) + "\t already exists")
	}
	
	// Write the state to the ledger
	err = stub.PutState(args[1], jsons_product)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	//
	////在机构结构中添加产品ID
	//var organization Organizaton
	////取出机构结构体
	//orginfo, err := stub.GetState(OrganizationID)
	//if err != nil {
	//	return shim.Error("failed to get the orginfo")
	//} else if orginfo == nil {
	//	return shim.Error("does not exists")
	//}
	//json.Unmarshal(orginfo, &organization)
	////更改机构结构体
	//organization.ProductMap[OrganizationID] = product
	//
	//org_add_product, err := json.Marshal(organization)
	////存入更改后的结构体
	//err = stub.PutState(OrganizationID, org_add_product)
	//if err != nil {
	//	return shim.Error(err.Error())
	//}
	return shim.Success(nil)
}

//getTransactionByID 获取某笔交易
func (t *SimpleChaincode) getTransactionByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 getTransactionByID")

	var Transactin_ID string //交易ID
	var transaction Transaction
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	Transactin_ID = args[1]

	TransactionInfo, err := stub.GetState(Transactin_ID)
	if err != nil {
		return shim.Error(err.Error())
	}
	//将byte的结果转换成struct
	err = json.Unmarshal(TransactionInfo, &transaction)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("  TransactionInfo  = %d  \n", TransactionInfo)

	return shim.Success(TransactionInfo)
}

//getProduct 获取产品信息
func (t *SimpleChaincode) getProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 getProduct")

	var Product_ID string //产品ID
	var product Product
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	Product_ID = args[1]

	ProductInfo, err := stub.GetState(Product_ID)
	if err != nil {
		return shim.Error(err.Error())
	}
	//将byte的结果转换成struct
	err = json.Unmarshal(ProductInfo, &product)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("  ProductInfo  = %d  \n", ProductInfo)
	return shim.Success(ProductInfo)
}

//getOrganization 获取机构信息
func (t *SimpleChaincode) getOrganization(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 getOrganization")

	var Organization_ID string // 商业银行ID
	var organization Organizaton

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode

	Organization_ID = args[1]

	OrganizationInfo, err := stub.GetState(Organization_ID)
	if err != nil {
		return shim.Error(err.Error())
	}
	//将byte的结果转换成struct
	err = json.Unmarshal(OrganizationInfo, &organization)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("  OrganizationInfo  = %d  \n", OrganizationInfo)

	return shim.Success(OrganizationInfo)
}

//getUser 获取用户信息
func (t *SimpleChaincode) getUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 getUser")

	var User_ID string // 用户ID
	var user User

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode

	User_ID = args[1]
	userinfo, err := stub.GetState(User_ID)
	if err != nil {
		return shim.Error(err.Error())
	}
	//将byte的结果转换成struct
	err = json.Unmarshal(userinfo, &user)

	fmt.Printf("  userinfo  = %d  \n", userinfo)

	return shim.Success(userinfo)
}

//writeUser  修改用户信息,全部更改?
func (t *SimpleChaincode) WriteUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 WriteUser")

	var ID string              //用户ID
	var Name string            //用户名字
	var IdentificationType int // 证件类型
	var Identification string  //证件号码
	var Sex int                //性别
	var Birthday string        //生日
	var BankCard string        //银行卡号
	var PhoneNumber string     //手机号

	
	// check if user 
	var user User
	userinfo , err := stub.GetState(ID) 
	if err != nil {
		return shim.Error("failed to get userinfo")
	}	else if userinfo == nil {
		return shim.Error("you should first create user!!!")
	}
	json.Unmarshal(userinfo, &user)
	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	ID = args[1]
	Name = args[2]
	IdentificationType, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：TotalNumber ")
	}
	Identification = args[4]
	Sex, err = strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：TotalNumber ")
	}
	Birthday = args[6]
	BankCard = args[7]
	PhoneNumber = args[8]

	user.ID = ID
	user.BankCard = BankCard
	user.Birthday = Birthday
	user.Identification = Identification
	user.IdentificationType = IdentificationType
	user.Name = Name
	user.PhoneNumber = PhoneNumber
	user.Sex = Sex

	jsons_users, err := json.Marshal(user) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}
	
	// Write the state to the ledger
	err = stub.PutState(args[1], jsons_users)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf(" CeateBank success \n")
	return shim.Success(nil)
}

//WriteOrganization         修改机构
func (t *SimpleChaincode) WriteOrganization(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 WriteOrganization")

	var OrganizationID string      //机构id
	var OrganizationName string    //机构名称
	var OrganizationType int       //机构类型
	
	
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// Initialize the chaincode
	OrganizationID = args[1]
	OrganizationName = args[2]

	OrganizationType, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}

	//check org if exists
	var organization Organizaton

	orginfo , err := stub.GetState(OrganizationID)
	if err != nil {
		return shim.Error("failed to get org info")
	} else if orginfo == nil {
		return shim.Error("org is not exists")
	}

	err = json.Unmarshal(orginfo, &organization)
	if err != nil {
		return shim.Error("cannot marshal orgazation")
	}
	organization.OrganizationID = OrganizationID
	organization.OrganizationName = OrganizationName
	organization.OrganizationType = OrganizationType

	jsons_organization, err := json.Marshal(organization) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}
	// Write the state to the ledger
	err = stub.PutState(args[1], jsons_organization)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("CreateOrganization \n")

	return shim.Success(nil)
}

//WriteProduct 修改产品
func (t *SimpleChaincode) WriteProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 WriteProduct")

	var ProductID string      //产品id
	var ProductName int       //产品名称
	var ProductType int       //产品类型
	var OrganizationID string //产品所属机构id
	var Portion int           //产品份额
	var product Product

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	ProductID = args[1]
	ProductName, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	ProductType, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	OrganizationID = args[4]
	Portion, err = strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	productinfo , err := stub.GetState(ProductID)
	if err != nil {
		return shim.Error("failed to get productinfo")
	} else if productinfo == nil {
		return shim.Error("product not exists")
	}
	err = json.Unmarshal(productinfo, &product)

	product.ProductID = ProductID
	product.ProductName = ProductName
	product.ProductType = ProductType
	product.OrganizationID = OrganizationID
	product.Portion = Portion

	jsons_product, err := json.Marshal(product) //转换成JSON返回的是byte[]
	if err != nil {
		return shim.Error(err.Error())
	}
	// Write the state to the ledger
	err = stub.PutState(args[1], jsons_product)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf(" CreateProduct success \n")
	return shim.Success(nil)
}

//Transation交易
func (t *SimpleChaincode) transation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("put order in ledger")

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments")
	}
	var transaction Transaction
	transactionBytes := args[1]
	err  = json.Unmarshal(bytes(transactionBytes), &transaction)
	if err != nil {
		return shim.Error("wrong get order")
	}
	err = stub.PutState(transaction.TransID, bytes(transactionBytes))
	if err != nil {
		return shim.Error(err.Error())
	}
	NameIndexKey,err :=  stub.CreateCompositeKey("TransactionType", []string{TransactionType.TransactionID,
		TransactionType.FromID, TransactionType.ToID})

	if err != nil {
		return shim.Error(err.Error())
	}

	value := []byte{0x00}
	stub.PutState(NameIndexKey, value)

	// ================
	return shim.Success(nil)

// Deletes an entity from state

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[1]

	// Get the state from the ledger
	Avalbytes, erro := stub.GetState(A)
	if erro != nil {
		return shim.Error(erro.Error())
	}
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
