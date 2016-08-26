from django import newforms as forms

class RegisterForm(forms.Form):
	username = forms.CharField(max_length=30)
	email = forms.CharField(max_length=75)
	password = forms.CharField(max_length=100, widget=forms.PasswordInput)