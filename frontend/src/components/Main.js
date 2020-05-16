import React, { useContext } from 'react'
import Context from '../context'

const Main = props => {
  const { state } = useContext(Context)
  if (!state.currentUser) {
    props.history.push("/login")
  }
  return (
  <p>Hello, {state.currentUser?.email}</p>
  );
}

export default Main
