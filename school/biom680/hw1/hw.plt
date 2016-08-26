set terminal png

set output "5ai.png"
set xlabel "Ca"
set ylabel "r"
plot [0:0.5] 0.05 notitle

set output "5aii.png"
set xlabel "t (h)"
set ylabel "Ca"
plot [0:10] 0.5-.05*x notitle

set output "5aiii.png"
set xlabel "t (h)"
set ylabel "r"
plot [0:10] 0.05 notitle

set output "5bi.png"
set xlabel "Ca"
set ylabel "r"
plot [0:0.5] 0.06*x notitle

set output "5bii.png"
set xlabel "t (h)"
set ylabel "Ca"
plot [0:10] .1*-log(x*.06) notitle

set output "5biii.png"
set xlabel "t (h)"
set ylabel "r"
plot [0:10] .1*-log(x*.06) notitle

set output "5ci.png"
set xlabel "Ca"
set ylabel "r"
plot [0:0.5] exp(0.4*x)

set output "5cii.png"
set xlabel "t (h)"
set ylabel "Ca"
plot [0:10] -.2*log(x*.4) notitle

set output "5ciii.png"
set xlabel "t (h)"
set ylabel "r"
plot [0:10] -.2*log(x**.4) notitle

set output "6i.png"
set xlabel "Ca"
set ylabel "r"
plot [0:10] (0.54*x)/(1.8+x) notitle

set output "6ii.png"
set xlabel "t (m)"
set ylabel "Ca"
plot [0:5] 10-(0.54*x)/(1.8+x) notitle

set output "6iii.png"
set xlabel "t (m)"
set ylabel "r"
plot [0:5] .972/(x**2+3.6*x+3.24) notitle