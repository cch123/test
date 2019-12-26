#[macro_export] // macro_export 会把这个 macro export 到 crate:: 根部
macro_rules! four {
    () => { 1 + 3 };
}
