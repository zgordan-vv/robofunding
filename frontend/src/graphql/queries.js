import gql from 'graphql-tag'

export const doLogin = gql`
    query doLogin($email: String!, $password: String!) {
        doLogin(input: {
            email: $email,
            password: $password,
        })
    }
`

export const getMe = gql`
    query getMe {
        getMe {
            _id
            email
        }
    }
`