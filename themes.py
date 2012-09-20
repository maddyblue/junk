THEME_MARCO = 'marco'
THEMES = [
	THEME_MARCO,
]

COLORS = {
	THEME_MARCO: [
		'marco',
	],
}

NAV_LEFT = 'left'
NAV_RIGHT = 'right'
NAV_TOP = 'top'
NAVS = [
	NAV_LEFT,
	NAV_RIGHT,
	NAV_TOP,
]

PAGE_TYPE_BLOG = 'blog'
PAGE_TYPE_GALLERY = 'gallery'
PAGE_TYPE_HOME = 'home'
PAGE_TYPE_TEXT = 'text'
PAGE_TYPES = [
	PAGE_TYPE_BLOG,
	PAGE_TYPE_GALLERY,
	PAGE_TYPE_HOME,
	PAGE_TYPE_TEXT,
]

SPECS = {
	THEME_MARCO: {
		PAGE_TYPE_HOME: {
			1: { # blog section at bottom
				'links': 3,
				'images': [
					(1000, 370),
					(310, 180),
					(310, 180),
					(310, 180),
				],
			},
			2: { # pic and text at bottom
				'links': 5,
				'images': [
					(1000, 370),
					(310, 180),
					(310, 180),
					(310, 180),
					(310, 180),
				],
				'text': 2,
			},
		},
		PAGE_TYPE_TEXT: {
			1: {
				'images': [
					(310, 450),
				],
				'text': 1,
				'lines': 1,
			},
			2: {
				'images': [
					(660, 165),
					(310, 180),
					(310, 180),
				],
				'text': 4,
				'lines': 2,
			},
			3: {
				'images': [
					(1000, 250),
					(310, 336),
				],
				'text': 4,
				'lines': 4,
			},
		},
		PAGE_TYPE_GALLERY: {
			1: {
				'pass': True,
			},
			2: {
				'lines': 2,
				'rowsz': 3,
				'rows': 3,
			},
		},
		PAGE_TYPE_BLOG: {
			1: {
				'lines': 1,
				'text': 1,
				'images': [
					(96, 92),
				],
				# todo: support image resizing on page layout change
				'postimagesz': (634, 172),
			},
		},
	},
}

def spec(theme, pagetype, layout):
	return SPECS.get(theme, {}).get(pagetype, {}).get(layout, {})

def layouts(theme, pagetype):
	return len(SPECS.get(theme, {}).get(pagetype, {}))

def types(theme):
	return SPECS.get(theme, {})
