#include <math.h>
#include <stdio.h>
#include <string.h>

#include <fftw3.h>
#include <sndfile.h>

#include "schalmei.h"
#include "freqs.h"

int main(int argc, char *argv[])
{
	#if 0 // {{{ dumpfiles
	SNDFILE *s;
	struct SF_INFO *si;
	double *in;
	int i, nfft, fs, n, numfreqs, *pks;
	char buf[128];
	char src[] = "../wav/haupt-principal-8/";
	char dst[] = "dump/";
	#define NUMFILES 15

	char files[NUMFILES][4] = {
		"c#3",
		"d#3",
		"f3",
		"g3",
		"a3",
		"b3",
		"c#4",
		"c#5",
		"c#6",
		"d#6",
		"f6",
		"g6",
		"a6",
		"b6",
		"c#7"
	};

	if((si = (struct SF_INFO *)malloc(sizeof(*si))) == NULL)
		return 1;

	nfft = pow(2, 17);

	for(i = 0; i < NUMFILES; i++)
	{
		buf[0] = '\0';
		strcat(buf, src);
		strcat(buf, files[i]);
		strcat(buf, ".wav");
		printf("%i %s -> ", i, buf);

		if((in = readfile(buf, s, si)) == NULL)
			return 1;

		buf[0] = '\0';
		strcat(buf, dst);
		strcat(buf, files[i]);
		strcat(buf, ".dump");
		printf("%s\n", buf);

		dumpfile(si->frames, in, si->samplerate, buf);
	}
	#endif // }}}

	#if 1 // {{{ readfiles
	#define NUMFILES 11
	#define NUMPEAKS 10
	#define DEGREES 3

	double *data, *pxx, *freqs, *peak_freqs, *percs, energy, last, cur, bases[NUMFILES], y[NUMFILES], *polyfits;
	int i, j, len, nfft, fs, numpeaks, numfreqs, *pks;
	FILE *f;
	char buf[128];
	char src[] = "dump/";

	char files[NUMFILES][4] = {
		"c#3",
		"d#3",
		"f3",
		"g3",
		"a3",
		//"b3",
		"c#4",
		"c#5",
		"c#6",
		"d#6",
		"f6",
		"g6"/*,
		"a6",
		"b6",
		"c#7"*/
	};

	double files_base[NUMFILES] = {
		137.27,
		154.77,
		173.61,
		195.14,
		218.36,
		//246,
		275.22,
		549.77,
		1098.53,
		1233.45,
		1387.88,
		1558.13/*,
		1748,
		1962,
		2197*/
	};

	nfft = pow(2, 17);
	numfreqs = nfft / 2 + 1;

	if((pxx = (double *)calloc((size_t)(numfreqs), sizeof(double))) == NULL)
		return 1;
	if((freqs = (double *)calloc((size_t)(numfreqs), sizeof(double))) == NULL)
		return 1;
	if((peak_freqs = (double *)calloc((size_t)(NUMFILES * NUMPEAKS), sizeof(double))) == NULL)
		return 1;
	if((percs = (double *)calloc((size_t)(NUMFILES * NUMPEAKS), sizeof(double))) == NULL)
		return 1;
	if((polyfits = (double *)calloc((size_t)(NUMPEAKS * DEGREES), sizeof(double))) == NULL)
		return 1;

	for(i = 0; i < NUMFILES; i++)
	{
		buf[0] = '\0';
		strlcat(buf, src, sizeof(buf));
		strlcat(buf, files[i], sizeof(buf));
		strlcat(buf, ".dump", sizeof(buf));
		printf("%15s: ", buf);

		if((data = readdumpfile(buf, &len, &fs)) == NULL)
			return 1;

		printf("%5i samples, %5iHz ", len, fs);

		j = psd(data, len, nfft, fs, pxx, freqs);

		if(j != 0)
		{
			printf("failed\n");
			continue;
		}

		pks = peaks(NUMPEAKS, numfreqs, pxx, freqs);

		for(j = 0, energy = 0; j < numfreqs; j++)
			energy += pxx[j];

		bases[i] = 0;
		last = 10000;
		for(j = 0; j < NUMPEAKS; j++)
		{
			peak_freqs[i * NUMPEAKS + j] = freqs[pks[j]];
			percs[i * NUMPEAKS + j] = pxx[pks[j]] / energy;

			//printf("\t%i: %7.2f, %5.2f\n", j, peak_freqs[i * NUMPEAKS + j], percs[i * NUMPEAKS + j]);

			cur = fabs(files_base[i] - freqs[pks[j]]);
			if(cur < last)
			{
				last = cur;
				bases[i] = freqs[pks[j]];
			}
		}

		free(data);
		free(pks);

		printf("base: %6.1f -> %7.2f\n", files_base[i], bases[i]);
	}

	for(i = 0; i < NUMPEAKS; i++)
	{
		for(j = 0; j < NUMFILES; j++)
		{
			y[j] = percs[j * NUMPEAKS + i];
			peak_freqs[i * NUMFILES + j] /= bases[j];
		}

		data = polyfit(NUMFILES, bases, y, DEGREES);

		for(j = 0; j < DEGREES; j++)
			polyfits[i * DEGREES + j] = data[j];

		free(data);
	}

	if((f = fopen("out.ck", "w")) == NULL)
		return NULL;

	fprintf(f, "%i => int peaks;\n", NUMPEAKS);
	fprintf(f, "%i => int recs;\n", NUMFILES);
	fprintf(f, "%i => int degs;\n", DEGREES - 1);

	fprintf(f, "\n[ ");
	chuckmatrix(f, NUMPEAKS, DEGREES, polyfits);
	fprintf(f, " ] @=> float percs[][];\n");

	fprintf(f, "\n[ ");
	chuckmatrix(f, NUMPEAKS, NUMFILES, peak_freqs);
	fprintf(f, " ] @=> float harms[][];\n");

	fprintf(f, "\n[ ");
	for(i = 0; i < NUMFILES; i++)
	{
		if(i > 0)
			fprintf(f, ", ");

		fprintf(f, "%17.15f", bases[i]);
	}
	fprintf(f, " ] @=> float bases[];\n");

	fprintf(f, chuck_file);

	fclose(f);

	#endif // }}}
}

