import React, {useContext, useEffect, useReducer} from 'react'
import ReactDOM from 'react-dom'
import { ApolloProvider } from 'react-apollo'
import Main from './components/Main'
import Project from './components/Project'
import Create from './components/Create'
import Login from './components/Login'
import Register from './components/Register'
import reducer from './reducer'
import * as serviceWorker from './serviceWorker'
import {BrowserRouter, Route, Switch} from 'react-router-dom'
import Context from './context'
import {gqlClient} from './client'

const Root = () => {
  const initialState = useContext(Context)
  const [state, dispatch] = useReducer(reducer, initialState)
  return (
    <ApolloProvider client={gqlClient}>
    <React.StrictMode>
        <BrowserRouter>
        <Context.Provider value={{state, dispatch}}>
        <Switch>
            <Route exact path='/' component={Main} />
            <Route path='/project/:id' component={Project} />
            <Route path='/create' component={Create} />
            <Route path='/login' component={Login} />
            <Route path='/register' component={Register} />
        </Switch>
        </Context.Provider>
        </BrowserRouter>
      </React.StrictMode>
    </ApolloProvider>
  )
}

ReactDOM.render(
  <Root />, document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
