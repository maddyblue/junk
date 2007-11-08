//Changes the box border when images are hovered over on the ads pages
function over(theImage){
	theVar = document.getElementById(theImage);
	theVar.className= "currentimage";
}
function out(theImage){
	theVar = document.getElementById(theImage);
	theVar.className= "notcurrentimage";
}
