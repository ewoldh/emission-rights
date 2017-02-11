package util

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"build-chaincode/entities"
	"encoding/json"
)

func TestGettersSuite(t *testing.T) {
	suite.Run(t, new(GettersTestSuite))
}

type GettersTestSuite struct {
	suite.Suite
	user         entities.User
	transaction  entities.Transaction
	company      entities.Company
	etaAccount   entities.ETAAccount
	users        []entities.User
	transactions []entities.Transaction
	companies    []entities.Company
	stub         *shim.MockStub
}

func (suite *GettersTestSuite) SetupTest() {
	suite.stub = shim.NewMockStub("chaincode", nil)
	suite.stub.MockTransactionStart("chaincode")
	suite.user = entities.User{
		UserID: "john",
		ETAAccountID: "1234",
		CompanyID: "1234",
		Salt: "salt",
		Hash: "hash",
	}
	suite.users = []entities.User{suite.user}
	userAsBytes, _ := json.Marshal(suite.user)
	suite.stub.PutState(suite.user.UserID, userAsBytes)
	stringArrayAsBytes, _ := json.Marshal([]string{suite.user.UserID})
	suite.stub.PutState("_users", stringArrayAsBytes)

	suite.transaction = entities.Transaction{
		TransactionID: "asdfasdf",
		Price: 5.25,
		Volume: 123,
		DateOfTransaction: 1486832703466,
		SellDate: 1486832703466,
		Seller: "asdfasdf",
		Buyer: "asdfasdf",
		Status: "asdfasdf",
		Transparent: true,
	}
	suite.transactions = []entities.Transaction{suite.transaction}
	transactionAsBytes, _ := json.Marshal(suite.transaction)
	suite.stub.PutState(suite.transaction.TransactionID, transactionAsBytes)
	stringArrayAsBytes, _ = json.Marshal([]string{suite.transaction.TransactionID})
	suite.stub.PutState("_transactions", stringArrayAsBytes)

	suite.company = entities.Company{
		CompanyID: "1234",
		CompanyName: "name",
		ApprovalStatus: "approved",
	}
	suite.companies = []entities.Company{suite.company}
	companyAsBytes, _ := json.Marshal(suite.company)
	suite.stub.PutState(suite.company.CompanyID, companyAsBytes)
	stringArrayAsBytes, _ = json.Marshal([]string{suite.company.CompanyID})
	suite.stub.PutState("_companies", stringArrayAsBytes)

	suite.etaAccount = entities.ETAAccount{
		ETAAccountID: "test",
		CompanyID: "test",
		Balance: 1234,
		AmountOfTransactions: 123,
	}

	etaAccountAsBytes, _ := json.Marshal(suite.etaAccount)
	suite.stub.PutState(suite.etaAccount.ETAAccountID, etaAccountAsBytes)
}

func (suite *GettersTestSuite) Test_CanGetUserWithoutError() {
	user, err := GetUserByID(suite.stub, suite.user.UserID)
	suite.NoError(err)

	suite.Equal(suite.user, user)
}

func (suite *GettersTestSuite) Test_CantGetUserWithWrongUserID() {
	user, err := GetUserByID(suite.stub, "wrongUserName")
	suite.EqualError(err, "Cannot unmarshall User with id wrongUserName, reason: unexpected end of JSON input")

	suite.EqualValues(user, entities.User{})
}

func (suite *GettersTestSuite) Test_CanGetTransactionWithoutError() {
	transaction, err := GetTransactionByID(suite.stub, suite.transaction.TransactionID)
	suite.NoError(err)

	suite.Equal(suite.transaction, transaction)
}

func (suite *GettersTestSuite) Test_CantGetTransactionWithWrongID() {
	transaction, err := GetTransactionByID(suite.stub, "wrongTranactionID")
	suite.EqualError(err, "Cannot unmarshall transaction with id wrongTranactionID, reason: unexpected end of JSON input")

	suite.EqualValues(transaction, entities.Transaction{})
}

func (suite *GettersTestSuite) Test_CanGetCompanyWithoutError() {
	company, err := GetCompanyByID(suite.stub, suite.company.CompanyID)
	suite.NoError(err)

	suite.Equal(suite.company, company)
}

func (suite *GettersTestSuite) Test_CantGetCompanyWithWrongID() {
	company, err := GetCompanyByID(suite.stub, "wrongCompanyID")
	suite.EqualError(err, "Cannot unmarshall company with id wrongCompanyID, reason: unexpected end of JSON input")

	suite.EqualValues(company, entities.Company{})
}

func (suite *GettersTestSuite) Test_CanGetETAAccountWithoutError() {
	etaAccount, err := GetETAAccountByUserID(suite.stub, suite.etaAccount.ETAAccountID)
	suite.NoError(err)

	suite.Equal(suite.etaAccount, etaAccount)
}

func (suite *GettersTestSuite) Test_CantGetETAAccountWithWrongID() {
	etaAccount, err := GetETAAccountByUserID(suite.stub, "wrongETAAccountID")
	suite.EqualError(err, "Could not retrieve user account with id: wrongETAAccountID, reason: " +
		"Cannot unmarshall User with id wrongETAAccountID, reason: " +
		"unexpected end of JSON input")

	suite.EqualValues(etaAccount, entities.ETAAccount{})
}

func (suite *GettersTestSuite) Test_CanGetAllUsersWithoutError() {
	allUsers, err := GetAllUsers(suite.stub)
	suite.NoError(err)

	suite.Equal(suite.users, allUsers)
}

func (suite *GettersTestSuite) Test_CanGetAllCompaniesWithoutError() {
	allCompanies, err := GetAllCompanies(suite.stub)
	suite.NoError(err)

	suite.Equal(suite.companies, allCompanies)
}

func (suite *GettersTestSuite) Test_CanGetAllTransactionsWithoutError() {
	allTransactions, err := GetAllTransactions(suite.stub)
	suite.NoError(err)

	suite.Equal(suite.transactions, allTransactions)
}
