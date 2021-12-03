const fs = require('fs')
const path = require('path')

const location = path.join(__dirname, './data.txt')

const data = fs.readFileSync(location, { encoding: 'utf-8' }).split(/\n\n/)

function intersection(setA, setB) {
  let _intersection = new Set()
  for (let elem of setB) {
    if (setA.has(elem)) {
      _intersection.add(elem)
    }
  }
  return _intersection
}

let sum = 0

data.forEach(val => {
  val = val.split(/\n/).filter(Boolean).map(x => [...x])
  let valSet = new Set(val[0])

  val.forEach(element => {
    const currSet = new Set(element)
    valSet = intersection(valSet, currSet)
  });
  sum += valSet.size
});

console.log(sum)
