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
func TestLockObject_TryLock(t *testing.T) {
	// new lock
	lock := NewLock("root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=UTC",
		"distributed_locks", "lock001", 10)

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
```