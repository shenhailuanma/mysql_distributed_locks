# mysql distributed locks

## Installation
```bash
go get github.com/shenhailuanma/mysql_distributed_locks
```

## Requirements

The mysql_distributed_locks designed that have no permission to create mysql table automatically.

Ensure there already prepared mysql table like this:
```sql
CREATE TABLE IF NOT EXISTS `distributed_locks`
(
    `name`         varchar(128) NOT NULL,
    `owner`        varchar(256) NOT NULL,
    `created_time` bigint       NOT NULL,
    `expire_time`  bigint       NOT NULL,
    PRIMARY KEY (name) USING HASH
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;
```

## Example

```go
package main

import (
	"fmt"
	mdlocks "github.com/shenhailuanma/mysql_distributed_locks"
)

func main()  {

	var databaseUrl = "root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=UTC"
	var databaseTable = "distributed_locks"
	var lockname = "testLock001"

	/**
	  New a lock object.

	  params:
	      databaseUrl: database url, format: "username:password@protocol(address)/dbname?param=value",
	                   details: https://github.com/go-sql-driver/mysql#dsn-data-source-name
	      databaseTable: table name in database
	      lockName: lock name
	      timeout: lock timeout in seconds

	  return: lock object
	*/
	lock := mdlocks.NewLock(databaseUrl, databaseTable, lockname, 10)

	// trylock
	err := lock.TryLock()
	if err != nil {
		fmt.Errorf("TryLock error: %s", err.Error())
	}

	// unlock
	err = lock.UnLock()
	if err != nil {
		fmt.Errorf("UnLock error: %s", err.Error())
	}
}


```