from math import prod
import os
from itertools import chain

input = "2333133121414131402"

if os.environ.get("REAL"):
  print("reading real file")
  with open(os.environ.get("HOME", "") + "/.cache/advent_of_code_inputs/2024/9.aocinput", "r") as file:
    input = file.read().strip()

def checksum(input):
  checksum = 0
  for (index, value) in enumerate(input):
    if value == ".":
      continue
    checksum += index * int(value)
  return checksum

def part1():
  final_list = []
  for (index, value) in enumerate(list(input)):
    memory_size = int(value)
    if index % 2 == 0:
      final_list.append([index//2 for _ in range(0, memory_size)])
    else:
      final_list.append(["." for _ in range(0, memory_size)])

  final_list = list(chain.from_iterable(final_list))

  i, j = 0, len(final_list) - 1
  while i < j:
    while final_list[i] != ".":
      i += 1
    while final_list[j] == ".":
      j -= 1

    if i < j:
      final_list[i] = final_list[j]
      final_list[j] = "."
      i += 1
      j -= 1

  return checksum(final_list)

def part2():
  final_list = []
  for (index, value) in enumerate(list(input)):
    memory_size = int(value)
    if index % 2 == 0:
      final_list.append([index//2 for _ in range(0, memory_size)])
    else:
      final_list.append(["." for _ in range(0, memory_size)])

  final_list = list(chain.from_iterable(final_list))

  free_spaces = []
  i = 0
  while i < len(final_list):
    if final_list[i] != ".":
      i +=1
      continue

    i2 = i
    while i2 < len(final_list) and final_list[i2 + 1] == ".":
      i2 += 1

    free_spaces.append((i, i2))
    i = i2 + 1

  j = len(final_list) - 1
  while j >= 0:
    if final_list[j] == ".":
      j -= 1
      continue

    val = final_list[j]
    j2 = j
    while j2 > 0 and final_list[j2-1] == val:
      j2 -= 1

    required = j - j2 + 1
    some = None
    for (idx, (i1, i2)) in enumerate(free_spaces):
      if i1 > j2:
        break
      space_size = i2 - i1 + 1
      if space_size >= required:
        some = idx
        break
    if some == None:
      j = j2 - 1
      continue

    (i1, i2) = free_spaces[some]
    new_i1 = i1 + required
    if new_i1 > i2:
      free_spaces.pop(some)
    else:
      free_spaces[some] = (new_i1, i2)

    for idx in range(i1, new_i1):
      final_list[idx] = val

    for idx in range(j2, j+1):
      final_list[idx] = "."

  return checksum(final_list)


print(part1())
print(part2())
