export class LumavateQuote {
    constructor() {
        this.QuoteText = '';
        this.FontSize = '12pt';
        this.QuotationMarks = true;
        this.Color = '#000';
        this.ShowCard = true;
        this.CardColor = '#FFF';
    }
    render() {
        return (h("div", { class: this.ShowCard ? "container-shadow" : "container", style: { backgroundColor: this.CardColor } },
            h("div", { class: "quote-container start" },
                h("div", { class: "spacer" }, this.QuotationMarks
                    ? h("svg", { xmlns: "http://www.w3.org/2000/svg", width: "12", height: "8", viewBox: "0 0 12 8" },
                        h("g", { fill: "none", "fill-rule": "evenodd", transform: "matrix(-1 0 0 1 15 -5)" },
                            h("path", { fill: this.Color, "fill-rule": "nonzero", d: "M10,5 L10,10 L12.75,10 L11,13 L13.25,13 L15,10 L15,5 L10,5 Z M3,10 L5.75,10 L4,13 L6.25,13 L8,10 L8,5 L3,5 L3,10 Z" }),
                            h("polygon", { points: "0 0 18 0 18 18 0 18" })))
                    : '')),
            h("div", { class: "quote-text", style: { color: this.Color, fontSize: this.FontSize } }, this.QuoteText),
            h("div", { class: "quote-container end" },
                h("div", { class: "spacer" }, this.QuotationMarks
                    ? h("svg", { xmlns: "http://www.w3.org/2000/svg", width: "12", height: "8", viewBox: "0 0 12 8" },
                        h("g", { fill: "none", "fill-rule": "evenodd", transform: "translate(-3 -5)" },
                            h("path", { fill: this.Color, "fill-rule": "nonzero", d: "M10,5 L10,10 L12.75,10 L11,13 L13.25,13 L15,10 L15,5 L10,5 Z M3,10 L5.75,10 L4,13 L6.25,13 L8,10 L8,5 L3,5 L3,10 Z" }),
                            h("polygon", { points: "0 0 18 0 18 18 0 18" })))
                    : ''))));
    }
    static get is() { return "lumavate-quote"; }
    static get properties() { return { "CardColor": { "type": String, "attr": "card-color" }, "Color": { "type": String, "attr": "color" }, "FontSize": { "type": String, "attr": "font-size" }, "QuotationMarks": { "type": Boolean, "attr": "quotation-marks" }, "QuoteText": { "type": String, "attr": "quote-text" }, "ShowCard": { "type": Boolean, "attr": "show-card" } }; }
    static get style() { return "/**style-placeholder:lumavate-quote:**/"; }
}
