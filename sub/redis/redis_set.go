package redis

import (
	jsoniter "github.com/json-iterator/go"
	"time"
)

func (s *RedisClient) SetString(key string, value string) error {
	err := s.SetStringWithExpiration(key, value, 0)
	return err
}

func (s *RedisClient) SetStringWithExpiration(key string, value string, expiration time.Duration) error {
	err := s.Redis.Set(s.Context, key, value, expiration).Err()
	return err
}

func (s *RedisClient) SetInt(key string, value int) error {
	err := s.SetIntWithExpiration(key, value, 0)
	return err
}

func (s *RedisClient) SetIntWithExpiration(key string, value int, expiration time.Duration) error {
	err := s.Redis.Set(s.Context, key, value, expiration).Err()
	return err
}

func (s *RedisClient) SetBool(key string, value bool) error {
	err := s.SetBoolWithExpiration(key, value, 0)
	return err
}

func (s *RedisClient) SetBoolWithExpiration(key string, value bool, expiration time.Duration) error {
	err := s.Redis.Set(s.Context, key, value, expiration).Err()
	return err
}

func (s *RedisClient) SetObject(key string, value interface{}) error {
	err := s.SetObjectWithExpiration(key, value, 0)
	return err
}

func (s *RedisClient) SetObjectWithExpiration(key string, value interface{}, expiration time.Duration) error {
	valueToSave, err := jsoniter.Marshal(value)

	if err == nil {
		err = s.Redis.Set(s.Context, key, string(valueToSave), expiration).Err()
	}

	return err
}
