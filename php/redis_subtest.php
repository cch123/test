<?php
    $redis = new Redis();
    $redis->connect('127.0.0.1', 6379);
    $redis->subscribe(array("channel-1"), function($redis, $chan, $msg){
        echo $msg;
    });
