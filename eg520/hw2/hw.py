from math import cos, sqrt, log

def f(n):
	return n ** 2 + 4 * cos(n)
#	return n ** 4 - 14 * n ** 3 + 60 * n ** 2 - 70 * n

def goldensection(a, b, close):
	p = (3 - sqrt(5)) / 2

	n = 0
	while (1 - p) ** n > close / (b - a):
		n += 1

	print 'N: ' + str(n)

	for i in range(n):
		newa = a + p * (b - a)
		newb = a + (1 - p) * (b - a)
		fa = f(newa)
		fb = f(newb)

		if fa < fb:
			b = newb
		else:
			a = newa

		print '%i & %.4f & %.4f & %.4f & %.4f & [%.4f, %.4f] \\\\' % ((i + 1), newa, newb, fa, fb, a, b)

def fib(n):
	r = 1
	s = 0

	while n > 0:
		n -= 1
		t = r + s
		s = r
		r = t

	return float(r)

def fibonacci(a, b, close):
	ep = 0.05

	n = 0
	while (1 + 2 * ep) / (close / (b - a)) > fib(n + 1):
		n += 1

	print 'N: ' + str(n)

	swap = True

	for i in range(n + 1, 1, -1):
		p = 1 - fib(i - 1) / fib(i)

		newa = a + p * (b - a)
		newb = a + (1 - p) * (b - a)

		if i == 2:
			if swap:
				newb = a + (1 - p - ep) * (b - a)
			else:
				newa = a + (p - ep) * (b - a)

		fa = f(newa)
		fb = f(newb)

		if fa < fb:
			swap = False
			b = newb
		else:
			swap = True
			a = newa

		print '%i & %.4f & %.4f & %.4f & %.4f & %.4f & [%.4f, %.4f] \\\\' % ((n - i + 2), p, newa, newb, fa, fb, a, b)

a = 1.
b = 2.
close = 0.2

#goldensection(a, b, close)
#fibonacci(a, b, close)

# Newton's method (problem 9.3)

def g(x1, x2):
	return [-400. * (x2 - x1 ** 2.) * x1 - 2. * (1 - x1), 200. * (x2 - x1 ** 2.)]

def Finv(x1, x2):
	c = 1. / ((1200. * x1 ** 2. - 400. * x2 + 2.) * 200. - (-400. * x1) * (x1 - 400.))

	return [
		[c * 200., c * 400. * x1],
		[c * (-x1 + 400.), c * (1200. * x1 ** 2. - 400. * x2 + 2.)]
	]

def Newton(x1, x2):
	gg = g(x1, x2)
	ff = Finv(x1, x2)
	y1 = gg[0] * ff[0][0] + gg[1] * ff[0][1]
	y2 = gg[0] * ff[1][0] + gg[1] * ff[1][1]
	newx1 = x1 - y1
	newx2 = x2 - y2
	print x1, x2, gg, ff, y1, y2, newx1, newx2
	return [newx1, newx2]

def fixedstepsize(x1, x2, alpha):
	gg = g(x1, x2)
	newx1 = x1 - alpha * gg[0]
	newx2 = x2 - alpha * gg[1]
	print x1, x2, gg, newx1, newx2
	return [newx1, newx2]

x = [0, 0]
alpha = 0.05

#x = Newton(x[0], x[1])
#x = Newton(x[0], x[1])

#x = fixedstepsize(x[0], x[1], alpha)
#x = fixedstepsize(x[0], x[1], alpha)

def ff(x):
	return 3. * (x - 2.) ** 3.

def df(x):
	return 9. * x

def ddf(x):
	return 9.

def quad(x):
	d01 = x[-1] - x[-2]
	d12 = x[-2] - x[-3]

	fp = (ff(x[-1]) - ff(x[-2])) / (x[-1] - x[-2])
	fp1 = (ff(x[-2]) - ff(x[-3])) / (x[-2] - x[-3])
	fpp = (fp - fp1) / (x[-1] - x[-2])

	#fp = df(x[-1])
	#fpp = ddf(x[-1])

	r = fp / fpp

	#print fx[-1], x[-1], d01, d12, f01, f12, r, x[-1] - r
	print 'x(k+1) = %f - %f / %f [%f] = %f. %f / %f = %f' % (x[-1], fp, fpp, r, x[-1] - r, df(x[-1]), ddf(x[-1]), df(x[-1]) / ddf(x[-1]))

	return x[-1] - r

x = [6., 5., 4.]

for i in range(10):
	x.append(quad(x))

for i in x:
	print 'f(%f) = %f' % (i, ff(i))