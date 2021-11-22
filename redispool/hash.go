package redispool

import "github.com/gomodule/redigo/redis"

/**
 * 将哈希表 hash 中域 field 的值设置为 value
 */
func (r *Redis)HSet( key string, field string, value string ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	reply,err = c.Do("HSET" , key , field , value)
	return reply,err
}

/**
 * 当且仅当域 field 尚未存在于哈希表的情况下， 将它的值设置为 value
 */
func (r *Redis)HSetNx( key string, field string, value string ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	reply,err = c.Do("HSETNX" , key , field , value)
	return reply,err
}

/**
 * 检查给定域 field 是否存在于哈希表 hash 当中
 */
func (r *Redis)HExists( key string ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	reply,err = c.Do("HEXISTS" , key)
	return reply,err
}

/**
 * 删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略
 * Key\field...
 */
func (r *Redis)HDel( args ...interface{} ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	reply,err = c.Do("HDEL" , args...)
	return reply,err
}

/**
 * 返回哈希表 key 中域的数量
 */
func (r *Redis)HLen( key string ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	reply,err = c.Do("HLEN" , key)
	return reply,err
}

/**
 * 返回哈希表 key 中域的数量
 */
func (r *Redis)HStrLen( key string , field string ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	reply,err = c.Do("HSTRLEN" , key , field)
	return reply,err
}

/**
 * 为哈希表 key 中的域 field 的值加上增量 increment
 */
func (r *Redis)HIncrBy( key string , field string , increment int ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	reply,err = c.Do("HINCRBY" , key , field , increment)
	return reply,err
}

/**
 * 为哈希表 key 中的域 field 的值加上增量 increment
 */
func (r *Redis)HIncrByFloat( key string , field string , increment float32 ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	reply,err = c.Do("HINCRBYFLOAT" , key , field , increment)
	return reply,err
}

/**
 * 同时将多个 field-value (域-值)对设置到哈希表 key 中
 * Key\[field]<value>...
 */
func (r *Redis)HMSet( args ...interface{} ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	reply,err = c.Do("HMSET" , args...)
	return reply,err
}

/**
 * 返回哈希表 key 中，一个或多个给定域的值
 * Key\field...
 */
func (r *Redis)HMGet( args ...interface{} ) (reply []string, err error){
	c := r.Pool.Get()
	defer c.Close()

	reply,err = redis.Strings(c.Do("HMGET" , args...))
	return reply,err
}

/**
 * 返回哈希表 key 中的所有域
 */
func (r *Redis)HKeys( key string ) (reply []string, err error){
	c := r.Pool.Get()
	defer c.Close()

	return redis.Strings(c.Do("HKEYS" , key))
}

/**
 * 返回哈希表 key 中所有域的值
 */
func (r *Redis)HVals( key string ) (reply []string, err error){
	c := r.Pool.Get()
	defer c.Close()

	return redis.Strings(c.Do("HVALS" , key))
}

/**
 * 返回哈希表 key 中，所有的域和值
 */
func (r *Redis)HGetAll( key string ) (reply map[string]string, err error){
	c := r.Pool.Get()
	defer c.Close()

	var result = make(map[string]string)
	list,err := redis.Strings(c.Do("HGETALL" , key))
	if err == nil{
		k := ""
		for _,v:=range list{
			if k=="" {
				k = v
			}else{
				result[k] = v
				k = ""
			}
		}
		return result,err
	}

	return result,err
}

/**
 * 迭代哈希键中的键值对
 * key\cursor\[match]<pattern>
 */
func (r *Redis)HScan( args ...interface{} ) (reply interface{}, err error){
	c := r.Pool.Get()
	defer c.Close()

	reply,err = c.Do("HSCAN" , args...)
	return reply,err
}
