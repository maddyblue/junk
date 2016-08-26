from django import newforms as forms

class UploadForm(forms.Form):
	image = forms.FileField()
	name = forms.CharField(max_length=100, required=False)