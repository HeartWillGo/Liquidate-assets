package main

import (
	"encoding/json"
	"fmt"

	"bytes"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)



//transaction json
type Transaction struct {


	SID string `json:"sID"`
	ReceiverSID string `json:"receiverSID"`
	OrigSID string `json:"origSID"`
	RequestSerial string `json:"requestSerial"`
	NextRequestSerial string `json:"nextRequestSerial"`

	TransDate string `json:"transDate"`
	TransStatus string `json:"transStatus"`
	//OrderNo 主键
	OrderNo string `json:"orderNo"`
	TransType string `json:"transType"`
	LoanType string `json:"loanType"`
	ParentOrderNo string `json:"parentOrderNo"`


	Amount string `json:"amount"`
	AccountType string `json:"accountType"`
	Account string `json:"account"`

	ProductCode string `json:"productCode"`
	OrganizationCode string `json:"organizationCode"`
	Portion string `json:"portion"`


	AccountName string `json:"accountName"`
	IDNo string `json:"idNo"`
	IDType string `json:"idType"`
	Sex string `json:"sex"`
	Birthday string `json:"birthday"`
	PhoneNo string `json:"phoneNo"`
	IDKey string `json:"idKey"`
	Message string `json:"message"`
	Desc string `json:"desc"`
}



//交易信息入链,创建索引信息
//args[0] functionname string
//args[1] IDNo string
//args = []string{"Transaction", "json格式的交易数据"}
func (t *SimpleChaincode) Transaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("0x0301 Transaction")

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments")
	}
	var transaction Transaction
	transactionBytes := args[1]
	err := json.Unmarshal([]byte(transactionBytes), &transaction)
	if err != nil {
		return shim.Error("wrong marshal get transaction")
	}
	err = stub.PutState(transaction.OrderNo, []byte(transactionBytes))
	if err != nil {
		return shim.Error(err.Error())
	}


	erf,_ := json.Marshal(transaction)
	fmt.Println("this is for ", string(erf))
	// 以下添加各种索引

	stub.GetTxTimestamp()
	value := []byte{0x00}

	// IDNo~OrderNo
	IDNo_OrderNo, err := stub.CreateCompositeKey("IDNo~OrderNo", []string{transaction.IDNo, transaction.OrderNo})
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState(IDNo_OrderNo, value)

	// IDNo~ProductCode~OrderNo
	IDNo_ProductCode_OrderNo, err := stub.CreateCompositeKey("IDNo~ProductCode~OrderNo", []string{transaction.IDNo,
		transaction.ProductCode, transaction.OrderNo})
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState(IDNo_ProductCode_OrderNo, value)

	// IDNo~OrganizationCode~OrderNo
	IDNo_OrganizationCode_OrderNo, err := stub.CreateCompositeKey("IDNo~OrganizationCode~OrderNo", []string{transaction.IDNo, transaction.OrganizationCode, transaction.OrderNo})
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState(IDNo_OrganizationCode_OrderNo, value)


	// ProductCode~OrderNo
	ProductCode_OrderNo, err := stub.CreateCompositeKey("ProductCode~OrderNo", []string{transaction.ProductCode, transaction.OrderNo})
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState(ProductCode_OrderNo, value)
	// ProductCode~IDNo
	ProductCode_IDNo, err := stub.CreateCompositeKey("ProductCode~IDNo", []string{transaction.ProductCode, transaction.IDNo})
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState(ProductCode_IDNo, value)


	// OrganizationCode~OrderNo
	OrganizationCode_OrderNo, err := stub.CreateCompositeKey("OrganizationCode~OrderNo", []string{transaction.OrganizationCode, transaction.OrderNo})
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState(OrganizationCode_OrderNo, value)

	// OrganizationCode~IDNo
	OrganizationCode_IDNo, err := stub.CreateCompositeKey("OrganizationCode~IDNo", []string{transaction.OrganizationCode, transaction.IDNo})
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState(OrganizationCode_IDNo, value)

	// OrganizationCode~IDNo~ProductCode
	OrganizationCode_IDNo_ProductCode, err := stub.CreateCompositeKey("OrganizationCode~IDNo~ProductCode", []string{transaction.OrganizationCode,
		transaction.IDNo, transaction.ProductCode})
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState(OrganizationCode_IDNo_ProductCode, value)


	// ================
	return shim.Success(nil)
}

// getTransactionByID 获取某笔交易
// args[0] functionname string
// args[1] IDNo string
func (t *SimpleChaincode) getTransactionByOrderNo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex0302 getTransactionByOrderNo")

	var OrderNo string //交易ID
	var transaction Transaction
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	OrderNo = args[1]

	TransactionInfo, err := stub.GetState(OrderNo)
	if err != nil {
		return shim.Error(err.Error())
	}
	//将byte的结果转换成struct
	err = json.Unmarshal(TransactionInfo, &transaction)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(string(TransactionInfo))
	return shim.Success(TransactionInfo)
}

//得到某一用户的所有交易,
//args[0] functionname string
//args[1] IDNo string
//args = []string {"getTransactionByUserID", "1"}
func (t *SimpleChaincode) getTransactionByIDNo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("0x0303 getTransactionByIDNo")
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	IDNo := args[1:]

	// Query the TransactionObject index by FromID
	// This will execute a key range query on all keys starting with 'IDNo'
	transactionOrderNoResultsIterator, err := stub.GetStateByPartialCompositeKey("IDNo~OrderNo", IDNo)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer transactionOrderNoResultsIterator.Close()

	// Iterate through result set and for each marble found, transfer to newOwner
	bArrayMemberAlreadyWritten := false
	var buffer bytes.Buffer
	buffer.WriteString("[")

	for transactionOrderNoResultsIterator.HasNext() {
		// Note that we don't get the value (2nd return variable), we'll just get the marble name from the composite key
		queryResponse, err := transactionOrderNoResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		objectType, compositeKeyParts, err := stub.SplitCompositeKey(queryResponse.Key)
		if err != nil {
			return shim.Error("we cannot splitcompositekey")
		}
		if objectType != "IDNo~OrderNo" {
			return shim.Error("object is not we want %s" + IDNo[0])
		}
		transactionid := compositeKeyParts[len(compositeKeyParts)-1]

		transactionBytes, err := stub.GetState(transactionid)
		if err != nil {
			return shim.Error("the transactionid is not put in the ledger")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(transactionid)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(transactionBytes))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}

	buffer.WriteString("]")
	fmt.Println("just make IDNo", buffer.String())
	return shim.Success(buffer.Bytes())
}

//args[0] functionname string
//args[1] IDNo string
//args = []string {"getTransactionByOrderNoRange", "startkey","endkey"}
func (t *SimpleChaincode) getTransactionByOrderNoRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("0x04 getTransactionByOrderNoRange")
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	startKey := args[1]
	endKey := args[2]

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())

}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// Only available on state databases that support rich query (e.g. CouchDB)
// =========================================================================================
func (t *SimpleChaincode) queryTransactionByIDNo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("0x0304")
	if len(args) < 2{
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	IDNo := args[1]
	fmt.Println("IDNo", IDNo)
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"orderNo\",\"idNo\":\"%s\"}}", IDNo)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}
