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

let arr = []

data.forEach(element => {
  const { row, col } = getID(element)
  const id = row * 8 + col
  if (arr[row] === undefined)
    arr[row] = []

  arr[row].push({ col, id })
});

arr.forEach(element => {
  if (element.length === 8)
    return

  element.sort((a, b) => a.col - b.col)
  for (let i = 1; i < element.length - 1; i++) {
    if (element[i].col !== element[i - 1].col + 1)
      console.log(element[i].id - 1)
  }
});
