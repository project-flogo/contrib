<!--
title: SQL Query
weight: 4616
-->

# SQLQuery Database Activity 
This activity provides your flogo application execute database queries. 


## Installation

```bash
flogo install github.com/project-flogo/activity/sqlquery
```

## Configuration

### Settings:
| Name               | Type   | Description
|:---                | :---   | :---    
| dbType             | string | The type of database (mysql, oracle, postgres, sqlite, sqlserver) - **REQUIRED**         
| driverName         | string | The database driver name - **REQUIRED**
| dataSourceName     | string | The database DataSource name - **REQUIRED**
| maxOpenConnections | int    | Max open connections (default is unlimited)
| maxIdleConnections | int    | Max idle connections (default is 2)
| query              | string | The SQL select query - **REQUIRED**
| disablePrepared    | bool   | Disable prepared statement usage
| labeledResults     | bool   | Return results labeled by column name

### Input:
| Name   | Type | Description
|:---    | :--- | :---    
| params | map  |  The query parameters

### Output:
| Name        | Type  | Description
|:---         | :---  | :---    
| columnNames | array |  The names of the result columns
| results     | array |  The results

## Examples

### Query
Simple query that gets all items with ID less than 10, retrieves all the columns.  In order to use *mysql*, you have to import the driver by adding `github.com/go-sql-driver/mysql` to 
the app imports section.  See [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) for more information on the driver.
```json
{
  "id": "dbquery",
  "name": "DbQuery",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/sqlquery",
    "settings": {
      "dbType": "mysql",
      "driverName": "mysql",
      "dataSourceName": "username:password@tcp(host:port)/dbName",
      "query": "select * from test where ID < 10"
    }
  }
}
```

### Named Query
Query with parameters.  Parameters are referenced using ':', e.g. `:id`, regardless of database
```json
{
  "id": "named_dbquery",
  "name": "Named DbQuery",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/sqlquery",
    "settings": {
      "dbType": "mysql",
      "driverName": "mysql",
      "dataSourceName": "username:password@tcp(host:port)/dbName",
      "query": "select * from test where ID < :id"
    },
    "input": {
      "params": {
        "id":10
      }
    }
  }
}
```
### Supported Drivers

- MySQL: [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- Oracle: [github.com/mattn/go-oci8](https://github.com/mattn/go-oci8)
- Postgres: [github.com/lib/pq](https://github.com/lib/pq) 
- SQLite: [github.com/mattn/go-sqlite3]( https://github.com/mattn/go-sqlite3)
- SQLServer: [github.com/denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb)

