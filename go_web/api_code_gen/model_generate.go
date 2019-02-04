package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
)

var (
	databaseName = flag.String("db", "", "database name")
	tableName    = flag.String("t", "", "table name")
	host         = flag.String("h", "", "mysql host")
	port         = flag.String("P", "", "mysql port")
	username     = flag.String("u", "", "username")
	password     = flag.String("p", "", "password")
)

func main() {
	flag.Parse()
	if *databaseName == "" || *tableName == "" || *host == "" {
		fmt.Println("Lacking of required flag: -db -t or host ?")
		flag.Usage()
		return
	}

	if *port == "" {
		*port = "3306"
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", *username, *password, *host, *port, *databaseName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("connect to mysql error\n", err)
		os.Exit(1)
	}

	rows, err := db.Query("select table_name, column_name, ordinal_position, data_type, column_type, column_comment from information_schema.columns where table_schema = ? and table_name = ?", *databaseName, *tableName)
	if err != nil {
		fmt.Println("MySQL says :\n", err)
		fmt.Println("Wrong username or password ?")
		os.Exit(1)
	}

	defer rows.Close()

	var cols []column

	for rows.Next() {
		col := column{}
		rows.Scan(&col.TableName, &col.ColumnName, &col.OrdinalPosition, &col.DataType, &col.ColumnType, &col.ColumnComment)

		tempArr := strings.Split(col.ColumnType, " ")
		if len(tempArr) > 1 {
			col.Key = fmt.Sprintf("%s %s", col.DataType, tempArr[1])
		} else {
			col.Key = col.DataType
		}

		cols = append(cols, col)
	}

	structDefinition := generateStructDefinitionStr(cols, *tableName)

	modelTemplate = strings.Replace(modelTemplate, `{{struct_def}}`, structDefinition, -1)

	modelTemplate = strings.Replace(modelTemplate, `{{struct_name}}`, getStructName(*tableName, 2), -1)

	modelTemplate = strings.Replace(modelTemplate, `{{table_name}}`, *tableName, -1)

	modelTemplate = strings.Replace(modelTemplate, `{{db_name}}`, *databaseName, -1)

	//fmt.Printf("%#v\n", cols)
	fmt.Println(modelTemplate)

}

var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

var typeMapping = map[string]string{
	"int":                "int", // int signed
	"integer":            "int",
	"tinyint":            "int8",
	"smallint":           "int16",
	"mediumint":          "int32",
	"bigint":             "int64",
	"int unsigned":       "uint", // int unsigned
	"integer unsigned":   "uint",
	"tinyint unsigned":   "uint8",
	"smallint unsigned":  "uint16",
	"mediumint unsigned": "uint32",
	"bigint unsigned":    "uint64",
	"bit":                "uint64",
	"bool":               "bool",   // boolean
	"enum":               "string", // enum
	"set":                "string", // set
	"varchar":            "string", // string & text
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string", // blob
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time", // time
	"datetime":           "time.Time",
	"timestamp":          "time.Time",
	"time":               "time.Time",
	"float":              "float32", // float & decimal
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string", // binary
	"varbinary":          "string",
}

type column struct {
	TableName       string `db:"TABLE_NAME"`
	ColumnName      string `db:"COLUMN_NAME"`
	OrdinalPosition int    `db:"ORDINAL_POSITION"`
	DataType        string `db:"DATA_TYPE"`
	ColumnType      string `db:"COLUMN_TYPE"`
	ColumnComment   string `db:"COLUMN_COMMENT"`
	Key             string
}

func getStructName(tableName string, reserveWordNum int) string {
	words := strings.Split(tableName, "_")
	var structName string

	if reserveWordNum > len(words) {
		reserveWordNum = len(words)
	}

	startIndex := len(words) - reserveWordNum
	for i := startIndex; i < len(words); i++ {
		structName += strings.Title(words[i])
	}

	return structName
}

func camelCase(in string) string {
	tokens := strings.Split(in, "_")
	for i := range tokens {
		tokens[i] = strings.Title(strings.Trim(tokens[i], " "))
	}
	return strings.Join(tokens, "")
}

func generateStructDefinitionStr(colList []column, tableName string) string {
	var content, header, tail, line string
	header = "type {{struct_name}} struct{\n"
	tail = "}\n"
	for _, col := range colList {
		// col name ===== field type ====== tag ======== comment
		line = "\t" + lintName(camelCase(col.ColumnName)) + "\t" +
			typeMapping[col.Key] + "\t" +
			fmt.Sprintf("`db:\"%v\"`", col.ColumnName) + "\t//" +
			col.ColumnComment

		content += line + "\n"
	}

	content = header + content + tail
	return content
}

