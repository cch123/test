#[macro_use]
extern crate actix_web;
use http::StatusCode;

use actix_web::{web, App, HttpServer,HttpRequest, HttpResponse};//, Responder};
use std::collections::HashMap;
use actix_http::KeepAlive;


fn index2(req: HttpRequest) -> HttpResponse {
    HttpResponse::Ok()
        .content_type("text/plain")
        .body("hello")

}

fn index(req: HttpRequest, path: web::Path<(String,)>) -> HttpResponse {
    HttpResponse::Ok()
        .content_type("text/plain")
        .body(format!("Hello fk!"))
}

fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            //.service(web::resource("/user/{name}").route(web::get().to(index)))
            //.service(web::resource("/index.html").to(|| "Hello world!"))
            .service(web::resource("/").to(index2))
            //.service(index3)
            //.service(index4)
    })
    .bind("0.0.0.0:9999")?
    .run()
}
