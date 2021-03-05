#!/usr/bin/env python3
# from typing import *
import random

YES = 'Yes'
NO = 'No'


def main():
    N = random.randint(1, 1000)
    M = random.randint(1, 1000)
    K = random.randint(0, N*M)
    print(N, M, K)


if __name__ == '__main__':
    main()
