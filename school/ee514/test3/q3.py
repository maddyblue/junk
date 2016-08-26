from numpy import arange

def s(t):
	if t % 2 < 1:
		return 1
	else:
		return -1

es = []
e = 0
tau = 0.5
time = 0.5
times = arange(0, 2, .01)
r = arange(0, 1, .01)

for t in times:
	e = 0
	for i in r:
		e += s(t - i)
	es.append(e)
