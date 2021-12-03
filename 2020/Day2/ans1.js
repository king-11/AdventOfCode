const path = require('path')
const dataPath = path.join(__dirname,'./data1.json')
let {data} = require(dataPath)

let sum = 0;

data.forEach((val) => {
  let [start,end,character,value] = val.split(/[-\s:]/).filter(Boolean)
  const regexExp = new RegExp(`^${character}\{${start},${end}\}$`,)
  const splitExp = new RegExp(`[^${character}]+`)
  console.log(value)
  value = value.split(splitExp).filter(Boolean).join("")
  console.log(value)
  const valid = regexExp.test(value)
  if(valid)
    sum++;
});

console.log(sum)

// let [start,end,character,value] = "2-7 p: ptphppvppppp".split(/[-\s:]/).filter(Boolean)
// const regexExp = new RegExp(`^${character}\{${start},${end}\}$`,)
// const splitExp = new RegExp(`[^${character}]+`)
// value = value.split(splitExp).filter(Boolean).join("")
// console.log(regexExp.test(value))
// console.log(start,end,character,value)
