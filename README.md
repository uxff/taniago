# Feature

taniago is a light golang file explorer, writen by golang, use beego.


# Requirement

```
go get -u github.com/mattn/go-sqlite3
go get -u github.com/beego/i18n
go get -u github.com/mattn/go-runewidth
```

# How to Use

```
$ git clone git@github.com:uxff/taniago.git
$ cd taniago
#
# need node and npm
$ npm install -g bower
$ bower install
#
# build
$ go build
#
# you need start mysql service, and config mysql in:
$ vim conf/app.conf
# add line:
 datasource=root:password@tcp(127.0.0.1:3306)/beegoauth?charset=utf8

```

# How to Run

```
$ ./taniago --dir /data/your/exist/site/dir --addr :6699
```

# Preview
![](https://raw.githubusercontent.com/uxff/taniago/master/20181127073913.png)
![](https://raw.githubusercontent.com/uxff/taniago/master/20181127074015.png)


# Describe and specific requirement

- make picset sites, link each other.
- make payment, pay for dirpath.
- save picset list to cache, to mysql.
- [important] get a face from sub, if no thumb.

# static link to nginx
nginx config as below, the static jpg,png,... output via nginx. this make the go program light.
```
server {
listen 80;
server_name yourdomain.com;
access_log /data/logs/nginx/access.yourdomain.log combined;
error_log /data/logs/nginx/error.yourdomain.log;
index index.html index.htm index.jsp index.php;
include other.conf;
root /data/wwwroot/yourfiledir;
#error_page 404 /404.html;

location / {
    proxy_pass http://127.0.0.1:6699;
    proxy_set_header Host $host;
}
location /fs {
    alias /data/wwwroot/yourfiledir;
}

}
```