void chuckmatrix(FILE * f, int r, int c, double *x)
{
	int i, j;

	for(i = 0; i < r; i++)
	{
		if(i > 0)
			fprintf(f, ", ");

		fprintf(f, "[ ");
		for(j = 0; j < c; j++)
		{
			if(j > 0)
				fprintf(f, ", ");

			fprintf(f, "%17.15f", x[i * c + j]);
		}
		fprintf(f, " ]");
	}
}

double * readdumpfile(char *fname, int *len, int *fs)
{
	double *x, *ret = NULL;
	FILE *f;
	int i;

	if((f = fopen(fname, "r")) == NULL)
		return NULL;

	if(2 != fscanf(f, "%i %i\n", len, fs))
		goto readdump_end;

	if((x = (double *)calloc((size_t)*len, sizeof(double))) == NULL)
		goto readdump_end;

	for(i = 0; i < *len; i++)
		if(1 != fscanf(f, "%lf\n", &x[i]))
			goto readdump_end;

	ret = x;

readdump_end:
	fclose(f);
	return ret;
}

int dumpfile(int len, double *x, int fs, char *fname)
{
	FILE *f;
	int i;

	if((f = fopen(fname, "w")) == NULL)
		return 1;

	fprintf(f, "%i %i\n", len, fs);

	for(i = 0; i < len; i++)
		fprintf(f, "%17.15f\n", x[i]);

	fclose(f);

	return 0;
}

double * polyfit(int len, double *x, double *y, int order)
{
	double *result = NULL, *v, *xt, *xt_x, *xt_x_inv, *xt_x_inv_xt;

	if(order < 1 || order > len)
		return NULL;

	if((v = vander(len, x, order)) == NULL)
		goto polyfit_end;
	if((xt = transpose(len, order, v)) == NULL)
		goto free_v;
	if((xt_x = mult(order, len, xt, len, order, v)) == NULL)
		goto free_xt;
	if((xt_x_inv = invert(order, xt_x)) == NULL)
		goto free_xt_x;
	if((xt_x_inv_xt = mult(order, order, xt_x_inv, order, len, xt)) == NULL)
		goto free_xt_x_inv;
	result = mult(order, len, xt_x_inv_xt, len, 1, y);

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
			printf("%12.10f ", x[i * cols + j]);
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

	if((temp = (double *)calloc((size_t)cols, sizeof(double))) == NULL)
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
			temp[k] = work[i * cols + k];
			work[i * cols + k] = work[j * cols + k];
			work[j * cols + k] = temp[k];
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

	free(work);
	free(temp);

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

	for(i = 0; i < n; i++)
		for(j = 0; j < order; j++)
			r[i * order + j] = pow(in[i], order - j - 1);

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

int * peaks(int n, int len, double *in, double *freqs)
{
	int i, idx, s, sprev, *indicies;
	double prev, cur, dif;
	int *pp;
	double *rr;

	if((pp = (int *)calloc(len, sizeof(int))) == NULL)
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
