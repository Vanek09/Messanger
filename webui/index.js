require('dotenv').config();
const tcp = require("./tcp")
const express = require("express")
const app = express()
const axios = require("axios")
const fs = require("fs")

app.use("/static", express.static("public"))
app.use(express.urlencoded({extended: true}))

app.set("view engine", "ejs")

app.get('/', (req, res) => {
    res.render("registration.ejs", {my_warning: null})
})

app.get('/user/:id', (req, res) => {
    const id = req.params.id
    axios.get(`${process.env.MONGO_API_URL}/getAll`).then(result => {
        if (result.data != null) {
            data = result.data
            // delete data["id"]
            for(i = 0; i < data.length; i++) {
                if (data[i]._id == id) {
                    delete data[i]
                }
            }
            res.render('todo.ejs', {users: data, messages: null, user: id, destination: null});
        } else {
            res.render('todo.ejs', {users: null, messages: null, user: id, destination: null});
        }
        // console.log(data)
        
    }).catch(({response}) => {
        console.log(response)
    })
})

app.get('/user/:id/:destination', (req,res) => {
    const id = req.params.id
    const destination = req.params.destination
    axios.get(`${process.env.MONGO_API_URL}/getAll`).then(result => {
        data = result.data
        // delete data["id"]
        console.log(`Recieved data: ${data}`)
        if (data != null) {
            for(i = 0; i < data.length; i++) {
                if (data[i]._id == id) {
                    delete data[i]
                }
            }
            tcp.recieveMessages(res, id, data, destination)
            // axios.get(`${process.env.MONGO_API_URL}/getMessages/${id}-${destination}`).then(result => {
            //     messages = result.data
            //     console.log(`DEBUG: ${messages}`)
            //     res_msg = []
            //     if(messages != null) {
            //         for(i = 0; i < messages.length; i++) {
            //             if( (messages[i].adress == id && messages[i].destination == destination) || (messages[i].adress == destination && messages[i].destination == id) ) {
            //                 res_msg.push(messages[i])
            //             }
            //         }
            //     }
            //     res.render('todo.ejs', {users: data, messages: messages, user: id, destination: destination});
            // }).catch(({response}) => {
            //     console.log(response)
            // })
        } else {
            axios.get(`${process.env.MONGO_API_URL}/getMessages/${id}-${destination}`).then(result => {
                res.render('todo.ejs', {users: null, messages: null, user: id, destination: destination});
            }).catch(({response}) => {
                console.log(response)
            })
        }
        
        // tcp.recieveMessages(res, id, data, destination)
        // console.log(data)
        
    }).catch(({response}) => {
        console.log(response)
    })
})

app.post('/send/:id/:destination', (req,res) => { 
    const id = req.params.id
    const destination = req.params.destination
    content = req.body.content
    axios.post(`${process.env.MONGO_API_URL}/sendMessage`, {
        message: req.body.content,
        adress: id,
        destination: destination
    })
    res.redirect(`/user/${id}/${destination}`)
})

app.post('/', (req, res) => {
    axios.get(`${process.env.MONGO_API_URL}/get/${req.body.user}`).then(result => {
        user = result.data
        if(user.hashed_pwd == req.body.pwd) {
            res.redirect(`user/${req.body.user}`)
        } else {
            res.render('registration.ejs', {my_warning: "Incorrect user or password"});
        }
    }).catch(({response}) => {
        console.log(response)
    })
})

app.post('/createUser', (req ,res) => {

    content = req.body
    try {
        axios.post(`${process.env.MONGO_API_URL}/put`, {
            _id: content.user,
            nickname: content.user,
            hashed_pwd: content.pwd
        }).then(result => {
            res.redirect(`user/${content.user}`)
        }).catch(({response}) => {
            console.log(response)
        })
    } catch {

    }
})


const PORT = process.env.PORT || 8080
// tcp.addUser("vanek01", "vanek01", "d8mjrm_Xa")
// tcp.getUsers()
// tcp.sendMessage("braurbeki", "vanek", "privet")
// tcp.recieveMessages("braurbeki", "vanek")
app.listen(PORT, () => {
    console.log(`App is running on port ${ PORT }`)
})