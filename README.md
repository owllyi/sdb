## [SDB](https://github.com/yemingfeng/sdb) ：纯 golang 开发、数据结构丰富、持久化的 NoSQL 数据库
------

### 为什么需要 SDB？

试想以下业务场景：

- 计数服务：对内容的点赞、播放等数据进行统计
- 评论服务：发布评论后，查看某个内容的评论列表
- 推荐服务：每个用户有一个包含内容和权重推荐列表

以上几个业务场景，都可以通过 MySQL + Redis 的方式实现。 这里的问题是：MySQL 更多的是充当持久化的能力，Redis 充当的是在线服务的读写能力。

那么只使用 Redis 行不行？ 答案是否定的，因为 Redis 无法保证数据不丢失。

那有没有一种存储能够支持高级的数据结构，并能够将数据进行持久化的呢？

答案是：非常少的。有些数据库要么是支持的数据结构不够丰富，要么是接入成本太高，要么是不可控。

为了解决上述问题，SDB 产生了。

------

### SDB 简单介绍

- 纯 golang 开发，核心代码不超过 1k，代码易读
- 数据结构丰富
    - [string](https://github.com/yemingfeng/sdb/blob/master/api/protobuf-spec/string.proto)
    - [list](https://github.com/yemingfeng/sdb/blob/master/api/protobuf-spec/list.proto)
    - [set](https://github.com/yemingfeng/sdb/blob/master/api/protobuf-spec/set.proto)
    - [sorted set](https://github.com/yemingfeng/sdb/blob/master/api/protobuf-spec/sorted_set.proto)
    - [bloom filter](https://github.com/yemingfeng/sdb/blob/master/api/protobuf-spec/bloom_filter.proto)
    - [hyper log log](https://github.com/yemingfeng/sdb/blob/master/api/protobuf-spec/hyper_log_log.proto)
    - [pub sub](https://github.com/yemingfeng/sdb/blob/master/api/protobuf-spec/pub_sub.proto)
- 持久化
    - 兼容 [pebble](https://github.com/cockroachdb/pebble)
      、[leveldb](https://github.com/syndtr/goleveldb)
      、[badger](https://github.com/dgraph-io/badger) 存储引擎
- 监控
    - 支持 prometheus + grafana 监控方案
- 限流
    - 支持每秒 qps 的限流策略
- 慢查询查看
    - 可查看慢查询的请求，进行分析

------

### 快速使用

#### 服务端使用

```shell
sh ./scripts/quick_start.sh
```

**默认使用 pebble 存储引擎**。启动后，端口会监听 9000 端口

#### 客户端使用

```go
package main

import (
	"github.com/yemingfeng/sdb/pkg/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Printf("faild to connect: %+v", err)
	}
	defer conn.Close()

	// 连接服务器
	c := pb.NewSDBClient(conn)
	setResponse, err := c.Set(context.Background(),
		&pb.SetRequest{Key: []byte("hello"), Value: []byte("world")})
	log.Printf("setResponse: %+v, err: %+v", setResponse, err)
	getResponse, err := c.Get(context.Background(),
		&pb.GetRequest{Key: []byte("hello")})
	log.Printf("getResponse: %+v, err: %+v", getResponse, err)
}
```

#### 更多客户端例子

- [string 操作](https://github.com/yemingfeng/sdb/blob/master/examples/string.go)
- [list 操作](https://github.com/yemingfeng/sdb/blob/master/examples/list.go)
- [set 操作](https://github.com/yemingfeng/sdb/blob/master/examples/set.go)
- [sorted set 操作](https://github.com/yemingfeng/sdb/blob/master/examples/sorted_set.go)
- [bloom filter 操作](https://github.com/yemingfeng/sdb/blob/master/examples/bloom_filter.go)
- [pub sub 操作](https://github.com/yemingfeng/sdb/blob/master/examples/pub_sub.go)

------

### [配置大全](https://github.com/yemingfeng/sdb/blob/master/configs/config.yml)

参数名 | 含义 | 默认值
---- | --- | ---
store.engine | 存储引擎，可选 pebble、level、badger | pebble
store.path | 存储目录 | ./db
server.grpc_port | grpc 监听的端口 | 9000
server.http_port | http 监控的端口，供 prometheus 使用 | 8081
server.rate | 每秒 qps 的限制 | 30000
server.slow_query_threshold | 慢查询记录的阈值，单位为 ms | 100

------

### 性能测试

测试脚本：[benchmark](https://github.com/yemingfeng/sdb/blob/master/examples/benchmark_sdb.go)

测试机器：MacBook Pro (13-inch, 2016, Four Thunderbolt 3 Ports)

处理器：2.9GHz 双核 Core i5

内存：8GB

**测试结果： peek QPS > 12k，avg QPS > 7k，set avg time < 70ms，get avg time <
0.2ms**

<img src="https://github.com/yemingfeng/sdb/raw/master/docs/benchmark.png" width=80% />

------

### 监控

#### 安装 docker 版本 grafana、prometheus（可跳过）

- 启动 [scripts/run_monitor.sh](https://github.com/yemingfeng/sdb/blob/master/scripts/run_monitor.sh)

#### 配置 grafana

- 打开 grafana：http://localhost:3000 （注意替换 ip 地址）
- 新建 prometheus datasources：http://host.docker.internal:9090 （如果使用 docker 安装则为这个地址。如果
  host.docker.internal
  无法访问，就直接替换 [prometheus.yml](https://github.com/yemingfeng/sdb/blob/master/scripts/prometheus.yml)
  文件的 host.docker.internal 为自己的 ip 地址就行）
- 将 [scripts/dashboard.json](https://github.com/yemingfeng/sdb/blob/master/scripts/dashboard.json)
  文件导入 grafana dashboard

最终效果可参考：性能测试的 grafana 图

------

### SDB 背后的思考

#### SDB 存储引擎选型

SDB 项目最核心的问题是数据存储方案的问题。

首先，我们不可能手写一个存储引擎。这个工作量太大，而且不可靠。 我们得在开源项目中找到适合 SDB 定位的存储方案。

SDB 需要能够提供高性能读写能力的存储引擎。 单机存储引擎方案常用的有：B+ 树、LSM 树、B 树等。

还有一个前置背景，golang 在云原生的表现非常不错，而且性能堪比 C 语言，开发效率也高，所以 SDB 首选使用纯 golang 进行开发。

那么现在的问题变成了：找到一款纯 golang 版本开发的存储引擎，这是比较有难度的。收集了一系列资料后，找到了以下开源方案：

- LSM 树
    - [go-leveldb](https://github.com/golang/leveldb/) ：是一个 unstable 的项目，无法使用
    - [syndtr-goleveldb](https://github.com/syndtr/goleveldb)
    - [badger](https://github.com/dgraph-io/badger)
    - [pebble](https://github.com/cockroachdb/pebble)
- B+ 树
    - [boltdb-bolt](https://github.com/boltdb/bolt) ：是废弃的项目，无法使用
    - [etcd-bolt](https://github.com/etcd-io/bbolt) ：主要是用于分布式环境下的数据同步，无法应对高并发的数据读写

综合来看，golangdb、badger、pebble 这三款存储引擎都是很不错的。

为了兼容这三款存储引擎，SDB
提供了抽象的[接口](https://github.com/yemingfeng/sdb/blob/master/internal/store/engine/interface.go)
，进而适配这三个存储引擎。

------

#### SDB 数据结构设计

SDB 已经通过上面三款存储引擎解决了数据存储的问题了。 但如何在 KV 的存储引擎上支持丰富的数据结构呢？

以 pebble 为例子，首先 pebble 提供了以下的接口能力：

- set(k, v)
- get(k)
- del(k)
- batch
- iterator

接下来，我以支持 List 数据结构为例子，剖析下 SDB 是如何通过 pebble 存储引擎支持 List 的。

List 数据结构提供了以下接口：LPush、LPop、LExist、LRange、LCount。

如果一个 List 的 key 为：[hello]，该 List 的列表元素有：[aaa, ccc, bbb]，那么该 List 的每个元素在 pebble 的存储为：

pebble key | pebble value
---- | ---
l/hello/{unique_ordering_key1} | aaa
l/hello/{unique_ordering_key2} | ccc
l/hello/{unique_ordering_key3} | bbb

List 元素的 pebble key 生成策略：

- 数据结构前缀：List 都以 **l** 字符为前缀，Set 是以 **s** 为前缀...
- List key 部分：List 的 key 为 hello
- unique_ordering_key：生成方式是通过雪花算法实现的，雪花算法保证局部自增
- pebble value 部分：List 元素真正的内容，如 aaa、ccc、bbb

为什么这么就能保证 List 的插入顺序呢？

这是因为 pebble 是 LSM 的实现，内部使用 key 的字典序排序。为了保证插入顺序，SDB 在 pebble key 中增加了 unique_ordering_key
作为排序的依据，从而保证了插入顺序。

有了 pebble key 的生成策略，一切都变得简单起来了。我们看看 LPush、LPop、LRange 的核心逻辑：

##### LPush

```go
func LPush(key []byte, values [][]byte) (bool, error) {
	batchAction := store.NewBatchAction()
	defer batchAction.Close()

	for _, value := range values {
		batchAction.Set(generateListKey(key, util.GetOrderingKey()), value)
	}

	return batchAction.Commit()
}
```

##### LPop

在写入到 pebble 的时候，key 的生成是通过 unique_ordering_key 的方案。 无法直接在 pebble 中找到 List 的元素在 pebble
key。在删除一个元素的时候，需要遍历 List 的所有元素，找到 value = 待删除的元素，然后进行删除。核心逻辑如下：

```go
func LPop(key []byte, values [][]byte) (bool, error) {
	batchAction := store.NewBatchAction()
	defer batchAction.Close()

	store.Iterate(&store.IteratorOption{Prefix: generateListPrefixKey(key)},
		func(key []byte, value []byte) {
			for i := range values {
				if bytes.Equal(values[i], value) {
					batchAction.Del(key)
				}
			}
		})

	return batchAction.Commit()
}
```

##### LRange

和删除逻辑类似，通过 iterator
接口进行遍历。 [这里对反向迭代做了额外的支持](https://github.com/yemingfeng/sdb/blob/master/internal/store/engine/pebble/store.go#L93)
允许 Offset 传入 -1，代表从后进行迭代。

```go
func LRange(key []byte, offset int32, limit int32) ([][]byte, error) {
	index := int32(0)
	res := make([][]byte, limit)
	store.Iterate(&store.IteratorOption{
		Prefix: generateListPrefixKey(key), Offset: int(offset), Limit: int(limit)},
		func(key []byte, value []byte) {
			res[index] = value
			index++
		})
	return res[0:index], nil
}
```

以上就实现了对 List 的数据结构的支持。

其他的数据结构大体逻辑类似，其中 [sorted_set](https://github.com/yemingfeng/sdb/blob/master/internal/service/sorted_set.go)
更加复杂些。可以自行查看。

#### SDB 通讯协议方案

解决完了存储和数据结构的问题后，SDB 面临了【最后一公里】的问题是通讯协议的选择。

SDB 的定位是支持多语言的，所以需要选择支持多语言的通讯框架。

grpc 是一个非常不错的选择，只需要使用 SDB proto 文件，就能通过 protoc 命令行工具自动生成各种语言的客户端，解决了需要开发不同客户端的问题。

#### SDB 集群方案

SDB 的集群方案其实是在规划中的，之前也考虑了 TiKV 集群方案和 Redis 集群方案。

但目前 SDB 把注意力放在持久化、数据结构上。增加更多的数据结构，并将易用性做到极致。**之后再实现集群方案。**

------

### 规划

- 支持更多的存储引擎
    - LSM
    - B+ Tree
- 支持对现有数据结构更多的操作
- 支持更丰富的数据结构
    - geo hash
    - 倒排索引
    - 向量检索
    - 广告定向
- 搭建 admin web ui

------

### 感谢

**感谢开源的力量，这里就不一一列举了，请大家移步 [go.mod](https://github.com/yemingfeng/sdb/blob/master/go.mod)**