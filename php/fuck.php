<?php
  $arr = [
      "user_id" => 1,
      "products" => [
          [
              "id" => 42,
              "quantity" => 3
          ],
          [
              "id" => 43,
              "quantity" => 2
          ]
      ]
  ];
echo json_encode($arr);
