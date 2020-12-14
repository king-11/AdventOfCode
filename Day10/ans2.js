const fs = require('fs')
const path = require('path')

const location = path.join(__dirname, './data.txt')

const data = fs.readFileSync(location, { encoding: 'utf-8' })

class Joltage {
  constructor(data) {
    this._data = data.split("\n").filter(Boolean).map(Number).sort((a, b) => a - b)
    this._data.unshift(0)
    this._n = this._data.length
    this._data.push(this._data[this._n-1] + 3)
    this._n++
    this._array = new Array(this._n).fill(0)
    this._array[this._n - 1] = 1
  }

  run() {
    this._array.reverse()
    this._data.reverse()

    this._array.forEach((val, idx) => {
      let current = idx + 1
      while (current < this._n && this._data[current] >= this._data[idx] - 3) {
        this._array[current] += val
        current++;
      }
    })

    this._array.reverse()
    this._data.reverse()
  }

  final() {
    return this._array[0]
  }
}

const p = new Joltage(data)
p.run()
console.log(p.final())
