package mysql_distributed_locks

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
	"time"
)

// mysql table struct
type distributedLock struct {
	Name        string `gorm:"primary_key"`
	Owner       string
	CreatedTime int64
	ExpireTime  int64
}

// mysql distributed lock object
type lockObject struct {
	name          string
	timeout       int // in seconds
	owner         string
	databaseUrl   string
	databaseTable string
	createdTime   int64
	expireTime    int64
	db            *gorm.DB
}

func NewLock(databaseUrl string, databaseTable string, lockName string, timeout int) *lockObject {
	return &lockObject{
		databaseUrl:   databaseUrl,
		databaseTable: databaseTable,
		name:          lockName,
		timeout:       timeout,
		owner:         uuid.NewV1().String(),
		db:            nil,
	}
}

func (lock *lockObject) TryLock() error {

	var err error

	// open database
	if lock.db == nil {
		lock.db, err = gorm.Open("mysql", lock.databaseUrl)
		if err == nil && lock.db != nil {
			lock.db.DB().SetMaxIdleConns(10)
			lock.db.DB().SetMaxOpenConns(100)
		}

		defer func() {
			err := lock.db.Close()
			if err != nil {
				fmt.Println("TryLock close db error:", err.Error())
			}
			lock.db = nil
		}()
	}

	// clean timeout lock
	lock.deleleExpiredLock()

	// prepare data
	var now = time.Now().Unix()
	lock.createdTime = now
	lock.expireTime = now + int64(lock.timeout)

	var newLock = distributedLock{
		Name:        lock.name,
		Owner:       lock.owner,
		CreatedTime: lock.createdTime,
		ExpireTime:  lock.expireTime,
	}

	// insert lock
	return lock.db.Table(lock.databaseTable).Create(&newLock).Error
}

func (lock *lockObject) UnLock() error {
	var err error

	// open database
	if lock.db == nil {
		lock.db, err = gorm.Open("mysql", lock.databaseUrl)
		if err == nil && lock.db != nil {
			lock.db.DB().SetMaxIdleConns(10)
			lock.db.DB().SetMaxOpenConns(100)
		}

		defer func() {
			err := lock.db.Close()
			if err != nil {
				fmt.Println("error:", err.Error())
			}
			lock.db = nil
		}()
	}

	// delete lock
	return  lock.db.Table(lock.databaseTable).Where("name = ? AND owner = ?", lock.name, lock.owner).Delete(distributedLock{}).Error
}

func (lock *lockObject) deleleExpiredLock() {

	if lock.db != nil {
		var now = time.Now().Unix()
		err := lock.db.Table(lock.databaseTable).Where("name = ? AND expire_time < ?", lock.name, now).Delete(distributedLock{}).Error
		if err != nil {
			fmt.Println("deleleExpiredLock error:", err.Error())
		}
	}

}