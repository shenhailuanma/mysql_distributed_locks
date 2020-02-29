package main

import (
	"fmt"
	"github.com/shenhailuanma/mysql_distributed_locks"
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
	lock := mysql_distributed_locks.NewLock(databaseUrl, databaseTable, lockname, 10)

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
