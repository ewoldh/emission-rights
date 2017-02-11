package util

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"errors"
	"encoding/json"
	"build-chaincode/entities"
)

func GetCurrentBlockchainUser(stub shim.ChaincodeStubInterface) (entities.User, error) {
	userIDAsBytes, err := stub.ReadCertAttribute("userID")
	if err != nil {
		return entities.User{}, errors.New("Could not retrieve user by certificate. Reason: " + err.Error())
	}

	return GetUser(stub, string(userIDAsBytes))
}

func GetUser(stub shim.ChaincodeStubInterface, userID string) (entities.User, error) {
	userAsBytes, err := stub.GetState(userID)
	if err != nil {
		return entities.User{}, errors.New("Could not retrieve information for this user")
	}

	var user entities.User
	if err = json.Unmarshal(userAsBytes, &user); err != nil {
		return entities.User{}, errors.New("Cannot get user, reason: " + err.Error())
	}

	return user, nil
}

func GetCompanyByID(stub shim.ChaincodeStubInterface, companyID string) (entities.Company, error) {
	companyAsBytes, err := stub.GetState(companyID)
	if err != nil {
		return entities.Company{}, errors.New("Could not retrieve information for this user")
	}

	var company entities.Company
	if err = json.Unmarshal(companyAsBytes, &company); err != nil {
		return entities.Company{}, errors.New("Cannot get user, reason: " + err.Error())
	}

	return company, nil
}

func GetAllUsers(stub shim.ChaincodeStubInterface) ([]entities.User, error) {
	usersIndex, err := GetIndex(stub, UsersIndexName)
	if err != nil {
		return []entities.User{}, errors.New("Could not retrieve userIndex, reason: " + err.Error())
	}

	var users []entities.User
	for _, userID := range usersIndex {
		userAsBytes, err := stub.GetState(userID)
		if err != nil {
			return []entities.User{}, errors.New("Could not retrieve user with ID: " + userID + ", reason: " + err.Error())
		}

		var user entities.User
		err = json.Unmarshal(userAsBytes, &user)
		if err != nil {
			return []entities.User{}, errors.New("Error while unmarshalling user, reason: " + err.Error())
		}

		users = append(users, user)
	}

	return users, nil
}

func GetAllCompanies(stub shim.ChaincodeStubInterface) ([]entities.Company, error) {
	companies := []entities.Company{}
	companyIndex, err := GetIndex(stub, CompaniesIndexName)
	if err != nil {
		return []entities.Company{}, errors.New("Error while getting companyIndex, reason: " + err.Error())
	}

	for _, companyID := range companyIndex {
		company, err := GetCompanyByID(stub, companyID)
		if err != nil {
			return []entities.Company{}, errors.New("Error while getting company, reason: " + err.Error())
		}
		companies = append(companies, company)
	}

	return companies, nil
}

func GetETAAccountByUserID(stub shim.ChaincodeStubInterface, userID string) (entities.ETAAccount, error) {
	user, err := GetUser(stub, userID)
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

func GetAllTransactions(stub shim.ChaincodeStubInterface) ([]entities.Transaction, error) {
	transactionsIndex, err := GetIndex(stub, TransactionsIndexName)
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

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
