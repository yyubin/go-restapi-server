package database

import "admin-server/admin/queries"

type Queries struct {
	*queries.NoticeQueries
	*queries.EmailQueries
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		NoticeQueries: &queries.NoticeQueries{DB: db}, // from Book model
		EmailQueries:  &queries.EmailQueries{DB: db},
	}, nil
}
