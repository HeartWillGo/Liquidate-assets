package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"os"
)

func generate_transdata(number int) {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	AlreadyWrite := false

	for i := 0; i < number; i++ {
		if AlreadyWrite == true {
			buffer.WriteString(",")
		}
		var transaction Transaction
		var idx string
		idx = strconv.Itoa(i)
		//精确到秒
		data :=  time.Now().Format("20060102150405")
		transaction.SID  =  data + idx
		transaction.ReceiverSID = data +idx
		transaction.OriginSID  = data +idx
		transaction.RequestSerial = data + idx
		transaction.NextRequestSerial = data + idx
		transaction.Proposaltime =  time.Now().Unix()

		transaction.Transactionid = "transactionid" + data + idx
		transaction.Transactiondate = time.Now().Unix()
		transaction.Parentorder = idx
		transaction.Suborder = idx
		transaction.Payid = idx
		transaction.Transtype = idx
		transaction.Fromtype = 1
		transaction.Fromid = "0000000000002" + strconv.Itoa(i/3)
		transaction.Totype = 1
		transaction.Toid = "0000000000003" + strconv.Itoa(i/3)
		transaction.Productid = "productid" + strconv.Itoa(i%3)
		transaction.Productinfo = "trohs si efil"
		transaction.Organizationid = "nagnip"
		random, err := strconv.Atoi(RandInt(2))
		if err != nil {
			fmt.Println(err.Error())
		}

		transaction.Amount = float64(random)
		random, err = strconv.Atoi(RandInt(1))
		if err != nil {
			fmt.Println(err.Error())
		}
		transaction.Price =  float64(random)
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

func generate_userData(number int) {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	AlreadyWrite := false
	for i := 0; i < number; i++ {
		if AlreadyWrite == true {
			buffer.WriteString(",")
		}
		var user User
		var idx string

		idx = strconv.Itoa(i)
		user.ID = "0000000000002" + strconv.Itoa(i/3)
		user.Name = RandStr(8)
		user.Identificationtype = 1
		user.Identification = "23342"
		user.Sex = 1
		user.Birthday = "20340912"
		user.Bankcard = "123243"
		user.Phonenumber = "999999"
		user.Token = idx

		transactionBytes, _ := json.Marshal(user)
		buffer.WriteString(string(transactionBytes))
		AlreadyWrite = true
	}
	buffer.WriteString("]")
	err := ioutil.WriteFile("../testdata/User.json", buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("not success is failed, am I right?")
	}
}

func generate_productData(number int) {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	AlreadyWrite := false
	for i := 0; i < number; i++ {
		if AlreadyWrite == true {
			buffer.WriteString(",")
		}
		var product Product
		var idx string
		idx = strconv.Itoa(i)

		product.Productid = "productid" + idx
		product.Productname = "zhaocaibao"
		product.Organizationid = "pingan"
		product.Producttype = 1
		product.Amount = 3
		product.Price = 33.0

		productBytes, _ := json.Marshal(product)
		buffer.WriteString(string(productBytes))
		AlreadyWrite = true
	}
	buffer.WriteString("]")
	err := ioutil.WriteFile("../testdata/Product.json", buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("not success is failed, am I right?")
	}
}

func generate_organizationData(number int) {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	AlreadyWrite := false
	for i := 0; i < number; i++ {
		if AlreadyWrite == true {
			buffer.WriteString(",")
		}
		var organization Organization
		var idx string
		idx = strconv.Itoa(i)
		organization.OrganizationID = "pingan" + idx
		organization.OrganizationName = "pingan"
		organization.OrganizationType = 1

		transactionBytes, _ := json.Marshal(organization)
		buffer.WriteString(string(transactionBytes))
		AlreadyWrite = true
	}
	buffer.WriteString("]")
	err := ioutil.WriteFile("../testdata/Organization.json", buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("not success is failed, am I right?")
	}
}

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
func getUsers() []User {
	raw, err := ioutil.ReadFile("../testdata/User.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []User
	json.Unmarshal(raw, &c)
	return c
}
func getProducts() []Product {
	raw, err := ioutil.ReadFile("../testdata/Product.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Product
	json.Unmarshal(raw, &c)
	return c
}
func getOrganizations() []Organization {
	raw, err := ioutil.ReadFile("../testdata/Organization.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Organization
	json.Unmarshal(raw, &c)
	return c
}