// lintName returns a different name if it should be different.
func lintName(name string) (should string) {
	// Fast path for simple cases: "_" and all lowercase.
	if name == "_" {
		return name
	}
	allLower := true
	for _, r := range name {
		if !unicode.IsLower(r) {
			allLower = false
			break
		}
	}
	if allLower {
		return name
	}

	// Split camelCase at any lower->upper transition, and split on underscores.
	// Check each word for common initialisms.
	runes := []rune(name)
	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word
		if i+1 == len(runes) {
			eow = true
		} else if runes[i+1] == '_' {
			// underscore; shift the remainder forward over any run of underscores
			eow = true
			n := 1
			for i+n+1 < len(runes) && runes[i+n+1] == '_' {
				n++
			}

			// Leave at most one underscore if the underscore is between two digits
			if i+n+1 < len(runes) && unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+n+1]) {
				n--
			}

			copy(runes[i+1:], runes[i+n+1:])
			runes = runes[:len(runes)-n]
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}

		// [w,i) is a word.
		word := string(runes[w:i])
		if u := strings.ToUpper(word); commonInitialisms[u] {
			// Keep consistent case, which is lowercase only at the start.
			if w == 0 && unicode.IsLower(runes[w]) {
				u = strings.ToLower(u)
			}
			// All the common initialisms are ASCII,
			// so we can replace the bytes exactly.
			copy(runes[w:], []rune(u))
		} else if w > 0 && strings.ToLower(word) == word {
			// already all lowercase, and not the first word, so uppercase the first character.
			runes[w] = unicode.ToUpper(runes[w])
		}
		w = i
	}
	return string(runes)
}

var modelTemplate = `
package mapper

import (
    "errors"
    "time"
    "upper.io/db.v3/lib/sqlbuilder"
)

const (
    {{struct_name}}TableNamePrefix = "{{table_name}}"
    {{struct_name}}DBName = "{{db_name}}"
)

{{struct_def}}


func {{struct_name}}GetOne(tableName string, fields []interface{}, where map[string]interface{}) ({{struct_name}}, error) {
    var res {{struct_name}}
    sess := db.Registry[{{struct_name}}DBName]

    q := sess.Select(fields...).From(tableName)

    if len(fields) == 1 {
        if fieldStr, ok := fields[0].(string); ok && fieldStr == "*" {
            q = sess.SelectFrom(tableName)
        }
    }

    for key, val := range where {
        q = q.And(key, val)
    }

    err := q.One(&res)
    if err != nil {
        return {{struct_name}}{}, err
    }
    return res, nil
}

func {{struct_name}}GetList(tableName string, fields []interface{}, where map[string]interface{}, orderBy []string) ([]{{struct_name}}, error) {
    var resList []{{struct_name}}
    sess := db.Registry[{{struct_name}}DBName]
    q := sess.Select(fields...).From(tableName)

    if len(fields) == 1 {
        if fieldStr, ok := fields[0].(string); ok && fieldStr == "*" {
            q = sess.SelectFrom(tableName)
        }
    }

    for key, val := range where {
        q = q.And(key, val)
    }

    for _, val := range orderBy {
        q = q.OrderBy(val)
    }

    err := q.All(&resList)
    if err != nil {
        return nil, err
    }

    return resList, nil
}

func {{struct_name}}Update(tableName string, set map[string]interface{}, where map[string]interface{}) (int, error) {
    sess := db.Registry[{{struct_name}}DBName]
    q := sess.Update(tableName).Set(set).Where(where)
    res, err := q.Exec()
    if err != nil {
        return 0, err
    }

    affected, err := res.RowsAffected()
    if err != nil {
        return 0, err
    }

    return int(affected), nil
}

func {{struct_name}}Create(tableName string, fieldMap map[string]string) (int, error) {

    sess := db.Registry[{{struct_name}}DBName]
    q := sess.InsertInto(tableName).Values(fieldMap)
    res, err := q.Exec()

    if err != nil {
        return -1, err
    }

    insertID, err := res.LastInsertId()

    if err != nil {
        return -1, err
    }

    return int(insertID), nil
}

func {{struct_name}}CreateBatch(tableName string, fieldMapList []map[string]string) (int, error) {

    sess := db.Registry[{{struct_name}}DBName]
    q := sess.InsertInto(tableName)
    for _, fieldMap := range fieldMapList {
        q.Values(fieldMap)
    }

    res, err := q.Exec()

    if err != nil {
        return 0, err
    }

    affected, err := res.RowsAffected()

    if err != nil {
        return 0, err
    }

    return int(affected), nil

}

// 特殊的查询需求
// 例如 select count(1), xx from xxx group by xxx
// 这种情况的话，需要自己处理 map 里的内容，做各种强制转换
func {{struct_name}}RawQuery(sql string) ([]map[string]string, error) {
    sess := db.Registry[{{struct_name}}DBName]
    rows, err := sess.Query(sql)

    if err != nil {
        return nil, err
    }

    var res []map[string]string
    iter := sqlbuilder.NewIterator(rows)
    err = iter.All(&res)
    return res, nil
}

// 特殊的更新需求，可能在业务里有
// update xxx set xxx case when之类的，要自己拼装的
func {{struct_name}}RawUpdate(sql string) (int, error) {
    sess := db.Registry[{{struct_name}}DBName]
    res, err := sess.Exec(sql)

    if err != nil {
        return 0, err
    }

    affected, err := res.RowsAffected()

    if err != nil {
        return 0, err
    }

    return int(affected), err
}

func {{struct_name}}Delete(tableName string, where map[string]interface{}) (int, error) {

    sess := db.Registry[{{struct_name}}DBName]
    q := sess.DeleteFrom(tableName)

    // 禁止不带条件的 where
    if len(where) == 0 {
        return 0, errors.New("no where is not allowed")
    }

    for key, val := range where {
        q = q.And(key, val)
    }

    res, err := q.Exec()
    if err != nil {
        return 0, err
    }

    affected, _ := res.RowsAffected()
    return int(affected), nil
}


`
