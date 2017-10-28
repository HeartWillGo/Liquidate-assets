package main

import (
	//"github.com/hyperledger/fabric/core/chaincode/shim"
	//pb "github.com/hyperledger/fabric/protos/peer"

	"encoding/json"
	"fmt"
	"time"
)

type UserAsset struct {
	StatisticDate   string                        `json:"statistic_date"`
	TradingEntityID string                        `json:"trading_entity_id"`
	TransactionNum  int                           `json:"transaction_num"`
	AssetType       string                        `json:"asset_type"`
	AssetInfo       string                        `json:"asset_info"`
	TradeStartTime  int64                         `json:"trade_start_time"`
	TradeEndTime    int64                         `json:"trade_end_time"`
	AssetBalance    float64                       `json:"asset_balance"`
	AssetIncome     float64                       `json:"asset_income"`
	AssetOutcome    float64                       `json:"asset_outcome"`
	OrganizatonMap  map[string]*OrganizationAsset `json:"organization_Map"`
	ProductMap      map[string]*ProductAsset      `json:"productmap"`
}
type OrganizationAsset struct {
	ID            string `json:"id"`
	StatisticDate string `json:"statistic_date"`

	Type           int     `json:"type"`
	TransactionNum int64   `json:"transactionnum"`
	TradestartTime int64   `json:"tradestarttime"`
	TradeendTime   int64   `json:"tradeendtime"`
	Balance        float64 `json:"balance"`
	Outcome        float64 `json:"outcome"`
	Income         float64 `json:"income"`

	ProductMap map[string]*ProductAsset `json:"productmap"`
	UserMap    map[string]*UserAsset    `json:"asset"`
}
type ProductAsset struct {
	ID             string                `json:"id"`
	StatisticDate  string                `json:"statistic_date"`
	Tradestarttime int64                 `json:"tradestarttime"`
	Tradeendtime   int64                 `json:"tradeendtime"`
	TransactionNum int64                 `json:"transactionum"`
	Balance        float64               `json:"balance"`
	Outcome        float64               `json:"outcome"`
	Income         float64               `json:"income"`
	UserMap        map[string]*UserAsset `json:"asset"`
}

type RecordTransaction struct {
	Key    string      `json:"Key"`
	Record Transaction `json:"Record"`
}

