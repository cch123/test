
var x = (function() {
  var count = 0;
  var timer , sec = 0;
  return {
    get: function() {
      return count ++;
    },
    set: function() {
      return count --;
    }

  };
})();

console.log(x.get())
console.log(x.get())
console.log(x.get())
console.log(x.get())
console.log(x.get())
console.log('----')
console.log(x.set())
console.log(x.set())
console.log(x.set())

