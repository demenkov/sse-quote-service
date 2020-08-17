package ssequote

import (
	"github.com/go-redis/redis"
	"reflect"
	"time"
)

var (
	Conf  *Config
	Redis *redis.Client
)

func InitPackage(p ...interface{}) {
	for _, r := range p {
		if c, ok := r.(*Config); ok {
			Conf = c
		}
	}

	Redis = RedisClient()
}

func Timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func in_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
