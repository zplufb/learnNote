### [转] Git在CentOS下搭建私有的git服务器

#### 1 引言

在日常的项目开发中，我们可以傻瓜式的使用github进行代码托管，进而进行团队的协同开发。但是很多时候我们开发的代码并不是开源的（特别是涉及到公司的业务上），这时候在服务器上面部署git就可以很好地解决这个问题——既保证了团队开发，又能闭源达到代码托管。本文将以CentOS为例子，记录在CentOS下部署git服务器的具体步骤

#### 2 步骤

##### 2.1 服务器端

###### 2.1.1 配置git

- 从yum上安装git

  ```bash
  yum –y install git
  ```

- 查看git版本：安装完成后，在服务器端输入

  ```
  git --version
  [root@VM_0_12_centos fanbi]# git --version
  git version 1.8.3.1  
  ```

- 创建用户：在服务器中创建一个git专属用户

  ```
  //以下为CentOS下的用户，账户密码自定义
  useradd git
  passwd  git  
  ```

-  禁止该git用户使用shell登录系统（为了安全性，一般都禁止）

  ```bash
  vi /etc/passwd
  #按i进入编辑模式，在最后一行将git用户修改成以下配置
git:x:1000:1000::/home/git:/usr/bin/git-shell
  #按ESC退出插入模式，输入 “:wq” 保存并且退出vi模式 
  ```
  

###### 2.1.2 配置远程仓库

- 创建一个空仓库，我们选择在路径：cd  /home/ 下先创建一个用户目录，在用户目录下创建一个git仓库

  ```bash
  cd /home   
  mkdir git  
  cd git   
  git init --bare LearnProject.git  
  ```

  到这里，空仓库已经创建成功 ，仓库的路径为：/home/git/LearnProject.git

- 为刚刚创建的用户git赋予权限

  输入以下命令，为git用户赋予权限

  ```bash
  //chown -R 用户名:组名 文件
  chown -R git:git LearnProject.git
  ```

- 打开 RSA 认证

  进入/etc/ssh目录，编辑 sshd_config

  ```bash
  cd  /etc/ssh 
  vi ssh_config
  ```

   按i进入插入模式，打开以下三个配置的注释（带#为注释）, 按ESC退出插入模式，按:wq保存

  ```bash
  RSAAuthentication yes
  PubkeyAuthentication yes
  AuthorizedKeysFile .ssh/authorized_keys
  ```

  保存完成后，重启sshd服务

  ```bash
  /etc/rc.d/init.d/sshd restart
  ```

  如果运行上述命令出现以下错误，可以尝试以下命令：

  ```
  service sshd restart
  ```

  至此，服务器端git配置完成；

##### 2.2 客户端

###### 2.2.1 配置git

- 安装git

  - 在Linux上参考服务器端安装git；

  - 在Windows下直接下载git安装，下载地址：https://git-scm.com/downloads

- 打开git bash创建git使用者

  ```bash
  git config --global user.name "fanbi"
  git config --global user.email "fanbi@163.com"
  ```

- 生成rsa密钥

  可参考：[Gitlab的SSH配置（linux和windows双版本）](https://www.cnblogs.com/fanbi/p/7772812.html)

  在git bash中，输入以下命令，为自己的邮箱创建一个密钥

  ```bash
  ssh-keygen -t rsa -C fanbi@163.com -b 4096 
  ```
  
  此时，我们创建了一个公钥和密钥，id_rsa是密钥，不能告诉任何人，id_rsa.pub是公钥，可以公开

###### 2.2.2 将本地的公钥添加到服务器上

- 用命令复制id_rsa.pub

  在gitbash下，输入以下命令（ip地址换成你git服务器的地址）

  ```bash
   ssh git@172.16.100.123 'cat >> .ssh/authorized_keys' < ~/.ssh/id_rsa.pub
  ```

- 手动将公钥添加到服务器中

  - 在服务器端，我们用vi编辑器访问authorized_keys，将id_rsa.pub内容粘贴进其中

    ```bash
    vi /home/git/.ssh/authorized_keys
    ```

###### 2.2.3 将本地的公钥添加到Github/GitLab服务器上（可选）

- 复制id_rsa.pub到Github/GitLab等设置，SSH密钥中黏贴即可。

###### 2.4.4 克隆远程仓库

- 克隆远程仓库

  ```bash
  //端口在22的情况下
  git clone git@172.16.100.123:/home/git/LearnProject.git
  ```

至此，我们就完成了远程仓库的克隆。

#### 3 参考

[1] https://blog.csdn.net/bbcckkl/article/details/81634761

