package methods

//import "github.com/google/uuid"

type Customer struct {
	// Primary key can only be assigned with name "Id", any other name does not work
	Id        int    `json:"id,omitempty,"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
}
