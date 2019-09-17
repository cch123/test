trait Graph {
    type Node; // associated type
    type Edge; // associated type

    fn get_visited(&self, n: &Self::Node) -> bool;
    fn set_visited(&mut self, n: &Self::Node);
    fn get_successors(&self, n: &Self::Node) -> Vec<Self::Node>;
}

impl<N, E> dyn Graph<Node=N, Edge=E>
{
    fn get_visited(&self, n: N) -> bool {
        true
    }

    fn set_visited(&mut self, n: N) {}

    fn get_successors(&self, n: N) -> Vec<N> {
        vec![]
    }
}

impl Graph for i32 {
    type Node = i32;
    type Edge = i32;
    fn get_visited(&self, n: &Self::Node) -> bool {
        true
    }

    fn set_visited(&mut self, n: &Self::Node) {}

    fn get_successors(&self, n: &Self::Node) -> Vec<Self::Node> {
        vec![]
    }
}

// 完全独立的泛型函数
fn dfs<G>(g: &mut G) -> Vec<G::Node>
where
    G: Graph,
{
    vec![]
}

fn main() {
    //let mut g = Graph<i32,i32>::new();
}
