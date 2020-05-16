import React, { useState, useEffect, useContext } from 'react'
import {gqlClient} from '../client'
import {doLogin, getMe} from '../graphql/queries'
import Context from '../context'
import './styles.css'

const Login = props => {
    const { state, dispatch } = useContext(Context)
    if (state.currentUser?._id) {
        props.history.push('/')
    }
    const getMeReq = async () => {
        const getMeResp = await gqlClient().query({query: getMe})
        const me = getMeResp.getMe
        console.log("ME:", me, me?._id)
        dispatch({ action: 'GET_ME', payload: me })
        if (me?._id) {
            props.history.push('/')
        }
    }
    useEffect(() => {getMeReq()}, [])
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [errors, setErrors] = useState([])
    const submit = async e => {
        e.preventDefault()
        try {
            await gqlClient().query({query: doLogin, variables: {
                email,
                password,
            }})
            console.log('logged in')
            setErrors([])
            //props.history.push('/')
        } catch(err) {
            if (err.response?.errors) {
                setErrors(err.response.errors)
            } else {
                console.error(err)
            }
        }
    }
    /*const getMeResp = await gqlClient().request(getMe)
    const me = getMeResp.getMe
    console.log("ME:", me)
    if (me._id !== "") {
        props.history.push('/')
    }*/
    return (
        <div>
            <form onSubmit={submit}>
                <input
                    type='email'
                    value={email}
                    onChange={e => setEmail(e.target.value)}
                    placeholder='Enter your email'
                />
                <input
                    type='password'
                    onChange={e => setPassword(e.target.value)}
                    placeholder='Enter your password'
                    value={password}
                />
                <button
                    type="submit"
                    active={email && password}
                />
            </form>
            {errors.map((error, i) => <p key={i} className="error">{error.message}</p>)}
        </div>
    );
}

export default Login