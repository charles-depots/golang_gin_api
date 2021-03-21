package cache

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	"sync"
	"testing"
	"time"
)

// Create the Redis client test
func InitRedisClientTest() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379", // Redis addr:port
		Password:     "root",           // Password
		DB:           0,                // The serial number starts at 0, the default is 0, can not be set
		MaxRetries:   3,                // The maximum number of retries when the command fails. Default is 0, i.e. no retry
		PoolSize:     10,               // Maximum number of connections in the connection pool. Default is number of CPUs * 10
		MinIdleConns: 5,                // Minimum number of free connections
	})

	if err := client.Ping().Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis error")
	}

	return client, nil
}

// @link: https://redis.io/topics/data-types-intro
// StringDemo:
func TestStringDemo(t *testing.T) {
	fmt.Println("------------------ Welcome to stringDemo  ----------------------")
	redisClient, err := InitRedisClientTest()
	if err != nil {
		fmt.Println("StringDemo redisClient is nil")
		return
	}

	key := "username"
	username := "charles"
	redisClient.Set(key, username, 10*time.Second)
	val := redisClient.Get(key)
	if val == nil {
		fmt.Errorf("StringDemo get error")
	}
	fmt.Println("", val)
}

// ListDemo:
func TestListDemo(t *testing.T) {
	fmt.Println("------------------ Welcome to listDemo  ------------------------")
	redisClient, err := InitRedisClientTest()
	if err != nil {
		fmt.Println("StringDemo redisClient is nil")
		return
	}

	latestUser := "login:last_login_user"
	result, err := redisClient.RPush(latestUser, "user1", "user2", "user3").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result:", result)

	result, err = redisClient.LPush(latestUser, "user0").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result:", result)

	length, err := redisClient.LLen(latestUser).Result()
	if err != nil {
		fmt.Println("ListDemo LLen is nil")
	}
	fmt.Println("length: ", length)

	mapOut, err := redisClient.LRange(latestUser, 0, 2).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for inx, item := range mapOut {
		fmt.Printf("%v:%v\n", inx, item)
	}
}

// HashDemo:
func TestHashDemo(t *testing.T) {
	fmt.Println("------------------ Welcome to HashDemo  ----------------------")
	redisClient, err := InitRedisClientTest()
	if err != nil {
		fmt.Println("StringDemo redisClient is nil")
		return
	}

	//Init a map[string]interface{}
	userInfo := make(map[string]interface{})
	userInfo["name"] = "charles"
	userInfo["age"] = 27
	userId := "user:10001"

	hash, err := redisClient.HMSet(userId, userInfo).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash)

	mapOut := redisClient.HGetAll(userId).Val()
	for inx, item := range mapOut {
		fmt.Printf("%v:%v\n", inx, item)
	}
}

// TestConnectPool
func TestConnectPool(t *testing.T) {
	fmt.Println("----------------- Welcome to connect pool test  -----------------")
	redisClient, err := InitRedisClientTest()
	if err != nil {
		fmt.Println("StringDemo redisClient is nil")
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				redisClient.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
				redisClient.Get(fmt.Sprintf("name%d", j)).Result()
			}

			fmt.Printf("PoolStats, TotalConns: %d, IdleConns: %d\n", redisClient.PoolStats().TotalConns, redisClient.PoolStats().IdleConns)
		}()
	}

	wg.Wait()
}
