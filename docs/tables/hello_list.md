# Table: hello_list

This table displays a hardcoded greeting.

## Examples

```
> select *, pg_typeof(json) from hello_list
+----+----------+-------------------+-----------+
| id | greeting | json              | pg_typeof |
+----+----------+-------------------+-----------+
| 1  | Hello    | {"hello":"world"} | jsonb     |
| 2  | Hello    | {"hello":"world"} | jsonb     |
| 3  | Hello    | {"hello":"world"} | jsonb     |
+----+----------+-------------------+-----------+

> select *, json->>'hello' as json_value from hello_list where id = 2
+----+----------+-------------------+------------+
| id | greeting | json              | json_value |
+----+----------+-------------------+------------+
| 2  | Hello    | {"hello":"world"} | world      |
+----+----------+-------------------+------------+
```

