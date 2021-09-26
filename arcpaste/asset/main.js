let raw = document.getElementById("raw-button");
raw.setAttribute("href", "./raw" + window.location.pathname);
let code_zone = document.getElementById("code-zone");

async function main() {
    let r = await fetch("http://localhost:8081/arcpaste" + window.location.pathname, {
        mode: "cors"
    });
    let text, lang;
    if (r.status === 404) {
        text = "not found";
        lang = "markdown";
    } else if (r.status === 500) {
        text = "internal server error";
        lang = "markdown";
    } else {
        let j = await r.json();
        text = j.content;
        lang = j.language;
    }
    code_zone.classList.add("language-" + lang);
    let lines = text.split("\n").length;
    code_zone.innerHTML = text;
    let lines_text = "<br>1";
    for (let i = 1; i < lines; i++) {
        lines_text = lines_text + `<br>${i + 1}`;
    }
    document.getElementById("lines").innerHTML = lines_text
    init()
}

main()