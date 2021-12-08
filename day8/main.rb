input = File.readlines("input.txt")

grand_total = 0

input.each do |line| 
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

    while unknown[6].size != 0
        unknown[6].each_with_index do |foo,i|
            if (foo - numbers[7]).size == 4 
                numbers[6] = unknown[6].delete_at(i)
            end
            if (foo - numbers[4]).size == 2
                numbers[9] = unknown[6].delete_at(i)
            end
        end

        if unknown[6].size == 1 
            numbers[0] = unknown[6].delete_at(0)
        end
    end

    while unknown[5].size !=0
        unknown[5].each_with_index do |foo,i|
            if (foo - numbers[7]).size == 2
                numbers[3] = unknown[5].delete_at(i)
            end

            if (numbers[6] - foo).size == 1 
                numbers[5] = unknown[5].delete_at(i)
            end

            if unknown[5].size == 1 
                numbers[2] = unknown[5].delete_at(0)
            end
        end
    end

    inverted = numbers.invert

    total = 0
    multi = 1000

    post.each do |bar|
        total += inverted[bar.chars.sort] * multi
        multi /= 10
    end

    grand_total += total
end

puts "part 2: #{grand_total}"