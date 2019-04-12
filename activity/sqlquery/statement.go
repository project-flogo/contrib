package sqlquery

import (
	"database/sql"
	"fmt"
	"github.com/project-flogo/core/data/coerce"
	"strings"
)

type DbType int8
type BindType int8
type StmtType int8

const (
	BtUnknown BindType = iota
	BtAt
	BtColon
	BtDollar
	BtQuestion

	DbUnknown DbType = iota
	DbMySql
	DbOracle
	DbPostgres
	DbSQLite
	DbSqlServer

	StUnknown StmtType = iota
	StSelect
	StInsert
	StUpdate
	StDelete
)

// ToTypeEnum get the data type that corresponds to the specified name
func ToDbType(typeStr string) (DbType, error) {

	switch strings.ToLower(typeStr) {
	case "mssql":
		return DbMySql, nil
	case "oracle":
		return DbOracle, nil
	case "postgres":
		return DbPostgres, nil
	case "sqlite":
		return DbSQLite, nil
	case "sqlserver":
		return DbSqlServer, nil
	default:
		return DbUnknown, fmt.Errorf("unknown type: %s", typeStr)
	}
}

// ToTypeEnum get the data type that corresponds to the specified name
func ToStmtType(typeStr string) (StmtType, error) {

	switch strings.ToLower(typeStr) {
	case "select":
		return StSelect, nil
	case "insert":
		return StInsert, nil
	case "update":
		return StUpdate, nil
	case "delete":
		return StDelete, nil
	default:
		return StUnknown, fmt.Errorf("unknown statement type: %s", typeStr)
	}
}

func GetBindType(dbType DbType) BindType {
	switch dbType {
	case DbPostgres:
		return BtDollar
	case DbMySql, DbSQLite:
		return BtQuestion
	case DbOracle:
		return BtColon
	case DbSqlServer:
		return BtAt
	default:
		return BtUnknown
	}
}

func NewSQLStatement(dbType DbType, sql string) (*SQLStatement, error) {

	sql = strings.TrimSpace(sql)
	sqlParts := strings.Fields(sql)

	if len(sqlParts) == 0 {
		return nil, fmt.Errorf("invalid sql '%s'", sql)
	}

	stmtType, err := ToStmtType(sqlParts[0])
	if err != nil {
		return nil, err
	}

	bt := GetBindType(dbType)
	parts := parse(sql,dbType, bt)
	numMap := make(map[string]int)

	//if "dollar" placeholder, calculate positions
	if bt == BtDollar {
		i := 0

		for _, part := range parts {
			if pPart, ok := part.(*paramPart); ok {
				if _, ok := numMap[pPart.param]; !ok {
					i = i + 1
					numMap[pPart.param] = i
				}
			}
		}
	}

	n := 0
	for i := 0; i < len(parts); i++ {
		n += len(parts[i].String())
	}

	var b strings.Builder
	b.Grow(n)
	for _, part := range parts {
		b.WriteString(part.Placeholder())
	}
	preparedSQL := b.String()

	return &SQLStatement{stmtType: stmtType, parts: parts, bindType: bt, preparedSQL: preparedSQL}, nil
}

// SQLStatement is a parsed DML SQL Statement
type SQLStatement struct {
	dbType         DbType
	bindType       BindType
	stmtType       StmtType
	parts          []Part
	placeholderIds map[string]int
	preparedSQL    string
}

func (s *SQLStatement) Type() StmtType {
	return s.stmtType
}

func (s *SQLStatement) HasParams() bool {
	return len(s.parts) > 1
}

func (s *SQLStatement) String() string {

	n := 0
	for i := 0; i < len(s.parts); i++ {
		n += len(s.parts[i].String())
	}

	var b strings.Builder
	b.Grow(n)
	for _, s := range s.parts {
		b.WriteString(s.String())
	}
	return b.String()
}

func (s *SQLStatement) ToStatementSQL(params map[string]interface{}) string {
	n := 0
	for i := 0; i < len(s.parts); i++ {
		n += len(s.parts[i].String())
	}

	var b strings.Builder
	b.Grow(n)
	for _, s := range s.parts {
		b.WriteString(s.ToValue(params))
	}
	return b.String()
}

func (s *SQLStatement) PreparedStatementSQL() string {
	return s.preparedSQL
}

func (s *SQLStatement) GetPreparedStatementArgs(params map[string]interface{}) []interface{} {

	var sParams []interface{}

	switch s.bindType {
	case BtAt, BtColon:
		//named
		for name, value := range params {
			sParams = append(sParams, sql.Named(name, value))
		}
	case BtQuestion:
		for _, part := range s.parts {
			if pPart, ok := part.(*paramPart); ok {

				if v, ok := params[pPart.param]; !ok {
					sParams = append(sParams, v)
				}
			}
		}
	case BtDollar:
		sParams = make([]interface{}, 3)
		for name, id := range s.placeholderIds {
			sParams[id-1] = params[name]
		}
	}

	return sParams
}

func parse(sqlStatement string, dbType DbType, bindType BindType) []Part {
	var i, j, start int

	var parts []Part

	for i = 0; i < len(sqlStatement); i++ {

		if sqlStatement[i] == '"' {
			for j = i + 1; j < len(sqlStatement); j++ {
				if sqlStatement[j] == '"' {
					break
				}
			}
			i = j
		} else if sqlStatement[i] == '\'' {
			for j = i + 1; j < len(sqlStatement); j++ {
				if sqlStatement[j] == '\'' {
					break
				}
			}
			i = j
		} else if sqlStatement[i] == ':' {
			parts = append(parts, &literalPart{literal: sqlStatement[start:i]})
			start = i
			for j = i; j < len(sqlStatement); j++ {
				if sqlStatement[j] == ' ' {
					break
				}
			}
			parts = append(parts, newParamPart(sqlStatement[start+1:j], dbType, bindType))
			i = j
			start = j
		}
	}

	if start < len(sqlStatement) {
		parts = append(parts, &literalPart{literal: sqlStatement[start:]})
	}

	return parts
}

type Part interface {
	ToValue(params map[string]interface{}) string
	Placeholder() string
	String() string
}

type literalPart struct {
	literal string
}

func (p *literalPart) ToValue(params map[string]interface{}) string {
	return p.literal
}

func (p *literalPart) Placeholder() string {
	return p.literal
}

func (p *literalPart) String() string {
	return p.literal
}

func newParamPart(param string, dbType DbType, bindType BindType) Part {

	part := &paramPart{param: param, dbType:dbType}
	switch bindType {
	case BtAt:
		part.placeholder = "@" + param
	case BtColon:
		part.placeholder = ":" + param
	default:
		part.placeholder = "?"
	}

	return part
}

type paramPart struct {
	dbType      DbType
	param       string
	placeholder string
}

func (p *paramPart) ToValue(params map[string]interface{}) string {

	v := params[p.param]
	return ToSql(p.dbType, v)
}

func (p *paramPart) Placeholder() string {
	return p.placeholder
}

func (p *paramPart) String() string {
	return ":" + p.param
}

func ToSql(dbType DbType, val interface{}) string {

	switch t := val.(type) {
	case bool:
		switch dbType {
		case DbSQLite, DbOracle:
			if t {
				return "1"
			} else {
				return "0"
			}
		case DbPostgres, DbSqlServer:
			if t {
				return "TRUE"
			} else {
				return "FALSE"
			}
		default:
			if t {
				return "true"
			} else {
				return "false"
			}
		}
	case int, int32, int64, float32, float64:
		s,_ := coerce.ToString(val)
		return s
	default:
		s, _ := coerce.ToString(val)
		return "'" + s + "'"
	}
}
