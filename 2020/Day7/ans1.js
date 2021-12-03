const fs = require('fs')
const path = require('path')

const location = path.join(__dirname, './data.txt')

const data = fs.readFileSync(location, { encoding: 'utf-8' })

let current = 'shiny gold'
let regex = new RegExp(`^([\\w\\s]+)bags contain.+${current}`, 'img')
// console.log(regex)
let array = [...data.matchAll(regex)]
let curr = new Set(array.map(x => x[1].trim()))
// console.log(curr)

while (true) {
  let n = curr.size
  curr.forEach(val => {
    let regex = new RegExp(`^([\\w\\s]+)bags contain.+${val}`, 'img')
    array = [...data.matchAll(regex)].map(x => x[1].trim())
    array.forEach(element => {
      curr.add(element)
    });
  });

  if (n === curr.size)
    break
}

console.log(curr.size)
