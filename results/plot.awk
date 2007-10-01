BEGIN {
	line = -1;
	FS = ", ";
}

FNR == 1 {
	fplt = FILENAME ".plt";
	print "set terminal png" > fplt;
	print "set xlabel \"Potential/V\"" > fplt;
	print "set ylabel \"Current/A\"" > fplt;

	print "set output \"" FILENAME ".png\"" > fplt;
	print "plot \\" > fplt;
	print "\"" FILENAME ".dat1\" with lines , \\" > fplt;
	print "\"" FILENAME ".dat2\" with lines, \\" > fplt;
	print "\"" FILENAME ".dat3\" with lines, \\" > fplt;
	print "\"" FILENAME ".dat4\" with lines, \\" > fplt;
	print "\"" FILENAME ".dat5\" with lines" > fplt;

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
}

/^[-0-9]/ {
	line += 1;
	fname = FILENAME ".dat" line % 5 + 1;
	print $1 " " $2 > fname;
}
