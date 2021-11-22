package redispool

import (
	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	Pool *redis.Pool
}

func Select(name string) *Redis {
	return &Redis{
		Pool:Pool(name),
	}
}

/**
 * [Key]<Value>\EX<time> 过期时间|PX<time> 毫秒级过期时间\NX 不存在时设置|XX 存在时设置
 */
func (r *Redis)Set( args ...interface{} ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("SET" , args...)
}

/**
 * 同时为多个键设置值
 * [Key]<Value>...
 */
func (r *Redis)MSet(args ...interface{} ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("MSET" , args...)
}

/**
 * 当且仅当所有给定键都不存在时， 为所有给定键设置值
 * [Key]<Value>...
 */
func (r *Redis)MSetNx(args ...interface{} ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("MSETNX" , args...)
}

/**
 * 追加值到现有值末尾
 */
func (r *Redis)Append(key string,value string ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("APPEND" , key , value)
}

/**
 * 将键 key 的值设为 value ， 并返回键 key 在被设置之前的旧值
 */
func (r *Redis)GetSet(key string,value string ) (reply string, err error){
	c := r.Pool.Get()
	defer c.Close()

	return redis.String(c.Do("GETSET" , key , value))
}

/**
 * 从偏移量 offset 开始， 用 value 参数覆写(overwrite)键 key 储存的字符串值
 */
func (r *Redis)SetRange(key string, offset int, value string ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("SETRANGE" , key , offset , value)
}


func (r *Redis)Get(key string) (reply string, err error){
	c := r.Pool.Get()
	defer c.Close()

	return redis.String(c.Do("GET" , key))
}

/**
 * 返回给定的一个或多个字符串键的值
 */
func (r *Redis)MGet(args ...interface{} ) (reply []string, err error){
	c := r.Pool.Get()
	defer c.Close()

	return redis.Strings(c.Do("MGET" , args...))
}

/**
 * 返回键 key 储存的字符串值的指定部分,负数偏移量表示从字符串的末尾开始计数
 */
func (r *Redis)GetRange(key string, star int, end int ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("GETRANGE" , key , star , end)
}

/**
 * 为键 key 储存的数字值加上一
 */
func (r *Redis)Incr(key string) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("INCR" , key)
}

/**
 * 为键 key 储存的数字值减去一
 */
func (r *Redis)Decr(key string , increment int) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("DECR" , key , increment)
}

/**
 * 为键 key 储存的数字值加上增量 increment
 */
func (r *Redis)IncrBy(key string , increment int) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("INCRBY" , key , increment)
}

/**
 * 将键 key 储存的整数值减去减量 decrement
 */
func (r *Redis)DecrBy(key string , increment int) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("DECRBY" , key , increment)
}

/**
 * 为键 key 储存的值加上浮点数增量 increment
 */
func (r *Redis)IncrByFloat(key string , increment float32) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("INCRBYFLOAT" , key , increment)
}


func (r *Redis)Strlen(key string) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("STRLEN" , key)
}

func (r *Redis)Del(key string) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("DEL" , key)
}

func (r *Redis)Exists(key string) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("EXISTS" , key)
}

/**
 * 为给定 key 设置生存时间
 */
func (r *Redis)Expire(key string , time int) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("EXPIRE" , key , time)
}

/**
 * 为给定 key 设置生存时间(毫秒)
 */
func (r *Redis)PExpire(key string , time int) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("PEXPIRE" , key , time)
}

/**
 * 为 key 设置生存时间至目标时间戳
 */
func (r *Redis)ExpireAt(key string , time int) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("EXPIREAT" , key , time)
}

/**
 * 为 key 设置生存时间至目标时间戳(毫秒)
 */
func (r *Redis)PExpireAt(key string , time int) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("PEXPIREAT" , key , time)
}

/**
 * 以秒为单位，返回给定 key 的剩余生存时间
 */
func (r *Redis)TTL(key string) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("TTL" , key)
}

/**
 * 以秒为单位，返回给定 key 的剩余生存时间(毫秒)
 */
func (r *Redis)PTTL(key string) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("PTTL" , key)
}

/**
 * 移除给定 key 的生存时间,变为持久KEY
 */
func (r *Redis)Persist(key string) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	return c.Do("PERSIST" , key)
}







