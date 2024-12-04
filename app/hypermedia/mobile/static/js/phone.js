// this should be placed in the expo /src/ folder.

import { phonecall } from 'react-native-communications';

const NAMESPACE_URI = "https://hypermedia.systems/hyperview/communications";

export default {
    action: "open-phone",
    callback: (behaviorElement) => {
        const number = behaviorElement.getAttributeNS(NAMESPACE_URI, "phone-number");
        if (number != null) {
            phonecall(number, false);
        }
    }
}