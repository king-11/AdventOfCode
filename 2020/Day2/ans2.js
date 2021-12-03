const path = require('path')
const dataPath = path.join(__dirname,'./data2.json')
let {data} = require(dataPath)

let sum = 0;

data.forEach((val) => {
  let [start,end,character,value] = val.split(/[-\s:]/).filter(Boolean)

  start--;end--;
  console.log(start,end,character,value)
  if(!((value[start] !== character && value[end] !== character) || (value[start] === character && value[end] === character)))
    sum++;
});

console.log(sum)

// let [start,end,character,value] = "2-7 p: ptphppvppppp".split(/[-\s:]/).filter(Boolean)

// start--;end--;
// console.log(value[end] === character)
// const valid = !((value[start] !== character && value[end] !== character) || (value[start] === character && value[end] === character))

// console.log(valid)