func computeAssetByUserID(statisticID string, transactionBytes []byte) UserAsset {
	var asset UserAsset
	var recordTransaction []RecordTransaction
	asset.TradingEntityID = statisticID
	fmt.Println(string(transactionBytes))
	err := json.Unmarshal(transactionBytes, &recordTransaction)
	if err != nil {
		fmt.Println("it is wrong")
		fmt.Println(err.Error())
	}

	asset.StatisticDate = fmt.Sprintf("%v", time.Now().Unix())
	asset.OrganizatonMap = make(map[string]*OrganizationAsset)
	AlreadyCreateProductMap := make(map[string]bool)

	for _, record := range recordTransaction {
		//用户的视角 user -----> organizaition
		//         organization ------> user
		//         organization ------
		tran := record.Record
		if tran.IDNo == asset.TradingEntityID {
			asset.AssetIncome += tran.Amount * tran.Portion

		} else if  tran.Toid == asset.TradingEntityID {
			asset.AssetOutcome += tran.Amount * tran.Portion
		}
		if (asset.AssetIncome - asset.AssetOutcome) < 0 {
			fmt.Println("negative")
		}
		asset.TradeStartTime = findMin(asset.TradeStartTime, tran.TransDate)
		asset.TradeEndTime =   findMax(asset.TradeEndTime, tran.TransDate)
		asset.TransactionNum += 1
		//机构

		_, ok := asset.OrganizatonMap[tran.OrganizationCode]
		if !ok {
			asset.OrganizatonMap[tran.OrganizationCode] = &OrganizationAsset{ID: tran.OrganizationCode}
		}
		//像这种情况orgAsset如果不存在，对其操作后，orgAsset是否立即存在
		//asset.OrganizatonMap[tran.OrganizationCode].TradestartTime = findMin(asset.OrganizatonMap[tran.OrganizationCode].TradestartTime, tran.TransDate)
		//产品
		//asset.OrganizatonMap[tran.OrganizationCode].TradeendTime = findMax(asset.OrganizatonMap[tran.OrganizationCode].TradeendTime, tran.TransDate)

		//记录机构的交易数据
		if statisticID == tran.IDNo {

			asset.OrganizatonMap[tran.OrganizationCode].Outcome += tran.Amount * tran.Portion
		} else if statisticID == tran.Toid {
			asset.OrganizatonMap[tran.OrganizationCode].Income += tran.Amount * tran.Portion
		}
		asset.OrganizatonMap[tran.OrganizationCode].TransactionNum += 1


		if AlreadyCreateProductMap[tran.OrganizationCode] == false {
			asset.OrganizatonMap[tran.OrganizationCode].ProductMap = make(map[string]*ProductAsset)
			AlreadyCreateProductMap[tran.OrganizationCode] = true

		}

		//记录产品的数据
		_, ok = asset.OrganizatonMap[tran.OrganizationCode].ProductMap[tran.ProductCode]
		if !ok {
			asset.OrganizatonMap[tran.OrganizationCode].ProductMap[tran.ProductCode] = &ProductAsset{ID: tran.ProductCode}
		}

		asset.OrganizatonMap[tran.OrganizationCode].ProductMap[tran.ProductCode].Tradeendtime =   findMax(asset.OrganizatonMap[tran.OrganizationCode].ProductMap[tran.ProductCode].Tradeendtime, tran.TransDate)
		asset.OrganizatonMap[tran.OrganizationCode].ProductMap[tran.ProductCode].Tradestarttime = findMin(asset.OrganizatonMap[tran.OrganizationCode].ProductMap[tran.ProductCode].Tradestarttime, tran.TransDate)


		if statisticID == tran.IDNo {
			asset.OrganizatonMap[tran.OrganizationCode].ProductMap[tran.ProductCode].Outcome += tran.Amount * tran.Portion
		} else if statisticID == tran.Toid {
			asset.OrganizatonMap[tran.OrganizationCode].ProductMap[tran.ProductCode].Income += tran.Amount * tran.Portion

		}
		asset.OrganizatonMap[tran.OrganizationCode].ProductMap[tran.ProductCode].TransactionNum += 1

	}
	asset.AssetBalance = asset.AssetIncome - asset.AssetOutcome
	return asset

}
func computeProductSaleInformation(transactionBytes []byte) ProductAsset {
	var productAsset ProductAsset
	var recordTransaction []RecordTransaction

	err := json.Unmarshal(transactionBytes, &recordTransaction)
	if err != nil {
		fmt.Println(err.Error())
	}
	productAsset.StatisticDate = fmt.Sprintf("%v", time.Now().Unix())

	for _, record := range recordTransaction {
		productAsset.TransactionNum += 1
		tran := record.Record
		productAsset.Balance += tran.Amount * tran.Portion
		productAsset.ID = tran.ProductCode
		productAsset.Tradestarttime = findMin(productAsset.Tradestarttime, tran.TransDate)
		productAsset.Tradeendtime = findMax(productAsset.Tradeendtime, tran.TransDate)
	}
	return productAsset

}

func computeProductAllUser(transactionBytes []byte) ProductAsset {

	var recordTransaction []RecordTransaction
	var productAsset ProductAsset
	productAsset.UserMap = make(map[string]*UserAsset)

	err := json.Unmarshal(transactionBytes, &recordTransaction)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, record := range recordTransaction {
		tran := record.Record

		_, ok := productAsset.UserMap[tran.Toid]
		if ok == false {
			productAsset.UserMap[tran.Toid] = &UserAsset{TradingEntityID: tran.Toid}
			productAsset.ID = tran.ProductCode
			productAsset.StatisticDate = fmt.Sprintf("%v", time.Now().Unix())

		}
		productAsset.TransactionNum += 1
		productAsset.Balance += tran.Amount * tran.Portion
		productAsset.UserMap[tran.Toid].AssetBalance += tran.Amount * tran.Portion
		productAsset.UserMap[tran.Toid].TransactionNum += 1
	}

	return productAsset

}

