package main

import (
	//"github.com/hyperledger/fabric/core/chaincode/shim"
	//pb "github.com/hyperledger/fabric/protos/peer"

	"encoding/json"
	"fmt"
	"time"
	"strconv"
)


type IDNoAsset struct {
	QueryTime string `json:"queryTime"`
	IDNo      string `json:"idNo"`
	TranNum   int `json:"tranNum"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Credit    float64 `json:"credit"`
	Debit     float64 `json:"debit"`
	Balance   float64 `json:"balance"`
	OrganizatonMap  map[string]*OrganizationAsset `json:"organizationMap"`
	ProductMap      map[string]*ProductAsset      `json:"productMap"`
}

//type IDNoAsset struct {
//	StatisticDate   string                        `json:"statistic_date"`
//	TradingEntityID string                        `json:"trading_entity_id"`
//	TransactionNum  int                           `json:"transaction_num"`
//	AssetType       string                        `json:"asset_type"`
//	AssetInfo       string                        `json:"asset_info"`
//	TradeStartTime  int64                         `json:"trade_start_time"`
//	TradeEndTime    int64                         `json:"trade_end_time"`
//	AssetBalance    float64                       `json:"asset_balance"`
//	AssetIncome     float64                       `json:"asset_income"`
//	AssetOutcome    float64                       `json:"asset_outcome"`
//	OrganizatonMap  map[string]*OrganizationAsset `json:"organization_Map"`
//	ProductMap      map[string]*ProductAsset      `json:"productmap"`
//}
type OrganizationAsset struct {
	OrganizationCode string `json:"organizationCode"`
	QueryTime        string `json:"queryTime"`
	TranNum          int `json:"tranNum"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	Credit           float64 `json:"credit"`
	Debit            float64 `json:"debit"`
	Balance          float64 `json:"balance"`
	ProductAssetMap map[string]*ProductAsset `json:"productMap"`
	IDNoAssetMap    map[string]*IDNoAsset    `json:"idNoAsset"`
}
//type OrganizationAsset struct {
//	Orga            string `json:"id"`
//	StatisticDate string `json:"statistic_date"`
//
//	Type           int     `json:"type"`
//	TransactionNum int64   `json:"transactionnum"`
//	TradestartTime int64   `json:"tradestarttime"`
//	TradeendTime   int64   `json:"tradeendtime"`
//	Balance        float64 `json:"balance"`
//	Outcome        float64 `json:"outcome"`
//	Income         float64 `json:"income"`
//
//	ProductMap map[string]*ProductAsset `json:"productmap"`
//	UserMap    map[string]*IDNoAsset    `json:"asset"`
//}
type ProductAsset struct {
	ProductCode string `json:"productCode"`
	QueryTime   string `json:"queryTime"`
	TranNum     int `json:"tranNum"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Credit      float64 `json:"credit"`
	Debit       float64 `json:"debit"`
	Balance     float64 `json:"balance"`
	IDNoAssetMap        map[string]*IDNoAsset `json:"idNoAsset"`
}

//type ProductAsset struct {
//	ID             string                `json:"id"`
//	StatisticDate  string                `json:"statistic_date"`
//	Tradestarttime int64                 `json:"tradestarttime"`
//	Tradeendtime   int64                 `json:"tradeendtime"`
//	TransactionNum int64                 `json:"transactionum"`
//	Balance        float64               `json:"balance"`
//	Outcome        float64               `json:"outcome"`
//	Income         float64               `json:"income"`
//	UserMap        map[string]*IDNoAsset `json:"asset"`
//}

type RecordTransaction struct {
	Key    string      `json:"Key"`
	Record Transaction `json:"Record"`
}

func ComputeIDNoAsset(idNo string, transactionBytes []byte) IDNoAsset {
	var idNoAsset IDNoAsset
	var recordTransaction []RecordTransaction
	idNoAsset.IDNo = idNo
	fmt.Println(string(transactionBytes))
	err := json.Unmarshal(transactionBytes, &recordTransaction)
	if err != nil {
		fmt.Println("it is wrong")
		fmt.Println(err.Error())
	}

	idNoAsset.QueryTime = fmt.Sprintf("%v", time.Now().Format("20060102150405"))
	idNoAsset.OrganizatonMap = make(map[string]*OrganizationAsset)
	AlreadyCreateProductMap := make(map[string]bool)

	for _, record := range recordTransaction {
		//用户的视角 user -----> organizaition
		//         organization ------> user
		//         organization ------
		tran := record.Record
		amount, _ := strconv.ParseFloat(tran.Amount, 64)


		idNoAsset.StartTime = findMin(idNoAsset.StartTime, tran.TransDate)
		idNoAsset.EndTime = findMax(idNoAsset.EndTime, tran.TransDate)

		idNoAsset.TranNum += 1
		//机构

		_, ok := idNoAsset.OrganizatonMap[tran.OrganizationCode]
		if !ok {
			idNoAsset.OrganizatonMap[tran.OrganizationCode] = &OrganizationAsset{OrganizationCode: tran.OrganizationCode}
		}

		idNoAsset.OrganizatonMap[tran.OrganizationCode].TranNum += 1

		if AlreadyCreateProductMap[tran.OrganizationCode] == false {
			idNoAsset.OrganizatonMap[tran.OrganizationCode].ProductAssetMap= make(map[string]*ProductAsset)
			AlreadyCreateProductMap[tran.OrganizationCode] = true

		}

		//像这种情况orgAsset如果不存在，对其操作后，orgAsset是否立即存在
		idNoAsset.OrganizatonMap[tran.OrganizationCode].StartTime = findMin(idNoAsset.OrganizatonMap[tran.OrganizationCode].StartTime, tran.TransDate)
		//产品
		idNoAsset.OrganizatonMap[tran.OrganizationCode].EndTime = findMax(idNoAsset.OrganizatonMap[tran.OrganizationCode].EndTime, tran.TransDate)


		//记录产品的数据
		_, ok = idNoAsset.OrganizatonMap[tran.OrganizationCode].ProductAssetMap[tran.ProductCode]
		if !ok {
			idNoAsset.OrganizatonMap[tran.OrganizationCode].ProductAssetMap[tran.ProductCode] = &ProductAsset{ProductCode: tran.ProductCode}
		}

		idNoAsset.OrganizatonMap[tran.OrganizationCode].ProductAssetMap[tran.ProductCode].StartTime = findMax(idNoAsset.OrganizatonMap[tran.OrganizationCode].ProductAssetMap[tran.ProductCode].StartTime, tran.TransDate)
		idNoAsset.OrganizatonMap[tran.OrganizationCode].ProductAssetMap[tran.ProductCode].EndTime = findMin(idNoAsset.OrganizatonMap[tran.OrganizationCode].ProductAssetMap[tran.ProductCode].EndTime, tran.TransDate)



		idNoAsset.OrganizatonMap[tran.OrganizationCode].ProductAssetMap[tran.ProductCode].TranNum += 1

		//记录机构的交易数据
		if tran.LoanType == "debit" {
			idNoAsset.Debit += amount
			idNoAsset.OrganizatonMap[tran.OrganizationCode].Debit += amount
			idNoAsset.OrganizatonMap[tran.OrganizationCode].ProductAssetMap[tran.ProductCode].Debit += amount

		} else if tran.LoanType == "credit" {
			idNoAsset.Credit += amount
			idNoAsset.OrganizatonMap[tran.OrganizationCode].Credit += amount
			idNoAsset.OrganizatonMap[tran.OrganizationCode].ProductAssetMap[tran.ProductCode].Credit += amount

		}


	}
	idNoAsset.Balance = idNoAsset.Credit - idNoAsset.Debit
	return idNoAsset

}
//func computeProductSaleInformation(transactionBytes []byte) ProductAsset {
//	var productAsset ProductAsset
//	var recordTransaction []RecordTransaction
//
//	err := json.Unmarshal(transactionBytes, &recordTransaction)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	productAsset.StatisticDate = fmt.Sprintf("%v", time.Now().Unix())
//
//	for _, record := range recordTransaction {
//		productAsset.TransactionNum += 1
//		tran := record.Record
//		productAsset.Balance += tran.Amount * tran.Portion
//		productAsset.ID = tran.ProductCode
//		productAsset.startTime = findMin(productAsset.startTime, tran.TransDate)
//		productAsset.Tradeendtime = findMax(productAsset.Tradeendtime, tran.TransDate)
//	}
//	return productAsset
//
////}
//
//func computeProductAllUser(transactionBytes []byte) ProductAsset {
//
//	var recordTransaction []RecordTransaction
//	var productAsset ProductAsset
//	productAsset.UserMap = make(map[string]*UserAsset)
//
//	err := json.Unmarshal(transactionBytes, &recordTransaction)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	for _, record := range recordTransaction {
//		tran := record.Record
//
//		_, ok := productAsset.UserMap[tran.Toid]
//		if ok == false {
//			productAsset.UserMap[tran.Toid] = &UserAsset{TradingEntityID: tran.Toid}
//			productAsset.ID = tran.ProductCode
//			productAsset.StatisticDate = fmt.Sprintf("%v", time.Now().Unix())
//
//		}
//		productAsset.TransactionNum += 1
//		productAsset.Balance += tran.Amount * tran.Portion
//		productAsset.UserMap[tran.Toid].AssetBalance += tran.Amount * tran.Portion
//		productAsset.UserMap[tran.Toid].TransactionNum += 1
//	}
//
//	return productAsset
//
//}
//
//func computeOrgnazitionAllProduct(transactionBytes []byte) OrganizationAsset {
//
//	var organizationAsset OrganizationAsset
//	var recordTransaction []RecordTransaction
//	err := json.Unmarshal(transactionBytes, &recordTransaction)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	organizationAsset.StatisticDate = fmt.Sprintf("%v", time.Now().Unix())
//	organizationAsset.ProductMap = make(map[string]*ProductAsset)
//	AlreadyCreateUser := false
//	for _, record := range recordTransaction {
//		tran := record.Record
//
//		_, ok := organizationAsset.ProductMap[tran.ProductCode]
//		if !ok {
//			organizationAsset.ProductMap[tran.ProductCode] = &ProductAsset{ID: tran.ProductCode}
//			organizationAsset.ID = tran.OrganizationCode
//		}
//		organizationAsset.TransactionNum += 1
//		organizationAsset.Balance += tran.Amount * tran.Portion
//
//		organizationAsset.ProductMap[tran.ProductCode].TransactionNum += 1
//		organizationAsset.ProductMap[tran.ProductCode].Balance += tran.Amount * tran.Portion
//
//		if AlreadyCreateUser == false {
//			//organizationAsset.ProductMap[tran.OrganizationCode].UserMap = make(map[string]*UserAsset)
//		}
//		AlreadyCreateUser = true
//		//_, ok = organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.Toid]
//		//if !ok {
//		//	organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.Toid] = &UserAsset{TradingEntityID:tran.Toid}
//		//}
//		//
//		//_, ok = organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.IDNo]
//		//if !ok {
//		//	organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.IDNo] = &UserAsset{TradingEntityID:tran.IDNo}
//		//}
//		//
//		//organizationAsset.ProductMap[tran.ProductCode].UserMap[tran.Toid].AssetIncome += tran.Amount * tran.Portion
//		//
//		//organizationAsset.ProductMap[tran.ProductCode].UserMap[tran.IDNo].AssetOutcome += tran.Amount * tran.Portion
//		//
//
//	}
//	organizationAssetBytes, err := json.Marshal(organizationAsset)
//	if err != nil {
//		fmt.Println("marshal userOperateProductMapBytes Wrong")
//	}
//	fmt.Println("organizationAssetBytes", string(organizationAssetBytes))
//
//	return organizationAsset
//}
//
//func computeOrgnazitionAllUser(transactionBytes []byte) OrganizationAsset {
//
//	var organizationAsset OrganizationAsset
//	var recordTransaction []RecordTransaction
//	err := json.Unmarshal(transactionBytes, &recordTransaction)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	organizationAsset.StatisticDate = fmt.Sprintf("%v", time.Now().Unix())
//	organizationAsset.UserMap = make(map[string]*UserAsset)
//	AlreadyCreateUser := false
//	for _, record := range recordTransaction {
//		tran := record.Record
//
//		_, ok := organizationAsset.UserMap[tran.Toid]
//		if !ok {
//			organizationAsset.UserMap[tran.Toid] = &UserAsset{TradingEntityID: tran.Toid}
//			organizationAsset.ID = tran.OrganizationCode
//		}
//		_, ok = organizationAsset.UserMap[tran.IDNo]
//		if !ok {
//			organizationAsset.UserMap[tran.IDNo] = &UserAsset{TradingEntityID: tran.IDNo}
//		}
//
//		organizationAsset.TransactionNum += 1
//		organizationAsset.Balance += tran.Amount * tran.Portion
//
//		organizationAsset.UserMap[tran.Toid].TransactionNum += 1
//		organizationAsset.UserMap[tran.Toid].AssetIncome += tran.Amount * tran.Portion
//
//		organizationAsset.UserMap[tran.IDNo].TransactionNum += 1
//		organizationAsset.UserMap[tran.IDNo].AssetOutcome += tran.Amount * tran.Portion
//
//		if AlreadyCreateUser == false {
//			//organizationAsset.ProductMap[tran.OrganizationCode].UserMap = make(map[string]*UserAsset)
//		}
//		AlreadyCreateUser = true
//		//_, ok = organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.Toid]
//		//if !ok {
//		//	organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.Toid] = &UserAsset{TradingEntityID:tran.Toid}
//		//}
//		//
//		//_, ok = organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.IDNo]
//		//if !ok {
//		//	organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.IDNo] = &UserAsset{TradingEntityID:tran.IDNo}
//		//}
//		//
//		//organizationAsset.ProductMap[tran.ProductCode].UserMap[tran.Toid].AssetIncome += tran.Amount * tran.Portion
//		//
//		//organizationAsset.ProductMap[tran.ProductCode].UserMap[tran.IDNo].AssetOutcome += tran.Amount * tran.Portion
//		//
//
//	}
//	organizationAssetBytes, err := json.Marshal(organizationAsset)
//	if err != nil {
//		fmt.Println("marshal userOperateProductMapBytes Wrong")
//	}
//	fmt.Println("organizationAssetBytes", string(organizationAssetBytes))
//
//	return organizationAsset
//}
//
func findMax(t1 string, t2 string) string {
	if t1 == "" {
		return t2
	}
	layout := "20060102150405"

	T1, _:= time.Parse(layout, t1)

	T2, _:= time.Parse(layout, t2)


	if T1.Before(T2) {
		return t2
	} else {
		return t1
	}
}
func findMin(t1 string, t2 string) string {
	if t1 == "" {
		return t2
	}
	layout := "20060102150405"
	T1, _:= time.Parse(layout, t1)
	T2, _:= time.Parse(layout, t2)


	if T1.After(T2) {
		return t2
	} else {
		return t1
	}
}
