package lazy_singleton

import "sync"

type LazySingleton struct {
}

var (
	lazySingleton *LazySingleton
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *LazySingleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &LazySingleton{}
		})
	}
	return lazySingleton
}
