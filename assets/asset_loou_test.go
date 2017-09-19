package main

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

import (
"fmt"
"testing"
"github.com/hyperledger/fabric/core/chaincode/shim"
)

var chaincodeName = "asset"

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

func checkState(t *testing.T, stub *shim.MockStub, transID string, expect string) {
	bytes := stub.State[transID]
	if bytes == nil {
		fmt.Println("State", transID, "failed to get value")
		t.FailNow()
	}
	fmt.Println("state value", string(bytes))
	if string(bytes) != expect {
		fmt.Println("State value", transID, "was not", expect, "as expected")
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

func TestTransaction(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex05", scc)

	checkInit(t, stub, [][]byte{[]byte("init"), []byte(""), []byte("432")})
	str := `{"parentorder":"0239", "suborder":"23", "payid":"string","transactionid":"123", "transtype":"2", "fromtype":23,"totype":2,"toid":"231","productid":"2343","productinfo":"badfa","organizationid":"2341","account":2314,"price":23}`
	//str := `{transactionid":"123", "transtype":"2", "fromtype":23,"fromid":"234","totype":2,"toid":"231","productid":"2343","productinfo":"badfa","organizationid":"2341","account":2314,"price":23}`

	fmt.Println("str", str)
	checkInvoke(t, stub, [][]byte{[]byte("invoke"),[]byte("transaction"), []byte(str)})

	checkState(t, stub, "123", "nil")
}

