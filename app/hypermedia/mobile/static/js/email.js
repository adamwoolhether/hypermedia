// this should be placed in the expo /src/ folder.

import { email } from 'react-native-communications';

const namespace = "https://hypermedia.systems/hyperview/communications";

export default {
    action: "open-email",
    callback: (behaviorElement) => {
        const emailAddress = behaviorElement.getAttributeNS(namespace, "email-address");
        if (emailAddress != null) {
            email([emailAddress], null, null, null, null);
        }
    }
}