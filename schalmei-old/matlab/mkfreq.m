function [d, dpp, dff] = mkfreq(len, fs, p, perc, f);

ratio = f / p(1);
pp = p * ratio;
[d, dpp, dff] = mksynth(len, fs, pp, perc, 1);
