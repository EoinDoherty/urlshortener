// plain old js
window.onload = function() {
    document.getElementById("submit-btn").addEventListener("click", () => {
        let input = document.getElementById("input-text").value;

        if (input.length > 0) {
            submitURL(input);
        }
    });
}

async function submitURL(url) {
    const id = await fetch("shorten", {
        "body": url,
        "method": "POST"
    }).then(data => data.text());
    console.log(id);
    
    const newURL = window.location + "s/" + id;
    const elem = document.getElementById("shortened-url");
    
    elem.innerHTML = newURL;
    elem.setAttribute("href", newURL);
}