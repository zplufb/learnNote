### Linux命令

##### 1 Common

- 常用
  - cp
  - mv
  - ls
  - cat
  - less
  - more
  - tail
  - rm
  - rmdir
  - find
  - grep
  - su
  - sudo
  - netstat 
    - netstat -nplt #查看系统后台跑的程序
  - chmod #修改权限
  - nohup
  - systemctl
  - service
- 不常用
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
- 软件
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

##### 2 Ubuntu特有

- apt-get

- apt

##### 3 Centos

- yum