fn clo1() {
    let mut v = 0;
    let mut inc = || {
        v += 1;
        println!("{}", v)
    };
    // 如果 move 的话，这种 copy 类型会将变量 copy 进去
    // 不再和外部变量发生关联
    // let mut inc = move || {v += 1;println!("{}", v)};
    inc();
    inc();
    let mut x = &mut v;
    println!("{}", x);
}

fn clo2() {
    let mut v = vec![1, 2, 3];
    let mut clo = |i: i32| {
        v.push(i);
    };
    // non copy 类型
    // 被 move 了之后无法在外部再使用
    //let mut clo = move |i: i32| {
    //    v.push(i);
    //};
    clo(10);
    clo(10);
    // 这里调用 print 会出错
    /*
        ...
    23 |     println!("{:?}", v);
       |                      ^ value borrowed here after move
       */

    // println!("{:?}", v);
}

fn main() {
    clo1();
    clo2();

    let mut x = 1;
    let cl = || {
        x += 1;
        println!("{}", x)
    };
    clo_param(cl);

    main2();
}

fn clo_param<F>(mut f: F)
where
    F: FnMut(),
{
    f();
    f();
}


struct Solution;
impl Solution {}

fn apply<F>(f: F)
where
    F: FnOnce(),
{
    f();
}

// apply2 和 apply 是等价的
// trait 作为参数和返回值都可以尝试写成 impl 形式
// 不过如果是 Fn、FnMut、FnOnce，这样写可能比较乱
// 所以才会用 where 形式
fn apply2(f: impl FnOnce()) {
    f();
}

// 返回闭包
fn get_cl() -> impl FnOnce(i32) {
    let x = move |mut i: i32| {
        i += 1;
    };
    x
}

fn get_cl2() -> impl Fn(i32) {
    let x = move |i: i32| println!("{}", i);
    x
}

fn get_cl3() -> impl FnMut() {
    let mut x = 1;
    // 如果不 move
    /*
        error[E0597]: `x` does not live long enough
      --> src/main.rs:28:9
       |
    24 | fn get_cl3() -> impl FnMut() {
       |                 ------------ opaque type requires that `x` is borrowed for `'static`
    ...
    27 |     let c = || {
       |             -- value captured here
    28 |         x += 1;
       |         ^ borrowed value does not live long enough
    ..
        */
    let c = move || {
        // 注意这里的 move
        x += 1;
        println! {"{}", x};
    };
    c
}

fn main2() {
    let mut x = "abc";
    let mut y = x.to_owned();

    let clo = || {
        println!("{}", x);

        y.push_str("ddd");
        std::mem::drop(y);
    };
    apply(clo);
    //apply2(clo);
}


use std::collections::HashMap;
/*
struct Router <C> where C : Fn(i32) -> i32 {
    routes : HashMap<String, Box<C>>,
}
*/
// programming rust
// page 318
struct Router {
    routes : HashMap<String, Box<dyn Fn(i32)->i32>>,
}

type C = Box<dyn Fn(i32) -> i32>;

//impl <C> Router<C> where C: Fn(i32) -> i32 {
impl Router {
    fn new() -> Self {
        Router {
            routes: HashMap::new(),
        }
    }

    fn add_route<Cl>(&mut self, url : &str, callback: Cl) where Cl : Fn(i32)->i32 + 'static {
        self.routes.insert(url.to_string(), Box::new(callback));
    }
}

fn apply4<A, B, C, G>(mut f: impl FnMut(B) -> G, a: A)
-> impl FnMut(&B) -> C
        where  
             G: FnMut(A) -> C,
             B: Copy,
             A: Clone {
    move |b| f(*b)(a.clone())
}

/*
不能嵌套 impl
fn apply3<A, B, C, G>(mut f: impl FnMut(B) -> impl FnMut(A)-> C, a: A)
-> impl FnMut(&B) -> C
        where  
             B: Copy,
             A: Clone {
    move |b| f(*b)(a.clone())
}
*/

fn main3() {
    let mut router = Router::new();
    router.add_route("/getid",|x| x + 1);
    router.add_route("/getidx",|x| x + 2);
}
