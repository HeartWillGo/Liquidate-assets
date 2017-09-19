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

	"github.com/hyperledger/fabric/core/chaincode/shim"
	asset "github.com/Liquidate-assets/liquidAsset"
)

var chaincodeName = "ex02"

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

func TestExample05_Init(t *testing.T) {
	scc := new(asset.SimpleChaincode)
	stub := shim.NewMockStub("ex05", scc)

	// Init A=123 B=234
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("hello"), []byte("432")})

	checkState(t, stub, "sumStoreName", "432")
}



