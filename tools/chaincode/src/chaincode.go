package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PushValuesContract struct {
	contractapi.Contract
}

type Asset struct {
	AssetID string 	`json:"assetId`
	Voltage float64 `json:voltage"`
	Current float64 `json:"current"`
	Time    string 	`json:"time"`
	Date	string	`json:"date"`
}

func (p *PushValuesContract) CreateAsset(ctx contractapi.TransactionContextInterface,
assetId string, voltage float64, current float64, time string, date string) (string, error){
	exists, err := p.AssetExists(ctx, assetId)
	//multiple return values revoked by calling AssetExists
	if err != nil {
		return "", fmt.Errorf("failed to check asset existence: %v", err)
	}
	if exists {
		return "", fmt.Errorf("asset %s already exists", assetId)
	}


	asset := Asset{
		AssetID: assetId,
		Voltage: voltage,
		Current: current,
		Time: time,
		Date: date, 
	}

	assetJSON, err := json.Marshal(asset)
	if err!= nil {
		return "", fmt.Errorf("failed to marshal asset: %v", err)
	}
	
	err = ctx.GetStub().PutState(assetId, assetJSON)
	if err != nil {
		return "", fmt.Errorf("failed to put asset on ledger: %v", err)
	}
	return fmt.Sprintf("Asset %s created", assetId), nil
}

func (p *PushValuesContract) AssetExists(ctx contractapi.TransactionContextInterface, assetId string) (bool, error){
	assetJSON, err:= ctx.GetStub().GetState(assetId)
	if err != nil{
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return assetJSON != nil, nil
}
func (p *PushValuesContract) ReadAsset(ctx contractapi.TransactionContextInterface, assetId string) (*Asset, error){
	assetJSON, err:= ctx.GetStub().GetState(assetId)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state:%v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("asset %s does not exist", assetId)
	}
	var asset Asset 
	err = json.Unmarshal(assetJSON, &asset)
	if err!= nil {
		return nil, fmt.Errorf("failed to unmarshal asset: %v", err)
	}
	return &asset, nil
}

func main(){
	chaincode, err:=contractapi.NewChaincode(&PushValuesContract{})
	if err != nil {
		fmt.Printf("Error creating PushValuesContract chaincode: %v", err)
		return
	}
	if err:= chaincode.Start(); err!= nil {
		fmt.Printf("Error starting PushValuesContract chaincode: %v", err)
	}
}