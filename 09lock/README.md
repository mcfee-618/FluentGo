## 互斥锁

互斥锁用 var mutex sync.Mutex 申明 ，通过 mutex.Lock() 进行加锁，通过 mutex.Unlock() 进行解锁。需要注意的是，互斥锁一定要记得解锁，否则很容易造成死锁。所以，一般会把 mutex.Unlock() 放到 defer 语句中，避免资源无法释放导致死锁。

使用锁的使用大概有如下两种使用方式：
1.全局变量
2.指针类型【推荐后者】

## 读写锁

go 语言读写锁一般用于读多写少的场景。go 读写锁遵循两个基本原则 ：

(1) 读的时候资源不是互斥的，可以多个 goroutine 一起读，也可以加多个读锁。
(2) 写的时候只能有一个 goroutine 进行写，写锁只能有一个，其他 goroutine 不能读也不能写

go 语言的读写锁的实现是 sync.RWMutex，它主要提供了下面 4 个方法：

func (rw *RWMutex) Lock() 写锁，

func (rw *RWMutex) Unlock() 写锁解锁

func (rw *RWMutex) RLock() 读锁

func (rw *RWMutex)RUnlock()　读锁解锁

## 原子操作CAS

原子指令是cpu执行的最小粒度，在执行过程中不会被中断。go 语言的原子操作都在 sync/atomic 标准库代码包下，这里我们只介绍最常用的原子增减和 CAS 操作。

CAS 全称是 “Compare And Swap”（比较并交换），是采用无锁化的思想实现变量的同步。 CAS 算法涉及到 3 个操作数：内存值 V、旧的预期值 A、要修改的新值 B，只有当旧的预期值 A 与内存值 V 相同时，将内存值 V 修改成新的值 B。这里一般是通过自旋去实现的。

 ```
void lock(lock_t *lock) {
    while (CompareAndSwap(&lock->flag, 0, 1) == 1)
        ; // spin
    }
 }
 ```
CAS 的实现过程中没有采用锁，也就不会创建互斥量和临界区。没有互斥资源，可以大大减少同步对程序性能的消耗。但如果 CAS 操作的值被频繁变更，CAS 操作可能不是非常容易成功。所以这里会使用 for 循环来进行多次尝试（自旋）。



