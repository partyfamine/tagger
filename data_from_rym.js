var tracks = document.getElementById("tracks").getElementsByClassName("track");
var trackTitles = [];
for (let i = 0; i < tracks.length; i++) {
    var fullTitleElem = tracks[i].getElementsByClassName("rendered_text")[0];
    if(fullTitleElem == undefined) {
        continue;
    }

    var childTitleElem = fullTitleElem.childNodes[0];
    if(childTitleElem.nodeValue == null) {
        continue;
    }
    var trackTitle = childTitleElem.nodeValue.trim();

    var featElem = tracks[i].getElementsByClassName("featured_credit")[0];
    if(featElem == undefined) {
        trackTitles.push(trackTitle);
        continue;
    }
    
    var childFeatElems = featElem.children;
    var featStr = "(";
    for (let i = 0; i < childFeatElems.length; i++) {
        featStr += childFeatElems[i].innerHTML;

        if(i == 0) {
            featStr += " ";
        } else if(i < (childFeatElems.length - 2)) {
            featStr += ", ";
        } else if(i == (childFeatElems.length - 2)) {
            featStr += " & ";
        } else {
            featStr += ")";
        }
    }
    trackTitles.push(trackTitle + " " + featStr)
}

var albumName = elementText("album_title");
var albumData = {
    name: albumName,
    artist: elementText("artist"),
    year: elementText("issue_year"),
    genre: elementText("genre"),
    trackNames: trackTitles
};

console.log('"' + albumName + '":' + JSON.stringify(albumData));

function elementText(className) {
    var fullElem = document.getElementsByClassName(className)[0];
    var childElem = fullElem.childNodes[0];
    return childElem.nodeValue.trim();
}
