const fs = require('fs')
const path = require('path')

const location = path.join(__dirname, './data.txt')

const data = fs.readFileSync(location, { encoding: 'utf-8' })

class Joltage {
  constructor(data) {
    this._data = data.split("\n").filter(Boolean).map(Number).sort((a, b) => a - b)
    this._n = this._data.length
    this._max = this._data[-1] + 3
    this._current = 0
    this._oneDiff = 0
    this._threeDiff = 0
  }

  run() {
    this._data.forEach((val) => {
      if (val - this._current === 1)
        this._oneDiff++;
      else if (val - this._current === 3)
        this._threeDiff++;

      this._current = val
    });

    this._threeDiff++;
  }

  final() {
    return this._oneDiff * this._threeDiff
  }
}

const p = new Joltage(data)

p.run()

console.log(p.final())
