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

        const type = props.types[props.type_name];

        this.state = !type ? {} : {
            val: type.fields.map((f) => getDefaultValue(f.type_id, f.is_repeated, f.type_name, props.enums, props.types)),
        };
    }


    handleInvokeMethod(e) {
        e.preventDefault();
        const type = this.props.types[this.props.type_name];

        this.props.onInvokeMethod(fieldsToVal(type.fields, this.state.val, this.props.types));
    }

    handleChange(val) {
        console.log(val);
        this.setState({
            val,
        });
    }

    render() {
        const type = this.props.types[this.props.type_name];
        return type ?
            <div className="form">
                <h4 className="form__title">{this.props.type_name}</h4>
                <form onSubmit={this.handleInvokeMethod.bind(this)}>
                    <Message type={type} val={this.state.val} onChange={this.handleChange.bind(this)}
                             types={this.props.types} enums={this.props.enums}
                    />
                    <div className="form__controls">
                        <button type="submit" className="button">Invoke</button>
                    </div>
                </form>
            </div>
            : <div>Unknown type: {this.props.type_name}</div>;
    }
}
