import axios from 'axios';

export const REQUEST_PACKAGES_AND_TYPES = 'REQUEST_PACKAGES_AND_TYPES';
export const REQUEST_PACKAGES_AND_TYPES_SUCCESS = 'REQUEST_PACKAGES_AND_TYPES_SUCCESS';
export const INVOKE_METHOD = 'INVOKE_METHOD';
export const INVOKE_METHOD_SUCCESS = 'INVOKE_METHOD_SUCCESS';
export const INVOKE_METHOD_ERROR = 'INVOKE_METHOD_ERROR';

export const loadPackages = () => {
    return dispatch => {
        const addr = '127.0.0.1:3001';
        dispatch({
            type: REQUEST_PACKAGES_AND_TYPES,
        });

        axios.get('/api/info?addr=' + addr)
            .then(({data: {packages, types}}) => {
                dispatch({
                    type: REQUEST_PACKAGES_AND_TYPES_SUCCESS,
                    packages,
                    types
                })
            });
    };

};

export const invokeMethod = (package_name, service_name, method_name, args) => {
    return dispatch => {
        dispatch({ type: INVOKE_METHOD, package_name, service_name, method_name, args});

        const addr = '127.0.0.1:3001';

        axios.post('/api/invoke', {
            addr,
            package_name,
            service_name,
            method_name,
            grpc_args: args,
        })
            .then(({data}) => {
                if (data.status === "ok") {
                    dispatch({
                        type: INVOKE_METHOD_SUCCESS,
                        service_name,
                        package_name,
                        method_name,
                        result: data.data,
                    })
                } else {
                    dispatch({
                        type: INVOKE_METHOD_ERROR,
                        service_name,
                        package_name,
                        method_name,
                        error: data.error,
                    })
                };
            })

    };
};