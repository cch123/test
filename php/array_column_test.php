<?php
  $arr = [
      [
          'id' => 'a',
          'quantity' => 2
      ],
      [
          'id' => 'b',
          'quantity' => 3
      ],
      [
          'id' => 'c',
          'quantity' => 5
      ]
  ];
  $x = array_column($arr, 'id', 'quantity');
  print_r($x);

