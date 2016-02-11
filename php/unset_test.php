<?php
  $arr = [
      1 => ['sdfsdf'],
      'sss' => ['sz'],
      'x' => []
  ];
  print_r($arr);
  unset($arr[1]);
  unset($arr['sss']);
  unset($arr['x']);
  var_dump($arr);
