// listener ctrl key + /
let mark = true;
let isShow = false;
let article = null;
document.addEventListener('keydown', function (e) {
    if (e.ctrlKey && e.keyCode === 191) {
        if (mark) {
            init()
        }
        mark = false;
        if (isShow) {
            setDisplay(true);
        } else {
            setDisplay(false);
        }
    }
});

function init() {
    console.log('init')
    loadCss();
    buildText();
}

function loadCss() {
    const head = document.getElementsByTagName('head')[0];
    const link = document.createElement('link');
    link.type = 'text/css';
    link.rel = 'stylesheet';
    link.href = chrome.runtime.getURL('css/style.css');
    head.appendChild(link);
    const documentClone = document.cloneNode(true);
    article = new Readability(documentClone).parse();
}

function buildText() {
    // 右上角添加save按钮
    let html = `<div>
            <div id="article-popup-content" class="article-card">
                <div class="article-card-details">
                    <p class="article-btn">Save</p>
                    <p contenteditable class="article-text-title" id="article-title">${article.title}</p>
                    <input type="text" class="article-type-input" id="article-type-input" placeholder="Type article tag here, separated by Spaces">
                    <div contenteditable class="article-text-body" id="article-text">${article.textContent}</div>
                </div>
            </div>
        </div>`
    document.body.innerHTML += html;

    document.addEventListener('click', (event) => {
        if (!document.getElementById('article-popup-content').contains(event.target)
            // && event.target !== document.getElementById('article-popup-button')
        ) {
            article.title = document.getElementById('article-title').innerHTML;
            article.textContent = document.getElementById('article-text').innerHTML;
            article.type = document.getElementById('article-type-input').value;
            setDisplay(true);
        }
    });

    document.getElementById('article-popup-content').addEventListener('click', (event) => {
        request('POST', 'save',
            {
                title: article.title,
                content: article.textContent,
                type: article.type?.trim().split(/\s+/),
                url: window.location.href
            },
            function (res) {
                console.log(res);
            }
        );
    });
}
function setDisplay(show) {
    if (show) {
        document.getElementById('article-popup-content').style.display = 'none';
        isShow = false;
    } else {
        document.getElementById('article-popup-content').style.display = 'block';
        isShow = true;
    }
}

function request(method, url, data, callback) {
    const xhr = new XMLHttpRequest();
    xhr.open(method, remote_path + url);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");


    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            callback(xhr.responseText);
        }
    }
    xhr.send(JSON.stringify(data));
}

