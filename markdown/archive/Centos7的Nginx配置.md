# Centos7的Nginx配置

## 1 语法规则

## 2 Centos7上Nginx配置

### 2.1 nginx位置

通过whereis 命令查询可得

```
[fanbi@VM_0_12_centos ~]$ whereis nginx
nginx: /usr/bin/nginx /usr/sbin/nginx /usr/lib64/nginx /etc/nginx /usr/share/nginx /usr/share/man/man8/nginx.8.gz
[fanbi@VM_0_12_centos ~]$

```

- /etc/nginx是其安装目录

- /usr/sbin/nginx是其运行程序

- usr/share/nginx是默认读取位置，若没其它配置，则读取其下面的index.html文件
- 具体配置和宝塔Linux面板上Nginx配置是一样的

## 3 宝塔Linux面板上Nginx配置 

### 3.1 server

#### 3.1.1 HTTP

```
   listen 80;
      server_name unionline.top;
      index index.php index.html index.htm default.php default.htm default.html;
      root /www/wwwroot/unionline.top;
      
    location ~ .*\.(gif|jpg|jpeg|png|bmp|swf)$
    {
        expires      30d;
        error_log off;
        access_log off;
    }
    
    location ~ .*\.(js|css)?$
    {
        expires      12h;
        error_log off;
        access_log off; 
    }
    access_log  /www/wwwlogs/unionline.top.log;
    error_log  /www/wwwlogs/unionline.top.error.log;
```
#### 3.1.2 HTTPS

  ```
  # HTTPS server
  server {
    listen       443;
    server_name  unionline.top;
    root   /www/wwwroot/unionline.top;
    index  index.html index.htm;
    
    ssl on;
  
    ssl_certificate      1_www.unionline.top_bundle.crt;
    ssl_certificate_key  2_www.unionline.top.key;
  
    ssl_session_cache    shared:SSL:1m;
    ssl_session_timeout  5m;
  
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2; #按照这个协议配置
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE;#按照这个套件配置
    ssl_prefer_server_ciphers on;
    
    #ERROR-PAGE-START  错误页配置，可以注释、删除或修改
    error_page 404 /404.html;
    error_page 502 /502.html;
    #ERROR-PAGE-END
    
    #PHP-INFO-START  PHP引用配置，可以注释或修改
    include enable-php-56.conf;
    #PHP-INFO-END
  
  	#默认是blog为主页
    location / {
            proxy_pass http://182.254.200.134:4000/;
   			#root /your_project_dir/blog/myHexoBlog/public;
            location ~ .*\.(css|js)$ {
               root /your_project_dir/blog/myHexoBlog/public;
               if (-f $request_filename) {
                    expires 1d;
                    break;
                    }
                }
        }
  
   
  
   location /blog/ {
            proxy_pass https://182.254.200.134:4000/;
        }
  
   location /html/ {
           	root /www/wwwroot/unionline.top/;
			index index.html ;
    
           location ~ .*\.(gif|jpg|jpeg|png|bmp|swf|css|js|eot|svg|ttf|woff|woff2)$ {
                root /www/wwwroot/unionline.top/;
               if (-f $request_filename) {
                    expires 1d;
                    break;
                    }
                }
        }
  }
  
  ```

####   3.3.3 重定向

```
  {
      listen 80;
      large_client_header_buffers 4 128k;
  	  server_name www.unionline.top;

      location / {
      #开启对http1.1支持
      proxy_http_version 1.1;
      
      #设置Connection为空串,以禁止传递头部到后端
      #http1.0中默认值Connection: close
      proxy_set_header Connection "";
      
      proxy_pass http://unionline.top;
    }
```

#### 3.3.4 http跳转https

```
server
{
  # 80端口是http正常访问的接口
  listen 80;
  server_name _;
  # 在这里，我做了https全加密处理，在访问http的时候自动跳转到https
  rewrite ^(.*)$ https://$host$1 permanent; #这行是关键
}
```

### 3.2 http

```
略
```

## 4 问题

### 4.1 Nginx 一个服务器多域名配置 以及 访问php文件直接下载而不运行

include enable-php.conf; 这句话起到了关键作用

```php

location ~ [^/]\.php(/|$)
        {
            try_files $uri =404;
            fastcgi_pass  unix:/tmp/php-cgi.sock;
            fastcgi_index index.php;
            include fastcgi.conf;
        }
```

### 4.2 403错误

通过查看对应的error.log，主要原因有

1. 项目路径无index.html index.php等等文件
2. 项目js/css等路径不正确
3. 权限不够
   - /www/server/nginx/conf/nginx.conf 中user拥有者和组别分别是www www，而项目是root或者fanbi的拥有着，改成user root root即可
   - 通过宝塔面包新建的项目默认用户是www，组别是www

## 5 参考

[1. 四种解决Nginx出现403 forbidden 报错的方法](https://blog.csdn.net/m0_37928917/article/details/88796807)

