package mysql

import (
	"github.com/prometheus/common/log"
)

func InitData() error {
	locker := NewLockDb("init", GetHostname(), DefaultLeaseAge)
	ok, err := locker.Lock()
	if err != nil {
		return err
	}

	if !ok {
		return nil
	}

	defer locker.UnLock()

	return run()
}

func run() error {
	log.Infof("%s begin init data", GetHostname())
	return nil
}
