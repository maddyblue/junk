def fac(n):
	if n <= 1:
		return 1
	else:
		return n * fac(n - 1)

def comb(n, k):
	return fac(n) / (fac(k) * fac(n - k))

def bernoulli(k, n, p):
	return comb(n, k) * p ** k * (1. - p) ** (n - k)