import React, { PureComponent } from 'react';
import Hyperview from 'hyperview';
import Swipeable from 'react-native-swipeable';

const NAMESPACE_URI = "https://hypermedia.systems/hyperview/swipeable";

export default class SwipeableRow extends PureComponent {
    static namespaceURI = NAMESPACE_URI;
    static localName = "row";

    getElements = (tagName) => {
        return Array.from(this.props.element.getElementsByTagNameNS(NAMESPACE_URI, tagName));
    };

    getButtons = () => {
        return this.getElements("button").map((buttonElement) => {
            return Hyperview.renderChildren(buttonElement, this.props.stylesheets, this.props.onUpdate, this.props.options);
        });
    };

    render() {
        const [main] = this.getElements("main");
        if (!main) {
            return null;
        }

        return (
            <Swipeable rightButtons={this.getButtons()}>
                {Hyperview.renderChildren(main, this.props.stylesheets, this.props.onUpdate, this.props.options)}
            </Swipeable>
        );
    }
}