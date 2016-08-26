[y,fs,bits] = wavread('wavs/a-sharp.wav');
y = y(:, 1);
n = 10;

spx = 1;
spy = 2;

%subplot(spy, spx, 1);
[p, perc, pxx, fxx] = fdisp(y, fs, n);
%subplot(spy, spx, 2);
[d, dpp, dff] = mksynth(1, fs, p, perc, 0);
%fit(dff, dpp, n);
%fit(fxx, pxx, n);

%subplot(2,1,2);
%semilogy(fxx, pxx);
%e = mkrange(1, fs, p, perc);
