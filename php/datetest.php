<?php
  $a = strtotime(date('now'));
  echo date('Y-m-d h:i:s', $a);

  if($a > strtotime('20151212')) {
      echo 'yes';
  }else {
      echo 'no';
  }
