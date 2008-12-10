import freqs
import numpy
import os
import os.path
import utilities
import warnings

from math import log

warnings.simplefilter('ignore', numpy.RankWarning)

chuck_main = """
fun float[] proc_harm(float freq)
{
	float ret[peaks];
	int i;
	0 => int idx;

	for(0 => i; i < recs; i++)
	{
		if(bases[i] < freq)
			i => idx;
	}

	for(0 => i; i < peaks; i++)
	{
		harms[i][idx] => ret[i];
	}

	return ret;
}

fun float[] proc_perc(float freq)
{
	float ret[peaks];
	int i;
	int j;

	for(0 => i; i < peaks; i++)
	{
		0 => ret[i];

		for(0 => j; j <= degs; j++)
		{
			percs[i][j] * Math.pow(freq, degs - j) +=> ret[i];
		}
	}

	return ret;
}

me.arg(0) => string freqstr;
if(freqstr.length() == 0) "440" => freqstr;
Std.atoi(freqstr) => float freq;

proc_harm(freq) @=> float harm[];
proc_perc(freq) @=> float perc[];

Pan2 p => dac;
SinOsc s[peaks];
for(0 => int i; i < peaks; i++)
{
	s[i] => p;
	harm[i] * freq => s[i].freq;
	perc[i] => s[i].gain;
	//<<< harm[i] * freq, "Hz,", perc[i], "gain" >>>;
}

while(true) 1::second => now;
"""

class Rank:

	PERC_FIT_DEGS = 3

	def __init__(self, directory, files, numpeaks, master_wavdir='../wav'):
		"""
		directory is the name of the directory containing the files
		master_wavdir is the name of the directory containing parameter directory
		files is a dict with keys as filenames and values as frequencies close to the aural (heard) frequency of the recording: the base frequency will be the closest one to this value in the resulting peaks list. Hence, if the base frequency does not
		appear in the first numpeaks peaks, the assigned base frequency will be incorrect.
		"""

		print 'processing rank %s:' %directory
		self.directory = directory
		self.numpeaks = numpeaks
		self.entries = []

		for fname, base in files.iteritems():
			self.entries.append(Entry(os.path.join(master_wavdir, directory, fname), numpeaks, base))

		self.entries.sort(cmp=lambda x, y: cmp(x.base, y.base))
		self.bases = [e.base for e in self.entries]

		perc_fit = []
		harm_fit = []
		harm_polyfit = []
		for i in range(numpeaks):
			p = [e.percs[i] for e in self.entries]
			perc_fit.append(numpy.lib.polyfit(self.bases, p, self.PERC_FIT_DEGS))

			h = [e.peaks_freqs[i] / e.base for e in self.entries]
			harm_fit.append(h)
			harm_polyfit.append(numpy.lib.polyfit(self.bases, h, len(self.bases)))

		self.perc_fit = numpy.array(perc_fit)
		self.harm_fit = numpy.array(harm_fit)
		self.harm_polyfit = numpy.array(harm_polyfit)

		self.synth = {}

	def __str__(self):
		ret = 'Rank %s:\n' %self.directory

		for e in self.entries:
			ret += str(e)

		ret += '\n'

		for i in range(len(self.perc_fit)):
			ret += '\tperc fit %i: %s\n' %(i, self.perc_fit[i])

		for i in range(len(self.harm_fit)):
			ret += '\tharm fit %i: %s\n' %(i, self.harm_fit[i])

		for i in range(len(self.harm_polyfit)):
			ret += '\tharm polyfit %i: %s\n' %(i, self.harm_polyfit[i])

		return ret

	def get_synth(self, synthfreqs):
		ret = {}
		for f in synthfreqs:
			if f not in self.synth:
				self.synth[f] = Synth(f, self.perc_fit, self.harm_fit, self.bases)

			ret[f] = self.synth[f]

		return ret

	def write_rank(self, length=1, fs=11025, basedir='../out', range_low=1, range_high=-1):
		outdir = os.path.join(basedir, str(self.numpeaks), self.directory)
		try:
			os.makedirs(outdir)
		except:
			pass

		self.get_synth(freqs.freqs[range_low:range_high])

		for key in freqs.keys[range_low:range_high]:
			outname = os.path.join(outdir, '%02i-%3s-%f.wav' %(key[0], key[1].replace(' ', '_'), key[2]))
			print outname
			utilities.write_wav(self.synth[key[2]].wav(length, fs), fs, outname)

	def write_freq(self, freq, length=1, fs=11025, basedir='../out'):
		outdir = os.path.join(basedir, str(self.numpeaks), self.directory)
		try:
			os.makedirs(outdir)
		except:
			pass

		self.get_synth([freq])

		outname = os.path.join(outdir, '%f.wav' %freq)
		w = self.synth[freq].wav(length, fs)
		utilities.write_wav(w, fs, outname)

	def chuck_rank(self, basedir='../out', range_low=1, range_high=-1):
		outdir = os.path.join(basedir, str(self.numpeaks), self.directory)
		try:
			os.makedirs(outdir)
		except:
			pass

		self.get_synth(freqs.freqs[range_low:range_high])

		for key in freqs.keys[range_low:range_high]:
			outname = os.path.join(outdir, '%02i-%3s-%f.ck' %(key[0], key[1].replace(' ', '_'), key[2]))
			utilities.write_chuck(self.synth[key[2]], outname)

	def chuck(self, outdir='../out/chuck'):
		try:
			os.makedirs(outdir)
		except:
			pass

		outname = os.path.join(outdir, '%s-%s.ck' %(self.directory, self.numpeaks))

		f = open(outname, 'w')
		f.write('%i => int peaks;\n' %self.numpeaks)
		f.write('%i => int recs;\n' %len(self.bases))
		f.write('%i => int degs;\n' %self.PERC_FIT_DEGS)

		f.write('\n[ [ %s ] ] @=> float percs[][];\n' %'], ['.join([ ', '.join(k) for k in [ [ '%20.25f' %j for j in i] for i in self.perc_fit ]]))
		f.write('\n[ [ %s ] ] @=> float harms[][];\n' %'], ['.join([ ', '.join(k) for k in [ [ '%10.5f' %j for j in i] for i in self.harm_fit ]]))
		f.write('\n[ %s ] @=> float bases[];\n' %', '.join([ '%f' %i for i in self.bases ]))

		f.write(chuck_main)
		f.close()

