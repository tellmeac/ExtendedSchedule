import {GoogleLoginResponse} from "react-google-login";

export interface UserAuthContent {
    tokenId: string
    avatar: string
    name: string
}

export function getUserAuthContentFromResponse(r: GoogleLoginResponse): UserAuthContent {
    const basic = r.getBasicProfile()
    return {
        tokenId: r.tokenId,
        avatar: basic.getImageUrl(),
        name: basic.getName()
    }
}