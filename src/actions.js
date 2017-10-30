export const REQUEST_PACKAGES_AND_TYPES = 'REQUEST_PACKAGES_AND_TYPES';
export const REQUEST_PACKAGES_AND_TYPES_SUCCESS = 'REQUEST_PACKAGES_AND_TYPES_SUCCESS';
export const INVOKE_METHOD = 'INVOKE_METHOD';
export const INVOKE_METHOD_SUCCESS = 'INVOKE_METHOD';

export const loadPackages = () => {
    return dispatch => {
        const addr = '127.0.0.1:3001';
        dispatch({
            type: REQUEST_PACKAGES_AND_TYPES,
        });

        fetch('/api/info?addr=' + addr)
            .then(r => r.json())
            .then(({packages, types}) => {
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

        fetch('/api/invoke', {
            method: 'POST',
            body: JSON.stringify({
                addr,
                package_name,
                service_name,
                method_name,
                grpc_args: args,
            })
        })
            .then(res => res.json)
            .then(data => {
                console.log(data);
            })

    };
};