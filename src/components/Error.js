import React, {Component} from 'react';
import injectSheet from 'react-jss';

const styles = {
    error: {
        background: '#e6736d',
        color: '#fff',
        padding: '20px',
    }
};


const Error = (props) =>
    <div className={props.classes.error}>
        {props.error}
    </div>;

export default injectSheet(styles)(Error);