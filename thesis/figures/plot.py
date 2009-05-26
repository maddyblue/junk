import matplotlib
matplotlib.use('AGG')
import matplotlib.pyplot as plt

def plot(dat, xl, yl, format, name):
	f = open(dat)
	x = []
	y = []
	for l in f:
		d = l.split()
		x.append(float(d[0]))
		y.append(float(d[1]))
	f.close()

	plt.plot(x, y, format)
	plt.xlabel(xl)
	plt.ylabel(yl)
	plt.savefig(name)
	plt.clf()

plot('hdv.dat', r'potential ($\mathrm{V}$)', r'output ($\mu \mathrm{A}$)', '+', 'hdv.png')
plot('216.avg', r'time ($\mathrm{s}$)', r'current ($\mathrm{A}$)', '-', '216.png')
