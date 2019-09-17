trait Graph {
    type Node; // associated type
    type Edge; // associated type
    
    fn get_visited(&self, n: &Self::Node) -> bool;
    fn set_visited(&mut self, n: &Self::Node);
    fn get_successors(&self, n: &Self::Node) -> [Self::Node];
}

fn main() {
    println!("Hello, world!");
}