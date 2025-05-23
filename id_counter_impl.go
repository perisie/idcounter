package idcounter

import (
	"fmt"
	"strconv"

	"github.com/perisie/kvstore"
)

type Id_counter_impl struct {
	kv_store kvstore.Kv_store
}

func New(kv_store kvstore.Kv_store) *Id_counter_impl {
	return &Id_counter_impl{
		kv_store: kv_store,
	}
}

func (i *Id_counter_impl) Get(key string) (int, error) {
	kv, err := i.kv_store.Get(key)
	if err != nil {
		return KEY_START_VALUE, err
	}
	if !kv.Exist() {
		kv, err = i.kv_store.Create(key, fmt.Sprint(KEY_START_VALUE))
		if err != nil {
			return KEY_START_VALUE, err
		}
	}
	return strconv.Atoi(kv.Value)
}

func (i *Id_counter_impl) Add(key string, amount int) (int, error) {
	if amount < 1 {
		return 0, fmt.Errorf("amount must be greater than 0")
	}
	id, err := i.Get(key)
	if err != nil {
		id, err = i.set(key, KEY_START_VALUE)
		if err != nil {
			return 0, err
		}
	}
	id_new := id + amount
	return i.set(key, id_new)
}

func (i *Id_counter_impl) set(key string, value int) (int, error) {
	value_str := strconv.Itoa(value)
	_, err := i.kv_store.Create(key, value_str)
	if err != nil {
		return KEY_START_VALUE, err
	}
	return value, nil
}
