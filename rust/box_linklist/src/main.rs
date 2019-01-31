#[derive(Debug)]
struct Node {
    num : i32,
    next : Option<Box<Node>>,
}

impl Node {
    fn new(val: i32) -> Self{
        return Node {
            num : val,
            next : None,
        }
    }
}

fn main() {
    let nums = vec![1,2,3,4];
    let mut dummy = Box::new(Node::new(0));
    let mut cur = &mut dummy;
    for n in nums {
        cur.next = Some(Box::new(Node::new(n)));
        cur = cur.next.as_mut().unwrap();
        // 注意这里，如果自己写
        // cur = &mut cur.next.unwrap();
        // 是通过不了编译的
        /*
        error[E0507]: cannot move out of borrowed content
  --> src/main.rs:24:20
   |
24 |         cur = &mut cur.next.unwrap();
   |                    ^^^^^^^^ cannot move out of borrowed content

error[E0716]: temporary value dropped while borrowed
  --> src/main.rs:24:20
   |
21 |         cur.next = Some(Box::new(Node::new(n)));
   |         -------- borrow used here, in later iteration of loop
...
24 |         cur = &mut cur.next.unwrap();
   |                    ^^^^^^^^^^^^^^^^^- temporary value is freed at the end of this statement
   |                    |
   |                    creates a temporary which is freed while still in use
   |
   = note: consider using a `let` binding to create a longer lived value

        */
    }
    println!("{:?}", dummy);
    println!("{:?}", dummy.next);
}
