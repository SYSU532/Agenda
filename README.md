# Agenda
[![Build Status](https://travis-ci.org/SYSU532/agenda.svg?branch=master)](https://travis-ci.org/SYSU532/agenda)

An command line agenda management app written with golang.

Supported Golang version:

- 1.9.x
- 1.10.x
- 1.11.x

### agenda安装与使用
安装指令：
```
$ go get -u github.com/SYSU532/agenda
```
查看帮助信息，以及可以使用的命令：
```
$ agenda -h
```

### 代码文件架构

* `agenda`程序代码的主体位于 `cmd`，`entity` 和 `log`文件夹中。

* 其中 `cmd` 中的代码为程序各命令的代码，充当命令执行前端功能，接收终端输入信息，将相关信息交接给后台数据库操作模块处理，打印相关输出信息。

* `entity`中为程序操作数据库的逻辑代码，划分为初始化数据库表结构的`sql.go`，专用于增删查改数据库的后台代码文件的`dbControl.go`，负责更新当前用户信息的`userinfo.go`。

* `log` 中为程序进行日志保存的代码，此处使用Go语言官方`log`模块，采用`go format`记录用户操作日志信息。

### 数据存储

程序产生的数据保存在程序执行目录的 data 文件夹中，其中包括三个文件：

- `agenda.db` 为 `Sqlite3` 数据库保存的数据。
- `agenda.log` 为 `agenda` 运行过程中的日志内容，以不同前缀名划分普通日志与错误日志，记录日志打印时间，
  利于查看相关操作对应的日志信息。
- `curUser.txt` 以`JSON`格式，保存了当前登陆的用户的用户名和使用 `base64` 编码后的密码。

### 命令列表

| Command      | Description                                           |
| ------------ | ----------------------------------------------------- |
| `register`   | 注册用户                                              |
| `login`      | 用户登录                                              |
| `logout`     | 登出当前用户                                          |
| `userlist`   | 列出当前存在的用户                                    |
| `userdelete` | 删除当前用户                                          |
| `cm`         | 以当前用户为创建者创建会议                            |
| `fm`         | 使用条件查找会议                                      |
| `addp`       | 添加会议参与者                                        |
| `rp`         | 从指定会议中移除指定参与者，当参与者数量为0时删除会议 |
| `quitm`      | 从作为参与者的指定会议中退出                          |
| `cancelm`    | 作为创建者取消创建的会议                              |
| `clearm`     | 取消所有作为创建者的会议                              |


### 命令具体用法
1. `register [-uUserName][–pPassword][–email=a@xxx.com][-oXXXXXXXXXXX]`

   注册用户

   可选参数flag 如下：

   - `-u`参数接收注册用户的用户名
   - `-p` 参数接收注册用户的密码
   - `-e` 参数接收注册用户的邮箱
   - `-o` 参数接收注册用户的电话

2. `login [-u username][-p password]`

   用户登录

   可选参数flag 如下：

   - `-u`参数接收登录用户的用户名
   - `-p` 参数接收登录用户的密码

3. `logout`

   登出当前用户，无参数

4. `userlist [-uUserName][-eEmail]`

   列出当前存在的用户

   可选参数flag 如下：

   - `-u`参数接收查找的用户名
   - `-e` 接收查找的邮箱
   - 用户只允许在登录状态下查看当前所有用户

5. `userdelete`

   - 删除当前用户，无参数

6. `cm [-t title][-p participator1] [-p participator2]` 

   以当前用户为创建者创建会议

   接收参数如下：

   - `-t` 后接创建的会议的标题
   - `-p` 后接创建的会议的参与者，可使用多个`-p` 来指定多个参与者

7. `fm [-t title][-s startTime] [-e endTime]` 

   使用条件查找会议

   参数如下：

   - `-t` 后接查找的会议的标题
   - `-s` 后接查找的会议的开始时间
   - `-t` 后接查找的会议的结束时间

8. `addp [-t title][-p participator1] [-p participator2] ...` 

   添加会议参与者

   参数如下：

   - `-t` 为需添加参与者的会议标题
   - `-p` 为要添加的参与者用户名，可使用多个以同时添加多个参与者

9. `rp [-t title][-p participator1][-p participator2] ...`

   从指定会议中移除指定参与者，当参与者数量为0时删除会议

   参数如下：

   参数如下：

   - `-t` 为需移除参与者的会议标题
   - `-p` 为要移除的参与者用户名，可使用多个以同时移除多个参与者

10. `quitm [-t title]`

    从作为参与者的指定会议中退出

    参数`-t` 为会议标题

11. `cancelm [-t title]`

    作为创建者取消创建的会议

    参数`-t` 为会议标题

12. `clearm` 

    取消所有作为创建者的会议

    无参数

### 特色功能
- `agenda`可用命令参数的输入不仅仅可以使用`pflag`格式输入，每当用户忘记输入关键参数时，会有终端提示再次输入，以`Stdin`接受输入参数，正确执行相关命令。

- `agenda`使用数据库`Sqlite3`持久化存储用户信息与会议信息，开放相应表结构存储实体信息，相关内容存放在`data`文件夹内部的`agenda.db`，可以通过安装`sqlitebrowser`可视化查看DB文件内部信息。
	* 安装`sqlitebrowser`命令：
	```
	$ sudo apt-get install sqlitebrowser
	```
	* 安装完成之后，注意本数据库同时只能与一个应用程序进行链接，保证操作的安全性，所以在使用可视化工具时，请在查看之后关闭与数据库的链接，才能正常进行`agenda`的使用。