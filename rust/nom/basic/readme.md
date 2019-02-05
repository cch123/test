## Warn

如果不用 nom 这个内置的 CompleteStr 类型，会导致各种 InComplete，比如 `a=1` parse 就会 InComplete。而必须在 `a=1` 后面加一个空格 `a=1 `。

还是挺别扭的。

https://github.com/Geal/nom/issues/839

https://github.com/Geal/nom/issues/657

看起来是因为 nom 本身可以支持流的场景，所以无法判断用户的输入是不是真的结束了，所以如果用户知道这段输入确实结束了的话，需要显式地告诉 nom 这里已经是完整的 string 了：CompleteStr。