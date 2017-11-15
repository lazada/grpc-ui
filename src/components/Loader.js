import React, {Component} from 'react';
import injectSheet from 'react-jss';

const styles = {
    loader: {
        'height': '400px',
        'position': 'relative',
    },
    title: {
        'font-weight': 'bold',
        'font-size': '32px',
        'position': 'absolute',
        'top': '50%',
        'left': '50%',
        'transform': 'translate(-50%, -50%)'
    },
};


class Loader extends Component {
    constructor(props) {
        super(props);

        this.state = {
          points: 1,
          timer: null,
        };
    }
    componentDidMount() {
        this.setState({
            timer: setInterval(() => {
                this.setState({
                    points: ((this.state.points + 1) % 3),
                });
            }, 500),
        })
    }
    componentWillUnmount() {
        clearInterval(this.state.timer)
    }
    render() {
        return (
            <div className={this.props.classes.loader}>
                <h2 className={this.props.classes.title}>Loading{Array(this.state.points + 1).fill(0).map(_ => '.')}</h2>
            </div>
        );
    }
}

export default Loader = injectSheet(styles)(Loader);