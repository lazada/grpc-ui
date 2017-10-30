import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import './app.sass';


const Int32Field = (props) => <div className="form-group">
    <label htmlFor={props.name}>{props.name}</label>
    <input className="form-control" name={props.name} id={props.name} type="number"/>
</div>;

const UnknownField = (props) => <div className="form-group">
    <label htmlFor={props.name}>{props.name}</label>
    <input className="form-control" name={props.name} id={props.name} placeholder={'type id: ' + props.type_id} type="text"/>
</div>;

const MessageField = (props) => {
    const type = props.types[props.type_name];
    return <Message {...type} types={props.types}/>
};

const Field = (props) => {
    switch (props.type_id) {
        case 5: // int32:
            return <Int32Field {...props}/>;
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
            <form onSubmit={props.onCall}>
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
    onMethodCall(e) {
        e.preventDefault();
        console.log(arguments);
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
                <div className="packages-list">
                    <div className="packages-list__container">
                        {packages}
                    </div>
                </div>
            </div>
        );
    }
}

ReactDOM.render(<App />, document.getElementById('root'));