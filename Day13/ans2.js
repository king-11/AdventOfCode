console.time('start')
const path = require('path')
const fs = require('fs')


let fileName = 'data.txt'
if (!isNaN(Number(process.argv[2]))) {
  fileName = `test${Number(process.argv[2])}.txt`
}

const file = path.join(__dirname, fileName)
const data = fs.readFileSync(file, { encoding: 'utf-8' })

function modularExponentiation(x, y, p) {
  let result = 1
  x = x % p
  if (x === 0) return 0
  while (y > 0) {
    if (y & 1)
      result = (result * x) % p
    y >>= 1
    x = (x * x) % p
  }

  return result
}

const absoluteModulo = (a, b) => ((a % b) + b) % b;


function Shuttle(data) {
  let [, mydata] = data.split("\n", 2)
  this._data = mydata.split(",").map((val, idx) => val !== 'x' ? { value: +val, index: +idx } : undefined).filter(Boolean)

  this._product = (this._data.reduce((total, current) => total * current.value, 1))
  this._sum = BigInt(0)

  const run = () => {
    this._data.forEach(({ value, index }) => {
      const remainder = absoluteModulo(value - index, value)
      const valueU = this._product / value
      const inverse = modularExponentiation(valueU, value - 2, value)

      // console.log((valueU * inverse) % value)
      this._sum = this._sum + BigInt(BigInt(remainder) * BigInt(valueU) * BigInt(inverse))
      // console.log(this._sum)
    })
  }

  const final = () => {
    console.log(this._data);
    return this._sum % BigInt(this._product)
  }

  return {
    run,
    final
  }
}

const p = new Shuttle(data)

p.run()

console.log(p.final());

console.timeEnd('start')
