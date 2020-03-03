# Nginx配置

## 1 语法规则

## 2 Centos7上Nginx配置

## 3 宝塔Linux面板上Nginx配置 

### 3.1 server

HTTP

```
   listen 80;
      server_name unionline.top;
      index index.php index.html index.htm default.php default.htm default.html;
      root /www/wwwroot/unionline.top;
```
HTTPS
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
  
    location /html/ {
      root   /www/wwwroot/unionline.top;
      index  index.html index.htm;
    }
    
    location /bbs/ {
          root   /www/wwwroot/bbs.unionline.top;
          index  index.php index.html index.htm;
      }
  
    
    #静态文件，nginx自己处理
    location ~ ^/(images|javascript|js|css|flash|media|static)/ {
        #过期30天，静态文件不怎么更新，过期可以设大一点，
        #如果频繁更新，则可以设置得小一点。
        expires 30d;
    }
  }
  
  ```

  重定向

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

### 3.2 http

```
略
```

## 4 问题

4.1 Nginx 一个服务器多域名配置 以及 访问php文件直接下载而不运行

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

