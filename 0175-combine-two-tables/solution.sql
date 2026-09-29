SELECT Person.firstName AS firstName, Person.lastName AS lastName, Address.city as city, Address.state as state
FROM Person AS Person
LEFT JOIN Address AS Address ON Person.personId = Address.personId;
