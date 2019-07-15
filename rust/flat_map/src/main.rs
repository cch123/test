// flat map 本质上实际是先 map，再 flatten
fn main() {
    let int_list = vec![vec![1, 2, 3], vec![3, 4, 5], vec![6, 7, 8]];

    // chars() returns an iterator
    let merged = int_list
        .iter()
        //.flat_map(|s| s.iter().map(|e| e + 1).collect::<Vec<_>>())
        //.flat_map(|elem| elem.iter().cloned().collect::<Vec<_>>())
        //.flatten()
        .flat_map(|elem| elem)
        .collect::<Vec<_>>();
    println!("{:#?}", merged)
}
