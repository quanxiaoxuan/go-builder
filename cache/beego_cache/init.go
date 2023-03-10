package beego_cache

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/cache"
)

var beegoCache cache.Cache

func init() {
	var err error
	beegoCache, err = cache.NewCache("memory", `{"interval":"360"}`)
	if beegoCache == nil || err != nil {
		fmt.Println("NewCache failed, err : ", err)
	}
}

func Put(k string, v interface{}, d time.Duration) error {
	return beegoCache.Put(k, v, d)
}

func Get(k string) interface{} {
	return beegoCache.Get(k)
}

func GetMulti(keys []string) interface{} {
	return beegoCache.GetMulti(keys)
}

func Delete(k string) error {
	return beegoCache.Delete(k)
}

func Decr(k string) error {
	return beegoCache.Decr(k)
}

func ClearAll() error {
	return beegoCache.ClearAll()
}

func IsExist(k string) bool {
	return beegoCache.IsExist(k)
}
