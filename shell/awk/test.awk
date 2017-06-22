{
    x = $1 % 1000
    #print x
    driver_id = $1
    order_id = $2
    print "update db_name.table_name"driver_id % 1000 , "set star = 5 where order_id = " driver_id ";"
    #print order_id
}

