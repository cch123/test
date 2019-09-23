#![feature(proc_macro_hygiene)]
#![feature(decl_macro)]

#[macro_use]
extern crate rocket;

mod config;
mod controller;
mod init;
mod logic;
mod model;

fn main() {
    init::init();
    rocket::ignite()
        .mount("/user", routes![controller::user::get_user_info])
        .mount("/order", routes![controller::order::get_order_info])
        .launch();
}
