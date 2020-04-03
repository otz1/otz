package cache

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

const defaultLiveDuration = time.Second * 5

func hash(keyword string) string {
	keyword = strings.ToLower(strings.TrimSpace(keyword))
	return fmt.Sprintf("kw_%s", keyword)
}

// StoreKeyword will store the given keyword in the cache
// ...
// the cache stores a map of all the keywords. the keywords
// have an expiry rate set
// the more we 'store a keyword' the longer the keyword will last for
// in the cache
func StoreKeyword(keyword string) {
	key := hash(keyword)

	_, err := redisDAO.Get(key).Result()
	if err == redis.Nil {
		redisDAO.Set(key, 1, time.Minute*10)
		return
	}
	redisDAO.Incr(key)
}

func delKeyword(keywords ...string) error {
	hashed := make([]string, len(keywords))
	for i, k := range keywords {
		hashed[i] = hash(k)
	}

	err := redisDAO.Del(hashed...).Err()
	if err != nil {
		sentry.CaptureException(err)
		log.Println(err)
		return errors.New("failed to delete key")
	}
	return nil
}

func getKeyword(keyword string) (int64, error) {
	val, err := redisDAO.Get(hash(keyword)).Result()
	if err != nil {
		sentry.CaptureException(err)
		log.Println(err)
		return 0, errors.New("faild to fetch key")
	}

	count, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		sentry.CaptureException(err)
		log.Println(err)
		return 0, errors.New("failed to parse count")
	}

	return count, nil
}

type Keyword struct {
	Word  string
	Count int64
}

// AllKeywords will return all the keywords
// and their values
// warning: this is slow. dont use this a lot!
// we only have this for the runner which runs periodically
func AllKeywords() ([]Keyword, error) {
	keys, err := redisDAO.Keys("kw:*").Result()
	if err != nil {
		sentry.CaptureException(err)
		log.Println(err)
		return []Keyword{}, errors.New("failed to get all keys")
	}

	res := []Keyword{}
	for _, k := range keys {
		val, err := redisDAO.Get(k).Result()
		if err != nil {
			sentry.CaptureException(err)
			log.Println(err)
			continue
		}
		count, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			sentry.CaptureException(err)
			log.Println(err)
			continue
		}
		res = append(res, Keyword{
			Word:  k,
			Count: count,
		})
	}

	return res, nil
}
