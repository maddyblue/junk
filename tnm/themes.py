THEME_GENESIS = 'genesis'
THEMES = [
	THEME_GENESIS,
]

COLORS = {
	THEME_GENESIS: [
		'genesis',
		'orange-grey-dark',
		'orange-grey-light',
		'green-grey-light',
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
PAGE_TYPE_CONTACT = 'contact'
PAGE_TYPE_GALLERY = 'gallery'
PAGE_TYPE_HOME = 'home'
PAGE_TYPE_TEXT = 'text'
PAGE_TYPES = [
	PAGE_TYPE_BLOG,
	PAGE_TYPE_CONTACT,
	PAGE_TYPE_GALLERY,
	PAGE_TYPE_HOME,
	PAGE_TYPE_TEXT,
]

LOREM = [
	'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at augue justo. Donec pretium ipsum ac tortor gravida vitae egestas nulla laoreet. Fusce quis leo non augue sollicitudin imperdiet. Cras eu fringilla massa. Duis tellus felis, pretium vitae vulputate tincidunt, vulputate sed erat. Fusce adipiscing dignissim interdum. Vestibulum sit amet arcu turpis. Praesent augue lectus, tincidunt sed pretium eu, laoreet id urna. Vivamus consequat ante eu urna vestibulum et posuere mi iaculis. Quisque blandit, nisl in tempor ornare, mi elit ornare augue, non feugiat felis justo semper nulla. Praesent euismod eleifend justo, a cursus mi hendrerit et. Phasellus tempus mollis hendrerit. Curabitur venenatis nulla a nulla imperdiet ultricies. Aliquam rutrum dui et purus ultricies id varius massa cursus. Vivamus nec auctor ligula. Vestibulum venenatis augue ac erat euismod ut euismod lacus lacinia.',
	'Cras scelerisque mattis mi, ut ultrices lorem cursus id. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Etiam in tincidunt ligula. Nulla pulvinar elementum enim, ac suscipit nisi ornare volutpat. Vivamus tincidunt tortor vel lorem ultricies eget fermentum diam iaculis. Nunc adipiscing orci eget lectus tristique tincidunt. Vestibulum imperdiet lectus mauris.',
	'Maecenas fermentum venenatis tortor rhoncus mollis. Sed tempus magna eget nisi dictum consequat. Sed ac nibh ut elit consequat vehicula. Quisque nec nulla a eros dapibus aliquet id at nunc. Etiam at est eget enim iaculis vehicula at at purus. Phasellus diam metus, ultricies et tristique eget, luctus sed dolor. Pellentesque ultrices luctus ante at gravida. In vel enim lectus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Phasellus venenatis porttitor ante suscipit mollis. Nulla at elementum massa. Aenean eu justo at nibh blandit commodo. Mauris eu quam vel est mattis semper id volutpat ipsum. Nam tempus sagittis neque vel fermentum. Etiam vestibulum eros sit amet neque venenatis vel malesuada ligula adipiscing.',
	'Nullam ac quam at nibh varius vestibulum nec at orci. Quisque vitae enim vel quam tincidunt aliquam. Aenean tincidunt pellentesque lacus id laoreet. Morbi vulputate nulla felis. Etiam elementum purus a elit dictum varius. Quisque consequat tempor cursus. Etiam lacus nisi, ultrices ut elementum eu, consectetur ultrices lorem. Aenean lobortis nibh in nunc blandit gravida. Cras semper eros in nisi consectetur viverra.',
	'Phasellus luctus dapibus arcu nec malesuada. Phasellus sodales lectus in neque accumsan id sodales lacus malesuada. Vestibulum porttitor interdum lacus, non porta augue bibendum et. Proin auctor lacinia urna, sed tempus purus viverra id. Suspendisse ac tellus sit amet ligula tristique tincidunt. Maecenas ut risus vel nibh malesuada faucibus. Vivamus in nunc fringilla felis aliquam porttitor. Fusce molestie justo at urna posuere vel eleifend lacus euismod. Ut nec gravida tortor. Mauris dolor neque, pretium a porttitor ac, congue ut dui. Suspendisse interdum commodo posuere. Duis in augue sem.',
]

TYPE_DEFAULTS = {
	'text': '<p>'.join(LOREM[:2]),
	'lines': 'Headline text',
}

SPECS = {
	THEME_GENESIS: {
		PAGE_TYPE_HOME: {
			1: { # no sidebar
				'lines': 3,
				'images': [
					(1000, 370),
					(310, 180),
					(310, 180),
					(310, 180),
				],
			},
			2: { # sidebar
				'lines': 3,
				'images': [
					(654, 243),
					(200, 116),
					(200, 116),
					(200, 116),
				],
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
					(650, 163),
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
		PAGE_TYPE_CONTACT: {
			1: { # right map
				'lines': 3,
				'maps': 1,
				'text': 1,
				'images': [
					(1000, 165),
				],
			},
			2: { # left map
				'lines': 4,
				'maps': 1,
				'text': 2,
			},
			3: { # top map
				'lines': 3,
				'maps': 1,
				'text': 1,
			},
			4: { # no map
				'lines': 3,
				'text': 1,
				'images': [
					(1000, 165),
				],
			},
		},
	},
}

def spec(theme, pagetype, layout):
	r = SPECS.get(theme, {}).get(pagetype, {}).get(layout, {})
	r['links'] = len(r.get('images', []))
	return r

def layouts(theme, pagetype):
	return len(SPECS.get(theme, {}).get(pagetype, {}))

def types(theme):
	return SPECS.get(theme, {})

def colors(theme):
	return COLORS.get(theme, {})
