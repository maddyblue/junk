import audioop
import numpy
import pylab
import wave

def get_psd(fname):
	w = wave.open(fname)
	wp = w.getparams()
	wd = w.readframes(wp[3])
	md = audioop.tomono(wd, wp[1], 1, 0)

	wav = []
	for i in range(wp[3]):
		wav.append(audioop.getsample(md, wp[1], i))
	wav = numpy.array(wav)

	(pxx, freqs) = pylab.psd(wav, 2**17, wp[2])
	return (pxx, freqs)

def sign(n):
	if n < 0:
		return -1
	elif n > 0:
		return 1
	return 0

def peaks(dat, n):
	sprev = prev = 0
	k = 0

	pp = []
	rr = []

	for i in range(len(dat)):
		cur = dat[i]
		dif = cur - prev
		s = sign(dif)

		if s == -1 and sprev == 1:
			pp[k:k+1] = [i - 1]
			rr[k:k+1] = [dat[i - 1]]
			k += 1
		elif s == 1 and i == len(dat):
			pp[k:k+1] = [i]
			rr[k:k+1] = [cur]

		prev = cur
		sprev = s

	pks = rr[:]
	pks_i = pp[:]

	p = []
	c = []

	for i in range(n):
		idx = rr.index(max(rr))
		p[i:i+1] = [pp[idx]]
		c[i:i+1] = [rr[idx]]
		rr[idx] = 0

	return (p, c, pks, pks_i)

(pxx, freqs) = get_psd('c3.wav')
(p, c, pks, pks_i) = peaks(pxx, 10)
