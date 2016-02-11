<?php
$tbl_src = "service_worksheet";
$tbl_dst = "g_service_worksheet";

$con = new mysqli("localhost", "root", "", "test");
$offset = 0;
while(true) {
    $res = $con->query("insert into g_service_worksheet (select * from service_worksheet limit {$offset}, 1000)");
    $offset += 100;
    print_r($res);
}


