chrome.browserAction.onClicked.addListener((tab)=>{
    chrome.tabs.query({}, (tabs)=>{
        var allPromises = [];
        tabs.forEach((tab)=>{
            allPromises.push(new Promise((resolve)=>{
                chrome.tabs.executeScript(
                    tab.id,
                    {file: 'scrape.js'},
                    (album)=>{
                        resolve(album)
                    }
                );
            }));
        });
        Promise.all(allPromises).then((results)=>{
            var allAlbums = {};
            results.forEach((result)=>{
                var album = result[0];
                allAlbums[album.name] = album;
            });
            downloadFileFromText("data.json", allAlbums);
        });
    });
});

function downloadFileFromText(filename, content) {
    var blob = new Blob([JSON.stringify(content)], {type: "application/json"});
    var url = URL.createObjectURL(blob);
    chrome.downloads.download({url: url, saveAs: true, filename: filename});
}
