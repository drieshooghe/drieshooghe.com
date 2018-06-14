window.onload = function () {

    class Line {
        icon: Element;
        pfx: Element;
        message: Element;

        constructor(icon: Element, pfx: Element, message: Element) {
            this.icon = icon;
            this.pfx = pfx;
            this.message = message;
        }

        type(): void {
            function restore(e: Element, i: number) {
                if (e) {
                    let sValue: string = e.getAttribute('data-content');
                    if (i < sValue.length) {
                        e.innerHTML += sValue[i];
                        setTimeout(function () { restore(e, i); }, 100);
                        i++;
                    }
                }
            }

            restore(this.icon, 0);
            restore(this.pfx, 0);
            restore(this.message, 0);
        }
    }

    function sleep(ms: Number) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

    function dispChars(eLine: Element) {
        
        // Get HTML items
        let promptLine = new Line(
            eLine.getElementsByClassName('js-prompt-icon')[0],
            eLine.getElementsByClassName('js-prompt-prefix')[0],
            eLine.getElementsByClassName('js-prompt-text')[0]
        );

        eLine.classList.remove('hidden');
        promptLine.type();
    }

    function displayLines(eLines: HTMLCollection) {
        for (let i = 0; i < eLines.length; i++) {     
            dispChars(eLines[i]);
        }
    }

    let eLines = document.getElementsByClassName('js-typewriting');
    displayLines(eLines);
}

