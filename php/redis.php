<?php
$redis = new Redis();
$connRes = $redis->connect("127.0.0.1", 6379);
$setRes = $redis->set("a",1);
var_dump($connRes);
var_dump($setRes);
