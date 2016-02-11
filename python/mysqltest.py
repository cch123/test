import pymysql.cursors

connection = pymysql.connect(host='localhost',
        user="",
        passwd="",
        db="",
        charset="utf8mb4",
        cursorclass=pymysql.cursors.DictCursor)

try:
    print 1
    with connection.cursor() as cursor:
        sql = """
                INSERT INTO `t_anti_products`
                (`email`, `password`)
                VALUES (%s, %s)
              """
        cursor.execute(sql, ('caochunhui@.com', 'nonono'))
    connection.commit()
    with connection.cursor() as cursor:
        # Read a single record
        sql = "SELECT `id`, `password` FROM `users` WHERE `email`=%s"
        cursor.execute(sql, ('caochunhui@.com',))
        result = cursor.fetchone()
        print result
finally:
    connection.close()

