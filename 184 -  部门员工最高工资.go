/*
Employee 表包含所有员工信息，
每个员工有其对应的Id, salary 和 department Id。
+----+-------+--------+--------------+
| Id | Name  | Salary | DepartmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 70000  | 1            |
| 2  | Henry | 80000  | 2            |
| 3  | Sam   | 60000  | 2            |
| 4  | Max   | 90000  | 1            |
+----+-------+--------+--------------+

Department 表包含公司所有部门的信息。
+----+----------+
| Id | Name     |
+----+----------+
| 1  | IT       |
| 2  | Sales    |
+----+----------+

*/


select 
    Department.name AS "Department",
    Employee.name AS "Employee",
    Salary
from 
    Employee  
    JOIN
    Department ON Employee.DepartmentId = Department.Id
where
    (Employee.DepartmentId,Salary) IN
    (select 
           DepartmentId,MAX(Salary)
    from
           Employee
    GROUP BY DepartmentId
)