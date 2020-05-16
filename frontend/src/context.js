import {createContext} from 'react'

const Context = createContext({
  sessionToken: "",
  currentUser: null,
})

export default Context