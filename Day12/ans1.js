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
    0: 'E',
    1: 'S',
    2: 'W',
    3: 'N'
  }
  this._n = this._data.length
  this._x = 0
  this._y = 0
  this._currDirection = 0
  const abs = (value) => value < 0 ? (-value) : value

  const move = ({ type, value }) => {
    if (type === 'F') type = this._directions[this._currDirection]
    if (['L', 'S', 'W'].includes(type)) value = -value

    if (['N', 'S'].includes(type)) this._y += value
    else if (['W', 'E'].includes(type)) this._x += value

    else {
      const turn = ((value / 90) % 4 + 4) % 4
      this._currDirection = (this._currDirection + turn) % 4
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
