def calPoints(operations: List[str]) -> int:
    total_sum = 0
    scoreRecord = []

    for val in range(len(operations)):
        match operations[val]:
            case "+":
                first_elem = scoreRecord.pop()
                second_elem = scoreRecord.pop()
                scoreRecord.append(first_elem)
                scoreRecord.append(second_elem) 
                scoreRecord.append(first_elem+second_elem)
            case "D":
                doubleScore = scoreRecord.pop()
                scoreRecord.append(doubleScore) 
                scoreRecord.append(doubleScore*2)
            case "C":
                if len(scoreRecord) != 0:
                    scoreRecord.pop()
            case _:
                scoreRecord.append(int(operations[val]))
    for sum in range(len(scoreRecord)):
        total_sum += scoreRecord[sum]
    return total_sum

def main():
    ops = ["1","2","+","C","5","D"]
    print(calPoints(ops))

if __name__ == "__main__":
    main()
