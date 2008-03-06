from pylab import *
from random import *

def fac(n):
	if n <= 1:
		return 1
	else:
		return n * fac(n - 1)

def comb(n, k):
	return fac(n) / (fac(k) * fac(n - k))

def bernoulli(k, n, p):
	return comb(n, k) * p ** k * (1. - p) ** (n - k)

def poisson(k, l):
	return exp(-l) * l**k / fac(k)

if __name__ == '__main__':
	prob = []

	#n = 9
	#n = 3
	n = 10
	#p = .111
	p = 1.01
	tot = 0
	for i in range(n + 1):
		#tot += bernoulli(i, n, p)
		#tot += .25
		tot += poisson(i, p)
		prob.append(tot)

	while True:
		t = []
		prev = 1
		s = range(1, 200)

		for i in s:
			cur = 0
			for j in range(prev):
				r = random()
				for k in range(len(prob)):
					if r < prob[k]:
						cur += k
						break

			print cur
			t.append(cur)
			prev = cur

		print

		if cur == 0:
			continue

		plot(s, t)
		xlabel('Time')
		ylabel('Population')
		#title('Binomial distribution with parameters 9 and .111')
		#savefig('binom')
		#title('Uniform distribution on 0, 1, 2, 3')
		#savefig('uniform')
		title('Poisson distribution with parameter 1.01')
		savefig('poisson')
		show()
		break
