var x = 1;
var z = 2;
var y = function(x){
  console.log(x);
};
y();
(function(x){
  y();
   function y(x) {
    console.log(x);
    var x = 2;
    console.log(x);
  }
})();
