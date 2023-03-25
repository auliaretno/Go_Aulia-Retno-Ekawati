--1
HGETALL users 
MATCH (u:users) --redis

--2
RETURN u.* --neo4j

--3
SELECT * FROM users; --cassandra