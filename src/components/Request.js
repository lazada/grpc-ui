import React, {Component} from 'react';
import {Message} from './fields';
import {getDefaultValue} from './types';

function fieldsToVal(fields, val, types) {
    return fields.map((f, i) => {
        let exportedVal = val[i];
        switch (f.type_id) {
            case 11:
                const type = types[f.type_name];
                exportedVal = fieldsToVal(type.fields, val[i], types);
        }
        return {
            number: f.number,
            val: exportedVal,
        }
    })
}

export default class Request extends Component {
    constructor(props) {
        super(props);

        this.state = {
            val: props.fields.map((f) => getDefaultValue(f.type_id, f.is_repeated, f.type_name, props.enums, props.types)),
        };
    }

    handleInvokeMethod(e) {
        e.preventDefault();
        this.props.onInvokeMethod(fieldsToVal(this.props.fields, this.state.val, this.props.types));
    }

    handleChange(val) {
        console.log(val);
        this.setState({
            val,
        });
    }

    render() {
        return (
            <Form
                fields={this.props.fields}
                val={this.state.val}
                types={this.props.types}
                enums={this.props.enums}
                onChange={this.handleChange.bind(this)}
                onInvoke={this.handleInvokeMethod.bind(this)}
            />
        );
    }
}

const Form = ({fields, val, onChange, onInvoke, types, enums}) =>
    <div className="form">
        <h4 className="form__title">Request</h4>
        <form onSubmit={onInvoke}>
            <Message {...{fields, val, onChange, types, enums}}/>
            <div className="form__controls">
                <button type="submit" className="button">Invoke</button>
            </div>
        </form>
    </div>;
