package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"fmt"
)

//交易内容


type SimpleChaincode struct {

}

//拟采用这种方式，只有一个orderer详情入账，至于他属于谁不重要，重要的是查询的复杂度不高就可以
type Ordered struct {
	//订单类
	MainOrderNo string `json:"parentorder"` //父订单号
	SubOrderNo  string `json:"suborder"`
	PayId 		string `json:"payid"`

	//交易双方
	TransactionID int    `json:"transactionid"` //交易id
	TransDate     string `json:"time"`       //交易时间
	TransType     int    `json:"transtype"`  //交易类型 0 表示申购，1，表示赎回， 2，表示入金
	FromType      int    `json:"fromtype"`   //发送方角色
	FromID        string `json:"fromid"`     //发送方 ID
	ToType        int    `json:"totype"`     //接收方角色
	ToID          string `json:"toid"`       //接收方 ID

	//交易内容
	ProductID      string `json:"productid"`      //交易产品id
	OrganizationID string `json:"organizationid"` //机构ID
	Account        int    `json:"account"`        //交易 份额
	Price          int    `json:"price"`        //交易价格
}


func (t *SimpleChaincode) Init (stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke (stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("0x02 invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		return shim.Error("unkown function call")
	}
	if len(args) < 2 {
		return shim.Error("not enough arguments")
	}
	switch {
	case args[0] == "Transaction":
		return t.Transaction(stub, args)

	case args[0] == "query":
		return t.query(stub, args)
	}

	return shim.Error("unknown action.")
}


func (t *SimpleChaincode) Transaction(stub shim.ChaincodeStubInterface) pb.Response{
	fmt.Println("put order in ledger")
	var



}