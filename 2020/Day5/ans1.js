const fs = require('fs')
const path = require('path')

const location = path.join(__dirname, './data.txt')

const data = fs.readFileSync(location, { encoding: 'utf-8' }).split(/\n/)

const getID = (val) => {
  const chars = [...val]

  const row = parseInt(chars.slice(0, 7).map(x => x === "B" ? 1 : 0).join(""), 2)
  const col = parseInt(chars.slice(7,).map(x => x === "R" ? 1 : 0).join(""), 2)

  return { row, col }
}

let max = -Infinity

data.forEach(element => {
  const { row, col } = getID(element)
  const id = row * 8 + col
  max = max < id ? id : max
});

console.log(max)
