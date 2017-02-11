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
	user entities.User
	stub           *shim.MockStub

}

func (suite *GettersTestSuite) SetupTest() {
	suite.stub = shim.NewMockStub("chaincode", nil)
	suite.stub.MockTransactionStart("chaincode")
	suite.user = entities.User{
		UserID: "john",
		CompanyID: "1234",
		Salt: "salt",
		Hash: "hash",
	}
	userAsBytes, _ := json.Marshal(suite.user)
	suite.stub.PutState(suite.user.UserID, userAsBytes)
}

func (suite *GettersTestSuite) Test_CanGetUserWithoutError() {
	user, err := GetUser(suite.stub, suite.user.UserID)
	suite.NoError(err)

	suite.Equal(suite.user, user)
}

func (suite *GettersTestSuite) Test_CantGetUserWithWrongUsername() {
	user, err := GetUser(suite.stub, "wrongUserName")
	suite.EqualError(err, "User with id: \"wrongUserName\" is not found in the chaincode")

	suite.EqualValues(user, entities.User{})
}
