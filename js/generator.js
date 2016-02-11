//这个例子的yield和python比较类似
var Fiber = require('fibers');

var inc = Fiber(function(start){
  var total = start;
  while(true) {
    total += start;
    Fiber.yield(total);
  }
})

for(var ii = 0; ii<=10; ii++) {
  console.log(inc.run(2));
}

for(var ii = 0; ii<=10; ii++) {
  console.log(inc.run(1))
}
