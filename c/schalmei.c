#include <math.h>
#include <stdio.h>
#include <string.h>

#include <fftw3.h>
#include <sndfile.h>

#include "schalmei.h"
#include "freqs.h"

int main(int argc, char *argv[])
{
	SNDFILE *s;
	struct SF_INFO *si;
	double *in, *pxx, *freqs;
	int i, nfft, fs, n, numfreqs;

	if((si = (struct SF_INFO *)malloc(sizeof(*si))) == NULL)
		return 1;

	if((in = readfile("../wav/haupt-principal-8/c3.wav", s, si)) == NULL)
		return 1;

	nfft = 64;
	fs = si->samplerate;
	n = si->frames;

	numfreqs = nfft / 2 + 1;
	if((pxx = (double *)calloc((size_t)(numfreqs), sizeof(double))) == NULL)
		return 1;
	if((freqs = (double *)calloc((size_t)(numfreqs), sizeof(double))) == NULL)
		return 1;

	psd(in, n, nfft, fs, pxx, freqs);

	for(i = 0; i < numfreqs; i++)
		printf("%10.5f: %20.10f\n", freqs[i], pxx[i]);

	return 0;
}

int psd(double *x, int n, int nfft, int fs, double *pxx, double *freqs)
{
	double *d, *p, *cur, *pcur, *window;
	double norm, ratio;
	#define NORM_FACTOR 2.0
	int i, j, loc, numfreqs, segments;

	/* zero pad if x is shorter than nfft */
	if(n < nfft)
	{
		if((d = (double *)calloc((size_t)nfft, sizeof(double))) == NULL)
			return 1;

		for(i = 0; i < n; i++)
			d[i] = x[i];

		for(; i < nfft; x++)
			d[i] = 0;

		n = nfft;
	}
	else
	{
		if((d = (double *)calloc((size_t)n, sizeof(double))) == NULL)
			return 1;

		for(i = 0; i < n; i++)
			d[i] = x[i];
	}

	numfreqs = nfft / 2 + 1;
	segments = n / nfft;

	window = hanning(nfft);

	norm = 0;
	for(i = 0; i < nfft; i++)
		norm += pow(window[i], NORM_FACTOR);

	if((p = (double *)calloc((size_t)(nfft * segments), sizeof(double))) == NULL)
	{
		free(d);
		free(window);
		return 1;
	}

	for(i = 0; i < numfreqs * segments; i++)
		p[i] = 0;

	for(i = 0; i < segments; i++)
	{
		loc = i * nfft;
		cur = &d[loc];
		pcur = &p[loc];

		for(j = 0; j < nfft; j++)
			cur[j] *= window[j];

		fft(nfft, cur, pcur);

		for(j = 0; j < numfreqs; j++)
			pcur[j] = pow(fabs(pcur[j]), 2) / norm;
	}

	ratio = (double)fs / (double)nfft;

	for(i = 0; i < numfreqs; i++)
	{
		freqs[i] = ratio * i;
		pxx[i] = 0;

		for(j = 0; j < segments; j++)
			pxx[i] += p[j * nfft + i];

		pxx[i] /= segments;
	}

	free(d);
	free(window);
	free(p);

	return 0;
}

double * hanning(size_t n)
{
	int i;
	double *d;

	if(n < 1)
		return NULL;

	if((d = (double *)calloc(n, sizeof(double))) == NULL)
		return NULL;

	for(i = 0; i < n; i++)
		d[i] = 0.5 - 0.5 * cos(2.0 * M_PI * i / (n - 1));

	return d;
}

double * readfile(char *fname, SNDFILE *s, struct SF_INFO *si)
{
	double *d, *in;
	sf_count_t c;
	int i, j;

	s = sf_open(fname, SFM_READ, si);

	if((d = (double *)calloc((size_t)(si->frames * si->channels), sizeof(double))) == NULL)
		return NULL;
	if((in = (double *)calloc((size_t)(si->frames), sizeof(double))) == NULL)
		return NULL;

	c = sf_read_double(s, d, (sf_count_t)(si->frames * si->channels));
	if(c != (si->frames * si->channels))
		return NULL;

	for(i = 0; i < si->frames; i++)
	{
		in[i] = 0;
		for(j = 0; j < si->channels; j++)
			in[i] += d[i * si->channels + j];
	}

	return in;
}

void fft(int n, double *in, double *out)
{
	fftw_plan p;

	p = fftw_plan_r2r_1d(n, in, out, FFTW_R2HC, FFTW_ESTIMATE);
	fftw_execute(p);
}
