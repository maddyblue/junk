function [wav, x] = mkwav(freq, len, fs)
%freq
degrees = 360;
x = 0:degrees/fs:degrees*len;
x = x/degrees;
wav = sin(pi*2*freq*x);
