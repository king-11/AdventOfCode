const fs = require('fs')
const path = require('path')

const location = path.join(__dirname, './data.txt')

const data = fs.readFileSync(location, { encoding: 'utf-8' })

function max(x, y) {
  return x > y ? x : y
}

function min(x, y) {
  return x < y ? x : y
}

class Ferry {
  constructor(data) {
    this._data = data.split("\n").filter(Boolean).map(x => x.split("").filter(Boolean))
    this._x = this._data.length
    this._y = this._data[0].length
    this._count = 0
  }

  check(x, y) {
    let start = [max(0, x - 1), max(0, y - 1)]
    let end = [min(this._x - 1, x + 1), min(this._y - 1, y + 1)]
    let count = 0

    for (let i = start[0]; i <= end[0]; i++) {
      for (let j = start[1]; j <= end[1]; j++) {
        if (i === x && j === y)
          continue

        if (this._data[i][j] === '#')
          count++
      }
    }

    return count
  }

  run() {
    while (true) {
      let curCount = 0
      let curData = this._data.map(x => [...x])
      this._data.forEach((val, idx) => {
        val.forEach((cur, idy) => {
          if (cur === '.')
            return

          const state = this.check(idx, idy)
          if (state === 0) {
            curData[idx][idy] = '#'
          }
          else if (state >= 4 && cur === '#') {
            curData[idx][idy] = 'L'
          }

          if (curData[idx][idy] === '#') curCount++
        });
      });

      if (curCount === this._count) {
        break
      }
      else {
        this._count = curCount
        this._data = curData
      }
    }
  }

  final() {
    return this._count
  }
}

const p = new Ferry(data)

p.run()

console.log(p.final());
