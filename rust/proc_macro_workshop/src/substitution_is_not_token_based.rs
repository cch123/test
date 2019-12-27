#[macro_export]
macro_rules! capture_expr_then_stringify {
    ($e:expr) => {
        stringify!($e)
    };
}


/*
输入一旦被转成 AST 节点，那被替换的结果就没法解构了
i.e. you cannot examine the contents or match against it ever again.
*/
#[macro_export]
macro_rules! capture_then_what_is {
    (#[$m:meta]) => {what_is!(#[$m])};
    // 注意，只有用 tt 或者 ident 来 capture 的内容才能二次 match
}

#[macro_export]
macro_rules! capture_then_what_is2 {
    (#[$m:tt]) => {what_is!(#[$m])};
    // 注意，只有用 tt 或者 ident 来 capture 的内容才能二次 match
}

#[macro_export]
macro_rules! what_is {
    (#[no_mangle]) => {"no_mangle attribute"};
    (#[inline]) => {"inline attribute"};
    ($($tts:tt)*) => {concat!("something else (", stringify!($($tts)*), ")")};
}