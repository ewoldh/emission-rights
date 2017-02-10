package invokeAndQuery

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"build-chaincode/entities"
	"build-chaincode/util"
	"errors"
)

func getSoldSalesByUserID(stub shim.ChaincodeStubInterface, userID string) ([]entities.Transaction, error) {
	transactions, err := util.GetAllTransactions(stub)
	if err != nil {
		return entities.Transaction{}, errors.New("empty transactions db" +err.Error())
	}

	var userTransaction []entities.Transaction
	for _, transaction:= range transactions {
		if transaction.Seller == userID {
			append(userTransaction, transaction)
		}
	}

	return userTransaction, nil
}

func getBoughtSalesByUserID(stub shim.ChaincodeStubInterface, userID string) ([]entities.Transaction, error) {
	transactions, err := util.GetAllTransactions(stub)
	if err != nil {
		return entities.Transaction{}, errors.New("empty transactions db" +err.Error())
	}

	var userTransaction []entities.Transaction
	for _, transaction:= range transactions {
		if transaction.Buyer == userID {
			append(userTransaction, transaction)
		}
	}

	return userTransaction, nil
}
