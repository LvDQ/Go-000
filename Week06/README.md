学习笔记

1. 参考 Hystrix 实现一个滑动窗口计数器。


ref: `https://www.cnblogs.com/li-peng/p/11050563.html`





微服务 

两大问题 

1. 数据一致性
2. 系统规模，可用性


- 隔离
服务隔离：
动静分离、读写分离
		CQRS 读写分离

物理隔离：
进程隔离&&集群隔离

热点隔离：
小表广播 && 主动与热


time after  time ticket 可以做定时


Cgroup可以做资源隔离



计算密集型 bcrypt
https://cch123.github.io/perf_opt/
这里面搜一下bcrypt


keepalive 默认75s左右发一次，很蠢



http有timeout  有超时空之


gettimeofday => VDSO
vdso https://blog.csdn.net/juana1/article/details/6904932

go的时间 monotonic  time.since


官方令牌桶算法：
https://pkg.go.dev/golang.org/x/time/rate

过载保护：
常见做法：李特尔法则 L=λW
CPU 内存作为信号量进行节流
最大吞吐接近于最大阈值（比如cpu压力90%）时  QPS*Latency  

可控延迟算法：CoDel


多纬度资源分配：
最大最小公平分享

（max-min）/ (max+min)

DRF  <- Fair Schedule


对于key对应的hashing热点
consistent hashing with bounded loads
HAproxy商业版有实现

——————————————————————————————————————————————————————————
熔断

Google SRE 熔断 

熔断丢弃数量：
max（0，（requests-K*accepts)/(requests+1))

能保证最大处理容量 而且极限趋近于accept*1倍，没有半开半闭

————————————————————————————————————————————————————————————————————
限流：客戶端留空 

退讓算法 



——————————————————————————————————————————————————————————————

logging
logging

https://dave.cheney.net/2015/11/05/lets-talk-about-logging
https://www.ardanlabs.com/blog/2013/11/using-log-package-in-go.html
https://www.ardanlabs.com/blog/2017/05/design-philosophy-on-logging.html
https://dave.cheney.net/2017/01/23/the-package-level-logger-anti-pattern

——————————————————————————————————————————————————————————————————

wire 主要用來解決依賴注入

orm 框架： ent