// this should be placed in the expo /src/ folder.

import { email } from 'react-native-communications';

const NAMESPACE_URI = "https://hypermedia.systems/hyperview/communications";

export default {
    action: "open-email",
    callback: (behaviorElement) => {
        const emailAddress = behaviorElement.getAttributeNS(NAMESPACE_URI, "email-address");
        if (emailAddress != null) {
            email([emailAddress], null, null, null, null);
        }
    }
}