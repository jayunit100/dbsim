package components

// PID used for the ID of each struct

type Pet struct {
	Name string
	Type string
	PID  int
}

func (p *Pet) Id() int {
	return p.PID
}

type Store struct {
	Address string
	Pets    map[string]Pet
	State   string
	PID     int
}

func (p *Store) Id() int {
	return p.PID
}

type Transaction struct {
	Store    *Store
	Pet      string
	Returned bool
	PID      int
}

func (p *Transaction) Id() int {
	return p.PID
}

type User struct {
	Name      string
	Purchases []*Transaction
	PID       int
}

func (p *User) Id() int {
	return p.PID
}
