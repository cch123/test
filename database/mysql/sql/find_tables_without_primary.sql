SELECT
  table_schema, table_name
FROM tables
WHERE table_schema = 'test'
AND table_name NOT IN (SELECT
  table_name
  FROM TABLE_CONSTRAINTS
  WHERE table_schema = 'test'
  AND constraint_name = 'primary');



SELECT
    *
FROM
    information_schema.TABLES t
        LEFT JOIN
    information_schema.STATISTICS s ON t.table_schema = s.table_schema
        AND t.table_name = s.table_name
        AND s.index_name = 'PRIMARY'
WHERE
    t.table_schema NOT IN ('mysql' , 'performance_schema',
        'information_schema',
        'sys')
        AND table_type = 'BASE TABLE'
        AND s.index_name IS NULL;
