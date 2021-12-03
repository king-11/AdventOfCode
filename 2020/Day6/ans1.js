const fs = require('fs')
const path = require('path')

const location = path.join(__dirname, './data.txt')

const data = fs.readFileSync(location, { encoding: 'utf-8' }).split(/\n\n/)

let sum = 0

data.forEach(val => {
  val = val.split(/\n/).filter(Boolean).map(x => [...x])
  const valSet = new Set()

  val.forEach(element => element.forEach(i => valSet.add(i)));

  // console.log(valSet)
  sum += valSet.size
});

console.log(sum)