func computeOrgnazitionAllProduct(transactionBytes []byte) OrganizationAsset {

	var organizationAsset OrganizationAsset
	var recordTransaction []RecordTransaction
	err := json.Unmarshal(transactionBytes, &recordTransaction)
	if err != nil {
		fmt.Println(err.Error())
	}
	organizationAsset.StatisticDate = fmt.Sprintf("%v", time.Now().Unix())
	organizationAsset.ProductMap = make(map[string]*ProductAsset)
	AlreadyCreateUser := false
	for _, record := range recordTransaction {
		tran := record.Record

		_, ok := organizationAsset.ProductMap[tran.ProductCode]
		if !ok {
			organizationAsset.ProductMap[tran.ProductCode] = &ProductAsset{ID: tran.ProductCode}
			organizationAsset.ID = tran.OrganizationCode
		}
		organizationAsset.TransactionNum += 1
		organizationAsset.Balance += tran.Amount * tran.Portion

		organizationAsset.ProductMap[tran.ProductCode].TransactionNum += 1
		organizationAsset.ProductMap[tran.ProductCode].Balance += tran.Amount * tran.Portion

		if AlreadyCreateUser == false {
			//organizationAsset.ProductMap[tran.OrganizationCode].UserMap = make(map[string]*UserAsset)
		}
		AlreadyCreateUser = true
		//_, ok = organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.Toid]
		//if !ok {
		//	organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.Toid] = &UserAsset{TradingEntityID:tran.Toid}
		//}
		//
		//_, ok = organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.IDNo]
		//if !ok {
		//	organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.IDNo] = &UserAsset{TradingEntityID:tran.IDNo}
		//}
		//
		//organizationAsset.ProductMap[tran.ProductCode].UserMap[tran.Toid].AssetIncome += tran.Amount * tran.Portion
		//
		//organizationAsset.ProductMap[tran.ProductCode].UserMap[tran.IDNo].AssetOutcome += tran.Amount * tran.Portion
		//

	}
	organizationAssetBytes, err := json.Marshal(organizationAsset)
	if err != nil {
		fmt.Println("marshal userOperateProductMapBytes Wrong")
	}
	fmt.Println("organizationAssetBytes", string(organizationAssetBytes))

	return organizationAsset
}

func computeOrgnazitionAllUser(transactionBytes []byte) OrganizationAsset {

	var organizationAsset OrganizationAsset
	var recordTransaction []RecordTransaction
	err := json.Unmarshal(transactionBytes, &recordTransaction)
	if err != nil {
		fmt.Println(err.Error())
	}
	organizationAsset.StatisticDate = fmt.Sprintf("%v", time.Now().Unix())
	organizationAsset.UserMap = make(map[string]*UserAsset)
	AlreadyCreateUser := false
	for _, record := range recordTransaction {
		tran := record.Record

		_, ok := organizationAsset.UserMap[tran.Toid]
		if !ok {
			organizationAsset.UserMap[tran.Toid] = &UserAsset{TradingEntityID: tran.Toid}
			organizationAsset.ID = tran.OrganizationCode
		}
		_, ok = organizationAsset.UserMap[tran.IDNo]
		if !ok {
			organizationAsset.UserMap[tran.IDNo] = &UserAsset{TradingEntityID: tran.IDNo}
		}

		organizationAsset.TransactionNum += 1
		organizationAsset.Balance += tran.Amount * tran.Portion

		organizationAsset.UserMap[tran.Toid].TransactionNum += 1
		organizationAsset.UserMap[tran.Toid].AssetIncome += tran.Amount * tran.Portion

		organizationAsset.UserMap[tran.IDNo].TransactionNum += 1
		organizationAsset.UserMap[tran.IDNo].AssetOutcome += tran.Amount * tran.Portion

		if AlreadyCreateUser == false {
			//organizationAsset.ProductMap[tran.OrganizationCode].UserMap = make(map[string]*UserAsset)
		}
		AlreadyCreateUser = true
		//_, ok = organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.Toid]
		//if !ok {
		//	organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.Toid] = &UserAsset{TradingEntityID:tran.Toid}
		//}
		//
		//_, ok = organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.IDNo]
		//if !ok {
		//	organizationAsset.ProductMap[tran.OrganizationCode].UserMap[tran.IDNo] = &UserAsset{TradingEntityID:tran.IDNo}
		//}
		//
		//organizationAsset.ProductMap[tran.ProductCode].UserMap[tran.Toid].AssetIncome += tran.Amount * tran.Portion
		//
		//organizationAsset.ProductMap[tran.ProductCode].UserMap[tran.IDNo].AssetOutcome += tran.Amount * tran.Portion
		//

	}
	organizationAssetBytes, err := json.Marshal(organizationAsset)
	if err != nil {
		fmt.Println("marshal userOperateProductMapBytes Wrong")
	}
	fmt.Println("organizationAssetBytes", string(organizationAssetBytes))

	return organizationAsset
}

func findMax(num1 int64, num2 int64) int64 {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
func findMin(num1 int64, num2 int64) int64 {
	if num1 < num2 && num1 != 0 {
		return num1
	} else {
		return num2
	}
}
