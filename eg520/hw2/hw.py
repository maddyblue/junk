from math import cos, sqrt

def f(n):
	return n ** 2 + 4 * cos(n)

a = 1
b = 2
p = (3 - sqrt(5)) / 2
close = 0.2

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

	print str(i + 1) + ', ' + str(newa) + ', ' + str(newb) + ', ' + str(fa) + ', ' + str(fb) + ', [' + str(a) + ', ' + str(b) + ']'