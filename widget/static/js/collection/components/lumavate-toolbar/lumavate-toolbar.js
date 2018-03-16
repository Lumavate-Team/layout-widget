export class LumavateToolbar {
    navigateHandler(event) {
        window.location.href = event.detail.url;
    }
    componentWillLoad() {
        this.innerItems = JSON.parse(this.items);
    }
    render() {
        return (h("div", { style: { backgroundColor: this.backgroundcolor ? this.backgroundcolor : "#fff" }, class: "container" }, this.innerItems.map((item) => h("lumavate-toolbar-button", { item: item }))));
    }
    static get is() { return "lumavate-toolbar"; }
    static get properties() { return { "backgroundcolor": { "type": String, "attr": "backgroundcolor" }, "items": { "type": String, "attr": "items" } }; }
    static get style() { return "/**style-placeholder:lumavate-toolbar:**/"; }
}
