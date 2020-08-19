package magazine

//Subscriber structure
type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

//Employee Struct
type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

//Address structure
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
