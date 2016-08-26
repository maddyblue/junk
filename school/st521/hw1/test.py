from math import exp
from random import choice

def fac(n):
	if n <= 1:
		return 1
	else:
		return n * fac(n - 1)

def comb(n, k):
	return fac(n) / (fac(k) * fac(n - k))

def poi(k, l):
	return exp(-l) * l**k / fac(k)

class Card:
	def __init__(self, suit, rank):
		self.suit = suit
		self.rank = rank

	def __str__(self):
		if self.suit == 1:
			s = 'Hearts'
		elif self.suit == 2:
			s = 'Spades'
		elif self.suit == 3:
			s = 'Clubs'
		elif self.suit == 4:
			s = 'Diamonds'
		else:
			s = 'Nulls'

		return str(self.rank) + ' of ' + s

deck = []

for i in range(1, 5):
	for j in range(1, 14):
		deck.append(Card(i, j))

hands = []
cards = 2

for i in range(len(deck)):
	for j in range(i + 1, len(deck)):
		for m in range(j + 1, len(deck)):
			for n in range(m + 1, len(deck)):
				for o in range(n + 1, len(deck)):
					hands.append([deck[i], deck[j], deck[m], deck[n], deck[o]])

countH = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]

for i in hands:
	c = 0
	for j in i:
		if j.rank == 1:
			c += 1

	countH[c] += 1

print countH, sum(countH)

def draw(deck, hand):
	c = choice(deck)

	while c in hand:
		c = choice(deck)

	return c

def ctest(n):
	for i in range(1, 52):
		for j in range(1, 52):
			if comb(i, j) == n:
				print i, j
