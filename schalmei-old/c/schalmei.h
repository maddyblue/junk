#ifndef SCHALMEI_H
#define SCHALMEI_H

int psd(double *, int, int, int, double *, double *);
double * readfile(char *, SNDFILE *, struct SF_INFO *);
double * readdumpfile(char *, int *, int *);
int dumpfile(int, double *, int, char *);
void fft(int, double *, double *);
double * hanning(size_t);
double * vander(int, double *, int);
int cmp(double, double);
int * peaks(int, int, double *, double *);
double * invert(int, double *);
double * transpose(int, int, double *);
void pf(int, int, double *);
double * mult(int, int, double *, int, int, double *);
double * polyfit(int, double *, double *, int);
void chuckmatrix(FILE *, int, int, double *);

char chuck_file[] = "\nfun float[] proc_harm(float freq)\n{\n	float ret[peaks];\n	int i;\n	0 => int idx;\n\n	for(0 => i; i < recs; i++)\n	{\n		if(bases[i] < freq)\n			i => idx;\n	}\n\n	for(0 => i; i < peaks; i++)\n	{\n		harms[i][idx] => ret[i];\n	}\n\n	return ret;\n}\n\nfun float[] proc_perc(float freq)\n{\n	float ret[peaks];\n	int i;\n	int j;\n\n	for(0 => i; i < peaks; i++)\n	{\n		0 => ret[i];\n\n		for(0 => j; j <= degs; j++)\n		{\n			percs[i][j] * Math.pow(freq, degs - j) +=> ret[i];\n		}\n	}\n\n	return ret;\n}\n\nme.arg(0) => string freqstr;\nif(freqstr.length() == 0) \"440\" => freqstr;\nStd.atoi(freqstr) => float freq;\n\nproc_harm(freq) @=> float harm[];\nproc_perc(freq) @=> float perc[];\n\nPan2 p => dac;\nSinOsc s[peaks];\nfor(0 => int i; i < peaks; i++)\n{\n	s[i] => p;\n	harm[i] * freq => s[i].freq;\n	perc[i] => s[i].gain;\n	//<<< harm[i] * freq, \"Hz,\", perc[i], \"gain\" >>>;\n}\n\nwhile(true) 1::second => now;";

#endif
