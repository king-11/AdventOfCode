const path = require('path')
const fs = require('fs')

const data = fs.readFileSync(path.join(__dirname, './data.txt'), {
  encoding: 'utf-8'
}).split("\n\n")

const reducer = (accumulator, currentValue) => { return { ...accumulator, ...currentValue } };

const cleanedData = data.map((val) => val.split(/[\s\n]/)
  .filter(Boolean)
  .map(x => {
    let [key, value] = x.split(":")
    return { [key]: value }
  })
  .reduce(reducer, {}))

// console.table(cleanedData)

const correct = []

cleanedData.forEach((val, idx) => {
  if (Object.keys(val).length === 8)
    correct.push({ ...val, idx })

  if (Object.keys(val).length === 7 && !val.hasOwnProperty('cid'))
    correct.push({ ...val, idx })
});

console.table(correct)
console.log(correct.length)
