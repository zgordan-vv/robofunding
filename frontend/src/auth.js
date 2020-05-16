import * as firebase from 'firebase/app'
import 'firebase/auth'

const firebaseAuth = firebase.initializeApp({
    apiKey: "AIzaSyCaJjzr1WxbQVDUUFW0jbG3LO2lGx2rz1Q",
    authDomain: "robofunding-test-auth.firebaseapp.com",
    databaseURL: "https://robofunding-test-auth.firebaseio.com",
    projectId: "robofunding-test-auth",
    storageBucket: "robofunding-test-auth.appspot.com",
    messagingSenderId: "756904015278",
    appId: "1:756904015278:web:beff36e508ee1d400d043f"
})

export default firebaseAuth