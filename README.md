# go-battleax 战斧

"战斧"旨在提供go开发效率与运维效率的提升，主要分为三大部分。
* 第一部分是提供RPC通用的封装，包括拦截器、限流熔断、日志等公共基础设施封装。
* 第二部分提供HTTP相关的。
* 第三部分是希望可以提供独立、高效率、可复用的组件封装。结合go的编程思想对可复用的基础方法进行统一封装。面向开发者。

## 规划：

* gen脚手架生成工具
* 基础组件
    * 配置管理，热更新、多源配置中心管理
    * 并发控制
        * 安全的Go协程
        * WaitGroup和ErrorGroup
        * 协程池
        * MapReduce
        * 流式处理器
        * 发布者订阅者
        * 生产者消费者
        * 安全锁
        * 原子操作
    * 时间管理
        * 定时器
        * 基础时间格式转换
* 存储引擎
    * ES
    * 大数据
    * Mongo
    * Mysql
    * Redis
    * MQ
        * RocketMq
        * Kafka
        * Nsq
        * Puslar
* 日志
    * 日志级别
    * 存储源
        * 文件
        * 控制台
        * MQ
        * 其他
    * 切割方式
        * 大小
        * 时间
    * 时效管理
        * 大小
        * 时间