
# golang redis.pool链接池 redis.lock分布式锁

基于redigo的redis pool和分布式锁的实现

1、config目录下是redis的配置项，可配置多个redis源

2、调用redispool.Select("redis源名称")获取redis操作链接

3、使用Lock()方法对分布式锁进行初始化，定义key、owner、timeout

4、redispool简单封装了String和Hash两种类型的redis方法，mget返回[]string,hmgetall返回map[string]string

5、具体使用可以参见main.go文件

# 性能参考：

1000次SET操作(参考参数:MaxIdle:10,MaxActive:50)

golang用时65-80ms

php-swoole用时100-120ms

golang相比php的协程性能高出30%-40%左右
