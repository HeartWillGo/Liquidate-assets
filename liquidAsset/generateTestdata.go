package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"strconv"
	"time"

	"os"
	"math/rand"
)

func generate_transdata(number int) {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	AlreadyWrite := false
	Tran := make([][]string, 0)
	Tran = append(Tran, []string{"P011", "O01", "debit","100"},
						[]string{"P011", "O01", "credit","100"},
						[]string{"P012", "O01", "debit","100"},
						[]string{"P012", "O01", "credit","100"},
						[]string{"P021", "O02", "debit","100"},
						[]string{"P021", "O02", "credit","100"},
						[]string{"P022", "O02", "debit","100"},
						[]string{"P022", "O02", "credit","100"},
						[]string{"P031", "O03", "debit","100"},
						[]string{"P031", "O03", "credit","100"},
						[]string{"P032", "O03", "debit","100"},
						[]string{"P032", "O03", "credit","100"})

	User := make([][]string, 0)
	User = append(User, []string{"01","1000"},
						[]string{"01","1001"},
						[]string{"02","2000"},
						[]string{"02","2001"},
						[]string{"03","3000"},
						[]string{"03","3001"})


	for i := 0; i < number; i++ {
		time.Sleep(1*time.Second)
		TranInfo := Tran[rand.Intn(len(Tran))]
		UserInfo := User[rand.Intn(len(User))]

		if AlreadyWrite == true {
			buffer.WriteString(",")
		}
		var transaction Transaction


		data :=  time.Now().Format("20060102150405")
		transaction.Birthday = ""
		transaction.LoanType = TranInfo[2]
		transaction.AccountName = "测试"
		transaction.Portion = ""
		transaction.IDNo = UserInfo[0]
		transaction.PhoneNo =  "18660821753"
		transaction.SID = "001"
		transaction.TransDate  = data
		transaction.OrigSID = "003"
		transaction.TransStatus  = "0"
		transaction.Amount = TranInfo[3]
		transaction.OrderNo = data+"5473214732942839073347079911"
		transaction.IDType = "1"
		transaction.NextRequestSerial  = data + "5459214592824143188183208384"
		transaction.AccountType = "pamaAcct"
		transaction.Sex = "M"
		transaction.Message = ""
		transaction.RequestSerial = data +"5457614592913507738619362713"
		transaction.ReceiverSID = "002"
		transaction.IDKey = "NDQwNTI4MjAwMDAxMDExMTEx"
		transaction.ProductCode = TranInfo[0]
		transaction.TransType = "橙子代收"
		transaction.OrganizationCode = TranInfo[1]
		transaction.Account = UserInfo[1]
		transaction.Desc = ""
		transactionBytes, _ := json.Marshal(transaction)
		buffer.WriteString(string(transactionBytes))
		AlreadyWrite = true
	}
	buffer.WriteString("]")
	err := ioutil.WriteFile("../testdata/transaction.json", buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("not success is failed, am I right?")
	}
}

//func generate_userData(number int) {
//	var buffer bytes.Buffer
//	buffer.WriteString("[")
//	AlreadyWrite := false
//	for i := 0; i < number; i++ {
//		if AlreadyWrite == true {
//			buffer.WriteString(",")
//		}
//		var user User
//		var idx string
//
//		idx = strconv.Itoa(i)
//		user.ID = "0000000000002" + strconv.Itoa(i/3)
//		user.Name = RandStr(8)
//		user.Identificationtype = 1
//		user.Identification = "23342"
//		user.Sex = 1
//		user.Birthday = "20340912"
//		user.Bankcard = "123243"
//		user.Phonenumber = "999999"
//		user.Token = idx
//
//		transactionBytes, _ := json.Marshal(user)
//		buffer.WriteString(string(transactionBytes))
//		AlreadyWrite = true
//	}
//	buffer.WriteString("]")
//	err := ioutil.WriteFile("../testdata/User.json", buffer.Bytes(), 0644)
//	if err != nil {
//		fmt.Println("not success is failed, am I right?")
//	}
//}
//
//func generate_productData(number int) {
//	var buffer bytes.Buffer
//	buffer.WriteString("[")
//	AlreadyWrite := false
//	for i := 0; i < number; i++ {
//		if AlreadyWrite == true {
//			buffer.WriteString(",")
//		}
//		var product Product
//		var idx string
//		idx = strconv.Itoa(i)
//
//		product.ProductCode = "productid" + idx
//		product.Productname = "zhaocaibao"
//		product.OrganizationCode = "pingan"
//		product.Producttype = 1
//		product.Amount = 3
//		product.Portion = 33.0
//
//		productBytes, _ := json.Marshal(product)
//		buffer.WriteString(string(productBytes))
//		AlreadyWrite = true
//	}
//	buffer.WriteString("]")
//	err := ioutil.WriteFile("../testdata/Product.json", buffer.Bytes(), 0644)
//	if err != nil {
//		fmt.Println("not success is failed, am I right?")
//	}
//}
//
//func generate_organizationData(number int) {
//	var buffer bytes.Buffer
//	buffer.WriteString("[")
//	AlreadyWrite := false
//	for i := 0; i < number; i++ {
//		if AlreadyWrite == true {
//			buffer.WriteString(",")
//		}
//		var organization Organization
//		var idx string
//		idx = strconv.Itoa(i)
//		organization.OrganizationID = "pingan" + idx
//		organization.OrganizationName = "pingan"
//		organization.OrganizationType = 1
//
//		transactionBytes, _ := json.Marshal(organization)
//		buffer.WriteString(string(transactionBytes))
//		AlreadyWrite = true
//	}
//	buffer.WriteString("]")
//	err := ioutil.WriteFile("../testdata/Organization.json", buffer.Bytes(), 0644)
//	if err != nil {
//		fmt.Println("not success is failed, am I right?")
//	}
//}

func getTrans() []Transaction {
	raw, err := ioutil.ReadFile("../testdata/transaction.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Transaction
	json.Unmarshal(raw, &c)
	return c
}
//func getUsers() []User {
//	raw, err := ioutil.ReadFile("../testdata/User.json")
//	if err != nil {
//		fmt.Println(err.Error())
//		os.Exit(1)
//	}
//
//	var c []User
//	json.Unmarshal(raw, &c)
//	return c
//}
//func getProducts() []Product {
//	raw, err := ioutil.ReadFile("../testdata/Product.json")
//	if err != nil {
//		fmt.Println(err.Error())
//		os.Exit(1)
//	}
//
//	var c []Product
//	json.Unmarshal(raw, &c)
//	return c
//}
//func getOrganizations() []Organization {
//	raw, err := ioutil.ReadFile("../testdata/Organization.json")
//	if err != nil {
//		fmt.Println(err.Error())
//		os.Exit(1)
//	}
//
//	var c []Organization
//	json.Unmarshal(raw, &c)
//	return c
//}
