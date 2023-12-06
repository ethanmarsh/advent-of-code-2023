msg = "Running Advent of Code Day 1 Solution - here we go!"
print(msg)

string_digits = [
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine"
]

def spelled_out_digit_to_int(str):
    if str == "one":
        return 1
    elif str == "two":
        return 2
    elif str == "three":
        return 3
    elif str == "four":
        return 4
    elif str == "five":
        return 5
    elif str == "six":
        return 6
    elif str == "seven":
        return 7
    elif str == "eight":
        return 8
    elif str == "nine":
        return 9
    else:
        print("error! can't convert spelled out digit to an int")

# Function to solve part 1
def find_both_digits(line):
    first_digit = ""
    last_digit = ""
    for character in line:
        if character.isdigit():
            last_digit = character
            if not first_digit:
                first_digit = character
    combined = first_digit + last_digit
    return int(combined)

def find_first_digit(line):
    first_digit = -1
    first_digit_index = -1
    first_str_digit = ""
    first_str_digit_index = 10000000
    for idx, character in enumerate(line):
        if character.isdigit():
            first_digit = int(character)
            first_digit_index = idx
            break
    
    for string_digit in string_digits:
        if string_digit in line:
            start = line.find(string_digit)
            if start < first_str_digit_index:
                first_str_digit_index = start
                first_str_digit = string_digit

    if first_digit_index < first_str_digit_index:
        return first_digit
    else:
        return spelled_out_digit_to_int(first_str_digit)
    
def find_last_digit(line):
    last_digit = -1
    last_digit_index = -1
    last_str_digit = ""
    last_str_digit_index = -1

    for idx, character in enumerate(line):
        if character.isdigit():
            last_digit = character
            last_digit_index = idx
    
    for string_digit in string_digits:
        if string_digit in line:
            start = line.rfind(string_digit)
            if start > last_str_digit_index:
                last_str_digit_index = start
                last_str_digit = string_digit

    if last_digit_index > last_str_digit_index:
        return last_digit
    else:
        return spelled_out_digit_to_int(last_str_digit)

file = open("day1input.txt", "r")
total = 0
for line in file:
    # Part 1
    # total += find_both_digits(line)
    # Part 2
    print(line)
    first_digit = find_first_digit(line)
    print("First digit: " + str(first_digit))
    last_digit = find_last_digit(line)
    print("Last digit: " + str(last_digit))
    combined = str(first_digit) + str(last_digit)
    print("Combined: " + combined)
    total += int(combined)
    print("-----------")

file.close
# print("Part 1 Expected output - 57346")
# print("Part 2 expected output - 57345")
print("Part 2 - Final output - " + str(total))