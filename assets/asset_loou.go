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
//type Ordered struct {
//	//订单类
//	//MainOrderNo string `json:"parentorder"` //父订单号
//	//SubOrderNo  string `json:"suborder"`
//	//PayId 		string `json:"payid"`
//
//	//交易双方
//	TransactionID string `json:"transactionid"` //交易id
//	//TransDate     string `json:"time"`       //交易时间
//	TransType     int    `json:"transtype"`  //交易类型 0 表示申购，1，表示赎回， 2，表示入金
//	//FromType      int    `json:"fromtype"`   //发送方角色
//	FromID        string `json:"fromid"`     //发送方 ID
//	//ToType        int    `json:"totype"`     //接收方角色
//	ToID          string `json:"toid"`       //接收方 ID
//
//	//交易内容
//	//ProductID      string `json:"productid"`      //交易产品id
//	//Productinfo    string `json:"productinfo"`	  //chanpinxiangqing
//	//OrganizationID string `json:"organizationid"` //机构ID
//	//Account        int    `json:"account"`        //交易 份额
//	//Price          int    `json:"price"`        //交易价格
//}
type Ordered struct {
	Parentorder    string `json:"parentorder"`
	Suborder       string `json:"suborder"`
	Payid          string `json:"payid"`
	Transactionid  string `json:"transactionid"`
	Transtype      string `json:"transtype"`
	Fromtype       int    `json:"fromtype"`
	Fromid         string `json:"fromid"`
	Totype         int    `json:"totype"`
	Toid           string `json:"toid"`
	Productid      string `json:"productid"`
	Productinfo    string `json:"productinfo"`
	Organizationid string `json:"organizationid"`
	Account        int    `json:"account"`
	Price          int    `json:"price"`
}
func (t *SimpleChaincode) Init (stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	fmt.Println("args is " , len(args), args)

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke (stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("0x02 invoke")

	function, args := stub.GetFunctionAndParameters()
	fmt.Println("function", function)
	fmt.Println("args", args)
	if function != "invoke" {
		return shim.Error("unkown function call")
	}
	if len(args) < 1 {
		return shim.Error("not enough arguments")
	}
	switch {
	case args[0] == "transaction":
		fmt.Println("get this")
		return t.putTransaction(stub, args)

	case args[0] == "query":
		return t.query(stub, args)
	}

	return shim.Error("unknown action.")
}


func (t *SimpleChaincode) putTransaction(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	fmt.Println("put order in ledger")

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments")
	}
	var order Ordered
	orderinfo := args[1]
	err  := json.Unmarshal([]byte(orderinfo), &order)
	if err != nil {
		return shim.Error("wrong get order")
	}
	err = stub.PutState(order.Transactionid, []byte(orderinfo))
	if err != nil {
		return shim.Error(err.Error())
	}
	NameIndexKey,err :=  stub.CreateCompositeKey("Ordertype", []string{order.Transactionid,
													order.Fromid, order.Toid})

	if err != nil {
		return shim.Error(err.Error())
	}

	value := []byte{0x00}
	stub.PutState(NameIndexKey, value)

	// ================
	return shim.Success(nil)

}

func (t *SimpleChaincode)getTransaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("0x02 get orTranction from ledger")

	if len(args) != 2 {
		return shim.Error("Incorrect number of argument")
	}

	transID := args[1]

	orderBytes, err:= stub.GetState(transID)
	if err != nil {
		return shim.Error("failed to get transid")
	}

	return shim.Success(orderBytes)
}

func (t *SimpleChaincode)getCompositeTransaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("0x02 get orTranction from ledger")

	if len(args) != 2 {
		return shim.Error("Incorrect number of argument")
	}




	// Query the color~name index by color
	// This will execute a key range query on all keys starting with 'color'
	transactionResultsIterator, err := stub.GetStateByPartialCompositeKey("Ordertype", args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	defer transactionResultsIterator.Close()

	// Iterate through result set and for each marble found, transfer to newOwner
	var i int
	for i = 0; transactionResultsIterator.HasNext(); i++ {
		// Note that we don't get the value (2nd return variable), we'll just get the marble name from the composite key
		responseRange, err := transactionResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		// get the color and name from color~name composite key
		objectType, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}
		returnedColor := compositeKeyParts[0]
		returnedMarbleName := compositeKeyParts[1]
		fmt.Printf("- found a marble from index:%s color:%s name:%s\n", objectType, returnedColor, returnedMarbleName)

		// Now call the transfer function for the found marble.
		// Re-use the same function that is used to transfer individual marbles
		}

		return shim.Success(nil)
	}
func (t *SimpleChaincode) getTransactionByDate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("0x02 get orTranction from ledger")

	if len(args) != 2 {
		return shim.Error("Incorrect number of argument")
	}

	fromID := args[1]


	// Query the color~name index by color
	// This will execute a key range query on all keys starting with 'color'
	transactionResultsIterator, err := stub.GetStateByPartialCompositeKey("Ordertype", []string{fromID})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer transactionResultsIterator.Close()

	// Iterate through result set and for each marble found, transfer to newOwner
	var i int
	for i = 0; transactionResultsIterator.HasNext(); i++ {
		// Note that we don't get the value (2nd return variable), we'll just get the marble name from the composite key
		responseRange, err := transactionResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		// get the color and name from color~name composite key
		objectType, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}
		returnedColor := compositeKeyParts[0]
		returnedMarbleName := compositeKeyParts[1]
		fmt.Printf("- found a marble from index:%s color:%s name:%s\n", objectType, returnedColor, returnedMarbleName)

		// Now call the transfer function for the found marble.
		// Re-use the same function that is used to transfer individual marbles
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	var err error
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}


	// Get the state from the ledger
	transaction, err := stub.GetState(args[1])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + args[1] + "\"}"
		return shim.Error(jsonResp)
	}

	if transaction == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + args[1] + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + args[1] + "\",\"Amount\":\"" + string(transaction) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(transaction)
}
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
