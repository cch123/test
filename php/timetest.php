<?php
date_default_timezone_set('Asia/Shanghai');
echo strtotime('2015-01-26');
echo "\n";
echo strtotime('2015-01-27');
echo "\n";
echo date('Y-m-d', 1422403200);
echo "\n";

echo strtotime('2012-10-29') - strtotime('2012-10-28');
echo "\n";
echo strtotime('2012-10-29')%86400;
echo "\n";
echo strtotime('2012-10-29')%86400 / 3600;
echo "\n";
echo strtotime('2015-01-27')%86400 / 3600;


