def part_one(lines)
  increase, previous = 0, 0

  lines.each_with_index do |line, i|
    increase += 1 if line > previous && i != 0

    previous = line
  end

  puts "part 1: #{increase}"
end

def part_two(lines)
  increase, previous = 0, 0

  lines.each_with_index do |line, i|
    next if i+2 > lines.size - 1

    batch = lines[i..i+2].sum

    increase +=1 if batch > previous && i != 0

    previous = batch
  end

  puts "part 2: #{increase}"
end

def main
  lines = File.readlines("input.txt").map(&:to_i)

  part_one(lines)
  part_two(lines)
end

main
