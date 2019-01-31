use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug)]
struct TreeNode {
    val: i32,
    next: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    pub fn new(v: i32) -> Self {
        return TreeNode { val: v, next: None };
    }
}

fn main() {
    f1();
    f2();
    f3();
    f4();
}

fn f1() {
    let dummy = Rc::new(RefCell::new(TreeNode::new(0)));
    let mut cursor = Rc::clone(&dummy);
    let nums = vec![1, 2, 3, 4];
    for n in nums {
        let node = Rc::new(RefCell::new(TreeNode::new(n)));
        cursor.borrow_mut().next = Some(node);
        let mut next;
        match cursor.borrow_mut().next.as_ref() {
            Some(n) => {
                next = Rc::clone(&n);
            }
            None => unreachable!(),
            /*
            这里必须有 unreachable，否则编译会报错
            error[E0382]: use of moved value: `next`
  --> src/main.rs:35:18
   |
35 |         cursor = next;
   |                  ^^^^ value used here after move
36 |     }
   |     - value moved here, in previous iteration of loop
            */
        }
        cursor = next;
    }
    println!("{:?}", dummy);
}

fn f2() {
    let dummy = Rc::new(RefCell::new(TreeNode::new(0)));
    let mut cursor = Rc::clone(&dummy);
    let nums = vec![1, 2, 3, 4];
    for n in nums {
        let node = Rc::new(RefCell::new(TreeNode::new(n)));
        cursor.borrow_mut().next = Some(node);
        let mut next;
        if let Some(n) = cursor.borrow_mut().next.as_ref() {
            next = Rc::clone(&n);
        } else {
            // 不写 unreachable 编译通过不了
            // why
            unreachable!()
        }
        cursor = next;
    }
    println!("{:?}", dummy);
}

fn f3() {
    let dummy = Rc::new(RefCell::new(TreeNode::new(0)));
    let mut cursor = Rc::clone(&dummy);
    let nums = vec![1, 2, 3, 4];
    for n in nums {
        let node = Rc::new(RefCell::new(TreeNode::new(n)));
        cursor.borrow_mut().next = Some(node);
        let next= cursor.borrow_mut().next.clone();
        cursor = next.unwrap();
    }
    println!("{:?}", dummy);
}

fn f4() {
    let dummy = Rc::new(RefCell::new(TreeNode::new(0)));
    let mut cursor = Rc::clone(&dummy);
    let nums = vec![1, 2, 3, 4];
    for n in nums {
        let node = Rc::new(RefCell::new(TreeNode::new(n)));
        cursor.borrow_mut().next = Some(node);
        // 和 f3 一样，这里如果不 clone，会报 cannot move out of borrowed content
        let next= cursor.borrow_mut().next.clone().unwrap();
        /*
        error[E0507]: cannot move out of borrowed content
  --> src/main.rs:92:19
   |
92 |         let next= cursor.borrow_mut().next.unwrap();
   |                   ^^^^^^^^^^^^^^^^^^^^^^^^ cannot move out of borrowed content
        */
        cursor = next;
    }
    println!("{:?}", dummy);
}