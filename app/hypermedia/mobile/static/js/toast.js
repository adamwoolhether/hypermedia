// this should be placed in the expo /src/ folder.

import Toast from 'react-native-root-toast';

const NAMESPACE_URI = "https://hypermedia.systems/hyperview/toast";

export default {
    action: "show-toast",
    callback: (behaviorElement) => {
        const text = behaviorElement.getAttributeNS(NAMESPACE_URI, "text");
        if (text != null) {
            // Toast has many options for configurtin the toast. We're doing simple here.
            Toast.show(text, {position: Toast.positions.TOP, duration: 2000});
        }
    }
}