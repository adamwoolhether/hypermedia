// this should be placed in the expo /src/ folder.
// better to just add this to the behaviors array in the HyperviewScreen.tsx file.

import React, { PureComponent } from 'react';
import Hyperview from 'hyperview';
import OpenPhone from './phone';
import OpenEmail from './email';
import ShowToast from './toast';
import SwipeableRow from "./swipeable";

export default class HyperviewScreen extends PureComponent {
    // ... omitted for brevity

    behaviors = [OpenPhone, OpenEmail, ShowToast];
    components = [SwipeableRow];

    render() {
        return (
            <Hyperview
                behaviors={this.behaviors}
                components={components}
                entrypointUrl={this.entrypointUrl}
            // more props...
            />
        );
    }
}