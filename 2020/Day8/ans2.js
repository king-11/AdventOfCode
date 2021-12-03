const fs = require('fs')
const path = require('path')

const location = path.join(__dirname, './data.txt')

const data = fs.readFileSync(location, { encoding: 'utf-8' }).split("\n").filter(Boolean)

class Program {
  constructor(data) {
    this._n = data.length
    this._acc = 0
    this._pc = 0
    this._current = 0
    this._done = new Array(this._n).fill(-1)
    this._instructions = data.map(line => {
      // console.log(line)
      const { groups } = /(?<instruction>nop|jmp|acc)\s(?<value>[+-]\d+)/.exec(line)
      groups.value = +groups.value
      return groups
    })
  }

  run() {
    while (this._pc < this._n) {
      const current = this._instructions[this._pc]
      if (this._done[this._pc] !== -1)
        break;

      this._done[this._pc] = this._current
      this._current++;

      switch (current.instruction) {
        case 'nop': break;
        case 'acc': this._acc += current.value; break;
        case 'jmp': this._pc += (current.value - 1); break;
      }

      this._pc++;
    }
  }

  check() {
    if (this._pc !== this._n)
      return [false, this._acc]

    return [true, this._acc]
  }
}

let accumulator = -1

data.forEach((element, index) => {
  const current = element.slice(0, 3)

  if (["nop", "jmp"].includes(current)) {
    let modifiedData = [...data]
    let val = [...modifiedData[index]]
    current === "nop" ? val.splice(0, 3, ..."jmp") : val.splice(0, 3, ..."nop")
    modifiedData[index] = val.join("")

    const p = new Program(modifiedData)
    p.run()
    if (p.check()[0] && accumulator === -1)
      accumulator = p.check()[1]
  }
});

console.log(accumulator)
