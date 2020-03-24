package main


//表1: Person
//+-------------+---------+
//| 列名         | 类型     |
//+-------------+---------+
//| PersonId    | int     |
//| FirstName   | varchar |
//| LastName    | varchar |
//+-------------+---------+
//PersonId 是上表主键


//表2: Address
//+-------------+---------+
//| 列名         | 类型    |
//+-------------+---------+
//| AddressId   | int     |
//| PersonId    | int     |
//| City        | varchar |
//| State       | varchar |
//+-------------+---------+
//AddressId 是上表主键

//要求： 编写一个 SQL 查询，满足条件：无论 person 是否有地址信息，都需要基于上述两表提供 person 的以下信息：
//FirstName, LastName, City, State

//答案：select FirstName, LastName, City, State from Person left join Address on Person.PersonId = Address.PersonId


//扩展：
//1.内连接
// 1）select 字段1,字段2... from 表1 inner join 表2 on 表1.公共字段 = 表2.公共字段
// 2) select 字段1,字段2... from 表1,表2 where 表1.公共字段 = 表2.公共字段

//2.外连接
// 1)左外连接： select 字段1,字段2... from 表1 left join 表2 on 表1.公共字段 = 表2.公共字段  (以左表为基准，如果右边的表没有数据，则用NULL表示)
// 2)右外连接： select 字段1,字段2... from 表1 right join 表2 on 表1.公共字段 = 表2.公共字段  (以右表为基准，如果左边的表没有数据，则用NULL表示)













