#[derive(Debug)]
struct ListNode {
    elem: i32,
    next: Option<Box<ListNode>>,
}
impl ListNode {
    pub fn new(v : i32) -> Self {
        return ListNode {
            elem:v,
            next: None,
        }
    }
}

fn build_list_from_vec(v: Vec<i32>) -> Option<Box<ListNode>> {
    let mut head = Some(Box::new(ListNode::new(0)));
    let mut cur = head.as_mut();
    for i in v {
        if let Some(x) = cur {
            x.next = Some(Box::new(ListNode::new(i)));
            cur = x.next.as_mut();
        }
    }
    head.unwrap().next
}

fn main() {
    println!("Hello, world!");
    let v = vec![1, 2, 3, 4];
    let mut l = build_list_from_vec(v);
    println!("{:?}", l);
}

