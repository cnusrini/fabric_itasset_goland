package main

import (
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer" 
	"encoding/json"
)

type( 
	AssetMgr struct{
}
OrgAsset struct{
	Id string `json:"id”` //the assetId
	AssetType string `json:"assetType"` //device type
	Status   string `json:"status"`     //status of asset
    Location string `json:”location”` //device location
    DeviceId string `json:”deviceId”` //DeviceId
    Comment  string `json:”comment”`  //comment
    From     string `json:”from”`     //from
    To       string `json:”to”`       //to
}
)

func (c *AssetMgr) Init(stub shim.ChainCodeStubInterface) pb.Response{
	return shim.Success(nil)
}

func (c *AssetMgr) Invoke(stub shim.ChainCodeStubInterface) pb.Response{
	return shim.Error("Invalid args")
}

func (c *AssetMgr) Order(stub shim.ChainCodeStubInterface, args[] string) pb.Response{
	return shim.Success(nil)
	
}

func (c *AssetMgr) Ship(stub shim.ChainCodeStubInterface, args[] string) pb.Response{
	return shim.Success(nil)
}

func (c *AssetMgr) Distribute(stub shim.ChainCodeStubInterface,args[] string) pb.Response{
	return shim.Success(nil)
}

		

func manin(){
	err := shim.Start(new(AssetMgr))

	if err != nil {
	
	fmt.Printf("Error creating new AssetMgr Contract: %s", err)
	
	}
}

