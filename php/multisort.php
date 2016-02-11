<?php
  $arr = [
      [
          'name' => 'xlex', 'age' => 44
      ],
      [
          'name' => 'aaa', 'age' => 55
      ]
  ];

  $names = array_column($arr, 'name');
  $ages = array_column($arr,'age');
  array_multisort($ages, SORT_ASC, $arr);
  print_r($arr);
  array_multisort($names, SORT_ASC, $arr);
  print_r($arr);
