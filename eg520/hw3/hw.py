def g(n):
	return [
		5. * n[0] + 2. * n[1] - 3.,
		2. * n[0] + n[1] - 1.
 ]

def Qconj(x, Q, b, dprev, i = 0):
	gg = g(x)

	if dprev != [0., 0.]:
		beta = (gg[1] * (Q[1][1] * dprev[1] + Q[0][1] * dprev[0]) + gg[0] * (Q[1][0] * dprev[1] + Q[0][0] * dprev[0])) / (dprev[1] * (Q[1][1] * dprev[1] + Q[0][1] * dprev[0]) + dprev[0] * (Q[1][0] * dprev[1] + Q[0][0] * dprev[0]))
	else:
		beta = 0.

	d = [-gg[0] + beta * dprev[0], -gg[1] + beta * dprev[1]]

	alpha = - (gg[0] * d[0] + gg[1] * d[1]) / (d[1] * (Q[1][1] * d[1] + Q[0][1] * d[0]) + d[0] * (Q[1][0] * d[1] + Q[0][0] * d[0]))
	xnew = [x[0] + alpha * d[0], x[1] + alpha * d[1]]
	print '$\\boldsymbol{g}^{(%i)} = %s^T, \\boldsymbol{d}^{(%i)} = %s, \\alpha_%i = %.4f, \\beta_%i = %.4f, \\boldsymbol{x}^{(%i)} = %s^T$ \\\\' % (i, repr(gg), i, repr(d), i, alpha, i, beta, i+1, repr(xnew))
	return (xnew, d)

x = [0., 0.]
Q = [[5., 2.], [2., 1.]]
b = [3., 1.]
d = [0., 0.]

(x, d) = Qconj(x, Q, b, d, 0)
(x, d) = Qconj(x, Q, b, d, 1)
(x, d) = Qconj(x, Q, b, d, 2)