<?php
$arr = [
    ["name" => 'a', "val" => "1"],
    ['name' => 'a', "val" => '2']
];

$keys = array_column($arr, 'name');
$new_arr = array_combine($keys, $arr);
$new_arr = array_values($new_arr);
print_r($new_arr);
