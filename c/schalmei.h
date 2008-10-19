#ifndef SCHALMEI_H
#define SCHALMEI_H

int psd(double *, int, int, int, double *, double *);
double * readfile(char *, SNDFILE *, struct SF_INFO *);
void fft(int, double *, double *);
double * hanning(size_t);

#endif
