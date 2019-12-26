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

#[macro_export]
macro_rules! linkedlist {
    () => {
        None
    };
    ($ ($x:expr), *) => {
        {
            let mut dummy_head = Box::new($crate::Node::new(0));
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
