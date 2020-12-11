const fs = require('fs')
const path = require('path')

const location = path.join(__dirname, './data.txt')

const data = fs.readFileSync(location, { encoding: 'utf-8' })

const preamble = 25

class CheckSum {
  constructor(preamble, data) {
    this._preamble = preamble
    this._mainData = data.split("\n").filter(Boolean).map(Number)
    this._n = this._mainData.length
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

  sumIt(val) {
    let [start, end] = [0, 0]
    let sum = this._mainData[0]
    while (end < this._n) {
      if (sum < val) {
        sum += this._mainData[++end]
      }
      else if (sum > val) {
        sum -= this._mainData[start++]
      }
      else {
        const ans = this._mainData.slice(start, end + 1).sort((a, b) => a - b);
        return (ans[0] + ans[ans.length - 1])
      }
    }

    return 0
  }
}

const p = new CheckSum(preamble, data)

const val = p.run()

console.log(p.sumIt(val))
