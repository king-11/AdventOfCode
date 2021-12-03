const path = require('path')
const fs = require('fs')

const file = path.join(__dirname, 'data.txt')
const data = fs.readFileSync(file, { encoding: 'utf-8' })

function Manhattan(data) {
  this._data = data.split("\n").filter(Boolean).map(val => {
    const { groups } = /^(?<type>[A-Z])(?<value>\d+)$/.exec(val)
    groups.value = (+groups.value)

    return { ...groups }
  })

  this._directions = {
    0: ({ x, y }) => { return { x, y } },
    1: ({ x, y }) => { return { x: y, y: -x } },
    2: ({ x, y }) => { return { x: -x, y: -y } },
    3: ({ x, y }) => { return { x: -y, y: x } },
  }

  this._x = 0
  this._y = 0

  this._curr = { y: 1, x: 10 }
  const abs = (value) => value < 0 ? (-value) : value

  const move = ({ type, value }) => {
    if (type === 'F') {
      this._x += value * this._curr.x
      this._y += value * this._curr.y
      return
    }

    if (['L', 'S', 'W'].includes(type)) value = -value

    if (['N', 'S'].includes(type)) this._curr.y += value
    else if (['W', 'E'].includes(type)) this._curr.x += value
    else {
      const turn = ((value / 90) % 4 + 4) % 4
      this._curr = this._directions[turn](this._curr)
    }
  }
  const run = () => {
    this._data.forEach(move);
  }

  const final = () => abs(this._x) + abs(this._y)

  return {
    run,
    final
  }
}

const p = new Manhattan(data)

p.run()

console.log(p.final());
