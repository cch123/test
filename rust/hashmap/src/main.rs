use std::collections::HashMap;

type Callback = Box<Fn(i32) -> i32>;

struct Router {
    routes : HashMap<String, Callback>,
    //routes : HashMap<String, Box<Fn(i32)->i32>>,
}

impl Router {
    fn new() -> Self {
        Router { routes : HashMap::new() }
    }
}

fn main() {
    let mut r = Router::new();
    r.routes.insert("abc".to_string(), Box::new(|x| x+2));
    r.routes.insert("abc".to_string(), Box::new(|x| x+2));
}
