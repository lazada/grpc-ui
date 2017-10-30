import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import './app.sass';


const UnknownField = (props) => <div className="field">
    <label className="field__label" htmlFor={props.name}>{props.name}</label>
    <input className="field__input" name={props.name} id={props.name} type="text"/>
</div>;

const MessageField = (props) => {
    const type = props.types[props.type_name];
    return <Message {...type} types={props.types}/>
};

const Field = (props) => {
    switch (props.type_id) {
        case 5: // int32:
            return <UnknownField {...props}/>;
        case 8: // boolean
            return <UnknownField {...props}/>;
        case 9: // int32:
            return <UnknownField {...props}/>;
        case 11: // message
            return <MessageField {...props}/>;
        default:
            return <UnknownField  {...props}/>;
    }
};

const Message = (props) =>
    <div>
        {props.fields.map((f) => <Field key={f.name} {...f} types={props.types} /> )}
    </div>;

const Method = (props) =>
    <div className="method">
        <div className="method__heading">
            <h4 className="method__name">{props.name}</h4>
        </div>
        <div className="method__body">
            <form>
                <Message {...props.types[props.in]} types={props.types}/>
                <button type="submit" className="btn btn-primary pull-right">Invoke</button>

                <Message {...props.types[props.out]} types={props.types}/>
            </form>
        </div>
    </div>;

const Package = (props) => (
    <div className="package">
        <h3 className="package__title">{props.name + ' / ' + props.service.name}</h3>
        {props.service.methods.map((method) => <Method key={method.name} {...method} types={props.types}/>)}
    </div>
);

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            packages: {},
            types: {},
        };
    }
    componentDidMount() {
        const addr = '127.0.0.1:3001';
        fetch('/api/info?addr=' + addr)
            .then(r => r.json())
            .then(({packages, types}) => {
                this.setState({
                    packages,
                    types,
                })
            });
    }
    render() {
        const packages = Object.keys(this.state.packages).map(package_name => {
            return this.state.packages[package_name].map((service) => {
                return <Package key={package_name} name={package_name} service={service} types={this.state.types}/>
            });
        });

        return (
            <div>
                <div className="navbar">
                    <div className="navbar__container">
                        <a href="" className="logo"/>
                    </div>
                </div>
                <div className="app">
                    <div className="app__container">
                        <div className="sidebar">
                            <ul class="sidebar__list">
                                {Object.keys(this.state.packages).map(p =>
                                    <li>
                                        {p}
                                        <ul>
                                            {this.state.packages[p].map((s) => (
                                                <li>
                                                    {s.name}
                                                    <ul>
                                                        {s.methods.map((m) =>
                                                            <li>{m.name}</li>
                                                        )}
                                                    </ul>
                                                </li>
                                            ))}
                                        </ul>
                                    </li>
                                )}
                            </ul>
                        </div>
                        <div className="packages-list">
                            {packages}
                        </div>
                    </div>
                </div>

            </div>
        );
    }
}

ReactDOM.render(<App />, document.getElementById('root'));