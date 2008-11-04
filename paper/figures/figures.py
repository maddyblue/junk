import audioop
import numpy
import pylab
import wave

wname = '../../wav/haupt-principal-8/c3.wav'
w = wave.open(wname)
wp = w.getparams()
wd = w.readframes(wp[3])
md = audioop.tomono(wd, wp[1], 1, 1)
sp = 1.0 / wp[2]
timevals = numpy.arange(0, wp[3] * sp, sp)

wav = []
for i in range(wp[3]):
	wav.append(audioop.getsample(md, wp[1], i))
wav = numpy.array(wav)

part = 2000
pylab.plot(timevals[:part], wav[:part])
pylab.xlabel('Time (s)')
pylab.ylabel('Relative amplitude')
pylab.savefig('time.png')

PSD_NFFT = 2**17
(pxx, fxx) = pylab.psd(wav, PSD_NFFT, wp[2])
pylab.close()

part = 7000
lxx = 10 * numpy.log10(pxx)
pylab.plot(fxx[:part], lxx[:part])
pylab.xlabel('Frequency (Hz)')
pylab.ylabel('Power (dB)')
pylab.savefig('psd.png')

n = 5 # num peaks

sprev = prev = 0

pp = []
rr = []

for i in range(len(pxx)):
	cur = pxx[i]
	dif = cur - prev
	s = cmp(dif, 0)

	# at a peak
	if s == -1 and sprev == 1:
		pp.append(i - 1)
		rr.append(pxx[i - 1])
	# peak at end of data
	elif s == 1 and i == len(pxx):
		pp.append(i)
		rr.append(cur)

	prev = cur
	sprev = s

p = []
c = []

for i in range(n):
	idx = rr.index(max(rr))
	p.append(pp[idx])
	c.append(rr[idx])
	rr[idx] = 0

peak_freqs = [ fxx[i] for i in p ]
peak_db = [ lxx[i] for i in p ]
pylab.plot(peak_freqs, peak_db, 'ro', ms=12)
pylab.xlabel('Frequency (Hz)')
pylab.ylabel('Power (dB)')
pylab.savefig('peaks.png')
