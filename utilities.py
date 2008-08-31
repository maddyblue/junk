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

	return numpy.array(p), numpy.array(c)

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

def mk_wav(freq, length, fs):
	"""
	Return a tuple of a time-series array and corresponding times at given frequency (Hz), length (sec), and sampling freqency (samples/sec).
	"""

	degrees = 360.0;
	x = numpy.arange(0, degrees * length, degrees / fs) / degrees
	wav = numpy.sin(numpy.pi * freq * x)

	return wav, x

def write_wavspec(freq, length, fs, fname):
	"""
	Wrapper for mk_wav + write_wav.
	"""

	(wav, x) = mk_wav(freq, length, fs)
	write_wav(wav, fs, fname)

def write_wav(wav, fs, fname):
	"""
	Write wav data at sampling frequency fs to file fname. wav must be a numpy.array. It may be in any encoding, and will be converted properly internally.
	"""

	w = wave.open(fname, 'wb')
	w.setnchannels(2)
	w.setsampwidth(2)
	w.setframerate(fs)

	# scale wave to 16-bit unsigned integers
	wav = wav - wav.min() # adjust so wav.min() = 0
	wav = (wav * 2**14 / wav.max()).astype('int16')

	frames = ''

	for d in wav:
		msb = d >> 8
		lsb = d - (msb << 8)
		frames += chr(lsb) + chr(msb)

	w.writeframes(frames)
	w.close()
