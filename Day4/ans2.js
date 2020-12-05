const path = require('path')
const fs = require('fs')

const data = fs.readFileSync(path.join(__dirname, './invalid.txt'), {
  encoding: 'utf-8'
}).split("\n\n")

const reducer = (accumulator, currentValue) => { return { ...accumulator, ...currentValue } };

const cleanedData = data.map((val) => val.split(/[\s\n]/)
  .filter(Boolean)
  .map(x => {
    let [key, value] = x.split(":")
    return { [key]: value }
  })
  .reduce(reducer, {}))

// console.log(cleanedData[0])

const checkBirth = ({ byr }) => byr.length === 4 && +byr >= 1920 && +byr <= 2002
const checkIssueYear = ({ iyr }) => iyr.length === 4 && +iyr >= 2010 && +iyr <= 2020
const checkExpYear = ({ eyr }) => eyr.length === 4 && +eyr >= 2020 && +eyr <= 2030
const checkHeight = ({ hgt }) => {
  let valid = true;
  valid &= /^[0-9]+(cm|in)$/.test(hgt)
  if (!valid) return false

  let val = +hgt.slice(0, -2)

  return (/cm$/.test(hgt) && val >= 150 && val <= 193) || (/in$/.test(hgt) && val >= 59 && val <= 76)
}
const checkHair = ({ hcl }) => /^#([a-f]|[0-9]){6,6}$/.test(hcl)
const checkEye = ({ ecl }) => ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'].indexOf(ecl) !== -1
const checkID = ({ pid }) => /^[0-9]{9,9}$/.test(pid)
const checkKeys = (val) => Object.keys(val).length === 8 || (Object.keys(val).length === 7 && !val.hasOwnProperty('cid'))

const check = (val) => checkKeys(val) && checkID(val) && checkEye(val) && checkHair(val) && checkHeight(val) && checkExpYear(val) && checkIssueYear(val) && checkBirth(val)

const correct = []

cleanedData.forEach((val, idx) => {
  if (check(val))
    correct.push(val)
});

console.table(correct)
console.log(correct.length)

// const test = {
//   ecl: 'grn',
//   cid: '315',
//   iyr: '2012',
//   hgt: '192cm',
//   eyr: '2023',
//   pid: '873355140',
//   byr: '1925',
//   hcl: '#cb2c03'
// }

// console.log(checkKeys(test))
