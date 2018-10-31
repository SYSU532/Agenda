#Agenda

An command line agenda management app written with golang.

Supported Golang version:

- 1.9.x
- 1.10.x
- 1.11.x

### 代码文件结构

程序的主体位于 cmd，entity 和 log 文件夹中，其中 cmd 中的代码为程序各命令的代码，entity 中为程序操作数据库的逻辑代码，SQL 命令以及对当前登录用户的记录逻辑，log 中为程序进行日志保存的代码。

### 数据存储

程序产生的数据保存在程序执行目录的 data 文件夹中，其中包括三个文件：

- agenda.db 为 Sqlite 数据库保存的数据
- agenda.log 为 agenda 运行过程中的日志内容
- curUser.txt 保存了当前登陆的用户的用户名和使用 base64 编码后的密码。

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



