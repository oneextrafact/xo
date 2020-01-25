// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// Tuple represents a Postgres Tuple type.
type Tuple struct {
	TupleName string // tuple_name
}

// PgTuples runs a custom query, returning results as Tuple.
func PgTuples(db XODB, schema string) ([]*Tuple, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`pg_catalog.format_type(t.oid, NULL) AS tuple_name ` +
		`FROM pg_catalog.pg_type t ` +
		`JOIN pg_catalog.pg_namespace n ON n.oid = t.typnamespace ` +
		`JOIN pg_catalog.pg_class c on t.typrelid = c.oid ` +
		`WHERE n.nspname = $1 AND c.relkind = 'c'`

	// run query
	XOLog(sqlstr, schema)
	q, err := db.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Tuple{}
	for q.Next() {
		t := Tuple{}

		// scan
		err = q.Scan(&t.TupleName)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}

	return res, nil
}