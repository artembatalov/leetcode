SELECT email
FROM Person as P
GROUP BY email
HAVING COUNT(*) > 1;
