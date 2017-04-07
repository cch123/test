SELECT 
	CONCAT(t.TABLE_SCHEMA,'.',t.TABLE_NAME) table_name,INDEX_NAME, CARDINALITY, 
    TABLE_ROWS, CARDINALITY/TABLE_ROWS AS SELECTIVITY
FROM
    information_schema.TABLES t,
 (
  SELECT table_schema,table_name,index_name,cardinality
  FROM information_schema.STATISTICS 
  WHERE seq_in_index=1 
 ) s
WHERE
    t.table_schema = s.table_schema 
        AND t.table_name = s.table_name AND t.table_rows != 0
        AND t.table_schema NOT IN ( 'mysql','performance_schema','information_schema','sys'
