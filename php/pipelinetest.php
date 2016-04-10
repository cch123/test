<?php

function hello1($parameter, $next) {
    echo "hello1\n";
    return $next($parameter);
}

function hello2($parameter, $next) {
    echo "hello2\n";
    return $next($parameter);
}

function hello3($parameter, $next) {
    echo "hello3\n";
    return $next($parameter);
}

function hello4($parameter, $next) {
    echo "hello4\n";
    return $next($parameter);
}

function hello5($parameter, $next) {
    echo "hello5\n";
    return $next($parameter);
}

function process() {
    echo "process\n";
}

$middleware = ["hello1", "hello2", "hello3", "hello4", "hello5"];

$count = 0;
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


call_user_func(
    array_reduce(
        array_reverse($middleware),
        getSlice(),
        firstSlice("process", "")
    )
    , ""
);
