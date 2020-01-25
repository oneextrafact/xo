// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// TupleField represents Tuple Field info.
type TupleField struct {
	FieldOrdinal int16  // field_ordinal
	FieldName    string // field_name
	DataType     string // data_type
	NotNull      bool   // not_null
}

// PgTupleFields runs a custom query, returning results as TupleField.
func PgTupleFields(db XODB, schema string, table string) ([]*TupleField, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`a.attnum                                         AS field_ordinal, ` +
		`a.attname, ` + // ::text                                  AS field_name
		`pg_catalog.format_type(a.atttypid, a.atttypmod)  AS data_type, ` +
		`a.attnotnull                                     AS not_null ` +
		`FROM pg_catalog.pg_attribute a ` +
		`JOIN pg_catalog.pg_type t ` +
		`ON a.attrelid = t.typrelid ` +
		`JOIN pg_catalog.pg_namespace n ` +
		`ON (n.oid = t.typnamespace) ` +
		`WHERE a.attnum > 0 AND NOT a.attisdropped AND n.nspname = $1 AND pg_catalog.format_type(t.oid, NULL) = $2 ` +
		`ORDER BY a.attnum`

	// run query
	XOLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*TupleField{}
	for q.Next() {
		tf := TupleField{}

		// scan
		err = q.Scan(&tf.FieldOrdinal, &tf.FieldName, &tf.DataType, &tf.NotNull)
		if err != nil {
			return nil, err
		}

		res = append(res, &tf)
	}

	return res, nil
}