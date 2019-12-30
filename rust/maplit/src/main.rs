use maplit::hashmap;
use unicorn::{Cpu, CpuX86, RegisterX86};

#[derive(Debug)]
enum Direction {
    East,
    West,
    South,
    North,
}

fn main() {
    let map = hashmap!{
        "a" => Direction::East,
        "b" => Direction::West,
    };
    dbg!(map);

    let map = hashmap!{
        "rr" => unicorn::RegisterX86::RAX,
    };
    dbg!(map);
}
