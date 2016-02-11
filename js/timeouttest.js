var t, sec = 0;
function countDown() {
  //clearTimeout(t);
  t = setTimeout(countDown, 1000);
  sec++;
  if(sec == 60) {
    clearTimeout(t);
  }
  console.log(sec);
}

function test() {
  t = setTimeout(countDown, 1000);
}

