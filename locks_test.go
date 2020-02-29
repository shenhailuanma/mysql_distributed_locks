package mysql_distributed_locks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var databaseUrl = "root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=UTC"
var errorDababaseUrl = "root@tcp(localhost:3306)/nodababase?charset=utf8&parseTime=True&loc=UTC"
var databaseTable = "distributed_locks_test"
var lockname = "testLock001"

func TestLockObject_TryLock(t *testing.T) {
	// new lock
	lock := NewLock(databaseUrl, databaseTable, lockname, 10)

	// trylock
	err := lock.TryLock()
	if err != nil {
		t.Error("TryLock error:", err.Error())
	}

	// unlock
	err = lock.UnLock()
	if err != nil {
		t.Error("UnLock error:", err.Error())
	}
}

func TestLockObject_DatabaseFailed(t *testing.T) {
	// new lock
	lock := NewLock(errorDababaseUrl, databaseTable, lockname, 10)

	// trylock
	err := lock.TryLock()
	assert.NotNil(t, err)

	lock.db = nil // reset for testing unlock open database failed
	err = lock.UnLock()
	assert.NotNil(t, err)
}
