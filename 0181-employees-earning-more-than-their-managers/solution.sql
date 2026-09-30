SELECT Current.name AS Employee
FROM Employee AS Current
LEFT JOIN Employee AS Added ON Current.managerId = Added.id
WHERE Added.salary < Current.salary;
