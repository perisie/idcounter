package idcounter

import "github.com/stretchr/testify/mock"

type Id_counter_mock struct {
	mock.Mock
}

func (i *Id_counter_mock) Get(key string) (int, error) {
	args := i.Called(key)
	return args.Int(0), args.Error(1)
}

func (i *Id_counter_mock) Add(key string, amount int) (int, error) {
	panic("implement me")
}
