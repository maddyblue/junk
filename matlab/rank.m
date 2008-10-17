rank_data;
start = 0;
load('freqs.mat');
inx = 2;
iny = 3;
inc = 4;
fdg = 0.1; % Gedakt pedal must be scaled down in amplitude in order to sound correctly

for ww = 1:length(files)
    if ww > length(files)/3
		break;
	end
	if start == 0
        x = files{ww};
        z = files{ww + 1};
		c = files{ww + 2};
        [bundle, batch] = process_rank(x, z, c, freqs, 1);
        big(:, :, ww) = bundle;
		series(:, :, ww) = batch;
	elseif ww == 12
        x = files{ww + inx};
        z = files{ww + iny};
		c = files{ww + inc};
        [bundle, batch] = process_rank(x, z, c, freqs, fdg);
        inx = inx + 2;
        iny = iny + 2;
		inc = inc + 2;
	else
		x = files{ww + inx};
        z = files{ww + iny};
		c = files{ww + inc};
        [bundle, batch] = process_rank(x, z, c, freqs, 1);
        inx = inx + 2;
        iny = iny + 2;
		inc = inc + 2;
    end
    
    start = 1;
    big(:, :, ww) = bundle;
	series(:, :, ww) = batch;
end