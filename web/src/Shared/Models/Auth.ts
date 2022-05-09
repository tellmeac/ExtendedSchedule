import {GoogleLoginResponse} from "react-google-login";

/**
 * Basic user's data
 */
export interface UserData {
    tokenId: string
    avatar: string
    name: string
}

/**
 * Returns user data from google login response
 * @param response
 */
export function getUserAuthContentFromResponse(response: GoogleLoginResponse): UserData {
    const basic = response.getBasicProfile()
    return {
        tokenId: response.tokenId,
        avatar: basic.getImageUrl(),
        name: basic.getName()
    }
}