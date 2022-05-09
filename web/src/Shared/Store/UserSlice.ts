import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {RootState} from "./Root";
import {UserAuthContent} from "../Models/Auth";

export interface UserState {
    data?: UserAuthContent
}

const initialUserState: UserState = {
    data: undefined,
}

export const userSlice = createSlice({
    name: "user",
    initialState: initialUserState,
    reducers: {
        updateUserData: (state, action: PayloadAction<UserAuthContent>) => {
            state.data = action.payload
        },
        resetUserData: (state) => {
            state.data = undefined
        },
    }
})

export const selectUserData = (state: RootState) => state.user.data

export const { updateUserData, resetUserData } = userSlice.actions