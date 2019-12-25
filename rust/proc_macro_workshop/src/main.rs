//#[macro_use] // 新版本 rust 没有 macro_use 也行
pub mod easy;

#[derive(Debug)]
pub struct Node {
    pub val: i32,
    pub next: Option<Box<Node>>,
}

impl Node {
    pub fn new(val: i32) -> Self {
        Node {next : None, val}
    }
}

macro_rules! linkedlist {
    () => {
        None
    };
    ($ ($x:expr), *) => {
        {
            let mut dummy_head = Box::new(Node::new(0));
            let mut cursor = &mut dummy_head;
            $(
                cursor.next = Some(Box::new($crate::Node::new($x)));
                cursor = cursor.next.as_mut().unwrap();
            )*
            drop(cursor);

            dummy_head.next
        }
    };
}

fn main() {
    let x = linkedlist![1,2,3,4,5];
    dbg!(x);

    println!("{}", four!())
}
