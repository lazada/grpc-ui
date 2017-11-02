import React, {Component} from 'react';
import {Message} from './fields';

export default class Request extends Component {
    constructor(props) {
        super(props);
        this.state = {
            val: props.fields.map((f) => {
                return getDefaultValue(f.type_id, f.is_repeated);
            })
        };
    }

    handleInvokeMethod(e) {
        e.preventDefault();
        const args = this.props.fields.map((f, i) => {
            return {
                number: f.number,
                val: this.state.val[i],
            };
        });
        this.props.onInvokeMethod(args);
    }

    handleChange(val) {
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
                onChange={this.handleChange.bind(this)}
                onInvoke={this.handleInvokeMethod.bind(this)}
            />
        );
    }
}

const Form = ({fields, val, onChange, onInvoke, types}) =>
    <div className="form">
        <h4 className="form__title">Request</h4>
        <form onSubmit={onInvoke}>
            <Message {...{fields, val, onChange, types}}/>
            <div className="form__controls">
                <button type="submit" className="button">Invoke</button>
            </div>
        </form>
    </div>;

export const getDefaultValue = (type_id, repeated) => {
    if (repeated) {
        return [];
    }
    switch (type_id) {
        case 8: //bool
            return 'false';
        case 11: //msg
            return [];
        default:
            return '';
    }
};