package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type column struct {
	ColumnName string
	Type       string
	Nullable   string
	Json       string
}

var db *sql.DB

//map for converting mysql type to golang types
var go_mysql_typemap = map[string]string{
	"int":                "int",
	"integer":            "int",
	"tinyint":            "int",
	"smallint":           "int",
	"mediumint":          "int",
	"bigint":             "int",
	"int unsigned":       "int",
	"integer unsigned":   "int",
	"tinyint unsigned":   "int",
	"smallint unsigned":  "int",
	"mediumint unsigned": "int",
	"bigint unsigned":    "int",
	"bit":                "int",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time",
	"datetime":           "time.Time",
	"timestamp":          "time.Time",
	"time":               "time.Time",
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

func main() {
	table_name := flag.String("t", "test_table", "table name")
	usr := flag.String("u", "root", "DB user name")
	pwd := flag.String("p", "", "DB password")
	db_name := flag.String("d", "test", "DB name")
	flag.Parse()
	conn := *usr + ":" + *pwd + "@tcp(127.0.0.1:3306)/" + *db_name
	var err error
	if conn != "" {
		db, err = sql.Open("mysql", conn)
		if err != nil {
			fmt.Println("[ERROR] Could not connect to database: ", err)
			os.Exit(1)
		}
	}

	generateModel(*table_name)

}

//function for generating golang struct
func generateModel(table_name string) {
	err, columns := getColumns(table_name)
	if err != nil {
		return
	}

	table_name = camelCase(table_name)
	depth := 1
	fmt.Print("type " + table_name + " struct {\n")
	for _, v := range columns {
		fmt.Print(tab(depth) + v.ColumnName + " " + v.Type + " " + v.Json)
		fmt.Print("\n")
	}
	fmt.Print(tab(depth-1) + "}\n")
}

// Function for fetching schema definition of passed table
func getColumns(table string) (errr error, columns []column) {
	rows, err := db.Query(`
        SELECT COLUMN_NAME,DATA_TYPE, IS_NULLABLE
        FROM information_schema.COLUMNS
        WHERE table_schema = DATABASE()
            AND TABLE_NAME = ? order by ORDINAL_POSITION`, table)
	if err != nil {
		fmt.Println("Error reading table information: ", err.Error())
		return err, nil
	}
	defer rows.Close()

	for rows.Next() {
		col := column{}
		err := rows.Scan(&col.ColumnName, &col.Type, &col.Nullable)

		if err != nil {
			fmt.Println(err.Error())
			return err, nil
		}

		col.Json = strings.ToLower(col.ColumnName)
		col.ColumnName = camelCase(col.ColumnName)
		col.Type = go_mysql_typemap[col.Type]
		if col.Nullable == "YES" {
			col.Json = fmt.Sprintf("`json:\"%s,omitempty\"`", col.Json)
		} else {
			col.Json = fmt.Sprintf("`json:\"%s\"`", col.Json)
		}
		columns = append(columns, col)
	}
	return err, columns
}

func camelCase(str string) string {
	name := strings.ToLower(str)
	var text string
	for _, p := range strings.Split(name, "_") {
		text += strings.ToUpper(p[0:1]) + p[1:]
	}
	return text
}

func tab(depth int) string {
	return strings.Repeat("\t", depth)
}
