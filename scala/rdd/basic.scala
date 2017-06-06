package com.cch.etl.v1

import org.apache.spark.{SparkConf, SparkContext}

object Parse {
    def main(args: Array[string]) : Unit = {
        val conf = new SparkConf().setAppName("CCH-SPARK")
        val sc = new SparkContext(conf)

        // 从数组创建 rdd
        val a = sc.parallelize(1 to 9, 3)

        // 从文件中读取内容创建 rdd
        val b = sc.textFile("rdd.txt")
    }
}

