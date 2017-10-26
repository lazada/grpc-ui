angular.module('grpc-ui', [])
    .controller('ServiceGrpcCtrl', ServiceGrpcCtrl);

function ServiceGrpcCtrl($scope, $http) {

    $scope.depth = 0;
    $scope.viewMode = 'recent';
    $scope.limit = 1;
    $scope.status = 'connecting';

    var ws = new WebSocket('ws://' + location.host + '/ws');

    ws.addEventListener('open', function () {
        $scope.status = 'loaded';
        $scope.$applyAsync();
    });

    ws.addEventListener('message', function(event) {
        var data = JSON.parse(event.data);
        var method = null;
        switch (data.type) {
            case 'invoke_resp':
                method = $scope.grpc_descr[data.package_name][data.service_name][data.method_name];
                if (data.error) {
                    method.output_errors.push(data.error);
                    method.output_log = [];
                } else {
                    method.output_result_formatted = JSON.stringify(JSON.parse(data.data), null, 4);
                }
                method.invoking = false;
                break;
            case 'task_created':
                method = $scope.grpc_descr[data.package_name][data.service_name][data.method_name];
                method.task_id = data.task_id;
                break;
            case 'stream_message':
                method = $scope.grpc_descr[data.package_name][data.service_name][data.method_name];
                method.output_result_formatted = '';
                method.output_log = method.output_log || [];
                method.output_log.unshift(JSON.stringify(JSON.parse(data.data), null, 4));
                if (method.output_log.length > 50) {
                    method.output_log.pop();
                }
                break;
            case 'stream_error':
                method = $scope.grpc_descr[data.package_name][data.service_name][data.method_name];

                break;
        }

        $scope.$applyAsync();
    });

    // max embedding depth (to avoid too long or even infinite loops)
    const maxDepth = 2;

    $scope.need2expand = function (currDepth) {
        return currDepth < maxDepth;
    };

    $scope.toggleMethod = function (method) {
        method.show = !method.show;
    };

    $scope.grpcInfoSubmit = function () {
        $scope.status = 'loading';
        $scope.error = null;

        $http.get('/api/info', {
            params: {
                addr: $scope.addr,
            }
        }).then(function (response) {
            $scope.status = 'loaded';

            var services = response.data.services;
            var grpc_descr = {};

            function parseInt_(method, value) {
                try {
                    return parseInt(value)
                } catch (ex) {
                    method.output_errors.push("parseInt( " + value + " ): " + String(ex))
                    return null
                }
            }

            function parseFloat_(method, value) {
                try {
                    return parseFloat(value)
                } catch (ex) {
                    method.output_errors.push("parseFloat( " + value + " ): " + String(ex))
                    return null
                }
            }

            function parseEnum(method, value) {
                try {
                    return parseInt(value.number)
                } catch (ex) {
                    method.output_errors.push("parseEnum( " + value + " ): " + String(ex))
                    return null
                }
            }

            function parseBool(method, value) {
                return value
            }

            function parseString(method, value) {
                return value
            }

            function parseJson(method, value) {
                try {
                    if (value.length) {
                        return JSON.parse(value)
                    } else {
                        return null
                    }
                } catch (ex) {
                    method.output_errors.push("JSON.parse( " + value + " ): " + String(ex))
                    return null
                }
            }

            function prefill_message_json(fields) {
                var result = {}
                for (var f in fields) {
                    var field = fields[f]
                    var data = get_input_data(field.type.id, field.type.fields)
                    result[field.name] = data.default
                }
                return JSON.stringify(result, null, 4)
            }

            function get_input_data(type_id, type_fields) {
                switch (type_id) {
                    case 1: // TYPE_DOUBLE
                    case 2: // TYPE_FLOAT
                        return {
                            control: "input_float",
                            parser: parseFloat_,
                            default: 0.0,
                        }

                    case 5: // TYPE_INT32
                    case 3: // TYPE_INT64
                    case 13: // TYPE_UINT32
                    case 4: // TYPE_UINT64
                    case 17: // TYPE_SINT32
                    case 18: // TYPE_SINT64
                    case 7: // TYPE_FIXED32
                    case 6: // TYPE_FIXED64
                    case 15: // TYPE_SFIXED32
                    case 16: // TYPE_SFIXED64
                        return {
                            control: "input_int",
                            parser: parseInt_,
                            default: 0,
                        }

                    case 14: // TYPE_ENUM
                        return {
                            control: "input_enum",
                            parser: parseEnum,
                            default: {"number": 0},
                        }

                    case 8: // TYPE_BOOL
                        return {
                            control: "input_checkbox",
                            parser: parseBool,
                            default: false,
                        }

                    case 9: // TYPE_STRING
                    case 12: // TYPE_BYTES
                        return {
                            control: "input_text",
                            parser: parseString,
                            default: "",
                        }

                    case 11: // TYPE_MESSAGE
                        return {
                            control: "textarea",
                            parser: parseJson,
                            default: prefill_message_json(type_fields),
                        }

                }
                return {}
            }

            function process_fields(fields) {
                var result = {}

                for (var f in fields) {
                    var field = fields[f]

                    var labels = []
                    var repeated = false
                    var map_descr = null
                    var message_descr = null
                    var enum_descr = null

                    if (field.type.options !== undefined && field.type.options.map_entry) {
                        labels.push("map");
                        repeated = true
                        map_descr = {}
                        for (var ff in field.type.fields) {
                            var mfield = field.type.fields[ff]
                            map_descr[mfield.name + '_id'] = mfield.type.id
                            map_descr[mfield.name + '_type'] = mfield.type.name
                        }
                    } else {
                        if (field.label === 3) {
                            labels.push("repeated")
                            repeated = true
                        }
                        if (field.type.fields) {
                            labels.push("message")
                            message_descr = process_fields(field.type.fields)
                        }
                        if (field.type.id == 14) { // TYPE_ENUM
                            enum_descr = field.enum.values
                        }
                    }

                    var input_descr = get_input_data(field.type.id, field.type.fields)
                    var input = {
                        controls: {},
                        values: [],
                    }
                    if (map_descr) {
                        var value = {}
                        for (var fi in field.type.fields) {
                            var ff = field.type.fields[fi]
                            var ii = get_input_data(ff.type.id, ff.type.fields)
                            input["controls"][ff.name] = {
                                control: ii.control,
                                parser: ii.parser,
                                default: ii.default,
                            }
                            value[ff.name] = ii.default
                        }
                        // input["values"].push(value)
                    } else {
                        input["controls"]["value"] = {
                            control: input_descr.control,
                            parser: input_descr.parser,
                            default: input_descr.default,
                        }
                        if (!repeated) {
                            input["values"].push({
                                value: input_descr.default,
                            })
                        }
                    }

                    result[field.name] = {
                        type: {
                            id: field.type.id,
                            name: field.type.name,
                        },
                        number: field.number,
                        input: input,
                    }

                    if (labels.length) {
                        result[field.name]["labels"] = labels
                    }
                    if (repeated) {
                        result[field.name]["repeated"] = true
                    }
                    if (map_descr) {
                        result[field.name]["map"] = map_descr
                    }
                    if (message_descr) {
                        result[field.name]["message"] = message_descr
                    }
                    if (enum_descr) {
                        result[field.name]["enum"] = enum_descr
                    }
                }

                return result
            }

            for (var s in services) {
                var service = services[s]

                if (!(service.package_name in grpc_descr)) {
                    grpc_descr[service.package_name] = {}
                }
                grpc_descr[service.package_name][service.name] = {}

                for (var m in service.methods) {
                    var method = service.methods[m]

                    grpc_descr[service.package_name][service.name][method.name] = {
                        can_invoke: true,
                        is_stream: false,
                    }

                    var message_map = {
                        "request": method.in,
                        "response": method.out,
                    }
                    for (var message_type in message_map) {
                        var message = message_map[message_type]
                        var fields = process_fields(message.fields)
                        var descr = {
                            name: message.name,
                        }
                        if (fields) {
                            descr["fields"] = fields
                        }
                        if (message_type == "response" && method.in_stream) {
                            descr["stream"] = true
                            grpc_descr[service.package_name][service.name][method.name].can_invoke = false
                        }

                        if (message_type == "response" && method.out_stream) {
                            grpc_descr[service.package_name][service.name][method.name].is_stream = true
                        }

                        grpc_descr[service.package_name][service.name][method.name][message_type] = descr
                    }
                }
            }

            $scope.grpcServices = services;
            $scope.grpc_descr = grpc_descr;

        }, function (response) {
            $scope.status = 'error';
            $scope.error = response.data;
        });
    };

    $scope.add_repeated_field = function (package_name, service_name, method_name, args_type, field_name) {
        var field = $scope.grpc_descr[package_name][service_name][method_name][args_type].fields[field_name]
        var value = {}
        for (var c in field.input.controls) {
            var control = field.input.controls[c]
            value[c] = control.default
        }
        field.input.values.push(value)
    };

    $scope.remove_repeated_field = function (package_name, service_name, method_name, args_type, field_name, index) {
        $scope.grpc_descr[package_name][service_name][method_name][args_type].fields[field_name].input.values.splice(index, 1)
    };

    $scope.invoke_method = function (package_name, service_name, method_name) {
        var method = $scope.grpc_descr[package_name][service_name][method_name];

        if (method.is_stream && method.invoking) {
            ws.send(JSON.stringify({
                type: "stop",
                task_id: method.task_id,
            }));

            method.invoking = false;
            method.output_errors = [];
            return
        }

        method.output_log = [];
        method.output_result_formatted = '';
        method.output_errors = [];
        method.invoking = true;

        ws.send(JSON.stringify({
            type: "invoke",
            data: {
                is_stream: method.is_stream,
                addr: $scope.addr,
                package_name: package_name,
                service_name: service_name,
                method_name: method_name,
                grpc_method: package_name + '.' + service_name + '.' + method_name,
                grpc_args: JSON.stringify(makeArgs(method)),
            },
        }));
    };

    function makeArgs(method) {
        var args = {};

        for (var field_name in method.request.fields) {
            var field = method.request.fields[field_name];

            var value = null
            if (field.map) {
                if (field.input.values.length) {
                    var k_parser = field.input.controls.key.parser
                    var v_parser = field.input.controls.value.parser
                    value = {}
                    for (var i in field.input.values) {
                        value[k_parser(method, field.input.values[i].key)] = v_parser(method, field.input.values[i].value)
                    }
                }
            } else {
                if (field.input.values.length) {
                    var parser = field.input.controls.value.parser
                    if (field.repeated) {
                        value = []
                        for (var i in field.input.values) {
                            var v = parser(method, field.input.values[i].value)
                            if (v) {
                                value.push(v)
                            }
                        }
                    } else {
                        value = parser(method, field.input.values[0].value)
                    }
                }
            }

            if (value) {
                args[field_name] = value
            }
        }

        return args
    }

}
