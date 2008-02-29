from django import newforms as forms

class PasswordForm(forms.Form):
	new_password = forms.CharField(max_length=100, widget=forms.PasswordInput)
	confirm_password = forms.CharField(max_length=100, widget=forms.PasswordInput)

	def clean(self):
		np = ''
		cp = ''

		for k, v in self.cleaned_data.iteritems():
			if k == 'new_password':
				np = v
			elif k == 'confirm_password':
				cp = v

		if np != cp:
			raise forms.ValidationError('Passwords do not match.')

		return self.cleaned_data