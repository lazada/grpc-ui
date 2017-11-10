import React, {Component} from 'react';

import Sidebar from './Sidebar';
import Obj from './Object.js';

import './app.sass';

import Adapter from '../Adapter.js';


class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            reflection: null,
            selected: null,
        }
    }
    componentDidMount() {
        const adapter = new Adapter();
       
        adapter.fetchReflection()
            .then(reflection => {
                this.setState({ reflection });
            });
        // TODO handle errors
    }
    render() {
        return (
            <div>
                <div className="navbar">
                    <div className="navbar__container">
                        <a href="" className="logo"/>
                    </div>
                </div>
                <div className="app">
                    <div className="app__container">
                        <div className="app__left">
                            <Sidebar
                                reflection={this.state.reflection}
                                onSelect={obj => {
                                    this.setState({ selected: obj });
                                }}
                            />
                        </div>
                        <div className="app__right">
                            {this.state.selected ? <Obj obj={this.state.selected}/> : null}
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}


export default App;