<?php
  $str = '{{fuck}}这里要把前面花括号的模式给替换掉';
  echo str_replace('{{fuck}}', 'blank', $str);
