# go-design-patterns
Golang Design Patterns

## 前言

从1995年`Erich Gamma`等人提出23种设计模式到现在，27年过去了，设计模式依旧是软件领域的热门话题。

>设计模式（Design Pattern）是一套被反复使用、多数人知晓的、经过分类编目的、代码设计经验的总结，使用设计模式是为了可重用代码、让代码更容易被他人理解并且保证代码可靠性。

从定义上看，设计模式其实是一种经验的总结，是针对特定问题的简洁而优雅的解决方案。既然是经验总结，那么学习设计模式最直接的好处就在于可以站在巨人的肩膀上解决软件开发过程中的一些特定问题。然而，学习设计模式的最高境界是习得其中解决问题所用到的思想，当你把它们的本质思想吃透了，也就能做到即使已经忘掉某个设计模式的名称和结构，也能在解决特定问题时信手拈来。

## 创建模式

- [建造者模式(Builder Pattern)](/01-builder-pattern)
    >将一个复杂对象的构建与它的表示分离, 使得同样的构建过程可以创建不同的表示

- [单例模式(Singleton Pattern)](/02-singleton-pattern)
    >单例模式是最简单的设计模式之一, 保证一个类仅有一个实例, 并提供一个全局的访问接口

- [工厂方法模式(Factory Method Pattern)](/03-factory-method-pattern)
    >使一个类的实例化延迟到其子类, 定义一个用于创建对象的接口, 让子类决定将哪一个类实例化

- [对象池模式(Object Pool Pattern)](/04-object-pool-pattern)
    >根据需求将预测的对象保存到channel中， 用于对象的生成成本大于维持成本

- [生成器模式(Generator Pattern)](/05-generator-pattern)
    >生成器模式可以允许使用者在生成要使用的下一个值时与生成器并行运行

- [抽象工厂模式(Abstract Factory Pattern)](/06-abstract-pattern)
    >提供一个创建一系列相关或相互依赖对象的接口, 而无需指定它们具体的类

- [原型模式(Prototype Pattern)](/07-prototype-pattern)
    >复制一个已存在的实例