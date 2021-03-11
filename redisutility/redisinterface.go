package redisutility

import (
	"api-gateway/config"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

var env config.Configuration

//TODO: add log.printf to set these in the logs

func SetConfiguration(config config.Configuration) {
	// injection of environment configuration
	env = config
}

func GetSecret() string {
	return env.Secret
}

// User is a simple user struct for this example
type UserToken struct {
	UserId  int 	 `json:"userid"`
	Token  string    `json:"token"`
}


func createConnection() redis.Conn {
	// newPool returns a pointer to a redis.Pool
	pool := newPool()
	// get a connection from the pool (redis.Conn)
	conn := pool.Get()
	// use defer to close the connection when the function completes
	return conn
	//// call Redis PING command to test connectivity
	//err := Ping(conn)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("successful connection")
	//}
}

func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// ping tests connectivity for redis (PONG should be returned)
func ping(c redis.Conn) error {
	// Send PING command to Redis
	// PING command returns a Redis "Simple String"
	// Use redis.String to convert the interface type to string
	s, err := redis.String(c.Do("PING"))
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)
	// Output: PONG

	return nil
}

func SetUserToken(userId int, token string) error {

	conn := createConnection()
	defer conn.Close()

	userIdvalue := strconv.FormatInt(int64(userId), 10)

	fmt.Println(env.Redis_UserId_Expire)
	fmt.Printf("Store token: %s record in redis", token)

	_, err := conn.Do("MSET", userIdvalue, token, token, userId)
	_, err = conn.Do("EXPIRE", userIdvalue, env.Redis_UserId_Expire)
	_, err = conn.Do("EXPIRE", token, env.Redis_UserToken_Expire)
	if err != nil {
		fmt.Println("error setting object usertoken")
		return err
	}

	return nil
}

func RefreshUserToken(token string) error {

	fmt.Printf("Token to be refreshed: %s ", token)

	conn := createConnection()
	defer conn.Close()

	userid, err := redis.String(conn.Do("GET", token))

	fmt.Println(env.Redis_UserId_Expire)
	fmt.Printf("Refresh token: %s record in redis", token)

	_, err = conn.Do("MSET", userid, token, token, userid)
	_, err = conn.Do("EXPIRE", userid, env.Redis_UserId_Expire)
	_, err = conn.Do("EXPIRE", token, env.Redis_UserToken_Expire)
	if err != nil {
		fmt.Println("error setting object usertoken")
		return err
	}

	return nil
}

func GetUserTokenForId(userId int) (string, error) {

	conn := createConnection()
	defer conn.Close()

	token, err := redis.String(conn.Do("GET", strconv.FormatInt(int64(userId), 10)))
	if len(token) > 0 {
		fmt.Println("User has token record in redis")
		return token, err
	} else {
		fmt.Println("User does not have token record in redis")
		return "", nil
	}
}

func GetUserIdForToken(userToken string) (string, error) {

	conn := createConnection()
	defer conn.Close()

	userid, err := redis.String(conn.Do("GET", userToken))
	if len(userid) > 0 {
		fmt.Println("User has token record in redis")
		return userid, err
	} else {
		fmt.Println("User does not have token record in redis")
		return "", nil
	}

}
