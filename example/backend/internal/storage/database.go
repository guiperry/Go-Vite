package storage

type Database struct {
	// Add your database implementation here
}

func NewDatabase() (*Database, error) {
	return &Database{}, nil
}

func (db *Database) Close() error {
	return nil
}
