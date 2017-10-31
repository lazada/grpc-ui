import React, {Component} from 'react';
import {connect} from 'react-redux';

import './Method.sass';
import axios from 'axios/index';


const UnknownField = (props) => <div className="field">
    {props.read_only ?
        <div>
            <label className="field__label" htmlFor={props.name}>{props.name} = {props.number} (string)</label>
        </div>
        :
        <div>
            <label className="field__label" htmlFor={props.name}>{props.name} = {props.number} (string)</label>
            <input className="field__input" name={props.name} id={props.name} type="text" value={props.val} onChange={props.onChange}/>
        </div>
    }
</div>;




class Request extends Component {
    constructor(props) {
        super(props);
        this.state = {
            fields: props.fields.map(f => {
                return {val: ''}
            }),
        };
    }
    handleInvokeMethod(e) {
        e.preventDefault();
        const args = this.props.fields.map((f, i) => {
            return {
                number: f.number,
                val: this.state.fields[i].val,
            };
        });
        this.props.onInvokeMethod(args);
    }

    handleChange(i, val) {
        let fields = this.state.fields.slice();
        fields[i] = {val:val};

        this.setState({
            fields: fields,
        });
        console.log(this.state.fields);
    }
    render() {
        return (
            <div className="message message--in">
                <h4 className="message__title">Request</h4>

                <form onSubmit={this.handleInvokeMethod.bind(this)}>
                    {this.props.fields.map((f, i) =>
                        <UnknownField  {...f} val={this.state.fields[i].val}
                                       onChange={(e) => this.handleChange(i, e.target.value)}/>
                    )}

                    {this.props.in ? <div className="message__controls">
                        <button type="submit" className="button">Invoke</button>
                    </div> : null }
                </form>
            </div>
        );
    }
}

const Response = (props) => {
    return <div className="message">
        <h4 className="message__title">Response</h4>

        {props.fields.map((f) =>
            <UnknownField  {...f} read_only={true}/>
        )}
    </div>;
};


class Method extends Component {
    constructor(props) {
        super(props);
        this.state = {
            loading: false,
            expanded: false,
            result: '',
            error: '',
        };

    }
    onHeadingClick() {
        this.setState({
            expanded: !this.state.expanded,
        })
    }
    handleInvokeMethod(args) {
        this.setState({
            error: '',
            response: '',
            loading: true,
        });

        axios.post('/api/invoke', {
            package_name: this.props.package_name,
            service_name: this.props.service_name,
            method_name: this.props.name,
            grpc_args: args,
        })
            .then(({data}) => {
                this.setState({
                    result: data.data,
                    loading: false,
                });
            })
            .catch((error) => {
                this.setState({
                    error: error.response.data.error,
                    loading: false,
                });
            })
    }
    render() {
        return <div className={`method ${this.state.loading ? 'method--loading' : ''}`}>
            <div className="method__heading" onClick={this.onHeadingClick.bind(this)}>
                <h4 className="method__name"> {this.props.name} <i className={this.state.expanded ? '' : 'fa fa-angle-down'}/></h4>
            </div>

            <div className="method__body" style={{display: this.state.expanded ? 'block' : 'none'}}>
                <Request {...this.props.types[this.props.in]} types={this.props.types} in={true} onInvokeMethod={this.handleInvokeMethod.bind(this)}/>

                {this.state.result ? <pre>{JSON.stringify(this.state.result)}</pre> : null}
                {this.state.error ? <pre>{JSON.stringify(this.state.error)}</pre> : null}
                <Response {...this.props.types[this.props.out]} types={this.props.types}/>
            </div>
        </div>
    }
}

export default Method;