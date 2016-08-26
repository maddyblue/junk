set terminal png

set output "2a.png"
set xlabel "y1"
set ylabel "y2"
plot [-5:0][-5:2] (-1+x)/2, x+1

set output "2b.png"
set xlabel "x1"
set ylabel "x2"
plot [0:10][0:10] 1+x, 2*x-2

set output "3a.png"
set xlabel "y1"
set ylabel "y2"
plot [0:6][0:6] 2+2*x, (9-2*x)/4, -3+x

set output "3b.png"
set xlabel "x1"
set ylabel "x2"
plot [-1:5][-1:5] (1+2*x)/2, (1-x)/4

set output "4a.png"
set xlabel "y1"
set ylabel "y2"
plot [0:6][0:3] 2-(3*x/7), (9-2*x)/4, -3+x

set output "4b.png"
set xlabel "x1"
set ylabel "x2"
plot [0:2.5][0:.5] (1-3*x/7)/2, (1-x)/4
