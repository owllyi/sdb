package service

import (
	"fmt"
	"github.com/yemingfeng/sdb/internal/store"
	"strconv"
)

const stringKeyPrefixTemplate = "s/%s"

func Set(key []byte, value []byte, sync bool) (bool, error) {
	return store.Set(generateStringKey(key), value, sync)
}

func Get(key []byte) ([]byte, error) {
	return store.Get(generateStringKey(key))
}

func Del(key []byte, sync bool) (bool, error) {
	return store.Del(generateStringKey(key), sync)
}

func Incr(key []byte, delta int32, sync bool) (bool, error) {
	lock(LString, key)
	defer unlock(LString, key)

	value, err := store.Get(key)
	if err != nil {
		return false, err
	}
	var valueInt = 0
	if len(value) > 0 {
		valueInt, err = strconv.Atoi(string(value))
		if err != nil {
			return false, err
		}
	}
	valueInt = valueInt + int(delta)

	return store.Set(key, []byte(strconv.Itoa(valueInt)), sync)
}

func generateStringKey(key []byte) []byte {
	return []byte(fmt.Sprintf(stringKeyPrefixTemplate, key))
}
