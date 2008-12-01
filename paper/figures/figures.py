import numpy
import pylab
import sys
sys.path.append('../../python')
import utilities

def mkgraphs(wname, prefix):
	print wname

	wav, wp = utilities.get_wav(wname)
	sp = 1.0 / wp[2]
	timevals = numpy.arange(0, wp[3] * sp, sp)

	part = 4000
	pylab.close()
	pylab.plot(timevals[:part], wav[:part])
	pylab.xlabel('Time (s)')
	pylab.ylabel('Relative amplitude')
	pylab.savefig('%stime.png' %prefix)

	PSD_NFFT = 2**17
	(pxx, fxx) = utilities.get_psd(wav, wp[2], PSD_NFFT)

	part = 1700
	lxx = 10 * numpy.log10(pxx)
	pylab.close()
	pylab.plot(fxx[:part], lxx[:part])
	pylab.xlabel('Frequency (Hz)')
	pylab.ylabel('Power (dB)')
	pylab.savefig('%spsd.png' %prefix)

	n = 5 # num peaks

	p, c = utilities.get_peaks(pxx, n)

	peak_freqs = [ fxx[i] for i in p ]
	peak_db = [ lxx[i] for i in p ]
	print peak_freqs
	pylab.plot(peak_freqs, peak_db, 'ro', ms=4)
	for i in range(n):
		pylab.text(peak_freqs[i], peak_db[i], '%5.1fHz' %peak_freqs[i], ha='center', position=(peak_freqs[i], peak_db[i] + 1))
	pylab.xlabel('Frequency (Hz)')
	pylab.ylabel('Power (dB)')
	pylab.savefig('%speaks.png' %prefix)

mkgraphs('recording.wav', 'recording-')
mkgraphs('synth.wav', 'synth-')
