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
	args := stub.GetStringArgs()

	if len(args)!= 3 {
		assetid := args[0]
		AssetType := args[1]
		DeviceId := args[2]

	assetData := OrgAsset{
		Id : "assetid",
		AssetType : "assetType",
		Status : "START",
		Location : "location",
		DeviceId : "deviceId",
		Comment : "comment",
		From : "N/A",
		To : "N/A",
	}
}

	assetBytes, _ := json.Marshal(assetData)
	assetErr := stub.PutState(assetid, assetBytes)
	if assetErr != nil {

		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
		
		}

	return shim.Success(nil)
}

func (c *AssetMgr) Invoke(stub shim.ChainCodeStubInterface) pb.Response{
	function, args := stub.GetFunctionAndParameters()
	if function == "Order"{
		c.Order(stub, args)
	} else if function == "Ship"{
		c.Ship(stub,args)
	} else if function == Distribute{
		c.Distribute(stub,args)
	} else if function == "query"{
		c.query(stub,args)
	} else if function == "getHistory"{
		c.getHistory(stub,args)
	}
	return shim.Error("Invalid args")
}

func (c *AssetMgr) Order(stub shim.ChainCodeStubInterface, args[] string) pb.Response{
	return c.UpdateAsset(stub, args, "ORDER", "SCHOOL", "OEM")
	
}

func (c *AssetMgr) UpdateAsset(stub shim.ChainCodeStubInterface) pb.Response{
	assetId := args[0]   
	comment := args[1]
	location := args[2] 

	assetBytes, err := stub.GetState(assetId)
	orgAsset := OrgAsset{}

	if currentStatus == "Order" && OrgAsset.Status == "START"{
		return shim.Error(err.Error())
	} else if currentStatus == "SHIP" && OrgAsset.Status == "Order"{.}
	else if currentStatus == "DISTRIBUTE" && OrgAsset.Status == "SHIP"{.}

	orgAsset.Comment = comment

orgAsset.Status = currentStatus

….
orgAsset0, _ := json.Marshal(orgAsset)

err = stub.PutState(assetId, orgAsset0)

…

return shim.Success(orgAsset0)orgAsset.Comment = comment

orgAsset.Status = currentStatus

….
orgAsset0, _ := json.Marshal(orgAsset)

err = stub.PutState(assetId, orgAsset0)

…

return shim.Success(orgAsset0)orgAsset.Comment = comment

orgAsset.Status = currentStatus

….
orgAsset0, _ := json.Marshal(orgAsset)

err = stub.PutState(assetId, orgAsset0)

…

return shim.Success(orgAsset0)orgAsset.Comment = comment

orgAsset.Status = currentStatus

….
orgAsset0, _ := json.Marshal(orgAsset)

err = stub.PutState(assetId, orgAsset0)

…

return shim.Success(orgAsset0)orgAsset.Comment = comment

orgAsset.Status = currentStatus

….
orgAsset0, _ := json.Marshal(orgAsset)

err = stub.PutState(assetId, orgAsset0)

…

return shim.Success(orgAsset0)orgAsset.Comment = comment

orgAsset.Status = currentStatus

….
orgAsset0, _ := json.Marshal(orgAsset)

err = stub.PutState(assetId, orgAsset0)

…

return shim.Success(orgAsset0)orgAsset.Comment = comment

orgAsset.Status = currentStatus

….
orgAsset0, _ := json.Marshal(orgAsset)

err = stub.PutState(assetId, orgAsset0)

…

return shim.Success(orgAsset0)orgAsset.Comment = comment

orgAsset.Status = currentStatus

….
orgAsset0, _ := json.Marshal(orgAsset)

err = stub.PutState(assetId, orgAsset0)

…

return shim.Success(orgAsset0)orgAsset.Comment = comment

orgAsset.Status = currentStatus

….
orgAsset0, _ := json.Marshal(orgAsset)

err = stub.PutState(assetId, orgAsset0)

…

return shim.Success(orgAsset0)orgAsset.Comment = comment

orgAsset.Status = currentStatus

….
orgAsset0, _ := json.Marshal(orgAsset)

err = stub.PutState(assetId, orgAsset0)

…

return shim.Success(orgAsset0)orgAsset.Comment = comment

orgAsset.Status = currentStatus

….
orgAsset0, _ := json.Marshal(orgAsset)

err = stub.PutState(assetId, orgAsset0)

…

return shim.Success(orgAsset0)
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

