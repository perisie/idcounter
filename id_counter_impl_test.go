package idcounter

import (
	"errors"
	"testing"

	"github.com/perisie/kvstore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_id_counter(t *testing.T) {
	id_counter := New(kvstore.Kv_store_fake_new())
	my_key := "my_key"

	id, _ := id_counter.Get(my_key)

	assert.Equal(t, 0, id)

	_, _ = id_counter.Add(my_key, 1)
	id, _ = id_counter.Add(my_key, 2)

	assert.Equal(t, 3, id)

	id, _ = id_counter.Add(my_key, -1337)

	assert.Equal(t, 0, id)
}

func Test_id_counter_error_store(t *testing.T) {
	kv_store := &kvstore.Kv_store_mock{}
	id_counter := New(kv_store)

	kv_store.On("Get", "").Return((*kvstore.Key_value)(nil), errors.New("error"))

	_, err := id_counter.Get("")
	assert.NotNil(t, err)
}

func Test_id_counter_error_kv_not_exist(t *testing.T) {
	kv_store := &kvstore.Kv_store_mock{}
	id_counter := New(kv_store)
	my_key := "my_key"

	kv_store.On("Get", my_key).Return(&kvstore.Key_value{}, nil)
	kv_store.On("Create", my_key, mock.AnythingOfType("string")).Return((*kvstore.Key_value)(nil), errors.New("error"))

	_, err := id_counter.Add(my_key, 1)
	assert.NotNil(t, err)
}

func Test_id_counter_error_set(t *testing.T) {
	kv_store := &kvstore.Kv_store_mock{}
	id_counter := New(kv_store)
	my_key := "my_key"

	kv_store.On("Create", my_key, mock.AnythingOfType("string")).Return((*kvstore.Key_value)(nil), errors.New("error"))

	_, err := id_counter.set(my_key, 1)
	assert.NotNil(t, err)
}
