function reddenPage() {
  var key = "";
  key = window.getSelection().toString();

let params = `scrollbars=no,resizable=no,status=no,location=no,toolbar=no,menubar=no,
width=600,height=300,left=100,top=100`;
  // console.log(doc)
  let w = window.open("http://localhost:8080/c?key=" + key, "YWSROCK", params)
  w.onblur = function() {
    this.close()
  }
}

chrome.action.onClicked.addListener((tab) => {
  if (!tab.url.includes("chrome://")) {
    chrome.scripting.executeScript({
      target: { tabId: tab.id },
      function: reddenPage
    });
  }
});
