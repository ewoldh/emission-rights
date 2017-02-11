package util

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"build-chaincode/entities"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestIndexSuite(t *testing.T) {
	suite.Run(t, new(IndexTestSuite))
}

type IndexTestSuite struct {
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

func (suite *IndexTestSuite) SetupTest() {
	suite.stub = shim.NewMockStub("chaincode", nil)
	suite.stub.MockTransactionStart("chaincode")
	ResetIndexes(suite.stub, shim.NewLogger("chaincode"))
}

func (suite *IndexTestSuite) Test_CanResetIndexesWithoutError() {
	err := ResetIndexes(suite.stub, shim.NewLogger("chaincode"))
	suite.NoError(err)
}

func (suite *IndexTestSuite) Test_CanGetIndexesWithoutError() {
	ResetIndexes(suite.stub, shim.NewLogger("chaincode"))
	_, err := GetIndex(suite.stub, UsersIndexName)
	suite.NoError(err)
}

func (suite *IndexTestSuite) Test_CanWriteIDToBlockchainWithoutError() {
	ResetIndexes(suite.stub, shim.NewLogger("chaincode"))
	_, err := WriteIDToBlockchainIndex(suite.stub, UsersIndexName, "idOfUser")
	suite.NoError(err)
	userIndex, err := GetIndex(suite.stub, UsersIndexName)
	suite.NoError(err)
	idExistInIndex := false
	for _, userID := range userIndex {
		if userID == "idOfUser" {
			idExistInIndex = true
		}
	}
	suite.True(idExistInIndex)

}

func (suite *IndexTestSuite) Test_CanCheckIfIDExistsInIndexesWithoutError() {
	ResetIndexes(suite.stub, shim.NewLogger("chaincode"))
	_, err := WriteIDToBlockchainIndex(suite.stub, UsersIndexName, "idOfUser")
	suite.NoError(err)
	_, err = GetIndex(suite.stub, UsersIndexName)
	suite.NoError(err)
	_, err = DoesIDExistInIndex(suite.stub, "idOfUser", UsersIndexName)
	suite.NoError(err)
}