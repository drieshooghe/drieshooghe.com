window.onload = function () {

    function sleep(ms: Number): Promise<void> {
        return new Promise(resolve => setTimeout(resolve, ms));
    }
    class Text {
        protected element: Element;

        constructor(element: Element) {
            this.element = element;
        }

        async type(): Promise<void> {
            if (!this.element) return;
            let sValue: string = this.element.getAttribute('data-content');

            for (let i = 0; i < sValue.length; ++i) {
                await sleep(50);
                this.element.innerHTML += sValue[i];
            }
        }

    }

    class Line {
        protected icon: Text;
        protected pfx: Text;
        protected message: Text;

        constructor(eLine: Element) {
            this.icon = new Text(eLine.getElementsByClassName('js-prompt-icon')[0]);
            this.pfx = new Text(eLine.getElementsByClassName('js-prompt-prefix')[0]);
            this.message = new Text(eLine.getElementsByClassName('js-prompt-text')[0]);
        }

        async type(): Promise<void> {
            await this.icon.type();
            await sleep(400);
            await this.pfx.type();
            await this.message.type();
        }
    }

    class Lines {
        protected lines: Line[];

        constructor (lines: HTMLCollection) {
            this.lines = [];
            for (let i = 0; i < lines.length; i++) {     
                this.lines.push(new Line(lines[i]));
            }
        }

        async type() {
            for (const line of this.lines) {
                await line.type();
            }
        }
    }

    const lines = new Lines(document.getElementsByClassName('js-typewriting'));
    lines.type().catch(console.error);
}