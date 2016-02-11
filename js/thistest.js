var name = 'top name';
function test() {
  var name = 'inner';
  console.log(this.name);
}
test()
this.sayHi = test
this.name = name
console.log(this)
