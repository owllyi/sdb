package level

import (
	"github.com/yemingfeng/sdb/internal/conf"
	"github.com/yemingfeng/sdb/internal/store/engine"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
)

type LevelStore struct {
	db *leveldb.DB
}

func NewLevelStore() *LevelStore {
	dbPath := conf.Conf.Store.Path + "/level"
	db, err := leveldb.OpenFile(dbPath, &opt.Options{Filter: filter.NewBloomFilter(10)})
	if err != nil {
		log.Fatalf("failed to open file: %+v", err)
	}
	log.Printf("db init %s complete", dbPath)

	return &LevelStore{db: db}
}

func (store *LevelStore) Set(key []byte, value []byte, sync bool) (bool, error) {
	err := store.db.Put(key, value, &opt.WriteOptions{Sync: sync})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (store *LevelStore) Get(key []byte) ([]byte, error) {
	value, err := store.db.Get(key, nil)
	if err == leveldb.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return value, err
}

func (store *LevelStore) Del(key []byte, sync bool) (bool, error) {
	if err := store.db.Delete(key, &opt.WriteOptions{Sync: sync}); err != nil {
		return false, nil
	}
	return true, nil
}

func (store *LevelStore) NewBatch() engine.Batch {
	return &LevelBatch{db: store.db, batch: new(leveldb.Batch)}
}

func (store *LevelStore) Iterator(opt *engine.PrefixIteratorOption, handle func([]byte, []byte)) {
	it := store.db.NewIterator(util.BytesPrefix(opt.Prefix), nil)
	defer func() {
		it.Release()
	}()

	if opt.Offset >= 0 {
		i := 0
		for it.First(); i < opt.Offset && it.Valid(); it.Next() {
			i++
		}

		i = 0
		for ; it.Valid(); it.Next() {
			handle(it.Key(), it.Value())
			i++
			if opt.Limit > 0 && i == opt.Limit {
				break
			}
		}
	} else {
		i := 0
		for it.Last(); i < -opt.Offset-1 && it.Valid(); it.Prev() {
			i++
		}

		i = 0
		for ; it.Valid(); it.Prev() {
			handle(it.Key(), it.Value())
			i++
			if opt.Limit > 0 && i == opt.Limit {
				break
			}
		}
	}
}

func (store *LevelStore) Close() error {
	return store.db.Close()
}
