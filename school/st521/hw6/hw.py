#from pylab import arange
from math import *

def fac(n):
	if n <= 1:
		return 1
	else:
		return n * fac(n - 1)

def poisson(k, l):
	return exp(-l) * l**k / fac(k)

#for l in arange(0.1, 15, .1):
for l in range(1, 40, 2):
	l /= 10.
	j = 0
	for i in range(1, 100, 2):
		j += poisson(i, l)
	print l, j