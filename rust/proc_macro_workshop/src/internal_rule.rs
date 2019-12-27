#[macro_export]
macro_rules! foo {
    (@as_expr $e:expr) => {$e};

    ($($tts:tt)*) => {
        foo!(@as_expr $($tts)*)
    };
}

#[macro_export]
macro_rules! crate_name_util {
    (@as_expr $e:expr) => {$e};
    (@as_item $i:item) => {$i};
    (@count_tts) => {0usize};
    // ...
}

/*
If exporting at least one internal macro is unavoidable
(e.g. you have many macros that depend on a common set of utility rules),
you can use this pattern to combine all internal macros into a single uber-macro.
*/
pub fn internal_rule() {
    foo!( 1 + 2 );
    crate_name_util!(@count_tts);
    crate_name_util!(@as_expr 1 + 2);
}
