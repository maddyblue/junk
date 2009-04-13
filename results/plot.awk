BEGIN {
	line = -1;
	FS = ", ";
	avg = 1;

	if(analysis == "Cyclic Voltammetry")
	{
		xlabel = "Potential/V";
	}
	else if(analysis == "i - t Curve")
	{
		xlabel = "Time/sec";
	}
	else if(analysis == "Differential Pulse Voltammetry")
	{
		xlabel = "Potential/V";
		avg = 0;
	}
}

FNR == 1 {
	fplt = FILENAME ".plt";
	print "set terminal png size 640, 480" > fplt;
	print "set xlabel \"" xlabel "\"" > fplt;
	print "set ylabel \"Current/A\"" > fplt;

	print "set output \"" FILENAME ".avg.png\"" > fplt;
	print "plot \"" FILENAME ".avg\" with lines" > fplt;
	print "set output \"" FILENAME ".png\"" > fplt;

	if(avg == 1)
	{
		print "plot \\" > fplt;
		print "\"" FILENAME ".dat1\" with lines , \\" > fplt;
		print "\"" FILENAME ".dat2\" with lines, \\" > fplt;
		print "\"" FILENAME ".dat3\" with lines, \\" > fplt;
		print "\"" FILENAME ".dat4\" with lines, \\" > fplt;
		print "\"" FILENAME ".dat5\" with lines" > fplt;
	}

	if(analysis == "Cyclic Voltammetry")
	{
		print "set output \"" FILENAME ".-1_1.png\"" > fplt;
		print "plot [-.1:.1] \"" FILENAME ".avg\" with lines" > fplt;

		print "set output \"" FILENAME ".-2_2.png\"" > fplt;
		print "plot [-.2:.2] \"" FILENAME ".avg\" with lines" > fplt;
	}
	else if(analysis == "i - t Curve")
	{
		print "set output \"" FILENAME ".+15.png\"" > fplt;
		print "plot [15:] \"" FILENAME ".avg\" with lines" > fplt;
		print "set output \"" FILENAME ".r5.png\"" > fplt;
		print "plot [:][-5e-10:5e-10] \"" FILENAME ".avg\" with lines" > fplt;

		if(high > 0)
		{
			print "set label 'base = " base "A' at " ctime ", " base " point lt 1 offset 1" > fplt;
			print "set label 'peak = " peak "A' at " ctime ", " peak " point lt 1 offset 1" > fplt;
			print "set label 'peak - base = " peak - base "' at " low ", " (base + peak) / 2 > fplt;
			print "set output \"" FILENAME ".c.png\"" > fplt;
			print "plot [" low - 2 ":" high + 5 "] \"" FILENAME ".avg\" with lines" > fplt;	
			print "unset label" > fplt;
		}
	}

	print "set terminal png size 200, 100" > fplt;
	print "set output \"" FILENAME ".tn.png\"" > fplt;
	print "set lmargin .2" > fplt;
	print "set rmargin .2" > fplt;
	print "set bmargin .2" > fplt;
	print "set tmargin .2" > fplt;
	print "unset xlabel" > fplt;
	print "unset ylabel" > fplt;
	print "unset tics" > fplt;
	print "plot \"" FILENAME ".avg\" notitle with lines" > fplt;
}

/^[-0-9]/ {
	if(avg == 0)
	{
		avgname = FILENAME ".avg";
		print $1 " " $2 > avgname;
	}
	else
	{
		line += 1;
		idx = line % 5;
		data[idx] = $2;
		time[idx] = $1;
		fname = FILENAME ".dat" idx + 1;
		print $1 " " $2 > fname;

		if(idx == 4)
		{
			avgname = FILENAME ".avg";
			range1 = FILENAME ".-1_1";
			range2 = FILENAME ".-2_2";

			avgtime = (time[0] + time[1] + time[2] + time[3] + time[4]) / 5;
			avgdata = (data[0] + data[1] + data[2] + data[3] + data[4]) / 5;
			print avgtime " " avgdata > avgname;

			if(avgtime >= -.2 && avgtime <= .2)
				print avgtime " " avgdata > range2;

			if(avgtime >= -.1 && avgtime <= .1)
				print avgtime " " avgdata > range1;
		}
	}
}
