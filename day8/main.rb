input = File.readlines("input.txt")

grand_total = input.sum do |line| 
    pre, post = line.split("|").map(&:split)

    numbers = {}
    unknown = {6 => [], 5 => []}

    pre.each do |code|
        split_code = code.chars.sort
        case split_code.size
        when 2
            numbers[1] = split_code
        when 3
            numbers[7] = split_code
        when 4
            numbers[4] = split_code
        when 7
            numbers[8] = split_code
        else
            unknown[split_code.size] << split_code
        end
    end

    unknown[6].each do |code|
        if (code - numbers[7]).size == 4 
            numbers[6] = code
        elsif (code - numbers[4]).size == 2
            numbers[9] = code
        else 
            numbers[0] = code
        end
    end

    unknown[5].each do |code|
        if (code - numbers[7]).size == 2
            numbers[3] = code
        elsif (numbers[6] - code).size == 1 
            numbers[5] = code
        else 
            numbers[2] = code
        end
    end

    inverted = numbers.invert

    multi = 10000

    post.sum do |bar|
        multi /= 10
        inverted[bar.chars.sort] * multi
    end
end

puts "part 2: #{grand_total}"