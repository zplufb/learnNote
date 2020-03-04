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

2.5 扩展大小

点击dev/sda1，右键Resize/Move，然后把新添加的未分配大小全部分给该区

2.6 如果中间存在交换区

先删除该交换区，然后按照2.5操作，具体见参考资料

#### 3 参考资料

- [1]: https://blog.csdn.net/weixin_39510813/article/details/78387334	"Vm虚拟机扩展Ubuntu系统磁盘空间"
- [Vm虚拟机扩展Ubuntu系统磁盘空间](https://blog.csdn.net/weixin_39510813/article/details/78387334)


