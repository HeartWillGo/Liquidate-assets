/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"testing"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"math/rand"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"strconv"
)
////交易内容
//type Transaction struct {
//	//交易的ID和
//	Transactionid   string `json:"transactionid"`
//	Transactiondate string `json:"transactiondate"`
//
//	Parentorder     string `json:"parentorder"`
//	Suborder        string `json:"suborder"`
//	Payid           string `json:"payid"`
//	//交易头部
//	Transtype       string `json:"transtype"`
//	Fromtype        int    `json:"fromtype"`
//	Fromid          string `json:"fromid"`
//	Totype          int    `json:"totype"`
//	Toid            string `json:"toid"`
//	//交易内容
//	Productid       string `json:"productid"`
//	Productinfo     string `json:"productinfo"`
//	Organizationid  string `json:"organizationid"`
//
//	Amount          int    `json:"amount"`
//	Price           int    `json:"price"`
//}

var chaincodeName = "aaset"

// chaincode_example05 looks like it wanted to return a JSON response to Query()
// it doesn't actually do this though, it just returns the sum value
func jsonResponse(name string, value string) string {
	return fmt.Sprintf("jsonResponse = \"{\"Name\":\"%v\",\"Value\":\"%v\"}", name, value)
}

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkState(t *testing.T, stub *shim.MockStub, name string, expect string) {
	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != expect {
		fmt.Println("State value", name, "was not", expect, "as expected")
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, args [][]byte, expect string) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Query", args, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", args, "failed to get result")
		t.FailNow()
	}
	if string(res.Payload) != expect {
		fmt.Println("Query result ", string(res.Payload), "was not", expect, "as expected")
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func TestInit(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex05", scc)

	str0 := `this we arrive the init function`
	checkInit(t, stub, [][]byte{[]byte("init"), []byte(str0)})

	checkState(t, stub, "qwertyuiop", str0)
}

func TestTransaction (t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex05", scc)

	str0 := `this we arrive the init function`
	checkInit(t, stub, [][]byte{[]byte("init"), []byte(str0)})
	create_json(8)



		//tran1 := `{"parentorder":"0239", "suborder":"23", "payid":"string","transactionid":`+
		//	string(i)+`, "transtype":"2", "fromtype":23,"totype":2,"toid":"231","productid":"2343"
		//	,"productinfo":"badfa","organizationid":"2341","account":2314,"price":23}`
		//checkInvoke(t, stub, [][]byte{[]byte("invoke"), []byte("Transaction"), []byte(tran1)})
}

//
func RandStr(strlen int) string {
	rand.Seed(time.Now().Unix())
	data := make([]byte, strlen)
	var num int
	for i := 0; i < strlen; i++ {
		num = rand.Intn(57) + 65
		for {
			if num>90 && num<97 {
				num = rand.Intn(57) + 65
			} else {
				break
			}
		}
		data[i] = byte(num)
	}
	return string(data)
}

func RandInt(strlen int) string {
	rand.Seed(time.Now().Unix())

	var data string
	var num int
	for i := 0; i < strlen; i++ {
		num = rand.Intn(57) + 65

		data += string(num)
	}
	return data
}


func create_json(number int ) {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	AlreadyWrite := false

	for i:=0; i < number; i++ {
		if AlreadyWrite == true {
			buffer.WriteString(",")
		}
		var transaction Transaction

		var idx string
		idx = strconv.Itoa(i)
		transaction.Transactionid =idx
		transaction.Transactiondate = time.Now().String()
		transaction.Parentorder =  idx
		transaction.Suborder =  idx
		transaction.Payid =  idx
		transaction.Transtype =  idx
		transaction.Fromtype = 1
		transaction.Fromid =  RandStr(8)
		transaction.Totype = 1
		transaction.Toid = RandStr(8)
		transaction.Productid = string(i % 3)
		transaction.Productinfo = "wegood"+"i%3"
		transaction.Organizationid = "pingan"
		transaction.Amount = i
		transaction.Price = 9

		transactionBytes, _ := json.Marshal(transaction)
		buffer.WriteString(string(transactionBytes))
		AlreadyWrite = true
	}
	buffer.WriteString("]")
	err := ioutil.WriteFile("output.json", buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("not success is failed, am I right?")
	}


}

