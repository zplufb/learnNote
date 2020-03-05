### Linux命令

#### 1 Common

##### 1.1 文件目录管理

###### 1.1.1 文件路径

- pwd #显示当前项目路径，print working directory简称

- cd
  - cd .. #切换上一级目录
  - cd ~ #切换到当前登录的Home目录
  - cd . #当前目录，相当于没切换
  - cd ~user #切换到指定用户的家目录，一般使用于root用户
- ls
  - ls -a #查看所有
  - ls -l #列表形式
  - ll = ls -al
  - ls -R \#递归显示。通常用于目录

###### 1.1 2 查看文件

- touch   新建空文件，也可以用来修改文件的时间戳（修改文件为当前的时间）

- stat   命令查看文件的信息

- cat #查看文件内容，concatenate  files的简称

  - cat -b #输出标准行号，不忽略空格

  - cat -n #输出标准行号，忽略空格

- head  #查看文件的前几行，默认十行 

  - head -n 5 #显示文件的前5行      

- tail   #查看文件的尾部几行，默认十行

  - tail -n 5 #显示文件后5行

- less   \#分页浏览查看文件

  - less 查看文件是，输入/可以搜索关键字

- more #显示百分比

- wc #统计文件行数，字符数，大小

###### 1.1.3 文件操作 

- cp
  - -r   \#递归复制目录
  - -p  \#复制权限
  - -v  \#显示复制过程中的详细信息

- mv #移动文件/目录或者重命名文件/目录
- rm
  -  -f #强制删除，不进行提醒 
  -  -r #递归删除目录
  -  -d #删除空目录
  - -rf #删除目录及以下文件

###### 1.1.4 文件夹操作

- mkdir #创建目录
  - -p 创建多级目录

- rmdir #删除空目录

##### 1.2 查找

- find
  - 
- grep

##### 1.3 管理

- su
  - su username＃切换用户名
- sudo
  - sudo su #切换到root账号
- netstat 
  - netstat -nplt #查看系统后台跑的程序
- kill
- chmod #修改权限
- nohup
- systemctl
  - systemctl start #启动服务
  - systemctl enable #设置开机启动
  - systemctl stop #关闭服务
  - systemctl Status #查看服务状态
  - systemctl daemon-reload #重载daemon，让新的服务文件生效
- service

##### 1.4 不常用

- alias
- du
- df
- uname
- whereis
  - whereis 文件名
- unzip
- tar
- nohup
- curl
- chgrp
  - chgrp -R fanbi filenamesOrDir #修改群组为fanbi
- chown
  - chown -R fanbi:fanbi filenamesOrIDir #修改文件拥有者和组别为fanbi fanbi
- useradd
- groupadd
- ps

##### 1.5 软件

- nginx
  - nginx -t #测试配置文件是否正确
  - nginx -s stop/reload
  - nginx #启动
- vim
  - i/a/o/O #插入/追加/下一行/上一行
  - r/R #替换一个/替换多个并覆盖
  - x #删除单个字符
  - dd #删除一行
  - hjkl #左下上右
- gitea
  - 配置说明：https://docs.gitea.io/zh-cn/config-cheat-sheet/

#### 2 Ubuntu特有

- apt-get

- apt

#### 3 Centos

- yum
  - yum install git #安装git
  - yum remove git #卸载git