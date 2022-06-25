{{range .GoQueries}}
const {{.ConstantName}} = {{$.Q}}-- name: {{.MethodName}} {{.Cmd}}
{{escape .SQL}}
{{$.Q}}

{{if .Arg.EmitStruct}}
type {{.Arg.Type}} struct { {{- range .Arg.UniqueFields}}
  {{.Name}} {{.Type}} {{if .Tag}}{{$.Q}}{{.Tag}}{{$.Q}}{{end}}
  {{- end}}
}
{{end}}

{{if .Ret.EmitStruct}}
type {{.Ret.Type}} struct { {{- range .Ret.Struct.Fields}}
  {{.Name}} {{.Type}} {{if .Tag}}{{$.Q}}{{.Tag}}{{$.Q}}{{end}}
  {{- end}}
}
{{end}}

{{if eq .Cmd ":one"}}
{{range .Comments}}//{{.}}
{{end -}}
func (q *Queries) {{.MethodName}}(ctx context.Context, {{.Arg.Pair}}) ({{.Ret.DefineType}}, error) {
  	{{- if $.EmitPreparedQueries}}
	row := q.queryRow(ctx, q.{{.FieldName}}, {{.ConstantName}}, {{.Arg.Params}})
	{{- else if $.EmitMethodsWithDBArgument}}
	row := db.QueryRowContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
	{{- else}}
	row := q.db.QueryRowContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
	{{- end}}
	{{- if ne .Arg.Pair .Ret.Pair }}
	var {{.Ret.Name}} {{.Ret.Type}}
	{{- end}}
	err := row.Scan({{.Ret.Scan}})
	return {{.Ret.ReturnName}}, err
}
{{end}}

{{if eq .Cmd ":many"}}
{{range .Comments}}//{{.}}
{{end -}}
{{- if $.EmitMethodsWithDBArgument -}}
func (q *Queries) {{.MethodName}}(ctx context.Context, db DBTX, {{.Arg.Pair}}) ([]{{.Ret.DefineType}}, error) {
{{- else -}}
func (q *Queries) {{.MethodName}}(ctx context.Context, {{.Arg.Pair}}) ([]{{.Ret.DefineType}}, error) {
{{- end -}}
    {{- if $.EmitPreparedQueries}}
    rows, err := q.query(ctx, q.{{.FieldName}}, {{.ConstantName}}, {{.Arg.Params}})
    {{- else if $.EmitMethodsWithDBArgument}}
    rows, err := db.QueryContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
    {{- else}}
    rows, err := q.db.QueryContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
    {{- end}}
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    {{- if $.EmitEmptySlices}}
    items := []{{.Ret.DefineType}}{}
    {{else}}
    var items []{{.Ret.DefineType}}
    {{end -}}
    for rows.Next() {
        var {{.Ret.Name}} {{.Ret.Type}}
        if err := rows.Scan({{.Ret.Scan}}); err != nil {
            return nil, err
        }
        items = append(items, {{.Ret.ReturnName}})
    }
    if err := rows.Close(); err != nil {
        return nil, err
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return items, nil
}
{{end}}

{{if eq .Cmd ":exec"}}
{{range .Comments}}//{{.}}
{{end -}}
{{- if $.EmitMethodsWithDBArgument -}}
func (q *Queries) {{.MethodName}}(ctx context.Context, db DBTX, {{.Arg.Pair}}) error {
{{- else -}}
func (q *Queries) {{.MethodName}}(ctx context.Context, {{.Arg.Pair}}) error {
{{- end -}}
    {{- if $.EmitPreparedQueries}}
    _, err := q.exec(ctx, q.{{.FieldName}}, {{.ConstantName}}, {{.Arg.Params}})
    {{- else if $.EmitMethodsWithDBArgument}}
    _, err := db.ExecContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
    {{- else}}
    _, err := q.db.ExecContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
    {{- end}}
    return err
}
{{end}}

{{if eq .Cmd ":execrows"}}
{{range .Comments}}//{{.}}
{{end -}}
{{- if $.EmitMethodsWithDBArgument -}}
func (q *Queries) {{.MethodName}}(ctx context.Context, db DBTX, {{.Arg.Pair}}) (int64, error) {
{{- else -}}
func (q *Queries) {{.MethodName}}(ctx context.Context, {{.Arg.Pair}}) (int64, error) {
{{- end -}}
    {{- if $.EmitPreparedQueries}}
    result, err := q.exec(ctx, q.{{.FieldName}}, {{.ConstantName}}, {{.Arg.Params}})
    {{- else if $.EmitMethodsWithDBArgument}}
    result, err := db.ExecContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
    {{- else}}
    result, err := q.db.ExecContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
    {{- end}}
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}
{{end}}

{{if eq .Cmd ":execlastid"}}
{{range .Comments}}//{{.}}
{{end -}}
{{- if $.EmitMethodsWithDBArgument -}}
func (q *Queries) {{.MethodName}}(ctx context.Context, db DBTX, {{.Arg.Pair}}) (int64, error) {
{{- else -}}
func (q *Queries) {{.MethodName}}(ctx context.Context, {{.Arg.Pair}}) (int64, error) {
{{- end -}}
    {{- if $.EmitPreparedQueries}}
    result, err := q.exec(ctx, q.{{.FieldName}}, {{.ConstantName}}, {{.Arg.Params}})
    {{- else if $.EmitMethodsWithDBArgument}}
    result, err := db.ExecContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
    {{- else}}
    result, err := q.db.ExecContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
    {{- end}}
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}
{{end}}

{{if eq .Cmd ":execresult"}}
{{range .Comments}}//{{.}}
{{end -}}
{{- if $.EmitMethodsWithDBArgument -}}
func (q *Queries) {{.MethodName}}(ctx context.Context, db DBTX, {{.Arg.Pair}}) (sql.Result, error) {
{{- else -}}
func (q *Queries) {{.MethodName}}(ctx context.Context, {{.Arg.Pair}}) (sql.Result, error) {
{{- end -}}
    {{- if $.EmitPreparedQueries}}
    return q.exec(ctx, q.{{.FieldName}}, {{.ConstantName}}, {{.Arg.Params}})
    {{- else if $.EmitMethodsWithDBArgument}}
    return db.ExecContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
    {{- else}}
    return q.db.ExecContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
    {{- end}}
}
{{end}}

{{end}}