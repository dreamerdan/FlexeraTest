
// main.go

package main


//Person structure

type Person struct {
	Id   	string      `json:Id`
	Name 	string      `json:Name`
	Age  	int         `json:Age`
	Acc  	Account     `json:Account`
	Email 	string      `json:Email`
	Address Addr   		`json:Address`
}


//Address Structure
type Addr struct {         
	Address1 string             `json:Address1`
	Address2 string             `json:Address2`
	Address3 string             `json:Address3`
	City     string             `json:City`
	PostCode string             `json:PostCode`
}
 

//Account Structure
type Account struct {
	AccountNumber int           `json:AccountNumber`
	Balance       float64       `json:Balance`
}


var People []Person
