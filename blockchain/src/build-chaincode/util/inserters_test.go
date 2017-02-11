package util

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"build-chaincode/entities"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
)

func TestInsertersSuite(t *testing.T) {
	suite.Run(t, new(InserterTestSuite))
}

type InserterTestSuite struct {
	suite.Suite
	stub *shim.MockStub
	user entities.User
}

func (suite *InserterTestSuite) SetupTest() {
	suite.stub = shim.NewMockStub("chaincode", nil)
	suite.stub.MockTransactionStart("chaincode")
	suite.user = entities.User{
		UserID: "john",
		ETAAccountID: "1234",
		CompanyID: "1234",
		Salt: "salt",
		Hash: "hash",
	}
	ResetIndexes(suite.stub, shim.NewLogger("chaincode"))
}

func (suite *InserterTestSuite) Test_CanStoreDataInChain() {
	userAsBytes, _ := json.Marshal(suite.user)
	err := StoreObjectInChain(suite.stub, suite.user.UserID, UsersIndexName, userAsBytes)
	suite.NoError(err)
}

func (suite *InserterTestSuite) Test_CanRetrieveStoredDataFromChain() {
	userAsBytes, _ := json.Marshal(suite.user)
	err := StoreObjectInChain(suite.stub, suite.user.UserID, UsersIndexName, userAsBytes)
	suite.NoError(err)
	user, err := GetUserByID(suite.stub, suite.user.UserID)
	suite.NoError(err)

	suite.Equal(suite.user, user)
}
func (suite *InserterTestSuite) Test_CanRetrieveUserIDFromIndexAfterStoringInChain() {
	userAsBytes, _ := json.Marshal(suite.user)
	err := StoreObjectInChain(suite.stub, suite.user.UserID, UsersIndexName, userAsBytes)
	suite.NoError(err)

	userIndex, err := GetIndex(suite.stub, UsersIndexName)
	suite.NoError(err)
	idExistInIndex := false
	for _, userID := range userIndex {
		if userID == suite.user.UserID {
			idExistInIndex = true
		}
	}
	suite.True(idExistInIndex)
}
//func (suite *InserterTestSuite) Test_CanStoreDataInChain() {
//	userAsBytes, _ := json.Marshal(suite.user)
//	err := StoreObjectInChain(suite.stub, suite.user.UserID, UsersIndexName, userAsBytes)
//	suite.NoError(err)
//}
