# 11. Sieve of Eratosthenes

# Sieve of Earatosthenes
def sieve(n):
    sieve = [True] * (n+1)
    sieve[0] = sieve[1] = False
    i = 2
    while(i*i <= n):
        if (sieve[i]):
            k = i*i
            while (k <= n):
                sieve[k] = False
                k += i
        i += 1
    return sieve


for i, v in enumerate(sieve(20)):
    print(i, v)


# 11.1 Factorization


# generate sieve
# each elements are the lower factors of the each indices
def arrayF(n):
    F = [0] * (n+1)
    i = 2
    while (i*i <= n):
        if (F[i] == 0):
            k = i * i
            while (k <= n):
                if F[k] == 0:
                    F[k] = i
                k += i
        i += 1
    return F


F = arrayF(20)
for i, v in enumerate(F):
    print(i, v)


# execute factorization with sieve
def factorization(x, F):
    primeFactors = []
    while (F[x] > 0):
        primeFactors += [F[x]]
        x = int(x / F[x])
        print(F[x], x)
    primeFactors += [x]
    return primeFactors


print(factorization(20, F))
