import re

red_regex = re.compile(r"(\d)+ red")
green_regex = re.compile(r"(\d)+ green")
blue_regex = re.compile(r"(\d)+ blue")
int_regex = re.compile(r"(\d)+")


def get_number(match) -> bool:
    if match is None:
        return 0

    red_count = int_regex.match(match.group(0))
    return int(red_count.group(0))


def get_game_power(line: str) -> int:
    game = line.split(":")[1].strip()
    plays = game.split(";")
    max_cubes = {"red": 0, "green": 0, "blue": 0}
    for play in plays:
        # Check for red
        match = red_regex.search(play)
        red = get_number(match)
        if red > max_cubes["red"]:
            max_cubes["red"] = red

        # Check for green
        match = green_regex.search(play)
        green = get_number(match)
        if green > max_cubes["green"]:
            max_cubes["green"] = green

        # Check for blue
        match = blue_regex.search(play)
        blue = get_number(match)
        if blue > max_cubes["blue"]:
            max_cubes["blue"] = blue

    return max_cubes["red"] * max_cubes["green"] * max_cubes["blue"]


if __name__ == "__main__":
    total = 0
    with open("puzzle_2_input.txt") as file:
        for i, line in enumerate(file):
            total += get_game_power(line)
    print(total)
