{{- $notVoid := (ne .Proc.ReturnType "void") -}}
{{- $tupleReturn := false -}}
{{- if ge (len .Proc.ReturnType) 5 }}
        {{$tupleReturn = (eq (slice .Proc.ReturnType 0 5) "SETOF")}}
{{- end}}
{{- $proc := (schema .Schema .Proc.ProcName) -}}
{{- if ne .Proc.ReturnType "trigger" -}}
// {{ goFuncName .Name }} calls the stored procedure '{{ $proc }}({{ .ProcParams }}) {{ .Proc.ReturnType }}' on db.
func {{ goFuncName .Name }}({{- if $tupleReturn}}db DB{{- else}}db DB{{- end }}{{ goparamlist .Params true true }}) ({{ if $notVoid }}result {{ retype .Return.Type }}, {{ end }}err error) {
	const stmt = `SELECT {{ if $tupleReturn}}* from {{ end }}{{ $proc }}({{ colvals .Params }})`

{{- if $tupleReturn }}
     //tuple return
{{- end }}
{{- if $notVoid }}


	Log(stmt{{ goparamlist .Params true false }})


	{{- if $tupleReturn}}
	if err = db.Select(&ret, stmt); err != nil {
		return
	}
	{{- else }}
	if err = db.QueryRow(stmt{{ goparamlist .Params true false }}).Scan(&result); err != nil {
		return
	}
	{{- end }}

	return
{{- else }}
	Log(stmt)
	_, err = db.Exec(stmt)
	return
{{- end }}
}
{{- end }}
