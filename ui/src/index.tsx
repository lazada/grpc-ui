import * as React from 'react';
import * as ReactDOM from 'react-dom';
import App from './App';
import Dev from './Dev';

import '@blueprintjs/core/dist/blueprint.css';
import 'normalize.css/normalize.css';
import './index.css';

const Comp = process.env.NODE_ENV === 'development' ? Dev : App;

ReactDOM.render(
  <Comp />,
  document.getElementById('root') as HTMLElement
);
