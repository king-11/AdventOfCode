const fs = require('fs')
const path = require('path')

const location = path.join(__dirname, './data.txt')

const data = fs.readFileSync(location, { encoding: 'utf-8' })

const preamble = 25

class CheckSum {
  constructor(preamble, data) {
    this._preamble = preamble
    this._mainData = data.split("\n").filter(Boolean).map(Number)
    this._checkData = this._mainData.slice(preamble)
  }

  isSum(val, idx) {
    const checkIn = this._mainData.slice(idx, idx + this._preamble).sort((a, b) => a - b)
    let [start, end] = [0, this._preamble - 1]
    while (start < end) {
      const sum = checkIn[start] + checkIn[end]
      if (sum > val)
        end--;
      else if (sum < val)
        start++;
      else
        return true
    }

    return false
  }

  run() {
    return this._checkData.find((val, idx) => !this.isSum(val, idx));
  }
}

const p = new CheckSum(preamble, data)

console.log(p.run())
