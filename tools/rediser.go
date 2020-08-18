package tools

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"strings"
)

type Scanner struct {
	Host   string
	Port   string
	redis  *redis.Client
	sysctx context.Context
}

func NewScanner(host, port string) (*Scanner, error) {
	s := &Scanner{
		Host:   host,
		Port:   port,
		sysctx: context.Background(),
	}
	if s.Host == "" || s.Port == "" {
		return nil, errors.New("redis 配置有误")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Host + ":" + s.Port,
	})
	_, err := rdb.Ping(s.sysctx).Result()
	if err != nil {
		return nil, err
	}
	s.redis = rdb
	return s, nil
}

func (receiver Scanner) ScanCode(pattern map[string]string) (codes map[string]string) {
	codes = make(map[string]string)
	for project, pattern := range pattern {
		temp := receiver.redis.Keys(receiver.sysctx, pattern)
		// 获取到符合类型的键
		smsKeys := temp.Val()
		for _, smsKey := range smsKeys {
			// 再根据键获取值
			val := receiver.redis.Get(receiver.sysctx, smsKey).Val()
			prefixPhone := strings.Split(smsKey, ":")
			l := len(prefixPhone)
			if l != 0 {
				codes[project+":"+prefixPhone[l-1]] = val
			}
		}
	}
	if len(codes) == 0 {
		codes["别看了"] = "啥也没有"
	}
	return
}
