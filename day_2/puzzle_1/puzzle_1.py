import re

red_regex = re.compile(r"(\d)+ red")
green_regex = re.compile(r"(\d)+ green")
blue_regex = re.compile(r"(\d)+ blue")
int_regex = re.compile(r"(\d)+")

allowed = {"red": 12, "green": 13, "blue": 14}


def check_allowed(match, max_allowed: int) -> bool:
    red_count = int_regex.match(match.group(0))
    if int(red_count.group(0)) > max_allowed:
        return False
    return True


def check_game(line: str) -> bool:
    game = line.split(":")[1].strip()
    plays = game.split(";")
    for play in plays:
        # Check for red
        match = red_regex.search(play)
        if match and not check_allowed(match, allowed["red"]):
            return False

        # Check for green
        match = green_regex.search(play)
        if match and not check_allowed(match, allowed["green"]):
            return False

        # Check for blue
        match = blue_regex.search(play)
        if match and not check_allowed(match, allowed["blue"]):
            return False

    return True


if __name__ == "__main__":
    total = 0
    with open("puzzle_1_input.txt") as file:
        for i, line in enumerate(file):
            if check_game(line):
                total += i + 1
            print(f"Line {i + 1}: {check_game(line)}")
    print(total)
