# BUG

- NPC 点关闭会报错
- 貌似怪物死亡后尸体消失太快的问题跟方向有关

# TODO

- 客户端 MirCore-Client Common.UserItem 剩下的字段补全

# 介绍

`yams` = Yet Another Mir Server.

为什么有了 mirgo 还要另开一个坑呢？因为 mirgo 的目的是仿照原本 LOMCN 开源的 C# Cristal M2 实现大部分功能，为了让更多感兴趣的朋友能参与进来，很多代码都遵循着原本 C# 的代码逻辑方便别人阅读，加上很多功能比如组队、行会等需要开 2 个客户端调试非常麻烦且浪费时间，这些功能我写着写着感觉很没意思，就放弃继续写下去了…

yams 客户端还是用的 C#，但服务端策划、玩法会完全按照我自己的想法去修改。

# 参考

[MMORPG AI 系统设计与实现](https://gameinstitute.qq.com/course/detail/10097)
