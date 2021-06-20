import asyncio
import aiohttp
import csv
from time import time

async def get(url):
    async with aiohttp.ClientSession() as session:
        async with session.get(url) as response:
            pass

async def main():
    jobs = []
    with open('pokemon.csv') as csv_file:
        csv_reader = csv.reader(csv_file, delimiter=',')
        for row in csv_reader:
            string_encode = row[1].encode("ascii", "ignore")
            pokemon = string_encode.decode().lower()

            url = "https://pokeapi.co/api/v2/pokemon/" + pokemon
            jobs.append(url)

    
    coroutines = [get(url) for url in jobs]
    await asyncio.gather(*coroutines)

if __name__ == "__main__":
    initial = time()
    asyncio.run(main())
    print(time() - initial)