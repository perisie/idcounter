package idcounter

type Id_counter interface {
	Get(key string) (int, error)
	Add(key string, amount int) (int, error)
}

const (
	KEY_START_VALUE = 0
)
