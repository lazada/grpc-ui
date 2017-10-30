import {REQUEST_PACKAGES_AND_TYPES, REQUEST_PACKAGES_AND_TYPES_SUCCESS, INVOKE_METHOD} from "./actions";


const initialState = {
    types: {},
    packages: [],
};

export default (state, action) =>{
    console.log(action);
    switch (action.type) {
        case REQUEST_PACKAGES_AND_TYPES:
            return Object.assign({}, state, {
                loading: true,
            });
        case REQUEST_PACKAGES_AND_TYPES_SUCCESS:
            return Object.assign({}, state, {
                loading: false,
                types: action.types,
                packages: action.packages,
            });
        case INVOKE_METHOD:
            return state;
        default:
            return initialState;
    }
}
