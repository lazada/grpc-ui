import React, {Component} from 'react';

import Sidebar from './Sidebar';
import Package from './Package';

import './app.sass';


export default class App extends Component {
    constructor() {
        super();
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
                        <div className="app__left">
                            <Sidebar packages={this.state.packages}/>
                        </div>
                        <div className="app__right">
                            <div className="packages-list">
                                {packages}
                            </div>
                        </div>

                    </div>
                </div>

            </div>
        );
    }
}