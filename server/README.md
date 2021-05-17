# Mall
> 服务器端


### 创建数据库 && 用户
```
create database `mall` charset=utf8mb4;
create user 'gomall'@'localhost' identified by '123456';
grant all privileges on `mall`.* to 'gomall'@'%' identified by '123456';
grant all privileges on `mall`.* to 'gomall'@'localhost' identified by '123456';
```