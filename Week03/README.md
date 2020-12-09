学习笔记

Week03 作业题目：
1.基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。



学习笔记

select {} 可以做阻塞主线程用


log.Fatal(err)里面会执行os.exit()
os.exit()会造成defer无法被执行到

所以log.Fatal最好还是在main函数里使用

————————————————————————————————————————————————————————

profiling 抓火焰图


————————————————————————————————————————————————————————

Goroutine 小原则：
1.调用者来决定要不要在后台执行

2.goroutine 的生命周期应该是要被管理的   可以用channel发信号的方式来优雅退出管理

3.要知道什么时候结束，能管控goroutine生命周期以及有手段让其退出


—————————————————————————————————————————————————————————

chan的用法：

i,ok <- chan

for data：= range chan{
	
}

select{
	case t.ch<-data:
		return nil
	case <-ctx.Done():
	return ctx.Err()"
}
这种用法如果销毁chan 则可以跳出循环



——————————————————————————————————————————————————————
作业
https://golang.org/ref/mem


——————————————————————————————————————————————————————

go 同步语义操作

Memory model

www.jianshu.com/p/5e44168f47a3


——————————————————————————————————————————————————————
MESI 协议 intel


——————————————————————————————————————————————————————

nginx 源码 不建议通篇阅读

红黑树





________________________________________________


看汇编 go tool compile -S file.go
压测 go test bench 、


atomic.Value
redis  BGSave Cow  

copy on Write   读多写少 面试必考  atomic value
 拷贝完整数据放入
```
func main() {
  var config atomic.Value
  config.Store(loadConfig())
  go func() {
    for {
      time.Sleep(time.Minute)
      config.Store(loadConfig()) // 写
    }
  }()
  for i := 0; i < 10; i++ {
    go func() {
      for r := range requests() {
        c := config.Load() // 读
        // do something
      }
    }()
  }
}
```



 ____________________________________________________

 go interface nil 坑



 ————————————————————————————————————————————————————
 fan-in  fan-out


 ————————————————————————————————————————————————————


 面试必考  master workers