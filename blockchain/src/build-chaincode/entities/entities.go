package entities

type ECertResponse struct {
	OK string `json:"OK"`
}

type TestData struct {
	Users  		[]User 	 `json:"users"`
	Things 		[]Thing  `json:"things"`
}

type TestDataElement interface {
	ID() string
}

type User struct {
	UserID   	string 	`json:"userID"`
	ETAAccountID	string  `json:"etaAccountID"`
	Username 	string 	`json:"username"`
	Salt     	string 	`json:"salt"`
	Hash     	string 	`json:"hash"`
	CompanyID	string  `json:"companyID"`
}

type Transaction struct {
	TransactionID  	string 	`json:"transactionID"`
	Price	 	float64 `json:"price"`
	Volume     	int64 	`json:"volume"`
	TransactionDate int64 	`json:"transactionDate"`
	Seller		string  `json:"seller"`
	Buyer		string  `json:"buyer"`
	RequestStatus	string	`json:"requestStatus"`
	Transparent	bool	`json:"transparent"`
}

type Company struct {
	CompanyID  	string 	`json:"companyID"`
	CompanyName 	string  `json:"companyName"`
	Department     	string 	`json:"department"`
	ApprovalStatus  string 	`json:"approvalStatus"`
}

type ETAAccount struct {
	ETAAccountID	      string    `json:"etaAccountID"`
	CompanyID 	      string    `json:"companyID"`
	Balance     	      int64 	`json:"balance"`
	AmountOfTransactions  int64 	`json:"amountOfTransactions"`
}

type BankAccount struct {
	BankAccountID	  int64   `json:"bankAccountID"`
	UserID    	  string  `json:"userID"`
	CompanyID 	  string  `json:"companyID"`
	AccountBalance    float64 `json:"accountBalance"`
}

type Thing struct {
	ThingID      	string 	`json:"thingID"`
	SomeProperty 	string 	`json:"someProperty"`
	UserID    	string 	`json:"userID"`
}

type UserAuthenticationResult struct {
	User        	User
	Authenticated 	bool
}

type Users struct {
	Users []User `json:"users"`
}

func (t *User) ID() string {
	return t.Username
}

func (t *Thing) ID() string {
	return t.ThingID
}