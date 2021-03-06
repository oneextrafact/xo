// Package models contains the types for schema 'public'.
package models

import "log"

// Code generated by xo. DO NOT EDIT.

// ProcParam represents a stored procedure param.
type ProcParam struct {
	ParamName     string // param_name
	ParamType     string // param_type
	ParamNullable bool   // param nullable
}

// PgProcParams runs a custom query, returning results as ProcParam.
func PgProcParams(db XODB, schema string, proc string) ([]*ProcParam, error) {
	var err error

	// sql query
	// const sqlstr = `SELECT ` +
	// 	`UNNEST(p.proargnames) as param_name, ` +
	// 	`UNNEST(STRING_TO_ARRAY(oidvectortypes(p.proargtypes), ', ')) ` + // ::varchar AS param_type
	// 	`FROM pg_proc p ` +
	// 	`JOIN ONLY pg_namespace n ON p.pronamespace = n.oid ` +
	// 	`WHERE n.nspname = $1 AND p.proname = $2`

	const sqlstr = `
	WITH
		args AS (
			SELECT
				UNNEST(p.proargnames) AS param_name,
				UNNEST(STRING_TO_ARRAY(oidvectortypes(p.proargtypes), ', ')) AS param_type,
				UNNEST(STRING_TO_ARRAY(pg_get_function_arguments(p.oid), ', ')) AS param_type_full
			FROM
				pg_proc p
					JOIN ONLY pg_namespace n ON p.pronamespace = n.oid
			WHERE
					n.nspname = $1
			AND p.proname = $2
		)
	SELECT
		param_name,
		param_type,
		param_type_full ILIKE '%DEFAULT%' AS param_nullable
	FROM
		args`

	// run query
	XOLog(sqlstr, schema, proc)

	q, err := db.Query(sqlstr, schema, proc)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	log.Println("\n\n\n", proc)

	// load results
	res := []*ProcParam{}
	for q.Next() {
		pp := ProcParam{}

		// scan
		err = q.Scan(&pp.ParamName, &pp.ParamType, &pp.ParamNullable)
		if err != nil {
			return nil, err
		}

		res = append(res, &pp)

		log.Println(proc, pp.ParamName, pp.ParamType, pp.ParamNullable)
	}

	return res, nil
}
