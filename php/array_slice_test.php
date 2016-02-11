<?php
   $a = 1;
   $arr = array(
       $a,
       $a,
       $a,
       $a,
       $a,
       $a,
   );
   print_r($arr);
   $arr = array_slice($arr, 4);
   print_r($arr);
