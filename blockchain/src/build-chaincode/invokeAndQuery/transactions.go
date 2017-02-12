package invokeAndQuery

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"build-chaincode/entities"
	"build-chaincode/util"
	"errors"
	"encoding/json"
)

func CreateTransaction(stub shim.ChaincodeStubInterface, transaction entities.Transaction) error {
	alreadyExist, err := util.DoesIDExistInIndex(stub, transaction.TransactionID, util.TransactionsIndexName)
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

	err = util.StoreObjectInChain(stub, transaction.TransactionID, util.TransactionsIndexName, transactionAsBytes)
	if err != nil {
		return errors.New("Could not create the transaction request, reason: " + err.Error())
	}

	err = util.StoreObjectInChain(stub, etaAccount.ETAAccountID, util.ETAAccountsIndexName, etaAccountAsBytes)
	if err != nil {
		return errors.New("Could not create the transaction request, reason: " + err.Error())
	}

	return nil
}

func GetSoldSalesByUser(stub shim.ChaincodeStubInterface) ([]entities.Transaction, error) {
	transactions, err := util.GetAllTransactions(stub)
	if err != nil {
		return []entities.Transaction{}, errors.New("empty transactions db" + err.Error())
	}

	userIDAsBytes, err := stub.ReadCertAttribute("userID")
	if err != nil {
		return []entities.Transaction{}, errors.New("error while checking certificate attribute, reason: " + err.Error())
	}

	userTransactions := []entities.Transaction{}
	for _, transaction := range transactions {
		if transaction.Seller == string(userIDAsBytes) {
			userTransactions = append(userTransactions, transaction)
		}
	}

	return userTransactions, nil
}

func GetBoughtSalesByUser(stub shim.ChaincodeStubInterface) ([]entities.Transaction, error) {
	transactions, err := util.GetAllTransactions(stub)
	if err != nil {
		return []entities.Transaction{}, errors.New("empty transactions db" + err.Error())
	}

	userIDAsBytes, err := stub.ReadCertAttribute("userID")
	if err != nil {
		return []entities.Transaction{}, errors.New("error while checking certificate attribute, reason: " + err.Error())
	}

	userTransactions := []entities.Transaction{}
	for _, transaction := range transactions {
		if transaction.Buyer == string(userIDAsBytes) {
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
