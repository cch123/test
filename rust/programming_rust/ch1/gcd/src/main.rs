fn gcd(mut n: u64, mut m: u64) -> u64 {
    assert!(n != 0 && m != 0);
    while m != 0 {
        if m < n {
            let t = m;
            m = n;
            n = t;
        }
        m = m % n;
    }
    return n;
}

fn main() {
    let n = gcd(1, 2);
    println!("{}", n);
}

// 这个 test 怎么才能存储到其它文件
#[test]
fn test_gcd() {
    assert_eq!(gcd(14, 15), 1);
    assert_eq!(gcd(2*3*5*7*11, 3*7*19), 3*7)
}