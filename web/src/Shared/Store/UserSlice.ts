import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {RootState} from "./Root";
import {UserData} from "../Models/Auth";

export interface UserState {
    data?: UserData
}

const initialUserState: UserState = {
    data: undefined,
}

export const userSlice = createSlice({
    name: "user",
    initialState: initialUserState,
    reducers: {
        updateUserData: (state, action: PayloadAction<UserData>) => {
            state.data = action.payload
        },
    }
})

export const selectUserData = (state: RootState) => state.user.data

export const { updateUserData } = userSlice.actions