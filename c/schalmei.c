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
	double *in, *pxx, *freqs, *out;
	int i, nfft, fs, n, numfreqs, *pks;

	/*

	#define M_xrows 3
	#define M_xcols 2
	#define M_yrows 2
	#define M_ycols 4

	double x[M_xrows * M_xcols] = {3, 4, 4, 5, 0, 3};
	double y[M_yrows * M_ycols] = {2, 6, -1, 0, 3, -2, 7, 1};

	pf(M_xrows, M_xcols, x);
	pf(M_yrows, M_ycols, y);

	out = mult(M_xrows, M_xcols, x, M_yrows, M_ycols, y);

	pf(M_xrows, M_ycols, out);

	//*/

	/* Vandermonde test
	if((in = (double *)calloc(4, sizeof(double))) == NULL)
		return 1;

	in[0] = 1;
	in[1] = 2;
	in[2] = 3;
	in[3] = 4;

	out = vander(4, in, 0);

	pf(4, 4, out);
	out = invert(4, out);
	pf(4, 4, out);

	if((in = (double *)calloc(16, sizeof(double))) == NULL)
		return 1;

	in = transpose(4, 4, out);
	pf(4, 4, in);

	return 0;

	//*/

	/*
	if((si = (struct SF_INFO *)malloc(sizeof(*si))) == NULL)
		return 1;

	if((in = readfile("../wav/haupt-principal-8/c3.wav", s, si)) == NULL)
		return 1;

	nfft = pow(2, 17);
	fs = si->samplerate;
	n = si->frames;

	numfreqs = nfft / 2 + 1;
	if((pxx = (double *)calloc((size_t)(numfreqs), sizeof(double))) == NULL)
		return 1;
	if((freqs = (double *)calloc((size_t)(numfreqs), sizeof(double))) == NULL)
		return 1;

	psd(in, n, nfft, fs, pxx, freqs);

	/*
	for(i = 0; i < numfreqs; i++)
		printf("%10.5f: %20.10f\n", freqs[i], pxx[i]);
	//*/

	/*
	n = 75;
	pks = peaks(n, numfreqs, pxx);

	for(i = 0; i < n; i++)
		printf("%2i: %10.5fHz (%08.5f)\n", i, freqs[pks[i]], pxx[pks[i]]);
	//*/

	#define X_LEN 4

	double x[X_LEN] = {1, 2, 3, 4};
	double y[X_LEN] = {5, 6, 7, 8};
	double *r;
	int order = 2;
	r = polyfit(X_LEN, x, y, order);
	pf(order, 1, r);

	return 0;
}

double * polyfit(int len, double *x, double *y, int order)
{
	double *result = NULL, *v, *xt, *xt_x, *xt_x_inv, *xt_x_inv_xt;

	if(order < 1 || order > len)
		return NULL;

	if((v = vander(len, x, order)) == NULL)
		goto polyfit_end;
	if((xt = transpose(len, len, v)) == NULL)
		goto free_v;
	if((xt_x = mult(len, len, xt, len, len, v)) == NULL)
		goto free_xt;
	if((xt_x_inv = invert(len, xt_x)) == NULL)
		goto free_xt_x;
	if((xt_x_inv_xt = mult(len, len, xt_x_inv, len, len, xt)) == NULL)
		goto free_xt_x_inv;
	result = mult(len, len, xt_x_inv_xt, len, 1, y);

	free(xt_x_inv_xt);
free_xt_x_inv:
	free(xt_x_inv);
free_xt_x:
	free(xt_x);
free_xt:
	free(xt);
free_v:
	free(v);

polyfit_end:
	return result;
}

double * mult(int xrows, int xcols, double *x, int yrows, int ycols, double *y)
{
	double *result;
	int i, j, m, n;

	if(xcols != yrows)
		return NULL;

	if((result = (double *)calloc((size_t)(xrows * ycols), sizeof(double))) == NULL)
		return NULL;

	for(i = 0; i < xrows; i++)
	{
		for(j = 0; j < ycols; j++)
		{
			result[i * ycols + j] = 0;

			for(m = 0; m < xcols; m++)
				result[i * ycols + j] += x[i * xcols + m] * y[m * ycols + j];
		}
	}

	return result;
}

void pf(int rows, int cols, double *x)
{
	int i, j;

	if(x == NULL)
	{
		printf("NULL\n");
		return;
	}

	for(i = 0; i < rows; i++)
	{
		for(j = 0; j < cols; j++)
			printf("%5.2f ", x[i * cols + j]);
		printf("\n");
	}
	printf("\n");
}

