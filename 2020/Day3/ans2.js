const path = require('path')
const dataPath = path.join(__dirname,'./data.json')
let {data} = require(dataPath)

const length = data[0].length

function hits(right,down) {
  let sum = 0, current = 0, j = 0;
  data.forEach((val, i) => {
    if(i === j){
      val[current%length] === "#" && i !== 0 ? sum++ : sum;
      current += right
      j += down
    }
  });
  console.log(sum)
  return sum
}

const slopes = [[1,1],[3,1],[5,1],[7,1],[1,2]]

let product = 1
slopes.forEach(val => {
  console.log(val)
  product *= hits(...val)
})

console.log(product)
