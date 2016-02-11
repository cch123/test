var t, sec = 0;
function countDown() {
  clearInterval(t);
  setTimeout(countDown, 1000);
  sec++;
  console.log(sec);
}

function test() {
  t = setInterval(countDown, 1000);
}

