#include <stdio.h>

#include <fftw3.h>
#include <sndfile.h>

#include "schalmei.h"
#include "freqs.h"

int main(int argc, char *argv[])
{
	SNDFILE *s;
	struct SF_INFO *si;
	double *in, *out;
	int i;

	if((si = (struct SF_INFO *)malloc(sizeof(*si))) == NULL)
		return 1;

	if((in = readfile("../wav/haupt-principal-8/c3.wav", s, si)) == NULL)
		return 1;

	if((out = (double *)calloc((size_t)(si->frames), sizeof(double))) == NULL)
		return 1;

	fft((int)(si->frames), in, out);

	return 0;
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
