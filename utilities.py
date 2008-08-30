import audioop
import numpy
import pylab
import wave

from freqs import keys, freqs
from math import log

PSD_NFFT = 2**17

def get_psd(fname, nfft=PSD_NFFT):
	"""
	Given a filename of a wav, returns a tuple of the power spectral density using Welch's method and associated frequencies.
	"""

	w = wave.open(fname)
	wp = w.getparams()
	wd = w.readframes(wp[3])
	md = audioop.tomono(wd, wp[1], 1, 0)

	wav = []
	for i in range(wp[3]):
		wav.append(audioop.getsample(md, wp[1], i))
	wav = numpy.array(wav)

	(pxx, fxx) = pylab.psd(wav, nfft, wp[2])
	return (pxx, fxx)

def get_peaks(dat, n):
	"""
	Return a list of the incidies of the n highest peaks from the array dat.
	"""

	sprev = prev = 0
	k = 0

	pp = []
	rr = []

	for i in range(len(dat)):
		cur = dat[i]
		dif = cur - prev
		s = cmp(dif, 0)

		if s == -1 and sprev == 1:
			pp[k:k+1] = [i - 1]
			rr[k:k+1] = [dat[i - 1]]
			k += 1
		elif s == 1 and i == len(dat):
			pp[k:k+1] = [i]
			rr[k:k+1] = [cur]

		prev = cur
		sprev = s

	p = []
	c = []

	for i in range(n):
		idx = rr.index(max(rr))
		p[i:i+1] = [pp[idx]]
		c[i:i+1] = [rr[idx]]
		rr[idx] = 0

	return p, c

def get_percs(psd, psd_freqs, peaks, peaks_energy):
	"""
	Returns a list of the percentages of each entry in peaks' energy contribution to the entire wave.
	"""

	s = sum(psd)
	perc = [i / s for i in peaks_energy]

	return perc

def get_note(frequencies):
	"""
	Takes a list of frequencies and returns a list of indicies corresponding to the closest entry in freqs.freqs to each entry.
	"""

	ret = []
	for f in frequencies:
		n = 0

		if f < freqs[0]:
			f = freqs[0]

		while n < (len(freqs) - 1):
			if f >= freqs[n] and f < freqs[n + 1]:
				low = log(freqs[n])
				mid = log(f)
				high = log(freqs[n + 1])

				if (high - mid) < (mid - low):
					n += 1

				break
			n += 1
		ret.append(n)

	return ret
