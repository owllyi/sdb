package service

import (
	"bytes"
	"fmt"
	"github.com/yemingfeng/sdb/internal/store"
	"github.com/yemingfeng/sdb/internal/store/engine"
	"github.com/yemingfeng/sdb/internal/util"
)

const listKeyPrefixTemplate = "l/%s"
const listKeyTemplate = listKeyPrefixTemplate + "/%d"

func LPush(key []byte, values [][]byte, sync bool) (bool, error) {
	lock(LList, key)
	defer unlock(LList, key)

	batch := store.NewBatch()
	defer batch.Close()

	for _, value := range values {
		batch.Set(generateListKey(key, util.GetOrderingKey()), value)
	}

	return batch.Commit(sync)
}

func LPop(key []byte, values [][]byte, sync bool) (bool, error) {
	lock(LList, key)
	defer unlock(LList, key)

	batch := store.NewBatch()
	defer batch.Close()

	store.Iterate(&engine.PrefixIteratorOption{Prefix: generateListPrefixKey(key)},
		func(key []byte, value []byte) {
			for i := range values {
				if bytes.Equal(values[i], value) {
					batch.Del(key)
				}
			}
		})

	return batch.Commit(sync)
}

func LRange(key []byte, offset int32, limit int32) ([][]byte, error) {
	index := int32(0)
	res := make([][]byte, limit)
	store.Iterate(&engine.PrefixIteratorOption{
		Prefix: generateListPrefixKey(key), Offset: int(offset), Limit: int(limit)},
		func(key []byte, value []byte) {
			res[index] = value
			index++
		})
	return res[0:index], nil
}

func LExist(key []byte, values [][]byte) ([]bool, error) {
	res := make([]bool, len(values))
	existMap := make(map[string]bool)
	store.Iterate(&engine.PrefixIteratorOption{Prefix: generateListPrefixKey(key)},
		func(key []byte, value []byte) {
			existMap[string(value)] = true
		})
	for i, value := range values {
		if existMap[string(value)] {
			res[i] = true
		}
	}
	return res, nil
}

func LDel(key []byte, sync bool) (bool, error) {
	lock(LList, key)
	defer unlock(LList, key)

	batch := store.NewBatch()
	defer batch.Close()

	store.Iterate(&engine.PrefixIteratorOption{Prefix: generateListPrefixKey(key)},
		func(key []byte, value []byte) {
			batch.Del(key)
		})

	return batch.Commit(sync)
}

func LCount(key []byte) (int32, error) {
	count := int32(0)
	store.Iterate(&engine.PrefixIteratorOption{Prefix: generateListPrefixKey(key)},
		func(key []byte, value []byte) {
			count++
		})
	return count, nil
}

func generateListKey(key []byte, orderingKey int64) []byte {
	return []byte(fmt.Sprintf(listKeyTemplate, key, orderingKey))
}

func generateListPrefixKey(key []byte) []byte {
	return []byte(fmt.Sprintf(listKeyPrefixTemplate, key))
}
