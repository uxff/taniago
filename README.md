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
# datasource=root:password@tcp(127.0.0.1:3306)/beegoauth?charset=utf8

```

# How to Run

```
$ ./taniago --dir /data/your/exist/site/dir --addr :6699
```


# Describe and specific requirement

- make picset sites, link each other.
- make payment, pay for dirpath.
- save picset list to cache, to mysql.
- [important] get a face from sub, if no thumb.



