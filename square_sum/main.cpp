#include <iostream>
#include <chrono>
using namespace std;

int main() {
    chrono::steady_clock::time_point begin = std::chrono::steady_clock::now();
    for(int i = 0; i < 10000; i++) {
        long long sum = 0;
        for(int j = 0; j <= i; j++) {
            sum += j * j;
        }
    }
    chrono::steady_clock::time_point end = std::chrono::steady_clock::now();
    printf("%lf\n", chrono::duration_cast<std::chrono::microseconds>(end - begin).count() /1000000.0);
}