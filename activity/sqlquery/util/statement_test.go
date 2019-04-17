package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMySql(t *testing.T) {
	dbType := DbMySql

	sql := "select * from table where t = :foo and s = :bar"
	s, err := NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where t = ? s = ?", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar"
	s, err  = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = ? s = ?", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar and q = \"tar\""
	s, err  = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = ? s = ? and q = \"tar\"", s.PreparedStatementSQL())
}

func TestParseOracle(t *testing.T) {
	dbType := DbOracle

	sql := "select * from table where t = :foo and s = :bar"
	s, err := NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where t = :foo and s = :bar", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar"
	s, err  = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar and q = \"tar\""
	s, err  = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar and q = \"tar\"", s.PreparedStatementSQL())
}

func TestParsePostgres(t *testing.T) {
	dbType := DbPostgres

	sql := "select * from table where t = :foo and s = :bar"
	s, err := NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where t = ? s = ?", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar"
	s, err  = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = ? s = ?", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar and q = \"tar\""
	s, err  = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = ? s = ? and q = \"tar\"", s.PreparedStatementSQL())
}

func TestParseSqlLite(t *testing.T) {
	dbType := DbSQLite

	sql := "select * from table where t = :foo and s = :bar"
	s, err := NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where t = ? s = ?", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar"
	s, err  = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = ? s = ?", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar and q = \"tar\""
	s, err  = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = ? s = ? and q = \"tar\"", s.PreparedStatementSQL())
}

func TestParseSqlServer(t *testing.T) {
	dbType := DbSqlServer

	sql := "select * from table where t = :foo and s = :bar"
	s, err := NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where t = @foo and s = @bar", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar"
	s, err  = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = @foo and s = @bar", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar and q = \"tar\""
	s, err  = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = @foo and s = @bar and q = \"tar\"", s.PreparedStatementSQL())
}

func TestFlattenSql(t *testing.T) {

	sql := "select * from table where t = :foo and s = :bar and r = :other"
	params := map[string]interface{} {"foo":true, "bar":2, "other":"test"}

	s, err := NewSQLStatement(DbMySql, sql)
	assert.Nil(t, err)
	flattenedSQL := s.ToStatementSQL(params)
	assert.Equal(t, "select * from table where t = true and s = 2 and r = 'test'", flattenedSQL)

	s, err = NewSQLStatement(DbOracle, sql)
	assert.Nil(t, err)
	flattenedSQL = s.ToStatementSQL(params)
	assert.Equal(t, "select * from table where t = 1 and s = 2 and r = 'test'", flattenedSQL)

	s, err = NewSQLStatement(DbPostgres, sql)
	assert.Nil(t, err)
	flattenedSQL = s.ToStatementSQL(params)
	assert.Equal(t, "select * from table where t = TRUE and s = 2 and r = 'test'", flattenedSQL)

	s, err = NewSQLStatement(DbSQLite, sql)
	assert.Nil(t, err)
	flattenedSQL = s.ToStatementSQL(params)
	assert.Equal(t, "select * from table where t = 1 and s = 2 and r = 'test'", flattenedSQL)

	s, err = NewSQLStatement(DbSqlServer, sql)
	assert.Nil(t, err)
	flattenedSQL = s.ToStatementSQL(params)
	assert.Equal(t, "select * from table where t = TRUE and s = 2 and r = 'test'", flattenedSQL)
}
