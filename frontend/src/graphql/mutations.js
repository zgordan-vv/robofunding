import gql from 'graphql-tag'

export const createUser = gql`
    mutation($email: String!, $password: String!, $passwordConfirmation: String!) {
        createUser(input: {
            email: $email,
            password: $password,
            passwordConfirmation: $passwordConfirmation
        }) {
            _id
        }
    }
`