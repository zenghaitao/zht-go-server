package redispool

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

type lockCfg struct {
	Pool *redis.Pool
	Key string
	Owner string
	Timeout int
}

func (r *Redis)Lock(key string, owner string, timeout int) *lockCfg{
	return &lockCfg{
		Pool: r.Pool,
		Key: key,
		Owner: owner,
		Timeout: timeout,
	}
}

func (r *lockCfg)Set() (ok bool, owner string, err error) {
	c := r.Pool.Get()
	defer c.Close()

	_, err = redis.String(c.Do("SET", r.Key, r.Owner, "EX", r.Timeout, "NX"))
	if err != nil {
		if err == redis.ErrNil {
			//加锁失败
			return false, "", nil
		}
		//加锁失败
		return false, "" , err
	}
	//加锁成功
	return true, r.Owner , nil
}

func (r *lockCfg)Release(owner string)(ok bool,err error)  {
	c := r.Pool.Get()
	defer c.Close()

	res ,err := c.Do("EXISTS",r.Key)
	if err != nil || res == 0{
		return true, nil
	}

	lua :=redis.NewScript(1,lockScript())
	res , err = lua.Do(c,r.Key,owner)

	if err != nil {
		return false, err
	}

	//不属于该释放锁的线程
	if res == 0 {
		err = errors.New("非法用户，无法释放该锁")
		return false, err
	}

	return true,nil
}

func lockScript() string{
	return "if redis.call(\"get\",KEYS[1]) == ARGV[1] then\n    return redis.call(\"del\",KEYS[1])\nelse\n    return 0\nend"
}