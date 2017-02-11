package util
//File to use for invoke queries that insert things into the blockchain

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"errors"
	"build-chaincode/entities"
	"encoding/json"
)

func AddTransaction(stub shim.ChaincodeStubInterface, transaction entities.Transaction) error {
	alreadyExist, err := DoesIDExistInIndex(stub, transaction.TransactionID, TransactionsIndexName)
	if err != nil {
		return errors.New("Could not check if ID uniques, reason: " + err.Error())
	}

	if (alreadyExist) {
		return errors.New("transaction id already stored in chaincode")
	}

	userID, err := stub.ReadCertAttribute("userID")
	if err != nil {
		return errors.New("error while checking certificate attribute, reason: " + err.Error())
	}

	etaAccount, err := GetETAAccountByUserID(stub, string(userID))
	if err != nil {
		return errors.New("error getting account, reason: " + err.Error())
	}

	if etaAccount.Balance < transaction.Volume {
		return errors.New("not enough balance to start trade")
	}

	etaAccount.Balance = etaAccount.Balance - transaction.Volume
	transaction.Seller = string(userID)

	transactionAsBytes, err := json.Marshal(transaction)
	if err != nil {
		return errors.New("error while marshalling transaction" + err.Error())
	}

	etaAccountAsBytes, err := json.Marshal(etaAccount)
	if err != nil {
		return errors.New("error while marshalling eta account" + err.Error())
	}

	err = StoreObjectInChain(stub, transaction.TransactionID, TransactionsIndexName, transactionAsBytes)
	if err != nil {
		return errors.New("Could not create the transaction request, reason: " + err.Error())
	}

	err = StoreObjectInChain(stub, etaAccount.ETAAccountID, ETAAccountsIndexName, etaAccountAsBytes)
	if err != nil {
		return errors.New("Could not create the transaction request, reason: " + err.Error())
	}

	return nil
}

func StoreObjectInChain(stub shim.ChaincodeStubInterface, objectID string, indexName string, object []byte) error {
	ID, err := WriteIDToBlockchainIndex(stub, indexName, objectID)
	if err != nil {
		return errors.New("Writing ID to index: " + indexName + "Reason: " + err.Error())
	}

	fmt.Println("adding: ", string(object))

	err = stub.PutState(string(ID), object)
	if err != nil {
		return errors.New("Putstate error: " + err.Error())
	}

	return nil
}