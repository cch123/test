#![feature(proc_macro_hygiene)]
#![feature(decl_macro)]

#[get("/info/<id>")]
pub fn get_order_info(id: i64) -> String {
    let mut res = "hello this is order detail".to_string();
    res.push_str(&id.to_string());
    res
}
