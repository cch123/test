SELECT
  table_schema, table_name
FROM tables
WHERE table_schema = 'test'
AND table_name NOT IN (SELECT
  table_name
  FROM TABLE_CONSTRAINTS
  WHERE table_schema = 'test'
  AND constraint_name = 'primary');
