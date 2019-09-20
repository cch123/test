#![feature(proc_macro_hygiene)]
#![feature(decl_macro)]

#[macro_use]
extern crate rocket;

mod controller;
mod logic;
mod model;
mod init;
mod config;

// 带参数的路由
// 用户实际访问的实际是
// mount 的第一个参数 + 这里括号里的字符串
#[get("/<name>/<age>")]
fn hello(name: String, age: u8) -> String {
    format!("Hello, {} year old named {}!", age, name)
}

#[get("/")]
fn world() -> String {
    "1".to_string()
}

#[get("/fuck")]
fn fuck() -> String {
    "oh ea".to_string()
}

fn main() {
    init::init();
    rocket::ignite()
        .mount("/hello", routes![hello])
        .mount("/world", routes![world, fuck])
        .launch();
}

