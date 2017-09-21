package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

)


//机构
type Organization struct {
	OrganizationID   string `json:"organizationid"`   //机构id
	OrganizationName string `json:"organizationname"` //机构名称
	OrganizationType int    `json:"organizationtype"` //机构类型
}

//createOrganization
func (t *SimpleChaincode) createOrganization(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 createOrganization")



	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	var organization Organization

	err := json.Unmarshal([]byte(args[1]), &organization)
	if err != nil {
		return shim.Error(err.Error())
	}



	OrganizationBytes, err := stub.GetState(organization.OrganizationID)
	if err != nil {
		return shim.Error("failed to get orginfo")
	} else if OrganizationBytes != nil {
		return shim.Error(string(OrganizationBytes) + "\t is already exists")
	}

	// Write the state to the ledger
	err = stub.PutState(organization.OrganizationID, []byte(args[1]))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
//getOrganization 获取机构信息
func (t *SimpleChaincode) getOrganization(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 getOrganization")

	var Organization_ID string // 商业银行ID
	var organization Organization

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

//TODO:修改机构
func (t *SimpleChaincode) WriteOrganization(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 WriteOrganization")

	fmt.Printf("CreateOrganization \n")

	return shim.Success(nil)
}


func (t *SimpleChaincode) getOrganizationProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("0x11 getOrganizationProduct")

	if len(args) != 2 {
		return shim.Error("Expecting 2, but get wrong")
	}
	resp := t.getTransactionByOrganizationid(stub, args)
	if resp.Status != shim.OK {
		return shim.Error("getUserAssetFailed")
	}
	productAsset := computeOrgnazitionAllProduct(resp.GetPayload())
	productAssetBytes, err := json.Marshal(productAsset)
	if err != nil {
		fmt.Println("marshal wrong")
	}
	fmt.Println(string(productAssetBytes))

	return shim.Success(productAssetBytes)

}

func (t *SimpleChaincode) getOrganizationAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}
func (t *SimpleChaincode) getOrganizationUser(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	return shim.Success(nil)
}