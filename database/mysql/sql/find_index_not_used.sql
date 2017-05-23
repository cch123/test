5.6:
performance_schem
SELECT
  object_schema,
  object_name,
  index_name
FROM table_io_waits_summary_by_index_usage
WHERE SUM_TIMER_WAIT = 0
AND SUM_TIMER_WRITE = 0
AND SUM_TIMER_INSERT = 0
AND SUM_TIMER_INSERT = 0
AND SUM_TIMER_UPDATE = 0
AND SUM_TIMER_DELETE = 0;

5.7:
sys
select * from schema_unused_indexes;
