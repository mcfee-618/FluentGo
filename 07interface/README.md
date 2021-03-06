## 初识interface

在Go中，接口是一组方法签名(声明的是一组方法的集合)。当一个类型为接口中的所有方法提供定义时，它被称为实现该接口。它与oop非常相似。接口指定类型应具有的方法，类型决定如何实现这些方法。

```
    type Animal interface {
        Speak() string
    }
```

## interface{} 类型

interface{} 类型，空接口，是导致很多混淆的根源。interface{} 类型是没有方法的接口。由于没有 implements 关键字，所以所有类型都至少实现了 0 个方法，所以 所有类型都实现了空接口。这意味着，如果您编写一个函数以 interface{} 值作为参数，那么您可以为该函数提供任何值。interface{} 在我们需要存储任意类型的数值的时候相当有用，有点类似于C语言的void*类型。

```
1. interface 是方法声明的集合
2. 任何类型的对象实现了在interface 接口中声明的全部方法，则表明该类型实现了该接口。
3. interface 可以作为一种数据类型，实现了该接口的任何对象都可以给对应的接口类型变量赋值。
```

## 多态

接口的多种不同的实现方式即为多态，多态性是允许你将父对象设置成为一个或更多的他的子对象相等的技术，赋值之后，父对象就可以根据当前赋值给它的子对象的特性以不同的方式运作。