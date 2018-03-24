extern crate iron;
extern crate router;
extern crate urlencoded;
#[macro_use] extern crate mime;

use iron::prelude::*;
use iron::status;
use router::Router;
use urlencoded::UrlEncodedBody;
use std::str::FromStr;

fn main() {
    let mut router = Router::new();

    router.get("/", get_form, "root");
    router.get("/sum", post_sum, "sum");
    println!("Serving on :3000...");
    Iron::new(get_form).http("localhost:3000").unwrap();
}

fn post_sum(request: &mut Request) -> IronResult<Response> {
    let mut response = Response::new();
    let form_data = match request.get_ref::<UrlEncodedBody>() {
        Err(_) => {
            response.set_mut(status::BadRequest);
            response.set_mut(mime!(Text/Html; Charset=Utf8));
            response.set_mut(format!("result is "));
            return Ok(response)
        }
        Ok(map) => map
    };

    let unparsed = match form_data.get("n") {
        None => {
            response.set_mut(status::BadRequest);
            response.set_mut(mime!(Text/Html; Charset=Utf8));
            response.set_mut(format!("result is "));
            return Ok(response)
        }
        Some(nums) => nums
    };

    let mut numbers = Vec::new();
    for p in unparsed {
        match u64::from_str(&p) {
            Err(_) => {
                return Ok(response)
            }
            Ok(n) => {
                numbers.push(n);
            }
        }
    }
    let mut sum = numbers[0];
    for n in &numbers[1..] {
        sum += *n;
    }

    response.set_mut(status::Ok);
    response.set_mut(mime!(Text/Html; Charset=Utf8));
    response.set_mut(format!("result is {}", sum));
    Ok(response)
}

fn get_form(_request: &mut Request) -> IronResult<Response> {
    let mut response = Response::new();
    response.set_mut(status::Ok);
    response.set_mut(mime!(Text/Html; Charset=Utf8));
    response.set_mut(r#"
    <title>sum calc </title>
    <form action="sum" method="post">
        <input type="text" name="n"/>
        <input type="text" name="n"/>
        <button type="submit">Compute sum</button>
    </form>
    "#);

    Ok(response)
}