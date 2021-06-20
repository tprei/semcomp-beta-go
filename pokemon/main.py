import requests
import csv
from time import time

if __name__ == "__main__":
    initial = time()

    with open('pokemon.csv') as csv_file:
        csv_reader = csv.reader(csv_file, delimiter=',')
        for row in csv_reader:
            string_encode = row[1].encode("ascii", "ignore")
            pokemon = string_encode.decode().lower()

            url = "https://pokeapi.co/api/v2/pokemon/" + pokemon
            r = requests.get(url)

    print(time() - initial)