from time import time

initial = time()
for i in range(0, 10000):
    total = 0
    for j in range(0, i+1):
        total += j * j
    
end = time()
print(end - initial)