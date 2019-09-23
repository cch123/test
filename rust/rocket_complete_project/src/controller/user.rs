#![feature(proc_macro_hygiene)]
#![feature(decl_macro)]

#[get("/info/<id>")]
pub fn get_user_info(id: i64) -> String {
    "hello here is Xargin".to_string()
}

#[get("/info/<id>")]
pub fn get_user_info(id: i64) -> String {
    "hello here is Xargin".to_string()
}
