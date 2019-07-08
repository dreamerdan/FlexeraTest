
// main.go

package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)
 
func main () {
	//Test Data
	People = []Person{
	Person{Id: "1", Name: "Dan", Age: 43, Acc: Account {AccountNumber: 123123343123, Balance: 100.00}, Email: 
"daniel@onethirdavenue.com", Address: Addr{Address1: "2 Langtry Lodge", Address2: "Parkgate", Address3: "Ballyclare", City: "Antrim", PostCode:"Bt39 0LL"}},
	Person{Id: "2", Name: "Andrew", Age: 43, Acc: Account {AccountNumber: 123123123123, Balance: 100.00}, Email: 
"Andrew@hotmal.com", Address: Addr{Address1: "2 somewhere close", Address2: "Belfast", Address3: "Belfast", City: "Belfast", PostCode:"Bt10 056"}},
	Person{Id: "3", Name: "Alice", Age: 33, Acc: Account {AccountNumber: 234345456567, Balance: 10023.00}, Email: 
"Alice@yahoo.com", Address: Addr{Address1: "2 Langtry Lodge", Address2: "Parkgate", Address3: "Ballyclare", City: "Antrim", PostCode:"Bt39 0LL"}},
	Person{Id: "4", Name: "Dave", Age: 50, Acc: Account {AccountNumber: 45623452134235, Balance: 103.00}, Email: 
"dudeman@gmail.com", Address: Addr{Address1: "2 Langtry Lodge", Address2: "Parkgate", Address3: "Ballyclare", City: "Antrim", PostCode:"Bt39 0LL"}},
	Person{Id: "5", Name: "Steve", Age: 62, Acc: Account {AccountNumber: 53487984735, Balance: 1020.00}, Email: 
"zztop@yahoo.com", Address: Addr{Address1: "2 Langtry Lodge", Address2: "Parkgate", Address3: "Ballyclare", City: "Antrim", PostCode:"Bt39 0LL"}},
	Person{Id: "6", Name: "Anne", Age: 22, Acc: Account {AccountNumber: 88394893487, Balance: -1300.00}, Email: 
"123@gmail.com", Address: Addr{Address1: "2 Langtry Lodge", Address2: "Parkgate", Address3: "Ballyclare", City: "Antrim", PostCode:"Bt39 0LL"}},
	Person{Id: "7", Name: "Gill", Age: 45, Acc: Account {AccountNumber: 549238723847, Balance: 3400.00}, Email: 
"345@yahoo.com", Address: Addr{Address1: "2 Langtry Lodge", Address2: "Parkgate", Address3: "Ballyclare", City: "Antrim", PostCode:"Bt39 0LL"}},                       
                        }
	 handleRequests()

}