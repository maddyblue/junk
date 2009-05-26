import matplotlib
matplotlib.use('AGG')
import matplotlib.pyplot as plt

def readfile(fname):
	f = open(fname)
	x = []
	y = []
	for l in f:
		d = l.split()
		x.append(float(d[0]))
		y.append(float(d[1]))
	f.close()
	return (x, y)

def plot(dat, xl, yl, format, name, notick=False):
	x, y = dat

	plt.plot(x, y, format)
	plt.xlabel(xl)
	plt.ylabel(yl)

	if notick:
		plt.axes().xaxis.set_ticks([])
		plt.axes().yaxis.set_ticks([])
		plt.xlabel(xl, fontsize=30)
		plt.ylabel(yl, fontsize=30)

	plt.savefig(name)
	plt.clf()

plot(readfile('hdv.dat'), r'potential ($\mathrm{V}$)', r'output ($\mu \mathrm{A}$)', '+', 'hdv.png')
plot(readfile('216.avg'), r'time ($\mathrm{s}$)', r'current ($\mathrm{A}$)', '-', '216.png')
plot(readfile('256.avg'), r'potential ($\mathrm{V}$)', r'current ($\mathrm{A}$)', '-', '256.png')
plot(readfile('63'), r'potential ($\mathrm{V}$)', r'current ($\mathrm{A}$)', '-', '63.png')

plot(([0, 1], [1, 1]), r'time ($\mathrm{s}$)', r'potential ($\mathrm{V}$)', '-', 'amperometry.png', True)
plot(([0, 1, 2, 3, 4], [0, 1, 0, 1, 0]), r'time ($\mathrm{s}$)', r'potential ($\mathrm{V}$)', '-', 'cv.png', True)

# dpv data

x = []
y = []
v = 0
step = 1
diff = 0.2
t = 0

for i in range(7):
	x.extend([t, t + 1, t + 1, t + 2])
	y.extend([v, v, v + step, v + step])
	v += diff
	t += 2

plot((x, y), r'time ($\mathrm{s}$)', r'potential ($\mathrm{V}$)', '-', 'dpv.png', True)
