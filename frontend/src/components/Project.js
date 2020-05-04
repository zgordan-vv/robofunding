import React from 'react'
import './styles.css'

const Project = props => {
    const {id} = props.match.params
    console.log(id)
    return (
        <p>Project {id}</p>
    );
}

export default Project