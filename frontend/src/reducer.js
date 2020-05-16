export default function reducer(state, {action, payload}) {
    console.log(action, payload)
    switch(action) {
        case 'GET_ME':
            console.log("dispatched getMe", payload)
            return {
                ...state,
                currentUser: payload,
            }
        default:
            return state
    }
}