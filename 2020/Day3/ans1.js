const path = require('path')
const dataPath = path.join(__dirname,'./data.json')
let {data} = require(dataPath)

const length = data[0].length

let sum = 0, current = 0;
data.forEach((val) => {
  val[current%length] === "#" ? sum++ : sum;
  current += 3
});

console.log(sum)
