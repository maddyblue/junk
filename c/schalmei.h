#ifndef SCHALMEI_H
#define SCHALMEI_H

int psd(double *, int, int, int, double *, double *);
double * readfile(char *, SNDFILE *, struct SF_INFO *);
void fft(int, double *, double *);
double * hanning(size_t);
double * vander(int, double *, int);
int cmp(double, double);
int * peaks(int, int, double *);
double * invert(int, double *);
double * transpose(int, int, double *);
void pf(int, int, double *);
double * mult(int, int, double *, int, int, double *);
double * polyfit(int, double *, double *, int);

#endif
