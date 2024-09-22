import math

def main():
    number = []

    while True:
        try:
            num = float(input())
        except ValueError as e:
            print(e)
            return

        number.append(num)

        mad = Mad(number)
        median_value = Median(number)
        min_value = median_value - mad
        max_value = (median_value - mad) * 1.5
        print(Round(min_value), Round(max_value))

def Round(x):
    if x >= 0:
        return int(x + 0.5)
    else:
        return int(x - 0.5)

def Median(number):
    sorted_numbers = sorted(number)
    n = len(sorted_numbers)
    
    if n % 2 == 1:
        return sorted_numbers[n // 2]
    else:
        return (sorted_numbers[n // 2] + sorted_numbers[(n // 2) - 1]) / 2

def Mad(number):
    median_value = Median(number)
    mad = [abs(n - median_value) for n in number]
    return Median(mad)

if __name__ == "__main__":
    main()
