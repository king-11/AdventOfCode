const path = require('path')
const fs = require('fs')

const file = path.join(__dirname, 'data.txt')
const data = fs.readFileSync(file, { encoding: 'utf-8' })

function Shuttle(data) {
  let [myTime, mydata] = data.split("\n", 2)
  this._data = mydata.split(",").filter(val => val !== 'x').map(Number)
  this._start = myTime
  delete myTime, mydata

  this._current = 1000

  const check = (val) => (val - (this._start % val)) % val

  const run = () => {
    this._current = this._data.reduce((minimum, val) => {
      const current = check(val)
      return minimum > current ? current : minimum
    }, 1000);

  }

  const final = () => {
    const val = this._data.find(val => this._current === check(val))
    return this._current * val
  }

  return {
    run,
    final
  }
}

const p = new Shuttle(data)

p.run()

console.log(p.final());
