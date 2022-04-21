package redis

import "encoding/json"

func (s *RedisClient) GetString(key string) (string, error) {
	val, err := s.Redis.Get(s.Context, key).Result()
	return val, err
}

func (s *RedisClient) GetInt(key string) (int, error) {
	val, err := s.Redis.Get(s.Context, key).Int()
	return val, err
}

func (s *RedisClient) GetBool(key string) (bool, error) {
	val, err := s.Redis.Get(s.Context, key).Bool()
	return val, err
}

func (s *RedisClient) GetObject(obj interface{}, key string) error {
	valBytes, err := s.Redis.Get(s.Context, key).Bytes()

	if err == nil {
		err = json.Unmarshal(valBytes, obj)
	}

	return err
}
