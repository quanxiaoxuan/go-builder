package go_cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var goCache *cache.Cache

func init() {
	// 创建一个默认过期时间为5分钟的缓存适配器，每60秒清除一次过期的项目
	goCache = cache.New(5*time.Minute, 60*time.Second)
}

func Set(k string, x interface{}, d time.Duration) {
	goCache.Set(k, x, d)
}

func SetDefault(k string, x interface{}) {
	goCache.SetDefault(k, x)
}

func Get(k string) (interface{}, bool) {
	return goCache.Get(k)
}

func Delete(k string) {
	goCache.Delete(k)
}

func Add(k string, x interface{}, d time.Duration) error {
	return goCache.Add(k, x, d)
}

func DeleteExpired() {
	goCache.DeleteExpired()
}
