package mysql

import (
	"context"
	"fmt"
	"github.com/oklog/ulid"
	"golang-gin-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"reflect"
	"time"
)

var (
	injectors    []func(db *gorm.DB)
	globalDB     *gorm.DB
	globalConfig config.Config
)

type ctxTransactionKey struct{}

// Initialize mysql connection
func Connect() {
	cfg := config.GetConfig().MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Passwd, cfg.Host, cfg.Database)
	fmt.Printf(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("mysql connection get error: %v", err)
	}
	registerCallback(db)
	callInjector(db)
	globalDB = db
}

// https://github.com/ulid/spec
// uuid sortable by time
func NewUlid() string {
	now := time.Now()
	return ulid.MustNew(ulid.Timestamp(now), ulid.Monotonic(rand.New(rand.NewSource(now.UnixNano())), 0)).String()
}

func registerCallback(db *gorm.DB) {
	// Auto load uuid
	err := db.Callback().Create().Before("gorm:create").Register("uuid", func(db *gorm.DB) {
		db.Statement.SetColumn("id", NewUlid())
	})
	if err != nil {
		fmt.Printf("regiser call back: %v", err)
	}
}

func RegisterInjector(f func(*gorm.DB)) {
	injectors = append(injectors, f)
}

// 如果使用跨模型事务则传参
func GetDB(ctx context.Context) *gorm.DB {
	iface := ctx.Value(ctxTransactionKey{})

	if iface != nil {
		tx, ok := iface.(*gorm.DB)
		if !ok {
			log.Panicf("unexpect context value type: %s", reflect.TypeOf(tx))
			return nil
		}

		return tx
	}

	return globalDB.WithContext(ctx)
}

func callInjector(db *gorm.DB) {
	for _, v := range injectors {
		v(db)
	}
}

func SetupTableModel(db *gorm.DB, model interface{}) {
	if globalConfig.MySQL.AutoMigrate {
		err := db.AutoMigrate(model)
		if err != nil {
			fmt.Printf("Auto migrate table get error: %v", err)
		}
	}
}

func WithOffsetLimit(db *gorm.DB, offset, limit int) *gorm.DB {
	if offset > 0 {
		db = db.Limit(offset)
	}

	if limit > 0 {
		db = db.Limit(limit)
	}

	return db
}
