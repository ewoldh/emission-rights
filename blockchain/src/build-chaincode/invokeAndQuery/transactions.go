package invokeAndQuery

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"build-chaincode/entities"
	"build-chaincode/util"
	"errors"
	"encoding/json"
)

func GetSoldSalesByUserID(stub shim.ChaincodeStubInterface, userID string) ([]entities.Transaction, error) {
	transactions, err := util.GetAllTransactions(stub)
	if err != nil {
		return []entities.Transaction{}, errors.New("empty transactions db" + err.Error())
	}

	userTransactions := []entities.Transaction{}
	for _, transaction := range transactions {
		if transaction.Seller == userID {
			userTransactions = append(userTransactions, transaction)
		}
	}

	return userTransactions, nil
}

func GetBoughtSalesByUserID(stub shim.ChaincodeStubInterface, userID string) ([]entities.Transaction, error) {
	transactions, err := util.GetAllTransactions(stub)
	if err != nil {
		return []entities.Transaction{}, errors.New("empty transactions db" + err.Error())
	}

	userTransactions := []entities.Transaction{}
	for _, transaction := range transactions {
		if transaction.Buyer == userID {
			userTransactions = append(userTransactions, transaction)
		}
	}

	return userTransactions, nil
}

func GetAllTransactionsOnSale(stub shim.ChaincodeStubInterface) ([]entities.Transaction, error) {
	transactionsIndex, err := util.GetIndex(stub, util.TransactionsIndexName)
	if err != nil {
		return []entities.Transaction{}, errors.New("Could not retrieve transactionIndex, reason: " + err.Error())
	}

	var transactions []entities.Transaction
	for _, transactionID := range transactionsIndex {
		transactionAsBytes, err := stub.GetState(transactionID)
		if err != nil {
			return []entities.Transaction{}, errors.New("Could not retrieve transaction with ID: " + transactionID + ", reason: " + err.Error())
		}

		var transaction entities.Transaction
		err = json.Unmarshal(transactionAsBytes, &transaction)
		if err != nil {
			return []entities.Transaction{}, errors.New("Error while unmarshalling transaction, reason: " + err.Error())
		}

		if (transaction.Status != "Sold") {
			transactions = append(transactions, transaction)
		}
	}

	return transactions, nil
}
