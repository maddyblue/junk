set terminal png

set output "1a.png"
set xlabel "x1"
set ylabel "x2"
plot [0:2] 4 - 2*x, "data" title "interior point" with linespoints

set output "1d.png"
set xlabel "y1"
set ylabel "w"
plot [-.5:0][-2:0] "dataw" title "interior point" with linespoints
