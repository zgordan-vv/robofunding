import React, { useState } from 'react'
import {gqlClient} from '../client'
import {createUser} from '../graphql/mutations'
import './styles.css'

const Register = props => {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [passwordConfirmation, setPasswordConfirmation] = useState('')
    const [errors, setErrors] = useState([])
    const submit = async (e) => {
        e.preventDefault()
        try {
            await gqlClient().request(createUser, {
                email,
                password,
                passwordConfirmation
            })
            setErrors([])
            props.history.push('/login')
        } catch(err) {
            if (err.response?.errors) {
                console.error(err.response.errors)
                setErrors(err.response.errors)
            } else {
                console.error(err)
            }
        }
    }
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
                <input
                    type='password'
                    value={passwordConfirmation}
                    onChange={e => setPasswordConfirmation(e.target.value)}
                    placeholder='Repeat your password'
                />
                <button
                    type="submit"
                    active={email && password && passwordConfirmation}
                />
            </form>
            {errors.map((error, i) => <p key={i} className="error">{error.message}</p>)}
        </div>
    );
}

export default Register