package rocksdb

import (
	"fmt"
	"os"

	"github.com/linxGnu/grocksdb"
)

type Config struct {
	Address  string
	Password string
	Database int
	PoolSize int
}

const DEFAULT_ROCKSDB_PATH = "tmp/rocksdb"

func getDefaultRocksDBOptions() *grocksdb.Options {
	bbto := grocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(grocksdb.NewLRUCache(3 << 30))

	opts := grocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)

	return opts
}

func getPath() (string, error) {
	dbPath := os.Getenv("ROCKSDB_PATH")
	if dbPath == "" {
		dbPath, _ = os.Getwd()
		err := os.MkdirAll(dbPath+"/"+DEFAULT_ROCKSDB_PATH, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("failed to create rocksdb path: %v", err)
		}
		dbPath = DEFAULT_ROCKSDB_PATH
	}

	return dbPath, nil
}

func NewRocksTransactionDB() (*grocksdb.TransactionDB, string, error) {
	opts := getDefaultRocksDBOptions()
	dbPath, err := getPath()
	if err != nil {
		return nil, "", err
	}

	transactionDBOpts := grocksdb.NewDefaultTransactionDBOptions()
	db, err := grocksdb.OpenTransactionDb(opts, transactionDBOpts, dbPath)
	return db, dbPath, err

}

func NewRocksDBDefault() (*grocksdb.DB, string, error) {
	opts := getDefaultRocksDBOptions()
	dbPath, err := getPath()
	if err != nil {
		return nil, "", err
	}

	db, err := grocksdb.OpenDb(opts, dbPath)

	return db, dbPath, err
}

func NewRocksDB(opts *grocksdb.Options, dbPath string) (*grocksdb.DB, error) {
	db, err := grocksdb.OpenDb(opts, dbPath)
	return db, err
}
