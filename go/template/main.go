package main

import (
	"fmt"
	"html/template"
	"os"
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
	// Used for :copyfrom
	Table string
}

func main() {
	var queries = []Query{
		{SQL: "select * from a", SourceName: "w.sql"},
	}

	tpl, err := template.ParseFiles("./querycode.tpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Printf("%+v\n", tpl.Tree)
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
