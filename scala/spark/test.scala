import org.apache.spark.sql.hive.HiveContext

val hiveCtx = new HiveContext(sc)

val studentRDD = hiveCtx.sql("select * from test.user").rdd

studentRDD

