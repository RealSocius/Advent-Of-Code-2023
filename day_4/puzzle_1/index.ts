const input = Bun.file("input"); 
const input_text = await input.text()
const input_lines = input_text.split("\n")

let total = 0
input_lines.forEach(line => {
    line = line.split(":")[1]
    total += checkWinningNumbers(line.split("|")[0].split(" "), line.split("|")[1].split(" "));
})

function checkWinningNumbers(winningNumbers: string[], ticketNumbers: string[]): number {

    ticketNumbers = ticketNumbers.filter(tn => !!tn)
    winningNumbers = winningNumbers.filter(wn => !!wn)
    
    
    const matches = ticketNumbers.reduce((n, item) => winningNumbers.includes(item) ? n+1 : n, 0)
    
    if (matches == 0)
        return 0

    return 2 ** (matches - 1)
}

console.log(total)