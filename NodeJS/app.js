const express = require('express')

const app = express()

app.use('/', (req, res) => {
    res.sendStatus(200)
})

    .listen(process.env.PORT || 3000, () => {
        console.log('•ᴗ• - servidor em execução');
    });