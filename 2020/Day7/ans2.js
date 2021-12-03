const fs = require('fs')
const path = require('path')

const location = path.join(__dirname, './data.txt')

const data = fs.readFileSync(location, { encoding: 'utf-8' })

const regex = (val) => new RegExp(`^${val} bags contain (.+)\\n`, 'gim')

const returnVals = (val) => val
  .split(",")
  .map(x => {
    const y = x.trim().split(" ", 3)

    if (y[0] === "no") return null

    return {
      name: `${y[1]} ${y[2]}`,
      count: +y[0]
    }
  })
  .filter(Boolean)

const reducer = (accumulator, val) => accumulator + val.count
const dp = new Map()

function dynamicProgramming({ name }) {
  const array = [...data.matchAll(regex(name))]
  const current = returnVals(array[0][1])

  if (dp.has(name)) return dp.get(name)

  if (current.length === 0) {
    dp.set(name, 0)
    return 0
  }

  let sum = current.reduce(reducer, 0)
  current.forEach(x => {
    sum += (x.count * dynamicProgramming(x))
  })
  dp.set(name, sum)

  return sum
}

console.log(dynamicProgramming({ name: "shiny gold" }));