double * transpose(int rows, int cols, double *x)
{
	double *result;
	int i, j;

	if((result = (double *)calloc((size_t)(cols * rows), sizeof(double))) == NULL)
		return NULL;

	for(i = 0; i < rows; i++)
		for(j = 0; j < cols; j++)
			result[j * rows + i] = x[i * cols + j];

	return result;
}

double * invert(int size, double *x)
{
	double *result, *work, *temp;
	double c;
	int cols = size * 2;
	int i, j, k, lead = 0;

	if((work = (double *)calloc((size_t)(cols * size), sizeof(double))) == NULL)
		return NULL;

	if((result = (double *)calloc((size_t)(cols * size), sizeof(double))) == NULL)
	{
		free(work);
		return NULL;
	}

	if((temp = (double *)calloc((size_t)size, sizeof(double))) == NULL)
	{
		free(work);
		free(result);
		return NULL;
	}

	for(i = 0; i < size; i++)
	{
		for(j = 0; j < size; j++)
		{
			work[i * cols + j] = x[i * size + j];
			work[i * cols + j + size] = 0;
		}

		work[i * cols + size + i] = 1;
	}

	for(j = 0; j < size; j++)
	{
		if(cols <= lead)
			break;

		i = j;
		while(work[i * cols + lead] == 0)
		{
			i++;
			if(size == i)
			{
				i = j;
				lead++;
				if(cols == lead)
					break;
			}
		}

		// swap rows i and j
		for(k = 0; k < cols; k++)
		{
			temp[i * cols + k] = work[i * cols + k];
			work[i * cols + k] = work[j * cols + k];
			work[j * cols + k] = temp[i * cols + k];
		}

		// divide row j by work[j][lead]
		c = work[j * cols + lead];
		for(k = 0; k < cols; k++)
			work[j * cols + k] /= c;

		for(i = 0; i < size; i++)
		{
			if(i != j)
			{
				// subtract work[i][lead] multiplied by row j from row i
				c = work[i * cols + lead];
				for(k = 0; k < cols; k++)
					work[i * cols + k] -= work[j * cols + k] * c;
			}
		}

		lead++;
	}

	for(i = 0; i < size; i++)
		for(j = 0; j < size; j++)
			result[i * size + j] = work[i * cols + j + size];

	return result;
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

		for(; i < nfft; i++)
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

double * vander(int n, double *in, int order)
{
	double *r;
	int i, j;

	if(order == 0)
		order = n;
	else if(order > n || order < 0)
		return NULL;

	if((r = (double *)calloc(order * n, sizeof(double))) == NULL)
		return NULL;

	for(i = 0; i < order; i++)
		for(j = 0; j < n; j++)
			r[j * n + i] = pow(in[j], order - i - 1);

	return r;
}

int cmp(double a, double b)
{
	if(a < b)
		return -1;
	else if(a > b)
		return 1;

	return 0;
}

int * peaks(int n, int len, double *in)
{
	int i, idx, s, sprev, *indicies;
	double prev, cur, dif;
	double *pp, *rr;

	if((pp = (double *)calloc(len, sizeof(double))) == NULL)
		return NULL;
	if((rr = (double *)calloc(len, sizeof(double))) == NULL)
		return NULL;
	if((indicies = (int *)calloc(n, sizeof(int))) == NULL)
		return NULL;

	sprev = prev = 0;
	idx = 0;

	for(i = 0; i < len; i++)
	{
		cur = in[i];
		dif = cur - prev;
		s = cmp(dif, 0);

		// at a peak
		if(s == -1 && sprev == 1)
		{
			pp[idx] = i - 1;
			rr[idx] = in[i - 1];
			idx++;
		}
		// peak at end of data
		else if(s == 1 && i == len)
		{
			pp[idx] = i;
			rr[idx] = cur;
			idx++;
		}

		prev = cur;
		sprev = s;
	}

	for(i = 0; i < n; i++)
	{
		// find max value
		sprev = 0;
		prev = pp[0];
		for(s = 0; s < idx; s++)
		{
			if(rr[s] > prev)
			{
				sprev = s;
				prev = rr[s];
			}
		}

		indicies[i] = pp[sprev];
		rr[sprev] = 0;
	}

	free(pp);
	free(rr);

	return indicies;
}
