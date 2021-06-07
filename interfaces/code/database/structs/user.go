package structs

type User struct {
	ID      int64
	Name    string
	Email   string
	Address *UserAddress
}
