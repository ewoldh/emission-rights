package entities

type ECertResponse struct {
	OK string `json:"OK"`
}

type TestData struct {
	Users  		[]User 	 `json:"users"`
	Companies 	[]Company  `json:"companies"`
	Transactions 	[]Transaction  `json:"transactions"`
	ETAAcounts 	[]ETAAccount  `json:"etaAccounts"`
	BankAccounts 	[]BankAccount  `json:"bankAccounts"`

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
	TransactionID   string 	`json:"transactionID"`
	Price           float64 `json:"price"`
	Volume          int64 	`json:"volume"`
	TransactionDate int64 	`json:"transactionDate"`
	Seller          string  `json:"seller"`
	Buyer           string  `json:"buyer"`
	Status          string	`json:"status"`
	Transparent     bool	`json:"transparent"`
}

type Company struct {
	CompanyID  	string 	`json:"companyID"`
	CompanyName 	string  `json:"companyName"`
	ApprovalStatus  string 	`json:"approvalStatus"`
}

type ETAAccount struct {
	ETAAccountID	      string    `json:"etaAccountID"`
	CompanyID 	      string    `json:"companyID"`
	Balance     	      int64 	`json:"balance"`
	AmountOfTransactions  int64 	`json:"amountOfTransactions"`
}

type BankAccount struct {
	BankAccountID	  string  `json:"bankAccountID"`
	UserID    	  string  `json:"userID"`
	CompanyID 	  string  `json:"companyID"`
	AccountBalance    float64 `json:"accountBalance"`
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
func (t *Transaction) ID() string {
	return t.TransactionID
}
func (t *Company) ID() string {
	return t.CompanyID
}
func (t *ETAAccount) ID() string {
	return t.ETAAccountID
}
func (t *BankAccount) ID() string {
	return t.BankAccountID
}
