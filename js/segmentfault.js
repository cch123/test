
function constfuncs() {
  var funcs = [];
  for (var i = 0; i < 10; i++) {
    funcs[i] = function () {
      return i;
    }()
  }
  console.log(i)
  return funcs;
}
var funcs = constfuncs();
console.log(funcs[1]);
console.log(funcs[2]);
