// todo
// 1. 如何控制连接数
// 2. 是否能做类似 ultron 的效果
// 控制线程数应该是
// thread::spawn(move || {
//});
/*
https://github.com/fdehau/tui-rs
https://github.com/ragona/clobber
https://computingforgeeks.com/termgraph-command-line-tool-draw-graphs-in-terminal-linux/
https://github.com/marcusolsson/tui-go
*/
use std::thread;
use async_std::task;

fn main() {
    let thread_num = 10;
    (0..thread_num).for_each(|i| {
        thread::spawn(move || {
            // step 1 create connection

            // step 2 spawn tasks

            /*
            task::spawn(async {
            });
            */
        });
    })
}
