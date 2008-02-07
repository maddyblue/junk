from math import exp

def fac(n):
	if n <= 1:
		return 1
	else:
		return n * fac(n - 1)

def comb(n, k):
	return fac(n) / (fac(k) * fac(n - k))

def poi(k, l):
	return exp(-l) * l**k / fac(k)

for i in range(0, 20):
	print i, poi(i, 10)

#for i in range(1, 52):
#	for j in range(1, 52):
#		if comb(i, j) == 6:
#			print i, j
