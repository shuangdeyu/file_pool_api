package service

import (
	"file_pool_api/conf"
	"file_pool_api/helpers"
	"github.com/yvasiyarov/php_session_decoder/php_serialize"
	"helper_go/comhelper"
	"log"

	//"github.com/go-redis/redis"
	"gopkg.in/redis.v5"
	"time"
)

const (
	// app相关
	CACHE_APP_TOKEN = "sys_token_"

	// 用户相关
)

var ExpireData = map[string]time.Duration{
	CACHE_APP_TOKEN: 120 * 3600000000000, // 120小时
}

var RedisConn *redis.Client

// 连接redis
func conn() *redis.Client {
	if RedisConn == nil {
		RedisConn = redis.NewClient(&redis.Options{
			Network:      "tcp",
			Addr:         conf.GetConfig().RedisHost,
			Password:     conf.GetConfig().RedisPassword,
			DB:           conf.GetConfig().RedisDatabase,
			DialTimeout:  conf.GetConfig().RedisTimeout * time.Second,
			PoolSize:     1000,
			PoolTimeout:  2 * time.Minute,
			IdleTimeout:  10 * time.Minute,
			ReadTimeout:  2 * time.Minute,
			WriteTimeout: 1 * time.Minute,
		})
		_, err := RedisConn.Ping().Result()
		log.Println(err)
		if err != nil {
			helpers.WriteAppErrorLog("redis连接出错")
		}
	}
	return RedisConn
}

// 获取缓存信息
func Get(key string) (string, error) {
	v, err := conn().Get(key).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}
func GetByMap(key string) map[string]interface{} {
	val, _ := Get(key)
	data, _ := comhelper.UnSerialize(val)
	if data != nil {
		var ret_data map[string]interface{}
		ret_data = make(map[string]interface{})
		for k, v := range data.(php_serialize.PhpArray) {
			d, ok := v.(php_serialize.PhpArray)
			if ok {
				var tmp_data map[string]interface{}
				tmp_data = make(map[string]interface{})
				for k1, v1 := range d {
					tmp_data[k1.(string)] = v1
				}
				ret_data[k.(string)] = tmp_data
			} else {
				ret_data[k.(string)] = v
			}
		}
		return ret_data
	}
	return nil
}

// 设置缓存信息
func Save(key string, postfix string, val string) bool {
	expire := _expire(key)
	real_key := key + postfix
	err := conn().Set(real_key, val, expire).Err()
	if err != nil {
		//panic(err)
		return false
	}
	return true
}

// 设置带缓存的超时时间
func Save_ex(key string, postfix string, val string, expire time.Duration) bool {
	real_key := key + postfix
	err := conn().Set(real_key, val, expire).Err()
	if err != nil {
		return false
	}
	return true
}

// 删除缓存信息
func Delete(key string) bool {
	err := conn().Del(key).Err()
	if err != nil {
		return false
	}
	return true
}

// 获取缓存（哈希）
func Hget(key, index string) (string, error) {
	v, err := conn().HGet(key, index).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

// 获取缓存（哈希）- 转换成map类型
func HgetByMap(key, index string) map[string]interface{} {
	val, _ := Hget(key, index)
	data, _ := comhelper.UnSerialize(val)
	if data != nil {
		var ret_data map[string]interface{}
		ret_data = make(map[string]interface{})
		for k, v := range data.(php_serialize.PhpArray) {
			d, ok := v.(php_serialize.PhpArray)
			if ok {
				var tmp_data map[string]interface{}
				tmp_data = make(map[string]interface{})
				for k1, v1 := range d {
					tmp_data[k1.(string)] = v1
				}
				ret_data[k.(string)] = tmp_data
			} else {
				ret_data[k.(string)] = v
			}
		}
		return ret_data
	}
	return nil
}

// 设置缓存（哈希）
func Hset(key, index, data string) bool {
	err := conn().HSet(key, index, data).Err()
	if err != nil {
		return false
	}
	return true
}

// 删除缓存（哈希）
func Hdelete(key, index string) bool {
	err := conn().HDel(key, index).Err()
	if err != nil {
		return false
	}
	return true
}

// 获取缓存时间,不在配置列表的代表不过期处理
func _expire(key string) time.Duration {
	if ExpireData[key] > 0 {
		return ExpireData[key]
	} else {
		return 0
	}
}
