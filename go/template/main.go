package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type QueryValue struct {
	EmitStruct  bool
	EmitPointer bool
	Name        string
	Struct      *Struct
	Typ         string
	SQLPackage  string
}

type Struct struct {
	Table   string
	Name    string
	Fields  []Field
	Comment string
}

type Field struct {
	Name    string // CamelCased name for Go
	DBName  string // Name as used in the DB
	Type    string
	Tags    map[string]string
	Comment string
}

type Query struct {
	Cmd          string
	Comments     []string
	MethodName   string
	FieldName    string
	ConstantName string
	SQL          string
	SourceName   string
	Ret          QueryValue
	Arg          QueryValue
	Table        string
}

func EscapeBacktick(s string) string {
	return strings.Replace(s, "`", "`+\"`\"+`", -1)
}

func main() {
	var queries = []Query{
		{SQL: "select * from a", SourceName: "w.sql", Cmd: ":one"},
	}

	funcMap := template.FuncMap{
		"escape": EscapeBacktick,
	}

	// parsefiles 会出错，而 parse 就不会
	//tpl, err := template.New("titletext").Funcs(funcMap).ParseFiles("./querycode.tpl")
	tpl, err := template.New("titletext").Funcs(funcMap).Parse(txt)

	//tpl, err := template.ParseFiles("./querycode.tpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	var tplData = struct {
		GoQueries   []Query
		OutputQuery bool
		Q           string
	}{
		GoQueries:   queries,
		OutputQuery: true,
		Q:           "getauthor",
	}
	err = tpl.Execute(os.Stdout, tplData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

var txt = `
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
func (q *Queries) {{.MethodName}}(ctx context.Context, {{.Arg.Pair}}) ([]{{.Ret.DefineType}}, error) {
    {{- if $.EmitPreparedQueries}}
    rows, err := q.query(ctx, q.{{.FieldName}}, {{.ConstantName}}, {{.Arg.Params}})
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
func (q *Queries) {{.MethodName}}(ctx context.Context, {{.Arg.Pair}}) error {
    {{- if $.EmitPreparedQueries}}
    _, err := q.exec(ctx, q.{{.FieldName}}, {{.ConstantName}}, {{.Arg.Params}})
    {{- else}}
    _, err := q.db.ExecContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
    {{- end}}
    return err
}
{{end}}

{{if eq .Cmd ":execrows"}}
{{range .Comments}}//{{.}}
{{end -}}
func (q *Queries) {{.MethodName}}(ctx context.Context, {{.Arg.Pair}}) (int64, error) {
    {{- if $.EmitPreparedQueries}}
    result, err := q.exec(ctx, q.{{.FieldName}}, {{.ConstantName}}, {{.Arg.Params}})
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
func (q *Queries) {{.MethodName}}(ctx context.Context, {{.Arg.Pair}}) (int64, error) {
    {{- if $.EmitPreparedQueries}}
    result, err := q.exec(ctx, q.{{.FieldName}}, {{.ConstantName}}, {{.Arg.Params}})
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
func (q *Queries) {{.MethodName}}(ctx context.Context, {{.Arg.Pair}}) (sql.Result, error) {
    {{- if $.EmitPreparedQueries}}
    return q.exec(ctx, q.{{.FieldName}}, {{.ConstantName}}, {{.Arg.Params}})
    {{- else}}
    return q.db.ExecContext(ctx, {{.ConstantName}}, {{.Arg.Params}})
    {{- end}}
}
{{end}}

{{end}}
`

func (v QueryValue) IsStruct() bool {
	return v.Struct != nil
}

func (v QueryValue) IsPointer() bool {
	return v.EmitPointer && v.Struct != nil
}

func (v QueryValue) isEmpty() bool {
	return v.Typ == "" && v.Name == "" && v.Struct == nil
}

func (v QueryValue) Pair() string {
	if v.isEmpty() {
		return ""
	}
	return v.Name + " " + v.DefineType()
}

func (v QueryValue) SlicePair() string {
	if v.isEmpty() {
		return ""
	}
	return v.Name + " []" + v.DefineType()
}

func (v QueryValue) Type() string {
	if v.Typ != "" {
		return v.Typ
	}
	if v.Struct != nil {
		return v.Struct.Name
	}
	panic("no type for QueryValue: " + v.Name)
}

func (v *QueryValue) DefineType() string {
	t := v.Type()
	if v.IsPointer() {
		return "*" + t
	}
	return t
}

func (v *QueryValue) ReturnName() string {
	if v.IsPointer() {
		return "&" + v.Name
	}
	return v.Name
}

func (v QueryValue) UniqueFields() []Field {
	seen := map[string]struct{}{}
	fields := make([]Field, 0, len(v.Struct.Fields))

	for _, field := range v.Struct.Fields {
		if _, found := seen[field.Name]; found {
			continue
		}
		seen[field.Name] = struct{}{}
		fields = append(fields, field)
	}

	return fields
}

func (v QueryValue) ColumnNames() string {
	if v.Struct == nil {
		return fmt.Sprintf("[]string{%q}", v.Name)
	}
	escapedNames := make([]string, len(v.Struct.Fields))
	for i, f := range v.Struct.Fields {
		escapedNames[i] = fmt.Sprintf("%q", f.DBName)
	}
	return "[]string{" + strings.Join(escapedNames, ", ") + "}"
}

func (v QueryValue) Scan() string {
	var out []string
	if v.Struct == nil {
		if strings.HasPrefix(v.Typ, "[]") && v.Typ != "[]byte" {
			out = append(out, "pq.Array(&"+v.Name+")")
		} else {
			out = append(out, "&"+v.Name)
		}
	} else {
		for _, f := range v.Struct.Fields {
			if strings.HasPrefix(f.Type, "[]") && f.Type != "[]byte" { //&& v.SQLPackage != SQLPackagePGX {
				out = append(out, "pq.Array(&"+v.Name+"."+f.Name+")")
			} else {
				out = append(out, "&"+v.Name+"."+f.Name)
			}
		}
	}
	if len(out) <= 3 {
		return strings.Join(out, ",")
	}
	out = append(out, "")
	return "\n" + strings.Join(out, ",\n")
}
