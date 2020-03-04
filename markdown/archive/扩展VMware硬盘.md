### Vm虚拟机扩展Ubuntu系统磁盘空间

#### 1 环境

1. VMware版本号：15.0.2 build-10952284
2. 系统：Ubuntu18.04

#### 2 步骤

2.1 打开VMWare虚机设置，点击扩展，输入扩展后的大小，确定

2.2 启动Ubuntu系统

2.3 下载gparted，可视化操作

```
sudo apt-get install gparted
```

2.4 启动gparted

```
sudo gparted
```

2.5 删除Ubuntu的分区

删完分区后，会得到未分配空间

![1583054405(1)](E:\work\learn\git\zplufb\learnNote\markdown\archive\images\1583054405(1).jpg)



![微信截图_20200301172102](E:\work\learn\git\zplufb\learnNote\markdown\archive\images\微信截图_20200301172102.png)

2. 6 扩展大小

点击dev/sda1，右键Resize/Move，然后把新添加的未分配大小全部分给该区，建议给相邻的盘，这边是Others（I）盘

![微信截图_20200301172358](E:\work\learn\git\zplufb\learnNote\markdown\archive\images\微信截图_20200301172358.png)

![微信截图_20200301172546](E:\work\learn\git\zplufb\learnNote\markdown\archive\images\微信截图_20200301172546.png)

![微信截图_20200301172631](E:\work\learn\git\zplufb\learnNote\markdown\archive\images\微信截图_20200301172631.png)

2.7 重新MBR（可选）

2.7.1 建议先备份分区表

![微信截图_20200301172200](E:\work\learn\git\zplufb\learnNote\markdown\archive\images\微信截图_20200301172200.png)

![微信截图_20200301172220](E:\work\learn\git\zplufb\learnNote\markdown\archive\images\微信截图_20200301172220.png)

![微信截图_20200301173104](E:\work\learn\git\zplufb\learnNote\markdown\archive\images\微信截图_20200301173104.png)

2. 8 如果中间存在交换区

   先删除该交换区，然后按照2.6操作，具体见参考资料

3. 9删除启动项（等价于重建MBR引导）

用魔方优化大师或者其他类似软件（建议先备份，后删除Ubunut启动项）

![微信截图_20200301174528](E:\work\learn\git\zplufb\learnNote\markdown\archive\images\微信截图_20200301174528.png)

#### 3 参考资料

- [1]: https://blog.csdn.net/weixin_39510813/article/details/78387334	"Vm虚拟机扩展Ubuntu系统磁盘空间"
- [Vm虚拟机扩展Ubuntu系统磁盘空间](https://blog.csdn.net/weixin_39510813/article/details/78387334)


