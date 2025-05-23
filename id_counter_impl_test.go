package idcounter

import (
	"errors"
	"os"
	"testing"

	"github.com/perisie/kvstore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_id_counter(t *testing.T) {
	id_counter := test_before(t)
	my_key := "my_key"

	_, err := id_counter.Get(my_key)
	assert.NotNil(t, err)

	id, err := id_counter.Add(my_key, 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, id)

	id, err = id_counter.Add(my_key, 2)
	assert.Nil(t, err)
	assert.Equal(t, 3, id)

	id, err = id_counter.Add(my_key, -1337)
	assert.NotNil(t, err)
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

func test_before(t *testing.T) Id_counter {
	err := os.RemoveAll("data")
	assert.Nil(t, err)

	kv_store := kvstore.Kv_store_mouse_new("data")
	id_counter := New(kv_store)
	return id_counter
}
