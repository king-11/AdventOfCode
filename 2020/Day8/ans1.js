const fs = require('fs')
const path = require('path')

const location = path.join(__dirname, './data.txt')

const data = fs.readFileSync(location, { encoding: 'utf-8' }).split("\n").filter(Boolean)


const n = data.length
let done = new Array(n).fill(-1)

let i = 0, acc = 0, x = 0;

while (i < n) {
  let [curr, val] = data[i].split(" ")
  val = +val

  if (done[i] !== -1)
    break;
  done[i] = x;
  x++;

  switch (curr) {
    case 'nop': i++; break;
    case 'acc': i++; acc += val; break;
    case 'jmp': i += val; break;
  }

}

console.log(acc)
