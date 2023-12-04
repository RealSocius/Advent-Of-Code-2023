const input = Bun.file("input"); 
const input_text = await input.text()
const input_lines = input_text.split("\n")


function processLines(allInputLines: string[], inputLines: string[]): number {
    let totalCards = 0
    inputLines.forEach((line, index) => {
        const parsedLine = parseLine(line)
        const matches = checkWinningNumbers(parsedLine.winningNumbers, parsedLine.ticketNumbers);
        totalCards += matches
        if (matches) {
            const ticketCopies = allInputLines.slice(index + 1, index + 1 + matches)
            totalCards += processLines(allInputLines.slice(index + 1), ticketCopies)
        }
    })
    return totalCards
}


function parseLine(line: string) {
    line = line.split(":")[1]
    const winningNumbers = line.split("|")[0].split(" ")
    const ticketNumbers = line.split("|")[1].split(" ")

    return {
        winningNumbers,
        ticketNumbers
    }
}

function checkWinningNumbers(winningNumbers: string[], ticketNumbers: string[]): number {
    ticketNumbers = ticketNumbers.filter(tn => !!tn)
    winningNumbers = winningNumbers.filter(wn => !!wn)
    
    const matches = ticketNumbers.reduce((n, item) => winningNumbers.includes(item) ? n+1 : n, 0)
    
    return matches
}

let total = input_lines.length
total += processLines(input_lines, input_lines)
console.log(total)