class Synth:
	def __init__(self, freq, perc, harm, bases):
		assert(len(perc) == len(harm))

		harm_idx = 0
		for i in range(len(bases)):
			if bases[i] < freq:
				harm_idx = i
		harm_col = harm[:,harm_idx]

		p = []
		h = []
		for i in range(len(perc)):
			p.append(numpy.polyval(perc[i], freq))
			h.append(harm_col[i] * freq)

		self.freq = freq
		self.perc = numpy.array(p)
		self.harm = numpy.array(h)

	def __str__(self):
		ret = '%7.2f: ' %self.freq

		for i in range(len(self.perc)):
			ret += '%7.2f(%4.1f), ' %(self.harm[i], self.perc[i] * 100)

		return ret[:-2]

	def wav(self, length, fs):
		ret = 0

		for i in range(len(self.perc)):
			(w, x) = utilities.mk_wav(self.harm[i], length, fs)
			ret += ret + (w * self.perc[i])

		return ret

class Entry:
	closeness = 0.01

	def __init__(self, fname, numpeaks, base=0):
		print '\t%s' %fname
		self.fname = fname
		self.wav, self.wp = utilities.get_wav(fname)
		self.psd, self.psd_freqs = utilities.get_psd(self.wav, self.wp[2])
		self.peaks, self.peaks_energy = utilities.get_peaks(self.psd, numpeaks)
		self.peaks_freqs = numpy.array([self.psd_freqs[i] for i in self.peaks])
		self.percs = self.peaks_energy / self.psd.sum()
		self.peaks_notes = utilities.get_note(self.peaks_freqs)

		if base:
			bi = None
			diff = 100

			for i in self.peaks_freqs:
				d = abs(base - i)
				if d < diff:
					diff = d
					bi = i

			self.base = bi
		else:
			self.base = self.psd_freqs[self.peaks[0]]

	def __str__(self):
		ret = '\tEntry %s:\n' %self.fname

		for idx in range(len(self.peaks)):
			i = self.peaks[idx]
			peak = self.psd_freqs[i]
			h = peak / self.base
			rh = round(h)
			ratio = rh / h

			if peak == self.base:
				disp = ' <- base harmonic'
			elif (1 - self.closeness) < ratio and (1 + self.closeness) > ratio:
				disp = ' <- harmonic %i' %rh
			else:
				disp = ''

			if not disp:
				h = self.base / peak
				rh = round(h)
				ratio = rh / h

				if (1 - self.closeness) < ratio and (1 + self.closeness) > ratio:
					disp = ' <- harmonic 1/%i' %rh

			ret += '\t%3i: %7.2f (%5.2f%%): %-3s (%7.2f)%s\n' %(idx, peak, self.percs[idx] * 100, freqs.keys[self.peaks_notes[idx]][1], freqs.keys[self.peaks_notes[idx]][2], disp)

		return ret
