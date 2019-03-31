/*
@Time : 2019/3/31 14:24 
@Author : shilinqing
@File : redis
@URL : https://github.com/hoisie/redis
@Method:

Redis Set:
判断成员元素是否是集合的成员: 	client.Sismember(URL_VISIT_SET, []byte(url))
将一个或多个成员元素加入到集合中
已经存在于集合的成员元素将被忽略: client.Sadd(URL_VISIT_SET, []byte(url))

Redis Scard 命令返回集合中元素的数量
Redis Sdiff 命令返回给定集合之间的差集。不存在的集合 key 将视为空集。
Redis Sinter 命令返回给定所有给定集合的交集。 不存在的集合 key 被视为空集。
Redis Spop 命令用于移除集合中的指定 key 的一个或多个随机元素，移除后会返回移除的元素。

Redis List:
加入集合: 									client.Lpush(QUEUE, []byte(url))
获取长度: 									client.Llen(URL_QUEUE)
命令用于移除列表的最后一个元素，返回值为移除的元素: 	client.Rpop(URL_QUEUE)
Redis Lpop 命令用于移除并返回列表的第一个元素。
Redis Lindex 命令用于通过索引获取列表中的元素。
Redis Rpush 命令用于将一个或多个值插入到列表的尾部(最右边)。
Redis Lrange 返回列表中指定区间内的元素，区间以偏移量 START 和 END 指定。


Redis 事务:
Redis Multi 命令用于标记一个事务块的开始。
Redis Exec 命令用于执行所有事务块内的命令。
Redis Discard 命令用于取消事务，放弃执行事务块内的所有命令。
Redis Unwatch 命令用于取消 WATCH 命令对所有 key 的监视。
Redis Watch 命令用于监视一个(或多个) key ，如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断
*/
package models

import (
"github.com/astaxie/goredis"
)

var (
	client goredis.Client
)

const(
	URL_QUEUE = "url_queue"
	URL_VISIT_SET = "url_visit_set"
)

/**
连接redis
"127.0.0.1:6389"
 */
func ConnectRedis(addr string, port string , password string){
	client.Addr = addr + ":" + port
	if password != ""{
		client.Password = password
	}
}






