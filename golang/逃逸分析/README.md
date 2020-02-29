参考：
* https://driverzhang.github.io/post/golang%E5%86%85%E5%AD%98%E5%88%86%E9%85%8D%E9%80%83%E9%80%B8%E5%88%86%E6%9E%90/
* https://cloud.tencent.com/developer/article/1165660

逃逸分析简介：
* go build -gcflags '-m'来观察逃逸。为了不让编译时自动内连函数，一般会加-l参数，最终为-gcflags '-m -l'
* 是确定指针动态范围的方法，简单说是确定变量在堆分配还是栈分配。