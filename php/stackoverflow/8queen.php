<?php
$arr = array(
    "C"  => array('A','A1','A2','B','B1'),
    "C1" => array('A','A1','A2','B','B1'),
    "A"  => array('B','B1','C','C1'),
    "B1" => array('A','A1','A2','C','C1'),
    "A2" => array('B','B1','C','C1'),
    "A1" => array('B','B1','C','C1')
);

$rec_arr = [];
$idx_arr = [];

function place_chess($val) {
}

function main($arr) {
    $start_arr = $arr["C"];
    foreach($start_arr as $key => $item) {
        $res_arr = [];
        $idx_arr = [];
    }
}
main($arr);

function check_valid($v) {
    if(in_array($v, $rec_arr)) {
        return FALSE;
    }
    return TRUE;
}

function pop_rec() {
    array_pop($rec_arr);
    return array_pop($idx_arr);
}
