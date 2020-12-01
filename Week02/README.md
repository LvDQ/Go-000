学习笔记

Week02 作业题目：
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


感觉其实上层不关心啊，ErrNoRows这种错误总不至于再追底层咋回事儿，拉出来单练好了


使用方法：

``` shell
# 未知错误，报堆栈
go run main.go 

# ErrNoRow
go run main.go ErrNoRows
```