// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// Enum represents a enum.
type Enum struct {
	EnumName string // enum_name
}

// PgEnums runs a custom query, returning results as Enum.
func PgEnums(db XODB, schema string) ([]*Enum, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`t.typname ` + // ::varchar AS enum_name
		`FROM pg_type t ` +
		`JOIN ONLY pg_namespace n ON n.oid = t.typnamespace ` +
		`JOIN ONLY pg_enum e ON t.oid = e.enumtypid ` +
		`WHERE n.nspname = $1`

	// run query
	XOLog(sqlstr, schema)
	q, err := db.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Enum{}
	for q.Next() {
		e := Enum{}

		// scan
		err = q.Scan(&e.EnumName)
		if err != nil {
			return nil, err
		}

		res = append(res, &e)
	}

	return res, nil
}
