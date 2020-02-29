package mysql_distributed_locks

import (
	"testing"
	"time"
)

func TestLockObject_TryLock(t *testing.T) {
	// new lock
	lock := NewLock("root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=UTC",
		"", "lock001", 10)

	// trylock
	err := lock.TryLock()
	if err != nil  {
		t.Error("TryLock error:", err.Error())
	}

	// unlock
	err = lock.UnLock()
	if err != nil  {
		t.Error("UnLock error:", err.Error())
	}
}

func TestLockObject_UnLock(t *testing.T) {
	// new lock
	lock := NewLock("root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=UTC",
		"", "lock002", 1)

	// try lock
	err := lock.TryLock()
	if err != nil  {
		t.Error("TryLock error:", err.Error())
	}

	time.Sleep(time.Second * 5)

	err = lock.UnLock()
	if err != nil  {
		t.Error("UnLock error:", err.Error())
	}
}