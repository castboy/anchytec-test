package main
import (
	"github.com/go-redis/redis"
	"fmt"
	"encoding/json"
	"time"
)

// redis-server.exe redis.windows.conf

type PersonRedis struct {
	Name string
	Age int
}

func (p PersonRedis) MarshalBinary() (data []byte, err error) {
	b, err := json.Marshal(p)
	return []byte(b), err
}

func main() {
	// Client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// Set
	person := PersonRedis{
		Name: "wmg",
		Age: 29,
	}

	err = client.Set("key", person, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// SetBit
	// `a` - 01100001  ; `b` - 01100010  ; offset from left to right.
	err = client.Set("andy", `a`, 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(client.SetBit("andy", 6, 1).String())
	fmt.Println(client.SetBit("andy", 7, 0).String())
	val, err = client.Get("andy").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("andy", val)


	// SetNX, if exist return 0 represent fail, not exist return 1 represent success
	fmt.Println(client.Exists("wmg").Val())
	fmt.Println(client.SetNX("wmg", "wmg", time.Second).Val())
	fmt.Println(client.SetNX("wmg", "wmg", time.Second).Val())

	// SetRange, cover string from offset with new string
	client.Set("k", "Hello World", 0)
	client.SetRange("k", 6, "Redis").String()
	fmt.Println(client.Get("k").String())

	// SetXX, success only key already exist.
	fmt.Println(client.SetXX("k", "ming", 0).Val())
}
