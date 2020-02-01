// XODB is the common interface for database operations that can be used with
// types from schema '{{ schema .Schema }}'.
//
// This should work with database/sql.DB and database/sql.Tx.
type XODB interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row

	ExecContext(string, ...interface{}) (sql.Result, error)
	QueryContext(string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(string, ...interface{}) *sql.Row
}

// Log provides the log func used by generated queries.
var Log = func(string, ...interface{}) { }
