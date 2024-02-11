package rocksdb_test

import (
	"fmt"
	"log"
	"math/big"
	"testing"
	"github.com/tranhuyducseven/rocksdb"

	"strconv"
	"time"

	"github.com/linxGnu/grocksdb"
)

func Test100KWritesAnd100kReads(t *testing.T) {
	db, _, err := rocksdb.NewRocksDBDefault()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	ro, wo := grocksdb.NewDefaultReadOptions(), grocksdb.NewDefaultWriteOptions()

	count := int(1e5)

	currentTime := time.Now()
	fmt.Println("Start time:", currentTime)

	for i := 0; i < count; i++ {
		// create a random ethereum address
		address := "0x" + strconv.Itoa(i) + "0000000000"
		addressIntValue := new(big.Int).SetBytes([]byte(address))
		key := addressIntValue.Bytes()
		value := big.NewInt(int64(i)).Bytes()

		if i%10000 == 0 {
			fmt.Println("Inserting data... ", i, " key = ", key, ", value = ", value)
		}

		err := db.Put(wo, key, value)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	currentTime = time.Now()
	fmt.Println("Finish inserting time:", currentTime)

	for i := 0; i < count; i++ {
		// create a random ethereum address
		address := "0x" + strconv.Itoa(i) + "0000000000"
		addressIntValue := new(big.Int).SetBytes([]byte(address))

		key := addressIntValue.Bytes()
		value, err := db.Get(ro, addressIntValue.Bytes())
		if i%10000 == 0 {
			fmt.Println("Getting data... ", i, " => key = ", key, ", value = ", value, " => ", new(big.Int).SetBytes(value.Data()).String())
		}
		if err != nil {
			log.Fatal(err)
			return
		}
		defer value.Free()
	}
	currentTime = time.Now()
	fmt.Println("Finish querying time:", currentTime)
}
