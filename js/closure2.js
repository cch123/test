var name = "trigkit4";
var segmentFault = {
  　name : "My SF",
  　getNameFunc : function(){
    　　return function(){
      　　　return this.name;
      　};
    }
};
console.log(segmentFault.getNameFunc()());  //弹出trigkit4
