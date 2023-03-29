package database

type Database struct {
	f1 string
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Get(str string) string {
	d.f1 = str
	return str
}
