<?php
  $str ="{'mobile' : [{15810540851}], 'content' : 'dachuwang test'}";
  $arr = json_decode($str);
  $new_arr = [
      'mobile' => array(1,2),
      'content' => 'ssss'
  ];
  var_dump($arr);
  var_dump(json_encode($new_arr));
