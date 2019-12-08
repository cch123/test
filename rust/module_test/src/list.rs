use crate::inner::go;
pub(super) fn print() {
    println!("{}", "this is in module crate::list");
    go::print_gogo();
}