trait Graph {
    type Node; // associated type
    type Edge; // associated type

    fn get_visited(&self, n: &Self::Node) -> bool;
    fn set_visited(&mut self, n: &Self::Node);
    fn get_successors(&self, n: &Self::Node) -> Vec<Self::Node>;
}

impl <T>Graph for T {
    type Node = T;
    type Edge = T;

    fn get_visited(&self, n: &Self::Node) -> bool {
        true
    }

    fn set_visited(&mut self, n: &Self::Node) {
    }

    fn get_successors(&self, n: &Self::Node) -> Vec<Self::Node> {
        vec![]
    }
}

/*
impl<N, E> dyn Graph<Node = N, Edge = E> {
    fn get_visited(&self, n: N) -> bool {
        true
    }

    fn set_visited(&mut self, n: N) {}

    fn get_successors(&self, n: N) -> Vec<N> {
        vec![]
    }
}
*/

// 完全独立的泛型函数
fn dfs<G>(g: &mut G) -> Vec<G::Node>
where
    G: Graph,
{
    vec![]
}

fn main() {
    let mut g: &dyn Graph<Node=i64, Edge=i64> = &10;
    println!("{}", g.get_visited(&12));
    let mut g: &dyn Graph<Node = String, Edge = String> = &"a".to_string();
    println!("{}", g.get_visited(&"zz".to_string()));
}
