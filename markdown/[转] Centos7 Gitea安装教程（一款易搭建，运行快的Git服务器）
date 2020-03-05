## Centos7 Gitea安装教程 - 一款易搭建，运行快的Git服务器 

### 1 引言

​	Gitea是从Gogs发展而来，同样的拥有极易安装，运行快速的特点，而且更新比Gogs频繁很多，维护的人也多，个人认为Gitea还是更好一些的，这里就说下安装方法。

​	Gitea是一个极易安装，运行非常快速，安装和使用体验良好的自建Git服务。采用Go作为后端语言，这使得只要生成一个可执行程序即可。并且他还支持跨平台，支持Linux、macOS和Windows以及各种架构，除了x86，amd64，还包括ARM和 PowerPC。

Github地址：https://github.com/go-gitea/gitea

### 2 功能

- 支持活动时间线
- 支持SSH以及HTTP/HTTPS协议
- 支持SMTP、LDAP和反向代理的用户认证
- 支持反向代理子路径
- 支持用户、组织和仓库管理系统
- 支持添加和删除仓库协作者
- 支持仓库和组织级别Web钩子（包括Slack集成）
- 支持仓库Git钩子和部署密钥
- 支持仓库工单（Issue）、合并请求（Pull Request）以及Wiki
- 支持迁移和镜像仓库以及它的Wiki
- 支持在线编辑仓库文件和Wiki
- 支持自定义源的Gravatar和Federated Avatar
- 支持邮件服务
- 支持后台管理面板
- 支持MySQL、PostgreSQL、SQLite3、MSSQL和TiDB（实验性支持）数据库
- 支持多语言本地化（21种语言）

### 3 安装

#### 3.1 安装`MySQL`/`Mariadb`数据库

```
安装完成后，安装MySQL，至少5.5.3版本。
```

#### 3.2 安装Git

```
#Debian和Ubuntu系统

apt-get -y install git

#CentOS系统

yum -y install git
```

#### 3.3 安装Gitea

最新版本下载地址：https://dl.gitea.io/gitea。

```
cd /usr/local/gitea

wget -O gitea https://dl.gitea.io/gitea/1.6.0/gitea-1.6.0-linux-amd64

chmod +x gitea

./gitea web
```

 接下来打开[http://ip](http://ip/):3000即可。

#### 3.4、域名访问

如果想用域名访问，可以用Nginx反代。反代配置为：

```
#在配置文件里添加

location / {

proxy_pass http://localhost:3000

proxy_redirect off;

proxy_set_header X-Real-IP $remote_addr;

proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

}
```

#### 3.5 使用服务来启动

新建一个rclone.service文件：

```
vi /usr/lib/systemd/system/gitea.service
```

写入：

```
[Unit]
Description=gitea

[Service]
User=root
ExecStart=/usr/local/gitea/gitea
Restart=on-abort

[Install]
WantedBy=multi-user.target
```

 重载daemon，让新的服务文件生效：

```
systemctl daemon-reload
```

现在就可以用systemctl来启动gitea了：

```
systemctl start gitea
```

设置开机启动：

```
systemctl enable gitea
```

停止、查看状态可以用：

```
systemctl stop gitea

systemctl status gitea
```

 接下来就是打开网址去初始化gitea配置。

### 4 参考

[1. Centos7 Gitea安装教程 - 一款易搭建，运行快的Git服务器](https://my.oschina.net/soben/blog/2966518)