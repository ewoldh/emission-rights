package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"os"
	"build-chaincode/util"
	"build-chaincode/entities"
	"build-chaincode/invokeAndQuery"
	"strconv"
)

var logger = shim.NewLogger("fabric-boilerplate")
//======================================================================================================================
//	 Structure Definitions
//======================================================================================================================
//	SimpleChaincode - A blank struct for use with Shim (An IBM Blockchain included go file used for get/put state
//					  and other IBM Blockchain functions)
//==============================================================================================================================
type Chaincode struct {
}

//======================================================================================================================
//	Invoke - Called on chaincode invoke. Takes a function name passed and calls that function. Passes the
//  		 initial arguments passed are passed on to the called function.
//======================================================================================================================

func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface, functionName string, args []string) ([]byte, error) {
	logger.Infof("Invoke is running " + functionName)

	if functionName == "init" {
		return t.Init(stub, "init", args)
	} else if functionName == "resetIndexes" {
		return nil, util.ResetIndexes(stub, logger)
	} else if functionName == "addUser" {
		return nil, t.addUser(stub, args[0], args[1])
	} else if functionName == "addTestdata" {
		return nil, t.addTestdata(stub, args[0])
	} else if functionName == "createTransaction" {
		transaction := entities.Transaction{}
		json.Unmarshal([]byte(args[0]), &transaction)
		return nil, invokeAndQuery.CreateTransaction(stub, transaction)
	} else if functionName == "createETAs" {
		amount, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return nil, errors.New("Couldn't convert string to int, reason: " + err.Error())
		}

		return nil, invokeAndQuery.CreateETAs(stub, amount)
	} else if functionName == "finaliseTransaction" {
		transactionID := args[0]
		timeOfTransaction, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return nil, errors.New("Couldn't convert string to int, reason: " + err.Error())
		}

		return nil, invokeAndQuery.FinaliseTrade(stub, transactionID, timeOfTransaction)
	}

	return nil, errors.New("Received unknown invoke function name")
}

//======================================================================================================================
//	Query - Called on chaincode query. Takes a function name passed and calls that function. Passes the
//  		initial arguments passed are passed on to the called function.
//=================================================================================================================================
func (t *Chaincode) Query(stub shim.ChaincodeStubInterface, functionName string, args []string) ([]byte, error) {
	logger.Infof("Query is running " + functionName)

	result, err := t.GetQueryResult(stub, functionName, args)
	if err != nil {
		return nil, err
	}

	return json.Marshal(result)
}

func (t *Chaincode) GetQueryResult(stub shim.ChaincodeStubInterface, functionName string, args []string) (interface{}, error) {
	if functionName == "getUser" {
		user, err := util.GetUserByID(stub, args[0])
		if err != nil {
			return nil, err
		}

		return user, nil
	} else if functionName == "authenticateAsUser" {
		user, err := util.GetUserByID(stub, args[0])
		if err != nil {
			logger.Infof("User with id %v not found.", args[0])
		}

		return t.authenticateAsUser(stub, user, args[1]), nil
	} else if functionName == "getAllSoldTransactionsByUserID" {
		soldSalesByUserID, err := invokeAndQuery.GetSoldSalesByUser(stub)
		if err != nil {
			return nil, errors.New("could not retrieve things by user id: " + args[0] + ", reason: " + err.Error())
		}

		return soldSalesByUserID, nil
	} else if functionName == "getAllBoughtTransactionsByUserID" {
		soldSalesByUserID, err := invokeAndQuery.GetBoughtSalesByUser(stub)
		if err != nil {
			return nil, errors.New("could not retrieve things by user id: " + args[0] + ", reason: " + err.Error())
		}

		return soldSalesByUserID, nil
	} else if functionName =="getAllTransactionsOnSale"{
		allTransactionsOnSale, err := invokeAndQuery.GetAllTransactionsOnSale(stub)
		if err != nil{
			return nil, errors.New("could not retrieve Sales transactions" + args[0] + "reason: " + err.Error())
		}

		return allTransactionsOnSale, nil
	} else if functionName == "getETAAccountByUserID" {
		etaAccountByUserID, err := invokeAndQuery.GetETAAccountByUserID(stub, args[0])
		if err != nil {
			return nil, errors.New("Could not retrieve an account with that user ID" + args[0] + ", reason: " + err.Error())
		}

		return etaAccountByUserID, nil
	} else if functionName == "getAllCompanies" {
		allCompanies, err := util.GetAllCompanies(stub)
		if err != nil {
			return nil, errors.New("Could not retrieve any companies, reason: " + err.Error())
		}

		return allCompanies, nil
	}

	return nil, errors.New("Received unknown query function name")
}

