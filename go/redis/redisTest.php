<?php
$a = new Redis();
$a->connect("localhost", "6379",0);
echo time();
echo "\n";
for ($i =0;$i<10000;$i++) {
    $r = $a->get("foo".$i);
}
echo time();
