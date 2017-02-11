package invokeAndQuery

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"build-chaincode/entities"
	"encoding/json"
	"build-chaincode/util"
	"errors"
)

func GetETAAccountByUserID(stub shim.ChaincodeStubInterface, userID string) (entities.ETAAccount, error) {
	user, err := util.GetUserByID(stub, userID)
	if err != nil {
		return entities.ETAAccount{}, errors.New("Could not retrieve user account" + err.Error())
	}

	etaAccountAsBytes, err := stub.GetState(user.ETAAccountID)
	if err != nil {
		return entities.ETAAccount{}, errors.New("Could not retrieve eta acoount of the user" + err.Error())
	}

	var etaAccount entities.ETAAccount
	err = json.Unmarshal(etaAccountAsBytes, &etaAccount)
	if err != nil {
		return entities.ETAAccount{}, errors.New("Could not unmarshall eta user acoount" + err.Error())
	}

	return etaAccount, nil
}

func CreateETAs(stub shim.ChaincodeStubInterface, amount int64) error {
	userIsAuthorised, err := stub.VerifyAttribute("role", []byte("authority"))
	if err != nil {
		return errors.New("error while checking certificate attribute, reason: " + err.Error())
	}

	if !userIsAuthorised {
		return  errors.New("this user doesn't have AUTORATAAAHHHH")
	}

	userIDAsBytes, err := stub.ReadCertAttribute("userID")
	if err != nil {
		return errors.New("error while checking certificate attribute, reason: " + err.Error())
	}

	etaAccount, err := GetETAAccountByUserID(stub, string(userIDAsBytes))
	if err != nil {
		return errors.New("error while getting eta acount, reason: " + err.Error())
	}

	etaAccount.Balance += amount
	etaAccountAsBytes, err := json.Marshal(etaAccount)
	if err != nil {
		return errors.New("error marshalling eta account, reason: " + err.Error())
	}

	err = util.StoreObjectInChain(stub, etaAccount.ETAAccountID, util.ETAAccountsIndexName, etaAccountAsBytes)
	if err != nil {
		return errors.New("Error storing eta account, reason: " + err.Error())
	}

	return nil
}