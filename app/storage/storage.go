// Experimental
// try simple storage package
// the type have: KV, Table

package storage

import (
	"fmt"
	"gitee.com/conero/uymas/v2"
)

// Kv the Kv style data
type Kv map[any]any

// Table the list of table
type Table []any

const (
	LiteralInt    = "int"   // golang type: int
	LiteralFloat  = "float" // golang type: float64
	LiteralNumber = "number"
	LiteralBool   = "bool"   // true/True; false/False
	LiteralString = "string" //'string' or "string" or string
	LiteralNull   = "null"
)

var (
	memoryStorageCache *Storage
)

// Literal the Literal variable. this is value from string
type Literal string

// Storage the storage of cache
type Storage struct {
	namespace string
	data      Kv
}

func (store *Storage) GetValue(key any) any {
	value, has := store.data[key]
	if has {
		return value
	}
	return nil
}

func (store *Storage) SetValue(key, value any) *Storage {
	store.data[key] = value
	return store
}

func (store *Storage) hasKey(key any) bool {
	_, has := store.data[key]
	return has
}

func (store *Storage) DelKey(key any) bool {
	if store.hasKey(key) {
		delete(store.data, key)
		return true
	}
	return false
}

func NewStorage(namespace string) *Storage {
	store := &Storage{
		namespace: namespace,
		data:      Kv{},
	}
	memoryStorageCache.SetValue(namespace, store)
	return store
}

func GetStorage(namespace string) *Storage {
	value := memoryStorageCache.GetValue(namespace)
	if value != nil {
		if store, isStore := value.(*Storage); isStore {
			return store
		}
	}
	return nil
}

func init() {
	namespace := fmt.Sprintf("%v_internal_sys_memory_", uymas.Name)

	store := &Storage{
		namespace: namespace,
		data:      Kv{},
	}
	memoryStorageCache = store
}