//======================================================================================================================
//  Main - main - Starts up the chaincode
//======================================================================================================================

func main() {
	// LogDebug, LogInfo, LogNotice, LogWarning, LogError, LogCritical (Default: LogDebug)
	logger.SetLevel(shim.LogInfo)

	logLevel, _ := shim.LogLevel(os.Getenv("SHIM_LOGGING_LEVEL"))
	shim.SetLoggingLevel(logLevel)

	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Error starting SimpleChaincode: %s", err)
	}
}

//======================================================================================================================
//  Init Function - Called when the user deploys the chaincode
//======================================================================================================================

func (t *Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

//======================================================================================================================
//  Invoke Functions
//======================================================================================================================

func (t *Chaincode) addUser(stub shim.ChaincodeStubInterface, index string, userJSONObject string) error {
	id, err := util.WriteIDToBlockchainIndex(stub, util.UsersIndexName, index)
	if err != nil {
		return errors.New("Error creating new id for user " + index)
	}

	err = stub.PutState(string(id), []byte(userJSONObject))
	if err != nil {
		return errors.New("Error putting user data on ledger")
	}

	return nil
}

func (t *Chaincode) addTestdata(stub shim.ChaincodeStubInterface, testDataAsJson string) error {
	var testData entities.TestData
	err := json.Unmarshal([]byte(testDataAsJson), &testData)
	if err != nil {
		return errors.New("Error while unmarshalling testdata")
	}

	for _, user := range testData.Users {
		userAsBytes, err := json.Marshal(user);
		if err != nil {
			return errors.New("Error marshalling testUser, reason: " + err.Error())
		}

		err = util.StoreObjectInChain(stub, user.UserID, util.UsersIndexName, userAsBytes)
		if err != nil {
			return errors.New("error in storing object, reason: " + err.Error())
		}
	}

	for _, company := range testData.Companies {
		thingAsBytes, err := json.Marshal(company);
		if err != nil {
			return errors.New("Error marshalling testThing, reason: " + err.Error())
		}

		err = util.StoreObjectInChain(stub, company.CompanyID, util.CompaniesIndexName, thingAsBytes)
		if err != nil {
			return errors.New("error in storing object, reason: " + err.Error())
		}
	}
	for _, transaction := range testData.Transactions {
		thingAsBytes, err := json.Marshal(transaction);
		if err != nil {
			return errors.New("Error marshalling testThing, reason: " + err.Error())
		}

		err = util.StoreObjectInChain(stub, transaction.TransactionID, util.TransactionsIndexName, thingAsBytes)
		if err != nil {
			return errors.New("error in storing object, reason: " + err.Error())
		}
	}

	for _, etaAccount := range testData.ETAAcounts {
		thingAsBytes, err := json.Marshal(etaAccount);
		if err != nil {
			return errors.New("Error marshalling testThing, reason: " + err.Error())
		}

		err = util.StoreObjectInChain(stub, etaAccount.ETAAccountID, util.ETAAccountsIndexName, thingAsBytes)
		if err != nil {
			return errors.New("error in storing object, reason: " + err.Error())
		}
	}

	for _, bankAccount := range testData.BankAccounts {
		thingAsBytes, err := json.Marshal(bankAccount);
		if err != nil {
			return errors.New("Error marshalling testThing, reason: " + err.Error())
		}

		err = util.StoreObjectInChain(stub, bankAccount.BankAccountID, util.BankAccountsIndexName, thingAsBytes)
		if err != nil {
			return errors.New("error in storing object, reason: " + err.Error())
		}
	}

	return nil
}

//======================================================================================================================
//		Query Functions
//======================================================================================================================

func (t *Chaincode) authenticateAsUser(stub shim.ChaincodeStubInterface, user entities.User, passwordHash string) (entities.UserAuthenticationResult) {
	if user == (entities.User{}) {
		fmt.Println("User not found")
		return entities.UserAuthenticationResult{
			User: user,
			Authenticated: false,
		}
	}

	if user.Hash != passwordHash {
		fmt.Println("Hash does not match")
		return entities.UserAuthenticationResult{
			User: user,
			Authenticated: false,
		}
	}

	return entities.UserAuthenticationResult{
		User: user,
		Authenticated: true,
	}
}

