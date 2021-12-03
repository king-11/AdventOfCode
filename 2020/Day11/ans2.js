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

  isValid(x, y) {
    return x >= 0 && x < this._x && y >= 0 && y < this._y
  }

  check(x, y) {
    let count = 0
    const directions = [
      //top
      { x: 0, y: -1 },
      //down
      { x: 0, y: 1 },
      //left
      { x: -1, y: 0 },
      //right
      { x: 1, y: 0 },
      //top right
      { x: 1, y: -1 },
      //top left
      { x: -1, y: -1 },
      //bottom left
      { x: -1, y: 1 },
      //bottom right
      { x: 1, y: 1 },
    ]

    directions.forEach(({ x: dx, y: dy }) => {
      let curX = x + dx, curY = y + dy
      while (this.isValid(curX, curY)) {
        if (this._data[curX][curY] === '#') {
          count++;
          break
        }
        else if (this._data[curX][curY] === 'L') {
          break
        }

        curX += dx
        curY += dy
      }
    });

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
          else if (state >= 5 && cur === '#') {
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

        // console.table(this._data)
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
