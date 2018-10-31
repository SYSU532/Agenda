Agenda 的各命令与参数简介如下：

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

5. `userdelete`

   删除当前用户，无参数

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











