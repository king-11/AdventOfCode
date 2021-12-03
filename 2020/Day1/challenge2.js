const path = require('path')
const dataPath = path.join(__dirname,'./challenge2.json')
let {data} = require(dataPath)

data = data.sort((a,b) => Number(a) < Number(b) ? -1 : 1)

const n = data.length

data.forEach((val, current) => {
  const finding = 2020 - val
  let [start, end] = [current+1, n-1]
  while(data[start] + data[end] !== finding && start < end){
    let sum = data[start] + data[end]
    if(sum > finding)
      end--;
    if(sum < finding)
      start++;
  }

  if(data[start]+data[end]+val === 2020){
    console.log(data[start],data[end],val)
  }
});
