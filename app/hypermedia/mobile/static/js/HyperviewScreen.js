// this should be places in the expo /src/ folder.

import React, { PureComponent } from 'react';
import Hyperview from 'hyperview';
import OpenPhone from './phone';
import OpenEmail from './email';

export default class HyperviewScreen extends PureComponent {
    // ... omitted for brevity

    behaviors = [OpenPhone, OpenEmail];

    render() {
        return (
            <Hyperview
                behaviors={this.behaviors}
                entrypointUrl={this.entrypointUrl}
            // more props...
            />
        );
    }
}