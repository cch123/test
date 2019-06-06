#[macro_use]
extern crate actix_web;
use http::StatusCode;

use actix_web::{error, guard, middleware, Error, HttpRequest, HttpResponse, Result};
use actix_web::{web, App, HttpServer, Responder};
use std::collections::HashMap;
use actix_web::test::read_response_json;

fn with_param(req: HttpRequest) -> HttpResponse {
    HttpResponse::Ok()
        .content_type("text/plain")
        .body("abcdefg");

    let mut m = HashMap::new();
    m.insert("fuck", "you");
    HttpResponse::Ok()
        .json(&m);

    HttpResponse::Ok().status(StatusCode::INTERNAL_SERVER_ERROR).json(m)
}

fn without_param(req: HttpRequest, path: web::Path<(String,)>) -> HttpResponse {
    HttpResponse::Ok()
        .content_type("text/plain")
        .body(format!("Hello fk!"))
}

fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(web::resource("/user/{name}").route(web::get().to(with_param)))
            .service(web::resource("/index.html").to(|| "Hello world!"))
            .service(web::resource("/").to(with_param))
    })
    .bind("0.0.0.0:8080")?
    .run()
}
