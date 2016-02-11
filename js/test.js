var foo = 0;
var timer = setInterval(function(){
  foo++;
}, 500);
while(1){
  if(foo == 10){
    clearInterval(timer);
    break;
  } else {
    console.log(foo);
  }
}

