BEGIN {
	line = -1;
	FS = ", ";
}

FNR == 1 {
	fplt = FILENAME ".plt";
	print "set terminal png size 640, 480" > fplt;
	print "set xlabel \"Potential/V\"" > fplt;
	print "set ylabel \"Current/A\"" > fplt;

	print "set output \"" FILENAME ".dat1.png\"" > fplt;
	print "plot \"" FILENAME ".dat1\" with lines" > fplt;

	print "set output \"" FILENAME ".dat2.png\"" > fplt;
	print "plot \"" FILENAME ".dat2\" with lines" > fplt;

	print "set output \"" FILENAME ".dat3.png\"" > fplt;
	print "plot \"" FILENAME ".dat3\" with lines" > fplt;

	print "set output \"" FILENAME ".dat4.png\"" > fplt;
	print "plot \"" FILENAME ".dat4\" with lines" > fplt;

	print "set output \"" FILENAME ".dat5.png\"" > fplt;
	print "plot \"" FILENAME ".dat5\" with lines" > fplt;

	print "set output \"" FILENAME ".avg.png\"" > fplt;
	print "plot \"" FILENAME ".avg\" with lines" > fplt;

	print "set output \"" FILENAME ".png\"" > fplt;
	print "plot \\" > fplt;
	print "\"" FILENAME ".dat1\" with lines , \\" > fplt;
	print "\"" FILENAME ".dat2\" with lines, \\" > fplt;
	print "\"" FILENAME ".dat3\" with lines, \\" > fplt;
	print "\"" FILENAME ".dat4\" with lines, \\" > fplt;
	print "\"" FILENAME ".dat5\" with lines" > fplt;

	print "set terminal png size 200, 100" > fplt;
	print "set output \"" FILENAME ".tn.png\"" > fplt;
	print "set lmargin .2" > fplt;
	print "set rmargin .2" > fplt;
	print "set bmargin .2" > fplt;
	print "set tmargin .2" > fplt;
	print "unset xlabel" > fplt;
	print "unset ylabel" > fplt;
	print "unset tics" > fplt;
	print "plot \\" > fplt;
	print "\"" FILENAME ".avg\" notitle with lines" > fplt;
}

/^[-0-9]/ {
	line += 1;
	idx = line % 5;
	data[idx] = $2;
	time[idx] = $1;
	fname = FILENAME ".dat" idx + 1;
	print $1 " " $2 > fname;

	if(idx == 4)
	{
		avgname = FILENAME ".avg";
		print ((time[0] + time[1] + time[2] + time[3] + time[4]) / 5) " " ((data[0] + data[1] + data[2] + data[3] + data[4]) / 5) > avgname;
	}
}
