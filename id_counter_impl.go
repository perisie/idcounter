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
		return 0, err
	}
	if !kv.Exist() {
		kv, err = i.kv_store.Create(key, fmt.Sprint(KEY_START_VALUE))
		if err != nil {
			return 0, err
		}
	}
	return strconv.Atoi(kv.Value)
}

func (i *Id_counter_impl) Add(key string, amount int) (int, error) {
	id, err := i.Get(key)
	if err != nil {
		return 0, err
	}
	id_new := id + amount
	if id_new < 0 {
		id_new = 0
	}
	return i.set(key, id_new)
}

func (i *Id_counter_impl) set(key string, value int) (int, error) {
	value_str := strconv.Itoa(value)
	_, err := i.kv_store.Create(key, value_str)
	if err != nil {
		return 0, err
	}
	return value, nil
}
