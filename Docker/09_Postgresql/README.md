https://medium.com/analytics-vidhya/getting-started-with-postgresql-using-docker-compose-34d6b808c47c
```
docker exec -it 09_postgresql_database_1 bash
psql --host=database --username=unicorn_user --dbname=rainbow_database
输入密码 magical_password
```

```
rainbow_database=# CREATE TABLE color_table(name TEXT);
CREATE TABLE
rainbow_database=# \d
              List of relations
 Schema |    Name     | Type  |    Owner     
--------+-------------+-------+--------------
 public | color_table | table | unicorn_user
(1 row)

rainbow_database=# SELECT * FROM color_table;
 name 
------
(0 rows)

rainbow_database=# INSERT INTO color_table VALUES ('pink');
INSERT 0 1
rainbow_database=# SELECT * FROM color_table;
 name 
------
 pink
(1 row)
```