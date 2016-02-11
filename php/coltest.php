<?php
$arr = array(
    array(
        'order_no' => 2014,
        'value' => 100
    ), array(
        'order_no' => 2014,
        'value' => 130
    )
);
$key = array_column($arr, 'order_no');
$arr = array_combine($key, $arr);
print_r($arr);
