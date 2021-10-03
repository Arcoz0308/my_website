let raw = document.getElementById("raw-button");
let code_zone = document.getElementById("code-zone");

async function loadContent() {
    raw.setAttribute("href", "./raw" + window.location.pathname);
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
    let lines_text = "1";
    code_zone.innerHTML = text.trim();
    for (let i = 1; i < lines; i++) {
        lines_text = lines_text + `<br>${i + 1}`;
    }
    document.getElementById("lines").innerHTML = lines_text;
    init();
}
async function loadNew() {
    code_zone.style.display = "none";
    document.getElementById("lines").style.textAlign = "center";
    document.getElementById("options").style.display = "block"
    const textarea = document.getElementById("textarea");
    const button = document.getElementById("save");
    textarea.style.display = "block";
    textarea.addEventListener("input", function () {
        this.style.height = "auto";
        this.style.height = this.scrollHeight + "px";

    }, false);
    raw.classList.add("disable")
}

if (window.location.pathname !== "/new") loadContent().then()
else loadNew().then()