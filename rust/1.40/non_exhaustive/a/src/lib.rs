#[non_exhaustive]
#[derive(Debug)]
pub struct Person {
	age : i32,
	height : i32,
}

pub fn New() -> Person {
	Person {age :1, height : 2}
}

