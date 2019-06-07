#[macro_use]
extern crate actix_web;
use http::StatusCode;

// use actix_web::{error, guard, middleware, Error, HttpRequest, HttpResponse, Result};
use actix_web::{web, App, HttpServer, HttpRequest, HttpResponse};//, Responder};
use std::collections::HashMap;
//use actix_web::test::read_response_json;

#[get("/fuck")]
fn index3(req: HttpRequest) -> HttpResponse {
    HttpResponse::Ok()
        .content_type("application/json")
        .body("oh jack, oh rose")
}

#[post("/setplay")]
fn index4(req: HttpRequest) -> HttpResponse {
    HttpResponse::Ok()
        .content_type("ook")
        .body("oh yeas")
}

fn index2(req: HttpRequest) -> HttpResponse {
    HttpResponse::Ok()
        .content_type("text/plain")
        .body("abcdefg");

    let mut m = HashMap::new();
    m.insert("fuck", "you");
    HttpResponse::Ok()
        .json(&m);

    HttpResponse::Ok().status(StatusCode::INTERNAL_SERVER_ERROR).json(m)
}

fn index(req: HttpRequest, path: web::Path<(String,)>) -> HttpResponse {
    HttpResponse::Ok()
        .content_type("text/plain")
        .body(format!("Hello fk!"))
}

fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(web::resource("/user/{name}").route(web::get().to(index)))
            .service(web::resource("/index.html").to(|| "Hello world!"))
            .service(web::resource("/").to(index2))
            .service(index3)
            .service(index4)
    })
    .bind("0.0.0.0:8080")?
    .run()
}
