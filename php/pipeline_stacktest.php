<?php
function hello($parameter, $next) {
    static $count = 0;

    echo "hello {$count}\n";
    $count++;
    return $next($parameter);
}

function getSlice() {
    global $count;
    echo "getslice\n";
    echo "count {$count}\n";
    $count++;
    //stack+pipe变成了新的stack
    //实际上就是从process和hello5开始“压栈”
    return function($stack, $pipe) {
        //echo "stack is {$stack}, pipe is {$pipe}\n";
        return function($passable) use ($stack, $pipe) {
            return call_user_func($pipe, $passable, $stack);
        };
    };
}

function firstSlice($dest) {
    return function() use ($dest) {
        return call_user_func($dest);
    };
}

$middleware = [];

for($i=0;$i<700;$i++) {
    $middleware[] = "hello";
}

function process() {
}

call_user_func(array_reduce($middleware, getSlice(), firstSlice("process")));
