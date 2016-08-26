PROG=		schalmei
SRCS=		schalmei.c freqs.c
CFLAGS=
COPTS+=		-g -I/usr/local/include
LDADD+=		-lsndfile -lfftw3 -lm
LDFLAGS+=	-L/usr/local/lib

.include <bsd.prog.mk>
