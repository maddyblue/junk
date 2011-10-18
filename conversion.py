# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>
#
# Permission to use, copy, modify, and distribute this software for any
# purpose with or without fee is hereby granted, provided that the above
# copyright notice and this permission notice appear in all copies.
#
# THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
# WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
# MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
# ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
# WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
# ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
# OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

import logging

from google.appengine.api import conversion
from google.appengine.ext import blobstore

import models

def entry_doc(entry, content, blobs, output_type='application/pdf'):
	asset = conversion.Asset('text/html', content.rendered)
	conversion_request = conversion.ConversionRequest(asset, output_type)

	for b in blobs:
		if b.type == models.BLOB_TYPE_IMAGE:
			data = blobstore.BlobReader(b.blob)
			conversion_request.add_asset(conversion.Asset(b.blob.content_type, data.read()))

	result = conversion.convert(conversion_request)

	logging.error('Conversion assets: %s', len(result.assets))

	if result and result.assets:
		return ''.join([i.data for i in result.assets])

	logging.error('Conversion error: %s', result.error_text)
	return None
