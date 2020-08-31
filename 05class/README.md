## 异常处理

go 语言中不采用这种异常捕获机制，而是通过下面几种方式进行异常处理。

* error接口：error 类型是一个接口类型，定义如下

```
    type error interface {
        Error() string
    }
```
一般来说， error 是 go 语言中最常见的处理错误的方式，通过返回 error，处理 error， 产生类似 exception 的效果，大部分接口时这样设计的：

```
    func Sqrt(f float64) (float64, error) {
        if f < 0 {
            return 0, errors.New("math: square root of negative number")
        }
        // 实现
    }
    result, err:= Sqrt(-1)

    if err != nil {
    fmt.Println(err)
    }
```
* defer语句：实现类似finally块的功能，可以使用关键字 defer 向函数注册退出调用，即主调函数退出时，defer后的函数才会被调用，defer 语句的作用是不管程序是否出现异常，均在函数退出时自动执行相关代码。（相当于 finally 代码块）。

即当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的（defer栈），当函数执行完毕后，再从defer栈，按照（先入后出）的方式出栈执行。


## panic-recover机制

panic-recover机制：error 只能针对预期内的错误，因为你是预判这段程序可能出现异常逻辑，才会去主动调用 errors.New () 生成一个 error 。但是对于一个方法来说，我们不可能预判到所有的异常情况，那假如某一个隐藏 bug 导致程序崩溃了怎么办呢？这里就需要引入 panic-recover机制了。假如代码运行时异常崩溃了，此时 go 会自动 panic，go 的每次 panic 都是非常消耗性能的，且 go 是单线程，所以，我们应该尽量去避免使用 panic。

* Panic：Panic是内建的停止控制流的函数。相当于其他编程语言的抛异常操作。当函数F调用了panic，F的执行会被停止，在F中panic前面定义的defer操作都会被执行，然后F函数返回。对于调用者来说，调用F的行为就像调用panic（如果F函数内部没有把panic recover掉）。如果都没有捕获该panic，相当于一层层panic，程序将会crash。panic可以直接调用，也可以是程序运行时错误导致，例如数组越界。

* Recover：Recover是一个从panic恢复的内建函数。Recover只有在defer的函数里面才能发挥真正的作用。如果是正常的情况（没有发生panic），调用recover将会返回nil并且没有任何影响。如果当前的goroutine panic了，recover的调用将会捕获到panic的值，并且恢复正常执行，一般的调用建议如下：
    * 在defer函数中，通过recever来终止一个gojroutine的panicking过程，从而恢复正常代码的执行，recover 仅在延迟函数中有效。
    * 可以获取通过panic传递的error

简单来讲：go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

一般defer recover这种机制经常用在常驻进程的应用，比如Web服务，在Go里面，每一个Web请求都会分配一个goroutine去处理，在没有做任何处理的情况下，假如某一个请求发生了panic，就会导致整个服务挂掉，这是不可接受的，所以在Web应用里面必须使用recover保证即使某一个请求发生错误也不影响其它请求.

## 面向对象-访问控制

Go语言没有像其它语言一样有public、protected、private等访问控制修饰符，它是通过字母大小写来控制可见性的，如果定义的常量、变量、类型、接口、结构、函数等的名称是大写字母开头表示能被其它包访问或调用（相当于public），非大写开头就只能在包内使用（相当于private，变量或常量也可以下划线开头）。


## 继承

其他面向对象的语言，是使用extends关键字来表示继承的，而go语言中继承的设计非常简单，不需要额外的关键字。go 语言可以通过结构体之间的组合来实现类似继承的效果。结构体嵌入多个匿名结构体（多重继承），如果两个匿名结构体有相同的字段和方法（同时结构体本身没有相同的字段和方法），在访问时，就必须指明匿名结构体名字。

## 多态

在 go 语言中，只要某个 struct 实现了某个 interface 的所有方法，那么我们就认为这个 struct 实现了这个类，这个类似python中的鸭子类型。

