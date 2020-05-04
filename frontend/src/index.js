import React from 'react'
import ReactDOM from 'react-dom'
import Main from './components/Main'
import Project from './components/Project'
import Create from './components/Create'
import * as serviceWorker from './serviceWorker'
import {BrowserRouter, Route, Switch} from 'react-router-dom'

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
     <Switch>
        <Route exact path='/' component={Main} />
        <Route path='/project/:id' component={Project} />
        <Route path='/create' component={Create} />
     </Switch>
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
