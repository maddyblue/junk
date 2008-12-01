from rank import Rank
import sys
import utilities

master_wavdir = 'wav'
wavfiles = {
	'haupt-principal-8': {
		'c1.wav': 138,
		'c2.wav': 277,
		'c3.wav': 550,
		'c4.wav': 1100,
		'c5.wav': 2200
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

	fname = '../wav/haupt-principal-8/c1.wav'

	if len(sys.argv) >= 2:
		fname = sys.argv[1]
		if len(sys.argv) == 3:
			n = float(sys.argv[2])

	ranks = []

	for wavdir, wavlist in wavfiles.iteritems():
		ranks.append(Rank(wavdir, wavlist, n))

	r = ranks[0]
	print r

	r.chuck()
