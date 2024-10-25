package db

import (
	"balance-api/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // Postgres driver for the database/sql package.
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InstanceType string

var (
	IRead  InstanceType = "read"
	IWrite InstanceType = "write"
)

type Database struct {
	readInstances []DatabaseInstance
	writeInstance DatabaseInstance
}

type DatabaseInstance struct {
	sqlDB *sql.DB
	db    *gorm.DB
}

func InitDB(cfg *config.DBConfig) *Database {
	sqlDB, gormDB, err := initDBInstance(
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.Schema,
		cfg.User,
		cfg.Password,
		cfg.Driver,
		cfg.DBMaxOpenConns,
		cfg.DBMaxIdleConns,
		cfg.DBConnMaxLife,
	)
	if err != nil {
		panic(err)
	}

	sqlReadDB, gormReadDB, err := initDBInstance(
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.Schema,
		cfg.User,
		cfg.Password,
		cfg.Driver,
		cfg.DBMaxOpenConns,
		cfg.DBMaxIdleConns,
		cfg.DBConnMaxLife,
	)
	if err != nil {
		panic(err)
	}
	return &Database{
		readInstances: []DatabaseInstance{
			{
				sqlDB: sqlReadDB,
				db:    gormReadDB,
			},
		},
		writeInstance: DatabaseInstance{
			sqlDB: sqlDB,
			db:    gormDB,
		},
	}
}

// initDBInstance - Initializes *sql.DB and *gorm.DB instances
func initDBInstance(
	host,
	port,
	name,
	schema,
	user,
	password,
	driver string,
	maxConnections,
	maxIdleConns,
	connMaxLifeSeconds int,
) (*sql.DB, *gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s search_path=%s user=%s password=%s sslmode=disable",
		host,
		port,
		name,
		schema,
		user,
		password,
	)

	sqlDB, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, nil, err
	}

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%v/%s?search_path=%v",
		user,
		password,
		host,
		port,
		name,
		schema,
	)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	sqlDB.SetMaxOpenConns(maxConnections)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifeSeconds) * time.Second)

	return sqlDB, gormDB, nil
}

func (db *Database) Exec(instType InstanceType) *gorm.DB {
	switch instType {
	case IRead:
		return db.readInstances[0].db //TODO: add round robin or randomizer
	case IWrite:
		return db.writeInstance.db
	}
	return nil
}

func (db *Database) NewTx(opts ...*sql.TxOptions) *gorm.DB {
	return db.Exec(IWrite).Begin()
}
