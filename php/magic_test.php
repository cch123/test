<?php
class Person {
    function __call($name, $arguments) {
        echo 'no such function: ' . $name;
    }
}

$p = new Person();
$p->test();
