def pmf(i):
	p = 0.75
	r = p

	for j in range(i - 1):
		r *= (1 - p)

	return r

t = 0
for i in range(1, 50):
	p = pmf(i)
	t += p
	print i, t, p
