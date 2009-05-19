set terminal png size 640, 480
set xlabel "Time/s"
set ylabel "Current/A"
set output "216.png"

plot [270:350] "216.avg" notitle with l
