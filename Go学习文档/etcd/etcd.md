#  etcd

![图片](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211215164459)

# 1. 什么是etcd

etcd 是一个分布式、高可靠的 kv 系统，基于 Go 语言实现，内部使用 raft 协议来保证数据一致性和高可靠性。



对于 etcd 的定位其实在名字里亦可窥见，etcd 这个名字拆开来看就是 **etc distributed** ，也是就是**分布式的 etc 嘛）**，大家都知道 linux 的 etc 目录就是单机的配置中心。所以 etcd 的应用定位其实就是分布式的配置中心系统嘛，目的就是为服务集群**提供一个的全局性配置中心**。



## etcd 的特点

- **简单**：接口简单，使用简单；
- **安全**：支持 TLS 证书认证；
- **强一致性**：通过 raft 协议保证数据的强一致性，实现 leader 到 follower 节点的正确复制；
- **高可用性**：集群能自动处理单点故障，网络分区等情况，对外提供高可靠、一致性的数据访问；
- **性能不错**：看官网的 benchmark 数据是 10000 次每秒的写入；

其实经过了 k8s 等超大型的项目验证，etcd 已经是一个非常成熟的项目了。



# 2.Leader 是怎么选举出来的？

![图片](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211215170103)


raft 是针对 paxos 的简化版本，拆解为三个**核心问题**：

1. Leader 选举；
2. 日志复制；
3. 正确性的保证；





### raft 的 Leader 选举

简单铺垫 raft 协议的理论基础。

在 raft 协议中角色分为三种：

- Leader ：绝对领导者，集群中只会有一个有效 Leader ，Leader 负责把日志复制到其他节点，并且周期性向 Follower 发出心跳维持统治；
- Follower ：跟随者，被动接收日志；
- Candidate ：候选者，中间状态，竞选 Leader 的状态中；

看下 raft 论文中的角色转化：

![图片](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211215172858)

从论文很清晰能得到几个要点：

1. 节点角色从 Follower 开始；
2. Follower 超时之后变成 Candidate ，竞选成功变成 Leader ，竞选失败变回 Follower ；
3. Leader 发现自己不配，比如收到更高任期的消息时候会自动变成 Follower ；

请求响应参数：

- term ：当前任期号，term 会随着 Leader 的变更而变更，任期是线性递增的；
- index ：日志的 id 号，线性递增的；

成为 Leader 有很多判断，这个此文不表，有一个最关键的点：**具备完备的数据**。详情请见[Raft共识算法](https://github.com/LwwL-123/Go_Study/blob/main/Go学习文档/共识算法/Raft.md)





本篇看一下 Etcd 的具体实现，梳理几个事情：

1. 初始状态是什么？
2. 怎么超时选举？
3. 怎么进行的角色切换，切换之后有什么差别？





### Etcd 的实现

下面从 Etcd 的具体实现来梳理一遍 Leader 选举的过程。



 **1**  **节点初始化**

每个节点在初始化的时候，构造自己的 raft 状态机过程必不可少，调用 newRaft 函数创建一个 raft 状态机对象。在函数中有一个关键操作，每个节点从成为 Follower 开始：

```go
func newRaft(c *Config) *raft {
    // ...
    // 加载出持久化了的 raft 的状态，比如 term，vote，commit index 等
    if !IsEmptyHardState(hs) {
        r.loadState(hs)
    }
    // 大家都是从 Follower 开始做起
    r.becomeFollower(r.Term, None)
    // ...
}
```

**划重点：从 Follower 做起**。来看下 Follower 角色的特色吧：

```go
func (r *raft) becomeFollower(term uint64, lead uint64) {
    // 定制状态机的处理逻辑
    r.step = stepFollower
    // 定制 tick 逻辑
    r.tick = r.tickElection
    // ...
}
```

如上 becomeFollower 最重要的是两个逻辑定制，tick 其实是由外层业务定时驱动的，t.tickElection 可太有意思了，里面藏着一个时刻反叛 Leader 的心：**时刻准备竞争领导权**。

```go
func (r *raft) tickElection() {
    // 每次计数加一
    r.electionElapsed++
    // 如果条件允许(在节点列表中，且不是leader)，并且已经超时，那么开始你的竞争之路吧
    if r.promotable() && r.pastElectionTimeout() {
        // MsgHup 消息有内部产生，开始竞选之路
        r.Step(pb.Message{From: r.id, Type: pb.MsgHup})
    }
}

func (r *raft) pastElectionTimeout() bool {
 // 计数超过了一个数值，那么则去竞争选举
 return r.electionElapsed >= r.randomizedElectionTimeout
}
```

划重点：Follower 角色 tick 计数超过一定阈值（ r.randomizedElectionTimeout ）的时候，就准备开始竞选。



 **2**  **谁都有可能超时 ！**

假设有集群有三个节点，A，B，C 各自 tick ，谁都有可能先超时，先到 pastElectionTimeout 这个条件，因为这个每个节点**阈值是随机生成的**。r.randomizedElectionTimeout 是一个随机值。

```go
func (r *raft) resetRandomizedElectionTimeout() {
    // 随机加了一个随机因子
 r.randomizedElectionTimeout = r.electionTimeout + globalRand.Intn(r.electionTimeout)
}
```

**划重点：每个节点选举超时阈值是随机的**。



 **3**  **从 MsgHup 消息开始**

tick 谁先超时，谁先发出一个 MsgHup 消息投递到 raft 状态机中，一切从此开始。继续往后看吧，raft 状态机是怎么处理的。

```go
r.Step(pb.Message{From: r.id, Type: pb.MsgHup})
```



 **4**  **raft 状态机处理 campaign 开始**

raft 状态机怎么处理 MsgHup 消息？

位置在 raft.Step 公共部分：

```go
// etcd/raft/raft.go
func (r *raft) Step(m pb.Message) error {
    // ...
    switch m.Type {
    case pb.MsgHup:
        // 本地产生的竞争选举的消息，说明超时了，下一步就是要开启选举了
        if r.state != StateLeader {
            if r.preVote {
                // 开启预投（可配）
                r.campaign(campaignPreElection)
            } else {
                // 开启竞选
                r.campaign(campaignElection)
            }
        } else {
            r.logger.Debugf("%x ignoring MsgHup because already leader", r.id)
        }
    }
```

关键事件：

1. 如果开启了 preVote 开关，那么走预投的流程；
2. 如果没有开启 preVote ，那么直接开启选举；

