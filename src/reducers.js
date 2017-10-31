import {REQUEST_PACKAGES_AND_TYPES, REQUEST_PACKAGES_AND_TYPES_SUCCESS, INVOKE_METHOD,
    INVOKE_METHOD_SUCCESS, INVOKE_METHOD_ERROR} from "./actions";


const initialState = {
    types: {},
    packages: [],
    methods: [],
};

export default (state, action) =>{
    console.log(action);
    switch (action.type) {
        case REQUEST_PACKAGES_AND_TYPES:
            return {...state, ...{
                loading: true,
            }};
        case REQUEST_PACKAGES_AND_TYPES_SUCCESS:
            return {...state, ...{
                loading: false,
                types: action.types,
                packages: action.packages
            }};
        case INVOKE_METHOD:
            let new_packages = {...state.packages}
            let services = state.packages[action.package_name];
            for (let j = 0; j < services.length; j++) {
                if (services[j].name === action.service_name) {
                    for(let k = 0; k < services[j].methods.length; k++) {
                        if (services[j].methods[k].name === action.method_name) {
                            services[j].methods[k].loading = true;
                        }
                    }
                }
            }
            return {...state, ...{
                packages: new_packages,
            }};
        case INVOKE_METHOD_SUCCESS:
            new_packages = {...state.packages};
            services = state.packages[action.package_name];
            for (let j = 0; j < services.length; j++) {
                if (services[j].name === action.service_name) {
                    for(let k = 0; k < services[j].methods.length; k++) {
                        if (services[j].methods[k].name === action.method_name) {
                            services[j].methods[k].loading = false;
                            services[j].methods[k].result = action.result;
                        }
                    }
                }
            }
            return {...state, ...{
                packages: new_packages,
            }};
        case INVOKE_METHOD_ERROR:
            new_packages = {...state.packages}
            services = state.packages[action.package_name];
            for (let j = 0; j < services.length; j++) {
                if (services[j].name === action.service_name) {
                    for(let k = 0; k < services[j].methods.length; k++) {
                        if (services[j].methods[k].name === action.method_name) {
                            services[j].methods[k].loading = false;
                            services[j].methods[k].error = action.error;
                        }
                    }
                }
            }
            return {...state, ...{
                packages: new_packages,
            }};
        default:
            return initialState;
    }
}
