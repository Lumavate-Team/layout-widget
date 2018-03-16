import '../../stencil.core';
import { EventEmitter } from '../../stencil.core';
export declare class LumavateToolbarButton {
    buttonClicked(event: UIEvent): void;
    navigate: EventEmitter;
    item: any;
    render(): JSX.Element;
}
