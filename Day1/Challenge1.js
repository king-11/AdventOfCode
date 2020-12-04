const path = require('path')
const dataPath = path.join(__dirname,'./challenge1.json')
let {data} = require(dataPath)

console.log(data)

data = data.sort((a,b) => Number(a) < Number(b) ? -1 : 1)

const n = data.length
let [start, end] = [0, n-1]

while(data[start] + data[end] !== 2020){
  let sum = data[start] + data[end]
  if(sum > 2020)
    end--;
  if(sum < 2020)
    start++;
}

console.log(data[start],data[end])
