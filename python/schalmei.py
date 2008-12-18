from rank import Rank
import sys
import utilities

master_wavdir = 'wav'
wavfiles = {
	'haupt-principal-8': {
		'c#3.wav': 138,
		'd#3.wav': 155,
		'f3.wav': 174,
		'g3.wav': 195,
		'a3.wav': 220,
		'b3.wav': 246,
		'c#4.wav': 277,
		'c#5.wav': 554,
		'c#6.wav': 1108,
		'd#6.wav': 1244,
		'f6.wav': 1396,
		'g6.wav': 1567,
		'a6.wav': 1760,
		'b6.wav': 1975,
		'c#7.wav': 2217,
		}
}

def f(fname):
	wav, wp = utilities.get_wav(fname)
	pxx, fxx = utilities.get_psd(wav, wp[2])
	p, c = utilities.get_peaks(pxx, n)
	pf = [ fxx[i] for i in p ]
	return wav, wp, pxx, fxx, p, c, pf

def figures(rank):
	rank.write_freq(550, 1, 44100)

	wav, wp, pxx, fxx, p, c, pf = f(fname)
	print '%s: %s' %(fname, pf)

	outname = '../paper/figures/synth.wav'
	s = utilities.mk_synth(pf[0], 1, wp[2], pf, c)
	utilities.write_wav(s, wp[2], outname)
	wav, wp, pxx, fxx, p, c, pf = f(outname)
	print '%s: %s' %(outname, pf)

	pxx, fxx = utilities.get_psd(s, wp[2])
	p, c = utilities.get_peaks(pxx, n)
	pf = [ fxx[i] for i in p ]
	print 'synth: %s' %pf

if __name__ == '__main__':
	n = 5;

	if len(sys.argv) == 2:
		n = int(sys.argv[1])

	if n < 2:
		n = 2

	ranks = []

	for wavdir, wavlist in wavfiles.iteritems():
		ranks.append(Rank(wavdir, wavlist, n))

	r = ranks[0]
	print r

	r.chuck()
