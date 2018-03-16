export class LumavateToolbarButton {
    buttonClicked(event) {
        if (this.item.url) {
            this.navigate.emit(this.item);
        }
        console.log(event);
    }
    render() {
        return (h("div", null,
            h("div", { onClick: (event) => this.buttonClicked(event) },
                h("i", { class: "material-icons", style: { color: this.item.color ? this.item.color : "#000" } }, this.item.icon),
                h("div", { style: { color: this.item.color ? this.item.color : "#000" }, class: "button-text" }, this.item.title))));
    }
    static get is() { return "lumavate-toolbar-button"; }
    static get properties() { return { "item": { "type": "Any", "attr": "item" } }; }
    static get events() { return [{ "name": "navigate", "method": "navigate", "bubbles": true, "cancelable": true, "composed": true }]; }
    static get style() { return "/**style-placeholder:lumavate-toolbar-button:**/"; }
}
